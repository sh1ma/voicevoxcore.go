package voicevoxcorego

// #include <stdint.h>
import "C"
import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Float | constraints.Integer
	comparable
}

type ctype interface {
	C.char | C.uchar |
		C.short | C.ushort |
		C.int | C.uint |
		C.long | C.ulong |
		C.longlong | C.ulonglong |
		C.float | C.double
}

func sliceToCtype[T number, Return ctype](n []T) Return {
	first := n[0]
	return (Return)(first)
}

func makeDataReceiver[Return any, Size any]() (*Return, *Size, []Return, Size) {
	var size Size
	data := make([]Return, 1)
	dataFirstPointer := &data[0]
	sizePointer := (*Size)(unsafe.Pointer(&size))

	return dataFirstPointer, sizePointer, data, size
}
