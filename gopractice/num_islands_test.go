package practice

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {

	// g := [][]byte{
	// 	{1, 0, 1, 0},
	// 	{1, 1, 0, 1},
	// 	{0, 0, 1, 1},
	// }

	g := [][]byte{
		{'1', '0', '1', '1', '1'},
		{'1', '0', '1', '0', '1'},
		{'1', '1', '1', '0', '1'},
	}

	// g := [["1","0","1","1","1"],["1","0","1","0","1"],["1","1","1","0","1"]]

	fmt.Println(NumIslands(g))

}