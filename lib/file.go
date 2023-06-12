
package lib

import (
  "errors"
  "fmt"
  "os"
  "path/filepath"
  "strings"
  "text/template"
  yaml "gopkg.in/yaml.v3"

  "github.com/imdario/mergo"
  "github.com/Masterminds/sprig/v3"
)

type FileContext map[string]any
 
type File struct {
  alias string
  path string
  dir string // containing directory
  MainTemplate *template.Template
  importMap map[string]*File
  globalState *GlobalState
  // TODO: reconsider how FileContext is handled.  Perhaps instead of passing it
  // down through files, always create a new clean FileContext at each level.
  context FileContext

  // lazy init variables, must be accessed via accessors
  repoRoot *string // accessor: mustGetRepoRoot()
}

func (f *File) Init(alias string, path string, globalState *GlobalState) *File {
  f.alias = alias
  f.path = filepath.Clean(path)
  // TODO: perhaps list specific paths rather than just this prefix.
  if _, foundPrefix := strings.CutPrefix(f.path, "/dev/"); foundPrefix {
    // TODO: or perhaps better the directory of the including file, defaulting
    // to CWD if this is the root template.
    f.dir = Must(os.Getwd())
  } else {
    f.dir = filepath.Dir(f.path)
  }
  f.globalState = globalState
  f.MainTemplate = template.New(f.alias).Funcs(sprig.FuncMap()).
                                         Funcs(f.globalState.FuncMap).
                                         Funcs(template.FuncMap{
    "import_as": f.tfImportAs,
    "export_from": f.tfExportFrom,
    "run": f.tfRun,
    "include": f.tfInclude,
    "hoist_file": f.tfHoistFile,
  })
  f.importMap = map[string]*File{}
  f.importMap["self"] = f
  f.context = FileContext{}
  return f
}

// TODO: replace alias with return value containing an object with a Run method
// that effectively calls "run".  Also remove the "run" function.
func (f *File) tfImportAs(alias string, importedFilePath string) string {
  if alias == "__main__" {
    panic("Not allowed to import_as __main__.  __main__ is reserved.")
  }

  if f.importMap[alias] != nil {
    panic("Import alias already used in this scope: \"" + alias + "\"")
  }

  f.importMap[alias] = f.loadFile(importedFilePath, alias)

  // TODO(clean): can the function have no return value instead of any empty
  // string?
  return ""
}

func (f *File) loadFile(loadingPath string, alias string) *File {
  // Handle "//" and relative paths (no change to abs paths).
  if fileNameNoPrefix, prefixFound := strings.CutPrefix(loadingPath, "//");
        prefixFound {
    loadingPath = filepath.Join(f.mustGetRepoRoot(), fileNameNoPrefix)
  } else if filepath.IsAbs(loadingPath) {
    // Nothing to do, already an absolute path.
  } else {  // It's a relative path (relative to this file's dir, not CWD).
    loadingPath = filepath.Join(f.dir, loadingPath)
  }

  loadingFile := new(File).Init(alias, loadingPath, f.globalState)
  Must(loadingFile.MainTemplate.Parse(string(Must(ReadFileOrStdin(loadingPath)))))
  MustExecuteTemplate(loadingFile.MainTemplate, nil)
  return loadingFile
}

func (f *File) mustGetRepoRoot() string {
  if f.repoRoot != nil { return *f.repoRoot }

  f.repoRoot = new(string)
  prevPath := f.path
  for {
    dir := filepath.Dir(prevPath)
    if prevPath == dir {
      panic("No WORKSPACE file found in any ancestor directory.")
    }
    if isRepoRoot(dir) {
      *f.repoRoot = dir
      return *f.repoRoot
    }
    prevPath = dir
  }
}

func isRepoRoot(path string) bool {
  // TODO(feature): support other repo root strategies than just 'WORKSPACE'.
  _, err := os.Stat(filepath.Join(path, "WORKSPACE"))
  return err == nil
}

func(f *File) tfExportFrom(importName string, templateName string) string {
  importFile := f.importMap[importName]
  if importFile == nil {
    panic(errors.New("Unknown import name: " + importName))
  }

  deepTemplate := importFile.MainTemplate.Lookup(templateName)
  if deepTemplate == nil {
    panic(errors.New("Unknown template: import=" + importName +
                     " template=" + templateName))
  }

  f.MainTemplate.Parse(fmt.Sprintf(
      `{{define "%v"}} {{- run "%v" "%v" . -}} {{end}}`,
      templateName, importName, templateName))

  return ""
}

func(f *File) tfRun(importName string, templateName string, context any) string {
  requestedFile := f.importMap[importName]
  if requestedFile == nil {
    panic(errors.New("Unknown import: " + importName))
  }

  requestedTemplate := requestedFile.MainTemplate.Lookup(templateName)
  if requestedTemplate == nil {
    panic(errors.New("Unknown template called: import=" + importName +
                     " template=" + templateName))
  }
  return MustExecuteTemplate(requestedTemplate, context)
}

func(f *File) tfInclude(templateName string, context any) string {
  return f.tfRun("self", templateName, context)
}

func(f *File) tfHoistFile(sailFilePath string,
                          paramTemplateName string) string {
  sailFile := f.loadFile(sailFilePath, "__sail_file__:" + sailFilePath)
  sailParams := map[string]any{}
  sailFile.context["params"] = sailParams

  defaultParamsTemplate := sailFile.MainTemplate.Lookup("default_params")
  if defaultParamsTemplate != nil {
    if err := yaml.Unmarshal([]byte(MustExecuteTemplate(defaultParamsTemplate, nil)),
                             &sailParams); err != nil {
      panic(errors.New("Sail's 'default_param' template must be valid YAML: " + err.Error()))
    }
  }

  if paramTemplateName != "" {
    paramTemplate := f.MainTemplate.Lookup(paramTemplateName)
    if paramTemplate == nil {
      panic(errors.New("Hoist param template not found: " + paramTemplateName))
    }
    overrideParams := map[string]any{}
    if err := yaml.Unmarshal([]byte(MustExecuteTemplate(paramTemplate, f.context)),
                             &overrideParams); err != nil {
      panic(errors.New("hoist param template must be valid YAML: " + err.Error()))
    }
    mergo.Merge(&sailParams, overrideParams, mergo.WithOverride)
  }

  sailTemplate := sailFile.MainTemplate.Lookup("sail")
  if sailTemplate == nil {
    panic(errors.New("hoist of sail file failed, 'sail' template not found in file: " +
                     sailFilePath))
  }
  return MustExecuteTemplate(sailTemplate, sailFile.context)
}

