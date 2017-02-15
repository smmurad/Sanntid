package main
/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "io.c"
*/
import "C"

import("fmt")


func main() {
	fmt.Println("fgakkogem")
	C.io_init()
}