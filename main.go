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
	inter := pyg.PygNewPython()
	defer inter.FinalizeEx()

	EvalStatus(inter.PreInitializeFromBytesArgs(os.Args))
	EvalStatus(inter.SetArgv(os.Args))
	EvalStatus(inter.InitializeFromConfig())

	inter.RunString("print('Hello World')")
	inter.FinalizeEx()
}
