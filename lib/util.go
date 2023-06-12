
package lib

import (
  "bytes"
  "errors"
  "fmt"
  "io"
  "math/big"
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

func cidrSubnet(cidr string, subnetBits int, subnetNumber int64) string {
  _, cidrNet, err := net.ParseCIDR(cidr)
  if err != nil {
    panic(err.Error())
  }

  cidrBits, totalBits := cidrNet.Mask.Size()
  remainingBits := totalBits - cidrBits - subnetBits
  if remainingBits < 0 {
    panic(errors.New(fmt.Sprintf(
      "Used more bits than available in the CIDR: cidr=%v subnetBits=%v",
      cidr, subnetBits)))
  }

  if subnetNumber >= (1 << subnetBits) {
    panic(errors.New(fmt.Sprintf(
      "CIDR subnetNumber too big for subnet size: subnetBits=%v subnetNumber=%v",
      subnetBits, subnetNumber)))
  }

  cidrNet.Mask = net.CIDRMask(cidrBits + subnetBits, totalBits)

  cidrIpInt := big.NewInt(0) 
  cidrIpInt.Or(big.NewInt(0).SetBytes(cidrNet.IP),
               big.NewInt(0).Lsh(big.NewInt(subnetNumber), uint(remainingBits)))
  cidrNet.IP = cidrIpInt.Bytes()
  return cidrNet.String()
}

