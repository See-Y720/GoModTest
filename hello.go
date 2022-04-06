package hello

import "fmt"

func Hello(name string) (retStr string) {
	retStr = fmt.Sprintf("hello,%s", name)
	return
}
