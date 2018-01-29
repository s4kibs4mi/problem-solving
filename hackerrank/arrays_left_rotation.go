package hackerrank

import "fmt"

/**
 * := Coded with love by Sakib Sami on 1/29/18.
 * := s4kibs4mi@gmail.com
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

// Problem Statement : https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem

func RunArraysLeftRotation() {
	var n, m, x int
	fmt.Scanf("%d %d", &n, &m)
	var numbers []int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		numbers = append(numbers, x)
	}
	for i := 0; i < m; i++ {
		numbers = append(numbers[1:], numbers[0])
	}
	isFirst := true
	for _, v := range numbers {
		if !isFirst {
			fmt.Print(" ")
		}
		fmt.Printf("%d", v)
		isFirst = false
	}
}
