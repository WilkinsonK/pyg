package main

import (
	"os"

	pyg "github.com/WilkinsonK/pyg/lib"
)

func EvalStatus(status pyg.PyStatus) {
	if status.IsException() {
		if status.IsExit() {
			os.Exit(status.ExitCode())
		}
		status.ExitStatusException()
	}
}

func main() {
	pygi := pyg.PygNewPython()
	defer pygi.FinalizeEx()

	EvalStatus(pygi.PreInitializeFromBytesArgs(os.Args))
	EvalStatus(pygi.SetArgv(os.Args))
	EvalStatus(pygi.SetConfigInteger(pygi.Config.Verbose, 2))
	EvalStatus(pygi.SetConfigString(pygi.Config.ProgramName, "PyGi"))
	EvalStatus(pygi.InitializeFromConfig())

	pygi.RunString("print('Hello World')")
	pygi.FinalizeEx()
}
