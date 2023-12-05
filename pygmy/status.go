package pygmy

// #cgo pkg-config: python3-embed
// #include "txllayer.h"
import "C"

func PyStatusNew(cStatus CPyStatus) PyStatus {
	return PyStatus{CInstance: cStatus}
}

// Initialize an error PyStatus.
// https://docs.python.org/3/c-api/init_config.html#c.PyStatus_Error
func PyStatusError(err string) PyStatus {
	cErr := C.CString(err)
	defer CFree(cErr)
	ret := C.PyStatus_Error(cErr)
	return PyStatusNew(ret)
}

// Initialize a PyStatus with an exit code.
// https://docs.python.org/3/c-api/init_config.html#c.PyStatus_Exit
func PyStatusExit(code int) PyStatus {
	ret := C.PyStatus_Exit((Cint)(code))
	return PyStatusNew(ret)
}

// Initialize an allocation failure PyStatus.
// https://docs.python.org/3/c-api/init_config.html#c.PyStatus_NoMemory
func PyStatusNoMemory() PyStatus {
	ret := C.PyStatus_NoMemory()
	return PyStatusNew(ret)
}

// Initialize a success PyStatus.
// https://docs.python.org/3/c-api/init_config.html#c.PyStatus_Ok
func PyStatusOk() PyStatus {
	ret := C.PyStatus_Ok()
	return PyStatusNew(ret)
}

// The error message.
// https://docs.python.org/3/c-api/init_config.html#c.PyStatus.err_msg
func (status *PyStatus) ErrMessage() string {
	return C.GoString(status.CInstance.err_msg)
}

// The name of the function which created an
// error.
// https://docs.python.org/3/c-api/init_config.html#c.PyStatus.func
func (status *PyStatus) FuncName() string {
	return C.GoString(status.CInstance._func)
}

// The exit code.
// https://docs.python.org/3/c-api/init_config.html#c.PyStatus.exitcode
func (status *PyStatus) ExitCode() int {
	return int(status.CInstance.exitcode)
}

// Display exception message and exit the process.
func (status *PyStatus) ExitStatusException() {
	C.Py_ExitStatusException(status.CInstance)
}

// Check if status is an error.
// https://docs.python.org/3/c-api/init_config.html#c.PyStatus_IsError
func (status *PyStatus) IsError() bool {
	return C.PyStatus_IsError(status.CInstance) != 0
}

// Check if a status is an exception or exit
// status. Must be handled if the result is true.
// https://docs.python.org/3/c-api/init_config.html#c.PyStatus_Exception
func (status *PyStatus) IsException() bool {
	return C.PyStatus_Exception(status.CInstance) != 0
}

// Check if status has an exit code.
// https://docs.python.org/3/c-api/init_config.html#c.PyStatus_IsExit
func (status *PyStatus) IsExit() bool {
	return C.PyStatus_IsExit(status.CInstance) != 0
}
