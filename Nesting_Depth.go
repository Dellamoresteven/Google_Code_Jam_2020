// @author Steven Dellamore
// @date July 5 2020
package main

import (
	"fmt"
	"strconv"
)

/*
SAMPLE INPUT
4
0000
101
111000
1
*/

/**
 * Both pass. O(N) complexity. I believe this is optimal way to do this.
 *
 * Time took: 8 min 37 seconds
 */

func main() {
	var numTests int
	fmt.Scanf("%d", &numTests)

	printP := func(num int, typ byte) {
		for i := 0; i < num; i++ {
			fmt.Printf("%c", typ)
		}
	}

	for i := 0; i < numTests; i++ {
		numOpen := 0
		var input string
		fmt.Scanf("%s", &input)
		fmt.Printf("Case #%d: ", i+1)
		for j := 0; j < len(input); j++ {
			num, _ := strconv.Atoi(string(input[j]))
			if num > numOpen {
				printP(num-numOpen, '(')
				numOpen += num - numOpen
			} else if num < numOpen {
				printP(numOpen-num, ')')
				numOpen -= numOpen - num
			}
			fmt.Print(num)
		}
		printP(numOpen, ')')
		fmt.Println()
	}
}
