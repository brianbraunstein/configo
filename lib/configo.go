
package lib

import (
  "fmt"
  "errors"
  "io"
  "os"
  "github.com/Masterminds/sprig/v3"
  "text/template"
  yaml "gopkg.in/yaml.v3"
)

func Must1(err error) {
  if err != nil { panic(err) }
}

func Must[T any](result T, err error) T {
  if err != nil { panic(err) }
  return result
}

func ReadFileOrStdin(path string) ([]byte, error) {
  if path == "-" {
    return io.ReadAll(os.Stdin)
  } else {
    return os.ReadFile(path)
  }
}

var funcs = template.FuncMap {
  "cheese": func(param string) string {
    return fmt.Sprintf("CHEESE %v CHEESE", param)
  },
}

func Run(templatePath string, dataPath string) {
  if templatePath == "-" && dataPath == "-" {
    panic(errors.New("both template and data file paths were set to stdin"));
  }

  var dataGolang any
  Must1(yaml.Unmarshal(Must(ReadFileOrStdin(dataPath)), &dataGolang))

  templateString := string(Must(ReadFileOrStdin(templatePath)))
  tpl := Must(template.New("base").Funcs(sprig.FuncMap()).Funcs(funcs).Parse(templateString))
  Must1(tpl.Execute(os.Stdout, dataGolang))
}

