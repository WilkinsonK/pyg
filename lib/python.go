package pyg

// #cgo pkg-config: python3-embed
// #include "txllayer.h"
import "C"
import (
	"fmt"
)

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

// Get a new instance of Python interpreter.
func PygNewPython() *Pyg {
	return PygNew(nil, nil)
}

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

func (pygi *Pyg) checkInitializedS(val string, err error) (string, error) {
	if err := pygi.checkInitialized(); err != nil {
		return "", err
	}
	return val, err
}

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

// Initialize Python from config configuration.
// https://docs.python.org/3/c-api/init_config.html#c.Py_InitializeFromConfig
func (pygi *Pyg) InitializeFromConfig() PyStatus {
	ret := C.CGO_Py_InitializeFromConfig(&pygi.Config.CInstance)
	return StatusNew(ret)
}

// Whether the Python interpreter has been
// initialized or not.
// https://docs.python.org/3/c-api/init.html#c.Py_IsInitialized
func (pygi *Pyg) IsInitialized() bool {
	return CInt2Bool(C.Py_IsInitialized())
}

// Return information about the sequence number
// and build date and time of the current Python
// interpreter.
// https://docs.python.org/3/c-api/init.html#c.Py_GetBuildInfo
func (pygi *Pyg) GetBuildInfo() string {
	return C.GoString(C.Py_GetBuildInfo())
}

// Return an indication of the compiler used to
// build the current Python version.
// https://docs.python.org/3/c-api/init.html#c.Py_GetCompiler
func (pygi *Pyg) GetCompiler() string {
	return C.GoString(C.Py_GetCompiler())
}

// Return the official copyright for the current
// Python version.
// https://docs.python.org/3/c-api/init.html#c.Py_GetCopyright
func (pygi *Pyg) GetCopyright() string {
	return C.GoString(C.Py_GetCopyright())
}

// Return the exec-prefix for installed
// platform-dependent files.
// https://docs.python.org/3/c-api/init.html#c.Py_GetExecPrefix
func (pygi *Pyg) GetExecPrefix() (string, error) {
	return pygi.checkInitializedS(WString2String(C.Py_GetExecPrefix()))
}

// Return the default module search path.
// https://docs.python.org/3/c-api/init.html#c.Py_GetPath
func (pygi *Pyg) GetPath() (string, error) {
	return pygi.checkInitializedS(WString2String(C.Py_GetPath()))
}

// Return the platform identifier for the current
// platform.
// https://docs.python.org/3/c-api/init.html#c.Py_GetPlatform
func (inter *Pyg) GetPlatform() string {
	return C.GoString(C.Py_GetPlatform())
}

// Return the prefix for installed platform-independent files.
// https://docs.python.org/3/c-api/init.html#c.Py_GetPrefix
func (pygi *Pyg) GetPrefix() (string, error) {
	return pygi.checkInitializedS(WString2String(C.Py_GetPrefix()))
}

// Return the full program name of the Python
// executable.
// https://docs.python.org/3/c-api/init.html#c.Py_GetProgramFullPath
func (pygi *Pyg) GetProgramFullPath() (string, error) {
	return pygi.checkInitializedS(WString2String(C.Py_GetProgramFullPath()))
}

// Return the program name set by Python
// configuration.
// https://docs.python.org/3/c-api/init.html#c.Py_GetProgramName
func (pygi *Pyg) GetProgramName() (string, error) {
	return pygi.checkInitializedS(WString2String(C.Py_GetProgramName()))
}

// Return the default "home", that is, the value
// set by configuration.
// https://docs.python.org/3/c-api/init.html#c.Py_GetPythonHome
func (pygi *Pyg) GetPythonHome() (string, error) {
	return pygi.checkInitializedS(WString2String(C.Py_GetPythonHome()))
}

// Return the version of this Python interpreter.
// https://docs.python.org/3/c-api/init.html#c.Py_GetVersion
func (pygi *Pyg) GetVersion() string {
	return C.GoString(C.Py_GetVersion())
}

// Preinitialize Python from preconfig
// preconfiguration.
// https://docs.python.org/3/c-api/init_config.html#c.Py_PreInitialize
func (pygi *Pyg) PreInitialize() PyStatus {
	ret := C.CGO_PreInitialize(&pygi.PreConfig.CInstance)
	return StatusNew(ret)
}

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

// This is a simplified interface to
// PyRun_SimpleFileExFlags() below, leaving
// closeit set to 0 and flags set to NULL.
// https://docs.python.org/3/c-api/veryhigh.html?highlight=pyrun_#c.PyRun_SimpleFile
func (pygi *Pyg) RunFile(fp *CFILE, fileName string) (int, error) {
	if err := pygi.checkInitialized(); err != nil {
		return -1, err
	}
	cFileName := C.CString(fileName)
	defer CFree(cFileName)

	return int(C.PyRun_SimpleFile(fp, cFileName)), nil
}

// Execute the command (PyConfig.run_command), the
// script (PyConfig.run_filename) or the module
// (PyConfig.run_module) specified on the command
// line or in the configuration.
// https://docs.python.org/3/c-api/init_config.html#c.Py_RunMain
func (pygi *Pyg) RunMain() (int, error) {
	if err := pygi.checkInitialized(); err != nil {
		return -1, err
	}
	return int(C.Py_RunMain()), nil
}

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

	return int(C.PyRun_SimpleString(cStr)), nil
}

func (pygi *Pyg) SetArgv(argv []string) PyStatus {
	return pygi.Config.SetBytesArgv(argv)
}
