package main

import (
	"fmt"
	"os"
	"testing"
)

// panic for unexpected error
func TestPacin(t *testing.T) {
	panic("a problem")
	_,err := 	os.Create("etc/go_config")
	if err != nil {
		panic(err)
	}
}

//defer
func TestDefer(t *testing.T) {
	f := CreateFile("defer.txt")
	defer CloseFile(f)
	WriteFile(f)
}

func CreateFile(p string) *os.File{
	fmt.Println("Creating")
	t,err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return t
}

func WriteFile(t *os.File) {
	fmt.Println("Writing")
	fmt.Fprintf(t,"data")
}

func CloseFile(t *os.File) {
	fmt.Println("closing")
	t.Close()
}
func MyPanic () {
	panic("a problem")
}

func TestRecover(t *testing.T) {
	defer func(){
		if r := recover();r != nil {
			fmt.Println("recovering")
		}
	}()
	MyPanic()
	fmt.Println("after mypanic")
}