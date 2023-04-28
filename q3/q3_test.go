package main

import "fmt"

func DominoPieces(m int, n int) (int, error) {

	if m <= 0 || n <= 0 {

		return 0, fmt.Errorf("m e n devem ser maiores que zero")
	}
			return (m * n) / 2, nil
		
	}
