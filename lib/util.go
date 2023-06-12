
package lib

import (
  "bytes"
  "errors"
  "fmt"
  "io"
  "net"
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

func cidrSubnet(cidr string, subnetBits int, subnetNumber int) string {
  _, cidrNet, err := net.ParseCIDR(cidr)
  if err != nil {
    panic(err.Error())
  }

  cidrBits, _ := cidrNet.Mask.Size()
  afterBits := cidrBits + subnetBits
  if afterBits > 32 {
    panic(errors.New(fmt.Sprintf(
      "CIDR bits + subnetBits exceeds 32: cidr=%v subnetBits=%v", cidr,
      subnetBits)))
  }

  if subnetNumber >= (1 << subnetBits) {
    panic(errors.New(fmt.Sprintf(
      "CIDR subnetNumber too big for subnet size: subnetBits=%v subnetNumber=%v",
      subnetBits, subnetNumber)))
  }

  cidrNet.Mask = net.CIDRMask(afterBits, 32)
  return cidrNet.String()
}

