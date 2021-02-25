package main

/*
#include <stdio.h>
void foobar() {
	printf("hello world i am printing with C code");
}
*/
import "C"
// comments above this line will be interpreted as C code. must be directly above, no spaced newline.
// import "C" must also be isolated from other imports, otherwise it will fail to recognize we want to call C code
import "fmt"

func main() {
	fmt.Printf("hello world i am printing with go\n")
	//C.printf("hello world i am printing with cgo")
	C.foobar()
}
