
package lib

import (
  "testing"
)

type testCase struct {
  cidr string
  subnetBits int
  subnetNumber int64
  expected string
}

func (tc *testCase) Run(t *testing.T) {
  actual := cidrSubnet(tc.cidr, tc.subnetBits, tc.subnetNumber)
  if tc.expected == actual { return }
  t.Errorf("\nExpected: %v" +
           "\nActual: %v" +
           "\nArgs: %v", tc.expected, actual, *tc)
}

func TestCidrSubnet(t *testing.T) {
  testCases := []testCase{
    {cidr: "10.0.0.3/16", subnetBits: 4, subnetNumber: 3,
     expected: "10.0.48.0/20"},
    {cidr: "10.0.0.3/16", subnetBits: 0, subnetNumber: 0,
     expected: "10.0.0.0/16"},
    {cidr: "10.0.0.3/16", subnetBits: 16, subnetNumber: 1 << 15 + 7,
     expected: "10.0.128.7/32"},
  }
  for _, tc := range testCases { tc.Run(t) }
}

