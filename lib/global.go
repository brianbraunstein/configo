
package lib

import (
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

// TODO: look into a solution to handle subnetNumber as an int64 that doesn't
// require lots of ugly conversion code from "any".  Generics/type constraints
// don't seem to be allowed for these template functions.
func (g *Global) tfCidrSubnet(cidr string, subnetBits int, subnetNumber int) string {
  return cidrSubnet(cidr, subnetBits, int64(subnetNumber))
}

func (g *Global) Run(templatePath string, dataPath string) {
  if templatePath == "-" && dataPath == "-" {
    panic(errors.New("both template and data file paths were set to stdin"));
  }

  g.FuncMap = template.FuncMap {
    "blank": func(ignored ...any) string { return "" },
    "toYaml": g.tfToYaml,
    "fromYaml": g.tfFromYaml,
    "cidrSubnet": g.tfCidrSubnet,
  }

  var dataGolang any
  if dataPath != "" {
    Must1(yaml.Unmarshal(Must(ReadFileOrStdin(dataPath)), &dataGolang))
  }

  templateString := string(Must(ReadFileOrStdin(templatePath)))

  if templatePath == "-" {
    templatePath = "./__stdin__"
  }
  f := new(File).Init("__main__", templatePath, &g.GlobalState)
  f.MainTemplate.Parse(templateString)
  Must1(f.MainTemplate.Execute(os.Stdout, dataGolang))
}

