package pygmy

// #cgo pkg-config: python3-embed
// #include "txllayer.h"
import "C"
import (
	"unsafe"

	wchar "github.com/GeertJohan/cgo.wchar"
)

type CPySSizeT = C.Py_ssize_t
type CPyConfig = C.CGO_PyConfig
type CPyObject = C.PyObject
type CPyPreConfig = C.CGO_PyPreConfig
type CPyStatus = C.PyStatus
type CPyWideStringList = C.PyWideStringList
type CPyWideString = *C.wchar_t
type Cwchar = C.wchar_t
type Cchar = C.char
type Cint = C.int
type Clong = C.long
type Cvoid = C.void

type PyArtifact[T comparable] struct {
	CInstance T
}

type PyConfig PyArtifact[CPyConfig]
type PyObject PyArtifact[CPyObject]
type PyPreConfig PyArtifact[CPyPreConfig]
type PyStatus PyArtifact[CPyStatus]

type Pygmy struct {
	PreConfig *PyPreConfig
	Config    *PyConfig
}

// Translate a boolean to a C type integer.
func Bool2CInt(obj bool) Cint {
	if obj {
		return 1
	} else {
		return 0
	}
}

func CFree[T comparable](obj *T) {
	C.free(unsafe.Pointer(obj))
}

// Translate a C type integer into a boolean.
func CInt2Bool(obj Cint) bool {
	return obj > 0
}

// Translate a wide char string into a Go string.
func WString2String(value CPyWideString) (string, error) {
	return wchar.FromWcharStringPtr(unsafe.Pointer(value)).GoString()
}

// Translate a Go string into a wide char string.
func String2WideString(value string) (CPyWideString, error) {
	w, err := wchar.FromGoString(value)
	return (CPyWideString)(w.Pointer()), err
}
