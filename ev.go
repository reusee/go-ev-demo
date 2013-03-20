package main

/*
#include "zmalloc.c"
#include "ae.c"
#include "ev.c"
*/
import "C"
import (
  "fmt"
  "unsafe"
)

func main() {
  loop := C.aeCreateEventLoop(C.int(1024))

  c := make(chan bool)
  n := 1000

  callback := func() {
    fmt.Printf("foo\n")
    c <- true
  }
  for i := 0; i < n; i++ {
    C.create_time_event(loop, C.longlong(i), unsafe.Pointer(&callback))
  }

  go C.aeMain(loop)
  for i := 0; i < n; i++ {
    <-c
  }
}
