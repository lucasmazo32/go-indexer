package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"sync"

	"indexer.com/indexer/internal/update_task"
	"indexer.com/indexer/internal/worker_pool"
)

var cpuprofile = flag.String("cpuprofile", "cpu.prof", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "mem.prof", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatalln("Could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatalln("Could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	path := os.Args[1]
	wg := &sync.WaitGroup{}
	p, err := worker_pool.NewSimplePool(16, 100)
	if err != nil {
		log.Fatalln(err)
	}
	p.Start()
	filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		lastChar := name[len(name)-1]
		if lastChar == '.' {
			wg.Add(1)
			p.AddWork(update_task.NewUpdateTask(name, wg))
			return nil
		}
		return nil
	})
	wg.Wait()
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
