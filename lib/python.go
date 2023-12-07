package pyg

// #cgo pkg-config: python3-embed
// #include "txllayer.h"
import "C"
import (
	"fmt"
)

// PygNew
//
// Get a new instance of Python interpreter with
// preset config and preconfig objects.
func PygNew(config *PyConfig, preconfig *PyPreConfig) *Pyg {
	if config == nil {
		config = ConfigNew()
		config.InitPythonConfig()
	}
	if preconfig == nil {
		preconfig = PreConfigNew()
		preconfig.InitPythonConfig()
	}
	return &Pyg{Config: config, PreConfig: preconfig}
}

// PygNewPython
//
// Get a new instance of Python interpreter.
func PygNewPython() *Pyg {
	return PygNew(nil, nil)
}

// PygNewIsolated
//
// Get a new isolated instance of Python
// interpreter.
func PygNewIsolated() *Pyg {
	preconfig, config := PreConfigNew(), ConfigNew()
	preconfig.InitIsolatedConfig()
	config.InitIsolatedConfig()
	return PygNew(config, preconfig)
}

func (pygi *Pyg) checkInitialized() error {
	if !pygi.IsInitialized() {
		return fmt.Errorf("interpreter not yet initialized")
	}
	return nil
}

func (pygi *Pyg) checkInitializedS(val string) (string, error) {
	if err := pygi.checkInitialized(); err != nil {
		return "", err
	}
	return val, nil
}

// FinalizeEx
//
// Undo all initializations made by
// Py_Initialize() and subsequent use of Python/C
// API functions, and destroy all sub-interpreters
// (see Py_NewInterpreter() below) that were
// created and not yet destroyed since the last
// call to Py_Initialize().
// https://docs.python.org/3/c-api/init.html#c.Py_FinalizeEx
func (pygi *Pyg) FinalizeEx() int {
	pygi.Config.Clear()
	pygi.PreConfig.Clear()
	return int(C.Py_FinalizeEx())
}

// InitializeFromConfig
//
// Initialize Python from config configuration.
// https://docs.python.org/3/c-api/init_config.html#c.Py_InitializeFromConfig
func (pygi *Pyg) InitializeFromConfig() PyStatus {
	ret := C.CGO_Py_InitializeFromConfig(&pygi.Config.CInstance)
	return StatusNew(ret)
}

// IsInitialized
//
// Whether the Python interpreter has been
// initialized or not.
// https://docs.python.org/3/c-api/init.html#c.Py_IsInitialized
func (pygi *Pyg) IsInitialized() bool {
	return CInt2Bool(C.Py_IsInitialized())
}

// GetBuildInfo
//
// Return information about the sequence number
// and build date and time of the current Python
// interpreter.
// https://docs.python.org/3/c-api/init.html#c.Py_GetBuildInfo
func (pygi *Pyg) GetBuildInfo() string {
	return C.GoString(C.Py_GetBuildInfo())
}

// GetCompiler
//
// Return an indication of the compiler used to
// build the current Python version.
// https://docs.python.org/3/c-api/init.html#c.Py_GetCompiler
func (pygi *Pyg) GetCompiler() string {
	return C.GoString(C.Py_GetCompiler())
}

// GetCopyright
//
// Return the official copyright for the current
// Python version.
// https://docs.python.org/3/c-api/init.html#c.Py_GetCopyright
func (pygi *Pyg) GetCopyright() string {
	return C.GoString(C.Py_GetCopyright())
}

// GetExecPrefix
//
// Return the exec-prefix for installed
// platform-dependent files.
// https://docs.python.org/3/c-api/init.html#c.Py_GetExecPrefix
func (pygi *Pyg) GetExecPrefix() (string, error) {
	return pygi.checkInitializedS(WString2String(C.Py_GetExecPrefix()))
}

// GetPath
//
// Return the default module search path.
// https://docs.python.org/3/c-api/init.html#c.Py_GetPath
func (pygi *Pyg) GetPath() (string, error) {
	return pygi.checkInitializedS(WString2String(C.Py_GetPath()))
}

// GetPlatform
//
// Return the platform identifier for the current
// platform.
// https://docs.python.org/3/c-api/init.html#c.Py_GetPlatform
func (pygi *Pyg) GetPlatform() string {
	return C.GoString(C.Py_GetPlatform())
}

// GetPrefix
//
// Return the prefix for installed platform-independent files.
// https://docs.python.org/3/c-api/init.html#c.Py_GetPrefix
func (pygi *Pyg) GetPrefix() (string, error) {
	return pygi.checkInitializedS(WString2String(C.Py_GetPrefix()))
}

// GetProgramFullPath
//
// Return the full program name of the Python
// executable.
// https://docs.python.org/3/c-api/init.html#c.Py_GetProgramFullPath
func (pygi *Pyg) GetProgramFullPath() (string, error) {
	return pygi.checkInitializedS(WString2String(C.Py_GetProgramFullPath()))
}

