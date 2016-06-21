package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
	"testing"
)

var cpuprofile = flag.String("cpuprofile", "eb.prof", "write cpu profile to file")

func main() {

	//
	// fmt.Println(isUnique("tesg"))
	// fmt.Println(isUniqueBit("test"))
	// fmt.Println(isPerm("tedst", "tesdt"))
	// url := "Mr John Smith      "
	// urlify(&url)
	// fmt.Println(url)
	// fmt.Println(urlifyWithNoPackage("Mr John Smith"))
	// fmt.Println(isPalindromePermutation("caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco"))
	// start := time.Now()
	// fmt.Println(oneAway("awayawayawayawayawayawayawayawayawayawayawayaway", "awayawayawayawayawayawayawayawayawayawayawayaway"))
	// end := time.Now()
	// fmt.Println("forst", end.Sub(start).Nanoseconds())
	// bench := &testing.BenchmarkResult{}
	fmt.Println("\n\n\n\n ")
	result := testing.Benchmark(func(*testing.B) {
		fmt.Println(stringCompressionOptimize1("aaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweee"))
	})

	fmt.Printf("iterations %+v\ntime %+v\nBytes %+v\nMemAlocs %+v\nMemBytes %+v\n", result.N, result.T, result.Bytes, result.MemAllocs, result.MemBytes)

	result = testing.Benchmark(func(*testing.B) {
		fmt.Println(stringCompressionOptimize("aaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweee"))
	})

	fmt.Printf("iterations %+v\ntime %+v\nBytes %+v\nMemAlocs %+v\nMemBytes %+v\n", result.N, result.T, result.Bytes, result.MemAllocs, result.MemBytes)
}

func stringCompression(s string) string {
	count := 0
	str := string(s[0])
	result := ""
	for k := range s {
		if string(s[k]) != str {
			result += fmt.Sprintf("%s%d", str, count)
			str = ""
			count = 0
		}

		str = string(s[k])
		count++
	}
	result += fmt.Sprintf("%s%d", str, count)

	if len(result) >= len(s) {
		return s
	}
	return result
}

func stringCompressionOptimize(s string) string {
	count := 0
	str := string(s[0])
	var data []byte
	for k := range s {
		if string(s[k]) != str {
			data = append(data, fmt.Sprintf("%s%d", str, count)...)
			str = ""
			count = 0
		}

		str = string(s[k])
		count++
	}
	data = append(data, fmt.Sprintf("%s%d", str, count)...)

	if len(data) >= len(s) {
		return s
	}
	return string(data)
}

func stringCompressionOptimize1(s string) string {
	var count int
	temp := s[0]
	var data []byte
	for k := range s {
		if s[k] != temp {
			data = append(data, fmt.Sprintf("%s%d", string(temp), count)...)
			count = 0
		}

		temp = s[k]
		count++
	}
	data = append(data, fmt.Sprintf("%s%d", string(temp), count)...)

	if len(data) >= len(s) {
		return s
	}
	return string(data)
}

func oneAway(s1, s2 string) bool {
	if math.Abs(float64(len(s1)-len(s2))) > 1 {
		return false
	}
	t := len(s1)
	e := len(s2)
	switch {
	case t == e:
		count := 0
		for k := range s1 {
			if s1[k] != s2[k] {
				count++
			}

			if count > 1 {
				return false
			}
		}
	case t > e:
		offset := false
		for k := range s1 {
			if k == e {
				break
			}
			if !offset {
				if s1[k] != s2[k] {
					offset = true
				}
			}
		}

	case t < e:
		offset := false
		for k := range s2 {
			if k >= t {
				break
			}
			if !offset && k < t-1 {
				if s1[k] != s2[k] {
					offset = true
				}
			}
			if offset || k == t-1 {
				if s1[k] != s2[k+1] {
					return false
				}
			}
		}
	}

	return true
}

// func isPalindromeBitwise(s string) bool {
// 	for i, v := range s {
// 		if s[i] <= 90 && s[i] >= 65 {
// 			v = v + 32
// 		}
// 		if v != 32 {
//
// 		}
// 	}
// 	//slower implementation
// 	// count := 0
// 	// for _, v := range pal {
// 	// 	if v%2 != 0 {
// 	// 		count++
// 	// 	}
// 	// 	if count > 1 {
// 	// 		return false
// 	// 	}
// 	// }
// 	if len(pal) > 1 {
// 		return false
// 	}
// 	return true
// }

func isPalindromePermutation(s string) bool {
	pal := make(map[rune]int)
	for i, v := range s {
		if s[i] <= 90 && s[i] >= 65 {
			v = v + 32
		}
		if _, ok := pal[v]; ok {
			delete(pal, v)
			continue
		}
		if v != 32 {
			pal[v]++
		}
	}
	//slower implementation
	// count := 0
	// for _, v := range pal {
	// 	if v%2 != 0 {
	// 		count++
	// 	}
	// 	if count > 1 {
	// 		return false
	// 	}
	// }
	if len(pal) > 1 {
		return false
	}
	return true
}

func urlify(a *string) {
	*a = strings.TrimSpace(*a)
	*a = strings.Replace(*a, " ", "%20", -1)
}

func urlifyWithNoPackage(a string) string {
	for k := range a {
		if a[k] == 32 {
			a = string(a[:k]) + "%20" + string(a[k+1:])
		}
	}
	return a
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
