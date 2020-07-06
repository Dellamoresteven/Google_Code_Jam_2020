// @author Steven Dellamore
// @date July 5 2020
package main

import "fmt"

/*
SAMPLE INPUT
3
4
1 2 3 4
2 1 4 3
3 4 1 2
4 3 2 1
4
2 2 2 2
2 3 2 3
2 2 2 3
2 2 2 2
3
2 1 3
1 3 2
1 2 3
*/

/**
 * This is definitely not the best way to do this. We could easily use a hash table to check for dups in O(N).
 * We could also use goroutines for this but I was having trouble sumbiting those to the Code jam website. This
 * was a practice one so I wont mess with it too much.
 *
 * Time took: 15 min 04 seconds
 */

func main() {
	var testCases int
	fmt.Scanf("%d", &testCases)
	var mSize int
	for i := 0; i < testCases; i++ {
		kE := 0
		rE := 0
		cE := 0
		fmt.Scanf("%d", &mSize)
		matrix := make([][]int, mSize)
		for j := 0; j < mSize; j++ {
			matrix[j] = make([]int, mSize)
			for k := 0; k < mSize; k++ {
				fmt.Scanf("%d", &matrix[j][k])
				if j == k {
					kE += matrix[j][k]
				}
			}
		}

		for j := 0; j < mSize; j++ {
			for k := 0; k < mSize; k++ {
				isDup := false
				for e := 0; e < k; e++ {
					if matrix[j][k] == matrix[j][e] {
						rE++
						isDup = true
					}
				}
				if isDup {
					break
				}
			}
		}

		// could use go goroutines here but for some reason the entry didn't like that
		for j := 0; j < mSize; j++ {
			for k := 0; k < mSize; k++ {
				isDup := false
				for e := 0; e < k; e++ {
					if matrix[k][j] == matrix[e][j] {
						cE++
						isDup = true
					}
				}
				if isDup {
					break
				}
			}
		}
		fmt.Printf("Case #%d: %d %d %d\n", i+1, kE, rE, cE)
	}
}
