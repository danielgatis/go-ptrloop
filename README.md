# go-ptrloop

[![Go Report Card](https://goreportcard.com/badge/github.com/danielgatis/go-ptrloop?style=flat-square)](https://goreportcard.com/report/github.com/danielgatis/go-ptrloop)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/danielgatis/go-ptrloop/master/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/danielgatis/go-ptrloop)

A helper to iterate over unsafe pointers.

## Install

```bash
go get -u github.com/danielgatis/go-ptrloop
```

And then import the package in your code:

```go
import "github.com/danielgatis/go-ptrloop/ptrloop"
```

### Example

An example described below is one of the use cases.

```go
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
```


```
❯ go run main.go
0 -> hello!
1 -> hello!
2 -> hello!
```


## License

Copyright (c) 2020-present [Daniel Gatis](https://github.com/danielgatis)

Licensed under [MIT License](./LICENSE)
