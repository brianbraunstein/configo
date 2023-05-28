package main

import (
  "flag"
  configo "bristyle.com/configo/lib"
)

func main() {
  templatePath := flag.String("template", "-", "template file path")
  dataPath := flag.String("data", "-", "data file path")
  flag.Parse()
  (&configo.Global{}).Run(*templatePath, *dataPath)
}

