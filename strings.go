package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "eb.prof", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("err", err)
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()

		fmt.Println(isUnique("tesg"))
		fmt.Println(isUniqueBit("test"))
		fmt.Println(isPerm("tedst", "tesdt"))
	}
}

func isPerm(a, b string) bool {

	if len(a) != len(b) {
		return false
	}

	set := make(map[rune]int)
	for _, v := range a {
		set[v]++
	}

	for _, v := range b {
		if _, ok := set[v]; !ok {
			return false
		}
		set[v]--
		if set[v] < 0 {
			return false
		}
	}

	return true
}

func isUnique(s string) bool {
	a := make(map[rune]bool)

	for _, v := range s {
		if ok := a[v]; ok {
			return false
		}
		a[v] = true
	}
	return true
}

func isUniqueBit(s string) bool {
	checker := 0
	for k := range s {
		v := s[k] - "a"[0]
		if (checker & (1 << v)) > 0 {
			return false
		}
		checker |= (1 << v)
	}
	return true
}
