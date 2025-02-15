
package gotest

import (
  // this is a test.
  "testing"

  // printing and formatting.
  "fmt"

  // other imports.
  "github.com/kraasch/godiff/godiff"
)

var (
  NL = fmt.Sprintln()
)

type TestList struct {
  TestName          string
  IsMulti           bool
  InputArr          []string
  ExpectedValue     string
}

type TestSuite struct {
  TestingFunction   func(in TestList) string
  Tests             []TestList
}

var suites = []TestSuite{}

func DoTest(t *testing.T, ss []TestSuite) {
  suites = ss
  TestAll(t)
}

func TestAll(t *testing.T) {
  for _, suite := range suites {
    for _, test := range suite.Tests {
      name := test.TestName
      t.Run(name, func(t *testing.T) {
        exp := test.ExpectedValue
        got := suite.TestingFunction(test)
        if exp != got {
          if test.IsMulti {
            t.Errorf("In '%s':\n", name)
            diff := godiff.CDiff(exp, got)
            t.Errorf("\nExp: '%#v'\nGot: '%#v'\n", exp, got)
            t.Errorf("exp/got:\n%s\n", diff)
          } else {
            t.Errorf("In '%s':\n  Exp: '%#v'\n  Got: '%#v'\n", name, exp, got)
          }
        }
      })
    }
  }
}

