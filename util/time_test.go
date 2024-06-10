package util

import (
	"fmt"
	"testing"
)

func TestMustParse(t *testing.T) {
	ti := MustParse("2024-04-22T13:33:50.626Z")
	fmt.Println(ti)
}

func TestParse(t *testing.T) {

}

func TestTryParse(t *testing.T) {

}
