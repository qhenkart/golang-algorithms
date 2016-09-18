package main

import (
	"fmt"
	"strings"
	"testing"
)

func main() {
	result := testing.Benchmark(func(*testing.B) {
		for i := 0; i < 1000; i++ {

			IsPalindrome("qwertyuiopoiuytrewq")
		}
	})
	fmt.Printf("iterations %+v\ntime %+v\nBytes %+v\nMemAlocs %+v\nMemBytes %+v\n\n\n\n\n", result.N, result.T, result.Bytes, result.MemAllocs, result.MemBytes)

	result = testing.Benchmark(func(*testing.B) {
		for i := 0; i < 1000; i++ {
			IsPalindrome2("qwertyuiopoiuytrewq")
		}
	})
	fmt.Printf("iterations %+v\ntime %+v\nBytes %+v\nMemAlocs %+v\nMemBytes %+v\n", result.N, result.T, result.Bytes, result.MemAllocs, result.MemBytes)

}

func IsPalindrome2(s string) bool {
	set := map[rune]struct{}{}
	for i, v := range s {
		if s[i] <= 90 && s[i] >= 65 {
			v = v + 32
		}

		if _, ok := set[v]; ok {
			delete(set, v)
			continue
		}
		if v != 32 {
			set[v] = struct{}{}
		}
	}

	if (len(s)%2 != 0 && len(set) == 1) || (len(s)%2 == 0 && len(set) == 0) {
		return true
	}
	return false
}
func IsPalindrome(s string) bool {
	set := make(map[rune]int)
	for _, char := range s {
		if char != 32 {
			set[char]++
		}
	}
	even := len(s)%2 == 0
	var oddCount int
	for _, v := range set {
		if v%2 != 0 {
			if oddCount < 1 && !even {
				oddCount++
				continue
			}
			return false
		}
	}
	return true
}

func URLify(s string) string {
	s = strings.Trim(s, " ")
	return strings.Replace(s, " ", "%20", -1)
}

func URLify2(s string) string {
	var url string
	for k, char := range s {
		if char == 32 {
			url = s[k+1:]
			continue
		}
		break
	}

	start := false
	for i := len(url) - 1; i >= 0; i-- {
		if !start && url[i] == 32 {
			url = url[:i]
			continue
		}
		start = true
		if url[i] == 32 {
			url = url[:i] + "%20" + url[i+1:]
		}
	}

	return url
}

func isUniq(s string) bool {
	set := map[rune]struct{}{}
	for _, char := range s {
		if _, ok := set[char]; ok {
			return false
		}
		set[char] = struct{}{}
	}
	return true
}

func isPerm(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	set := make(map[rune]int)

	for _, char := range s {
		set[char]++
	}

	for _, char := range t {
		set[char]--
		if set[char] < 0 {
			return false
		}
	}
	return true

}
