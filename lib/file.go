
package lib

import (
  "bytes"
  "errors"
  "fmt"
  "io"
  "os"
  "path/filepath"
  "strings"
  "text/template"

  "github.com/Masterminds/sprig/v3"
)
 
type File struct {
  alias string
  fileName string  // TODO: rename path
  dir string // containing directory
  MainTemplate *template.Template
  importMap map[string]*File
  globalState *GlobalState

  // lazy init variables, must be accessed via accessors
  repoRoot *string // accessor: mustGetRepoRoot()
}

func (f *File) Init(alias string, fileName string, globalState *GlobalState) *File {
  f.alias = alias
  f.fileName = filepath.Clean(fileName)
  if _, foundPrefix := strings.CutPrefix(f.fileName, "/dev/"); foundPrefix {
    // TODO: or perhaps better the directory of the including file, defaulting
    // to CWD if this is the root template.
    f.dir = Must(os.Getwd())
  } else {
    f.dir = filepath.Dir(f.fileName)
  }
  f.globalState = globalState
  f.MainTemplate = template.New(f.alias).Funcs(sprig.FuncMap()).
                                         Funcs(f.globalState.FuncMap).
                                         Funcs(template.FuncMap{
    "import_as": f.tfImportAs,
    "export_from": f.tfExportFrom,
    "include": f.tfInclude,
    "run": f.tfRun,
  })
  f.importMap = map[string]*File{}
  f.importMap["self"] = f
  return f
}

func(f *File) tfImportAs(alias string, importedFilePath string) string {
  if alias == "__main__" {
    panic("Not allowed to import_as __main__.  __main__ is reserved.")
  }

  if f.importMap[alias] != nil {
    panic("Import alias already used in this scope: \"" + alias + "\"")
  }

  // Handle "//" and relative paths (no change to abs paths).
  if fileNameNoPrefix, prefixFound := strings.CutPrefix(importedFilePath, "//");
        prefixFound {
    importedFilePath = filepath.Join(f.mustGetRepoRoot(), fileNameNoPrefix)
  } else if filepath.IsAbs(importedFilePath) {
    // Nothing to do, already an absolute path.
  } else {  // It's a relative path (relative to this file's dir, not CWD).
    importedFilePath = filepath.Join(f.dir, importedFilePath)
  }

  importedFile := new(File).Init(alias, importedFilePath, f.globalState)
  importedFile.MainTemplate.Parse(string(Must(ReadFileOrStdin(importedFilePath))))
  Must1(importedFile.MainTemplate.Execute(io.Discard, nil))
  f.importMap[alias] = importedFile

  // TODO(clean): can the function have no return value instead of any empty
  // string?
  return ""
}

func (f *File) mustGetRepoRoot() string {
  if f.repoRoot != nil { return *f.repoRoot }

  f.repoRoot = new(string)
  prevPath := f.fileName
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

  buf := bytes.Buffer{}
  if err := requestedTemplate.Execute(&buf, context); err != nil {
    // Newline required to make the error message readable.
    panic(errors.New("\n" + err.Error()))
  }
  return buf.String()
}

func(f *File) tfInclude(templateName string, context any) string {
  return f.tfRun("self", templateName, context)
}

