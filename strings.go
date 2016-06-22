package strings

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
	// start := time.Now()
	// fmt.Println(oneAway("awayawayawayawayawayawayawayawayawayawayawayaway", "awayawayawayawayawayawayawayawayawayawayawayaway"))
	// end := time.Now()
	// fmt.Println("forst", end.Sub(start).Nanoseconds())
	// fmt.Println(stringCompressionOptimize1("aaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweee"))
	// fmt.Println(stringCompressionOptimize("aaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweeeaaaaaaaabbweee"))
	// fmt.Println(isPalindromePermutation("caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco caAco"))
	// fmt.Println("\n\n\n\n ")
	// matrix := [][]int{
	// 	[]int{1, 2, 3, 5, 5, 2, 3, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{4, 5, 6, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{1, 2, 3, 5, 5, 2, 3, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{4, 5, 6, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{1, 2, 3, 5, 5, 2, 3, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{4, 5, 6, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// 	[]int{7, 8, 9, 3, 4, 5, 6, 1, 3, 2, 3, 2, 3, 2, 3, 2, 3, 4, 3, 2, 1, 3},
	// }
	// matrix := [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{4, 5, 6},
	// 	[]int{7, 8, 9},
	// }
	// matrix = [][]int{
	// 	[]int{1, 2, 3, 4},
	// 	[]int{5, 6, 7, 8},
	// 	[]int{9, 10, 11, 12},
	// 	[]int{13, 14, 15, 16},
	// }
	// rotateMatrix(matrix))
	result := testing.Benchmark(func(*testing.B) {
	})

	fmt.Printf("iterations %+v\ntime %+v\nBytes %+v\nMemAlocs %+v\nMemBytes %+v\n", result.N, result.T, result.Bytes, result.MemAllocs, result.MemBytes)
	//
	// result = testing.Benchmark(func(*testing.B) {
	// })
	//
	// fmt.Printf("iterations %+v\ntime %+v\nBytes %+v\nMemAlocs %+v\nMemBytes %+v\n", result.N, result.T, result.Bytes, result.MemAllocs, result.MemBytes)
}

func zeroMatrix(matrix [][]int) {
	type coord struct {
		i int
		j int
	}
	parent := len(matrix)
	child := len(matrix[0])
	var zeros []coord
	for i := 0; i < parent; i++ {
		for j := 0; j < child; j++ {
			if matrix[i][j] == 0 {
				zeros = append(zeros, coord{i: i, j: j})
			}
		}
	}
}

func rotateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {

			offset := i - first
			top := matrix[first][i] // save top

			//left -> top
			matrix[first][i] = matrix[last-offset][first]
			//bottom -> left
			matrix[last-offset][first] = matrix[last][last-offset]
			//right -> bottom
			matrix[last][last-offset] = matrix[i][last]
			//top -> right
			matrix[i][last] = top

		}
	}

	return matrix
}

func rotateMatrix1(matrix [][]int) [][]int {
	l := len(matrix) - 1
	result := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		for j := 0; j <= l; j++ {
			// fmt.Printf("swapping %d with %d\n", matrix[len(matrix)-1-j][i], matrix[i][j])
			result[i] = append(result[i], matrix[l-j][i])
		}
	}

	return result
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
