
package lib

import (
  "fmt"
  "errors"
  "os"
  "text/template"
  yaml "gopkg.in/yaml.v3"
)

type Global struct {}

func (c *Global) tfCheese(param string) string {
  return fmt.Sprintf("CHEESE %v CHEESE", param)
}

func (g *Global) Run(templatePath string, dataPath string) {
  if templatePath == "-" && dataPath == "-" {
    panic(errors.New("both template and data file paths were set to stdin"));
  }

  globalFuncs := template.FuncMap {
    "cheese": g.tfCheese,
    "blank": func(ignored ...any) string { return "" },
  }

  var dataGolang any
  Must1(yaml.Unmarshal(Must(ReadFileOrStdin(dataPath)), &dataGolang))

  templateString := string(Must(ReadFileOrStdin(templatePath)))

  f := new(File).Init(templatePath, "__main__", globalFuncs)
  f.MainTemplate.Parse(templateString)
  Must1(f.MainTemplate.Execute(os.Stdout, dataGolang))
}