// GetProgramName
//
// Return the program name set by Python
// configuration.
// https://docs.python.org/3/c-api/init.html#c.Py_GetProgramName
func (pygi *Pyg) GetProgramName() (string, error) {
	return pygi.checkInitializedS(WString2String(C.Py_GetProgramName()))
}

// GetPythonHome
//
// Return the default "home", that is, the value
// set by configuration.
// https://docs.python.org/3/c-api/init.html#c.Py_GetPythonHome
func (pygi *Pyg) GetPythonHome() (string, error) {
	return pygi.checkInitializedS(WString2String(C.Py_GetPythonHome()))
}

// GetVersion
//
// Return the version of this Python interpreter.
// https://docs.python.org/3/c-api/init.html#c.Py_GetVersion
func (pygi *Pyg) GetVersion() string {
	return C.GoString(C.Py_GetVersion())
}

// PreInitialize
//
// Preinitialize Python from preconfig
// preconfiguration.
// https://docs.python.org/3/c-api/init_config.html#c.Py_PreInitialize
func (pygi *Pyg) PreInitialize() PyStatus {
	ret := C.CGO_PreInitialize(&pygi.PreConfig.CInstance)
	return StatusNew(ret)
}

// PreInitializeFromArgs
//
// Preinitialize Python from preconfig
// preconfiguration. Parse argv command line
// arguments (wide strings) if parse_argv of
// PyPreConfig is non-zero.
// https://docs.python.org/3/c-api/init_config.html#c.Py_PreInitializeFromArgs
func (pygi *Pyg) PreInitializeFromArgs(argv []CPyWideString) PyStatus {
	ret := C.CGO_Py_PreInitializeFromArgs(
		&pygi.PreConfig.CInstance,
		Clong(len(argv)),
		&argv[0])
	return StatusNew(ret)
}

// PreInitializeFromBytesArgs
//
// Preinitialize Python from preconfig
// preconfiguration. Parse argv command line
// arguments (bytes strings) if parse_argv of
// PyPreConfig is non-zero.
// https://docs.python.org/3/c-api/init_config.html#c.Py_PreInitializeFromBytesArgs
func (pygi *Pyg) PreInitializeFromBytesArgs(argv []string) PyStatus {
	cArgV := make([]*Cchar, len(argv))
	for idx, arg := range argv {
		carg := C.CString(arg)
		defer CFree(carg)
		cArgV[idx] = carg
	}

	ret := C.CGO_Py_PreInitializeFromBytesArgs(
		&pygi.PreConfig.CInstance,
		Clong(len(argv)),
		&cArgV[0])
	return StatusNew(ret)
}

// RunFile
//
// This is a simplified interface to
// PyRun_SimpleFileExFlags() below, leaving
// closeit set to 0 and flags set to NULL.
// https://docs.python.org/3/c-api/veryhigh.html?highlight=pyrun_#c.PyRun_SimpleFile
func (pygi *Pyg) RunFile(fileName string) (int, error) {
	if err := pygi.checkInitialized(); err != nil {
		return -1, err
	}
	cFileName := C.CString(fileName)
	cMode := C.CString("r")
	defer CFree(cFileName)
	defer CFree(cMode)

	fd, err := C.fopen(cFileName, cMode)
	if err != nil {
		return -1, err
	}
	defer C.fclose(fd)

	i, err := C.PyRun_SimpleFile(fd, cFileName)
	return int(i), err
}

// RunMain
//
// Execute the command (PyConfig.run_command), the
// script (PyConfig.run_filename) or the module
// (PyConfig.run_module) specified on the command
// line or in the configuration.
// https://docs.python.org/3/c-api/init_config.html#c.Py_RunMain
func (pygi *Pyg) RunMain() (int, error) {
	if err := pygi.checkInitialized(); err != nil {
		return -1, err
	}

	i, err := C.Py_RunMain()
	return int(i), err
}

// RunString
//
// This is a simplified interface to
// PyRun_SimpleStringFlags() below, leaving the
// PyCompilerFlags* argument set to NULL.
// https://docs.python.org/3/c-api/veryhigh.html?highlight=pyrun_#c.PyRun_SimpleString
func (pygi *Pyg) RunString(str string) (int, error) {
	if err := pygi.checkInitialized(); err != nil {
		return -1, err
	}
	cStr := C.CString(str)
	defer CFree(cStr)

	i, err := C.PyRun_SimpleString(cStr)
	return int(i), err
}

func (pygi *Pyg) SetArgv(argv []string) PyStatus {
	return pygi.Config.SetBytesArgv(argv)
}

func (pygi *Pyg) SetConfigBoolean(prop *Cint, value bool) PyStatus {
	return pygi.Config.SetBoolean(prop, value)
}

func (pygi *Pyg) SetConfigInteger(prop *Cint, value int) PyStatus {
	return pygi.Config.SetInteger(prop, value)
}

func (pygi *Pyg) SetConfigString(prop *CPyWideString, value string) PyStatus {
	return pygi.Config.SetString(prop, String2WideString(value))
}
