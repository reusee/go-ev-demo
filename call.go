package main

import "C"
import (
  "unsafe"
)

//export call_func
func call_func(fun unsafe.Pointer) {
  (*(*func())(fun))()
}
