
package lib

import (
  "bytes"
  "errors"
  "io"
  "os"
  "text/template"
)

// TODO: rename
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

func MustExecuteTemplate(t *template.Template, context any) string {
  buf := bytes.Buffer{}
  if err := t.Execute(&buf, context); err != nil {
    panic(errors.New("\n" + err.Error()))
  }
  return buf.String()
}

