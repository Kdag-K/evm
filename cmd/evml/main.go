package main

import (
	//	_ "net/http/pprof"
	//	"runtime"

	cmd "github.com/Kdag-K/evm/cmd/evml/commands"
)

func main() {

	//	runtime.SetBlockProfileRate(1)
	//	runtime.SetMutexProfileFraction(1)

	rootCmd := cmd.RootCmd
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
