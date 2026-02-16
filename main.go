package main

import "fmt"

func main() {
	var m [3][3]float64

	fmt.Println("Enter the matrix elements 3x3: ")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&m[i][j])
		}
	}

	det := m[0][0]*(m[1][1]*m[2][2]-m[1][2]*m[2][1]) -
		m[0][1]*(m[1][0]*m[2][2]-m[1][2]*m[2][0]) +
		m[0][2]*(m[1][0]*m[2][1]-m[1][1]*m[2][0])

	fmt.Println("Determinant =", det)
}
