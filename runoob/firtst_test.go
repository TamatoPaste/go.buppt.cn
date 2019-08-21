package main

import (
	"fmt"
	"testing"
)

func TestFisrtTry(t *testing.T) {
	var a = [3][5]int{{1, 2, 3, 4, 5}, {0, 9, 8, 7, 6}, {3, 4, 5, 6, 7}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
	t.Log("hello, test !")
}
