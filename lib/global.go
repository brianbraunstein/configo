
package lib

import (
  "fmt"
  "errors"
  "os"
  "strings"
  "text/template"
  yaml "gopkg.in/yaml.v3"
)

type GlobalState struct {
  FuncMap template.FuncMap
}

type Global struct {
  GlobalState
}

func (g *Global) tfCheese(param string) string {
  return fmt.Sprintf("CHEESE %v CHEESE", param)
}

func (g *Global) tfToYaml(goData any) string {
  yamlData, err := yaml.Marshal(goData)
  if err != nil { panic(err) }
  return strings.TrimSuffix(string(yamlData), "\n")
}

func (g *Global) tfFromYaml(yamlData string) map[string]any {
  goData := map[string]any{}
  err := yaml.Unmarshal([]byte(yamlData), &goData)
  if err != nil { panic(err) }
  return goData
}

func (g *Global) Run(templatePath string, dataPath string) {
  if templatePath == "-" && dataPath == "-" {
    panic(errors.New("both template and data file paths were set to stdin"));
  }

  g.FuncMap = template.FuncMap {
    "cheese": g.tfCheese,
    "blank": func(ignored ...any) string { return "" },
    "toYaml": g.tfToYaml,
    "fromYaml": g.tfFromYaml,
  }

  var dataGolang any
  if dataPath != "" {
    Must1(yaml.Unmarshal(Must(ReadFileOrStdin(dataPath)), &dataGolang))
  }

  templateString := string(Must(ReadFileOrStdin(templatePath)))

  if templatePath == "-" {
    templatePath = "./__stdin__"
  }
  f := new(File).Init(templatePath, &g.GlobalState)
  f.MainTemplate.Parse(templateString)
  Must1(f.MainTemplate.Execute(os.Stdout, dataGolang))
}

