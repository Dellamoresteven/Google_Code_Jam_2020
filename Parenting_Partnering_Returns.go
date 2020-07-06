// @author Steven Dellamore
// @date July 5 2020
package main

import (
	"fmt"
)

/*
SAMPLE INPUT
4
3
360 480
420 540
600 660
3
0 1440
1 3
2 4
1 5
5
99 150
1 100
100 301
2 5
150 250
2
0 720
720 1440
*/

/*
1
3
0 1440
1 3
2 4
1 5
*/

/**
 * Took me a bit more time. I first had a solution that worked if everything was ordered, but
 * I found out the input is not ordered. So I then used sort.SliceStable to order my by start time
 * but had to keep into account the original index so I could go back to get the answer. Unfortunately
 * google used a < 1.8 version of golang so I could not call sort.SliceStable, so I made my own sort with
 * a custom comparator which took a little time to figure out how to pass lambdas around in go. But now i
 * know!! woo. I'm not super happy with the solution but it works and I believe it is optimal.
 *
 * Time took: 35 min 15 seconds
 */

type pair struct {
	start      int
	end        int
	oIndex     int
	taskHolder string
}

func main() {
	var numTests int
	fmt.Scanf("%d", &numTests)

	for i := 0; i < numTests; i++ {
		var numItems int // Num Activities
		var aList []pair // List of my Activities
		fmt.Scanf("%d", &numItems)
		// each activity
		for k := 0; k < numItems; k++ {
			var start int
			var end int
			fmt.Scanf("%d %d", &start, &end)
			aList = append(aList, pair{start, end, k, ""})
		}
		insertionsort(aList, func(j int) bool {
			return aList[j-1].start > aList[j].start
		})
		// Cannot use because google is using a version of golang < 1.8 :(
		// sort.SliceStable(aList, func(i, j int) bool {
		// 	return aList[i].start < aList[j].start
		// })
		possible := true
		var ansString string
		c := 0 // 0 is the start
		j := 0
		for k := 0; k < numItems; k++ {
			start := aList[k].start
			end := aList[k].end
			if c <= start {
				c = end
				aList[k].taskHolder = "C"
			} else if j <= start {
				j = end
				aList[k].taskHolder = "J"
			} else {
				fmt.Printf("Case #%d: IMPOSSIBLE\n", i+1)
				possible = false
			}
		}
		if possible {
			insertionsort(aList, func(j int) bool {
				return aList[j-1].oIndex > aList[j].oIndex
			})
			for _, p := range aList {
				ansString += p.taskHolder
			}
			fmt.Printf("Case #%d: %s\n", i+1, ansString)
		}
	}
}

// Had to create my own sort because we are on a version < 1.8
func insertionsort(items []pair, comp func(j int) bool) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if comp(j) {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
