
package lib

import (
  "bytes"
  "errors"
  "fmt"
  "io"
  "text/template"

  "github.com/Masterminds/sprig/v3"
)
 
type File struct {
  fileName string
  MainTemplate *template.Template
  importMap map[string]*File
  globalFuncs template.FuncMap
}

func (f *File) Init(fileName string, name string, globalFuncs template.FuncMap) *File {
  f.fileName = fileName
  f.globalFuncs = globalFuncs
  f.MainTemplate = template.New(name).Funcs(sprig.FuncMap()).
                                            Funcs(globalFuncs).
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

func(f *File) tfImportAs(name string, fileName string) string {
  if f.importMap[name] != nil {
    panic("Import name already used in this scope: \"" + name + "\"")
  }

  importedFile := new(File).Init(fileName, fileName, f.globalFuncs)
  importedFile.MainTemplate.Parse(string(Must(ReadFileOrStdin(fileName))))
  Must1(importedFile.MainTemplate.Execute(io.Discard, nil))
  f.importMap[name] = importedFile

  return ""
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

