package test2

import (
	"testing"
)

type Aint int

type A interface {
	Table(int) int
}

func (ai *Aint) Table(a int) int {
	return int(*ai) + 2
}

func TestA(t *testing.T) {

}