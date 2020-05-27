package ptrloop

import (
	"unsafe"
)

type cbf func(unsafe.Pointer, int) bool

// Loop Given a pointer a length and a function, iterate over the pointer executing the function while it returns true
func Loop(ptr unsafe.Pointer, length int, f cbf) {
	size := unsafe.Sizeof(ptr)

	for i := 0; i < length; i++ {
		stop := !f(ptr, i)

		if stop {
			break
		}

		ptr = unsafe.Pointer(uintptr(ptr) + size)
	}
}
