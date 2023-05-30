
package lib

import (
  "io"
  "os"
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
