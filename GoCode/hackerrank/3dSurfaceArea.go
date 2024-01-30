package main

// https://www.hackerrank.com/challenges/3d-surface-area/problem
func surfaceArea(A [][]int32) int32 {
	// Write your code here
	H := len(A)
	W := len(A[0])
	var area int32

	for i := 0; i < H; i++ {
		var sum int32
		for j := 0; j < W; j++ {
			sum += A[i][j]
		}
		area += int32(2*W) + int32(4*sum)
	}
	for i := 0; i < H-1; i++ {
		var overlap int32
		for j := 0; j < W; j++ {
			if A[i][j] < A[i+1][j] {
				overlap += A[i][j]
			} else {
				overlap += A[i+1][j]
			}
		}
		area -= overlap * 2
	}
	for j := 0; j < W-1; j++ {
		var overlap int32
		for i := 0; i < H; i++ {
			if A[i][j] < A[i][j+1] {
				overlap += A[i][j]
			} else {
				overlap += A[i][j+1]
			}
		}
		area -= overlap * 2
	}

	return area
}
