package main

import (
  "flag"
  "fmt"
  configo "github.com/brianbraunstein/configo/lib"
)

func main() {
  templatePath := flag.String("template", "-", "template file path")
  dataPath := flag.String("data", "", "data file path")
  version := flag.Bool("version", false, "Print version and exit")
  flag.Parse()
  if *version {
    fmt.Println(configoVersion())
    return
  }
  (&configo.Global{}).Run(*templatePath, *dataPath)
}

