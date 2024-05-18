package main

import (
	"os"
	"path/filepath"

	"github.com/golang-collections/collections/stack"
)

var DIR_STACK *stack.Stack = stack.New()

func Pushd(dir string) {
	DIR_STACK.Push(dir)
	os.Chdir(dir)
}

func Popd() string {
	dir := DIR_STACK.Pop().(string)
	os.Chdir(DIR_STACK.Peek().(string))
	return dir
}

func RelDir(dir string) string {
	return filepath.Join(DIR_STACK.Peek().(string), filepath.Dir(dir))
}
