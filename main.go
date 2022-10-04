package main

import (
	"runtime"

	"cheat/src"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	src.RunAsRoot()
	src.PrintBanner()
	src.NewSignalHandler()
	hackEngine := src.NewEngine()
	// TODO: Add key handling
	hackEngine.Render()
}

// mmap will be useful later || mmap será útil mais tarde
// data, err = syscall.Mmap(int(file.Fd()), 0, int(size), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
// if err != nil {
// 	panic("Unable to mmap file")
// }
