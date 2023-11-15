package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(narcissistic(1634))
	fmt.Println(narcissistic(153))
	fmt.Println(narcissistic(111))
	fmt.Println(parityOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	fmt.Println(parityOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))
	fmt.Println(parityOutlier([]int{11, 13, 15, 19, 9, 13, -21}))
	fmt.Println(findNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue"))
	fmt.Println(blueOcean([]int{1, 2, 3, 4, 6, 10}, []int{1}))
	fmt.Println(blueOcean([]int{1, 5, 5, 5, 5, 3}, []int{5}))

	// blueOcean([1,2,3,4,6,10], [1])
}

func narcissistic(n int) bool {
	strNum := strconv.Itoa(n)
	numDigits := len(strNum)

	sum := 0
	for _, char := range strNum {
		digit, _ := strconv.Atoi(string(char))
		sum += int(math.Pow(float64(digit), float64(numDigits)))
	}

	return sum == n
}

func parityOutlier(arr []int) string {
	var evens, odds []int
	var outlier int

	for _, num := range arr {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}
	if len(evens) == 0 || len(odds) == 0 {
		return "false"
	}
	if len(evens) == 1 {
		outlier = evens[0]
	} else {
		outlier = odds[0]
	}

	return strconv.Itoa(outlier)
}

func findNeedle(arr []string, k string) int {
	for i, v := range arr {
		if k == v {
			return i
		}
	}
	return 0
}

func blueOcean(arr1 []int, arr2 []int) []int {
	var res []int
	mapCheck := make(map[int]bool, 0)
	for _, v := range arr1 {
		mapCheck[v] = true
	}
	for _, v := range arr2 {
		if mapCheck[v] {
			delete(mapCheck, v)
		}
	}
	for v := range mapCheck {
		res = append(res, v)
	}
	sort.Sort(sort.IntSlice(res))

	return res
}
