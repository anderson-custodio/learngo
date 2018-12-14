// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	// by using 'shift' we can create the slide effect for placeholders.
	for shift := 0; ; shift++ {
		// we need to clear the screen here.
		// or the previous character will be left on the screen
		//
		// alternative: you can fill the rest of the missing placeholders
		//              with empty lines
		screen.Clear()
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			l := len(clock)

			// this sets the beginning and the ending placeholders.
			// to prevent the indexing error: we use the remainder operator.
			s, e := shift%l, l

			// to slide placeholders from the right part of the screen.
			//
			// here, we assume that as if the clock's length is double of its length.
			// this makes things easy to manage: that's why: l*2 is there.
			//
			// whenever, the current shift factor's double remainder is greater than
			// the length of the clock - 1, it changes the starting and ending positions.
			if shift%(l*2) > l-1 {
				s, e = 0, shift%l+1
			}

			// print empty lines for the right-to-left slide effect.
			//
			// this creates the effect of moving placeholders from right to left.
			for j := 0; j < l-e; j++ {
				fmt.Print("     ")
			}

			// draw the digits starting from 's' to 'e'
			for i := s; i < e; i++ {
				next := clock[i][line]
				if clock[i] == colon && sec%2 == 0 {
					next = "   "
				}

				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
	}
}
