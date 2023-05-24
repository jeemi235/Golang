package main

import (
	"os"
	"runtime/trace"
)

func main() {
	// f, err := os.Create("cpu_profile.prof")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// err = pprof.StartCPUProfile(f)
	// if err != nil {
	// 	panic(err)
	// }
	// defer pprof.StopCPUProfile()

	// code to be profiled

	// f, err := os.Create("mem_profile.prof")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// err = pprof.WriteHeapProfile(f)
	// if err != nil {
	// 	panic(err)
	// }
	// code to be profiled

	// f, err := os.Create("block_profile.prof")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// err = pprof.Lookup("block").WriteTo(f, 0)
	// if err != nil {
	// 	panic(err)
	// }

	// code to be profiled

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//code to be profiled

}

// go tool pprof cpu_profile.prof
