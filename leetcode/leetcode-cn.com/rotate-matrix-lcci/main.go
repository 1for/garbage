package main

import "fmt"

func main(){
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	
	rotate(matrix)
	fmt.Println(matrix);
}

func rotate(matrix [][]int) {
	n := len(matrix);
	var tmp, i, j int;

	for i=0; i< n/2; i++ {
		for j=i; j<(n-1-i); j++ {
			tmp = matrix[i][j];
			matrix[i][j] = matrix[n-1-j][i];
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j];
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i];
			matrix[j][n-1-i] = tmp;
		}
	}
}
