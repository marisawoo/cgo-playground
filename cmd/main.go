package main

/*
#include <stdio.h>
#include <stdlib.h>
void printMsg(char *str) {
	printf("%s", str);
}
*/
import "C"
// comments above this line will be interpreted as C code. must be directly above, no spaced newline.
// import "C" must also be isolated from other imports, otherwise it will fail to recognize we want to call C code
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("hello world i am printing with Go\n")
	cstr := C.CString("hello world i am printing with C")
	C.printMsg(cstr)
	C.free(unsafe.Pointer(cstr)) // C string allocated in C memory, i.e. we need to free
}
