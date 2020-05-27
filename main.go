package main

/*
#include <stdlib.h>
#include <strings.h>

char** createArray(int n) {
  char** array = (char **) malloc(n * sizeof(char *));

  for (int i = 0; i < n; ++i) {
    array[i] = strdup("hello!");
  }

  return array;
}
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/danielgatis/go-ptrloop/ptrloop"
)

func main() {
	n := 3
	array := C.createArray(C.int(n))

	ptrloop.Loop(unsafe.Pointer(array), int(n), func(ptr unsafe.Pointer, i int) bool {
		s := *(**C.char)(ptr)
		fmt.Printf("%v -> %v\n", i, C.GoString(s))
		return true
	})
}
