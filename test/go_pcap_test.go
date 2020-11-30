package test

import (
	"fmt"
	"testing"
)

func TestUseFmt(t *testing.T) {
}

func TestFor(t *testing.T) {
	ss := []string{"aq", "bd", "cc", "bb"}
	for _, v := range ss {
		fmt.Println(v)
	}
}

func TestError(t *testing.T) {
	fmt.Errorf("new error %s", "aaa")
}
