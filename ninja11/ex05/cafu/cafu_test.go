package cafu

import (
  "fmt"
  "testing"
)

func TestGetPrice(t *testing.T) {
  got := GetPrice()
  expected := 50

  fmt.Println("GetPrice:", got)

  if got < expected {
    t.Error("got", got, "expected", expected, "or higher!")
  }
}
