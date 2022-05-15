package sample

import (
	"fmt"
	"os"
	"strings"

	"github.com/speecan/moo/game"
)

// EstimateHuman is played by human
func EstimateHuman(difficulty int) game.Estimate {
	return func(fn game.Question) (res []int) {
		var input string
		fmt.Print("?: ")
		fmt.Fscanln(os.Stdin, &input)
		guess := game.Str2Int(strings.Split(input, ""))
		fn(guess)
		return guess
	}
}

// EstimateWithRandom is idiot algo.
// returns estimate number with simply random
func EstimateWithRandom(difficulty int) game.Estimate {
	return func(fn game.Question) (res []int) {
		r := game.GetMooNum(difficulty)
		fn(r)
		return r
	}
}

// EstimateWithRandom2 is idiot algo.
// exclude duplicate queries
func EstimateWithRandom2(difficulty int) game.Estimate {
	query := make([][]int, 0)
	isDuplicated := func(i []int) bool {
		for _, v := range query {
			if game.Equals(v, i) {
				return true
			}
		}
		return false
	}
	return func(fn game.Question) (res []int) {
		var r []int
		for {
			r = game.GetMooNum(difficulty)
			if !isDuplicated(r) {
				break
			}
		}
		fn(r)
		query = append(query, r)
		return r
	}
}

func EstimateWithRandom3(difficulty int) game.Estimate {
	query := make([][]int, 0)
	hb := make([][]int, 0)

	ischeck := func(i []int) bool {
		for j := 0; j < len(query); j++ {
			countH := func(v []int) int {
				c := 0
				for j := 0; j < len(v); j++ {
					if v[j] == i[j] {
						c++
					}
				}
				return c
			}
			countB := func(v []int) int {
				c := 0
				for _, x := range i {
					for _, y := range v {
						if x == y {
							c++
						}
					}
				}
				return c
			}
			if countH(query[j]) != hb[j][0] || countB(query[j])-countH(query[j]) != hb[j][1] {
				return false
			}

		}
		return true
	}
	return func(fn game.Question) (res []int) {
		var r []int
		for {
			r = game.GetMooNum(difficulty)
			if ischeck(r) {
				break
			}
		}
		x, y := fn(r)
		var temp = []int{x, y}
		hb = append(hb, temp)
		query = append(query, r)
		// fmt.Println(hb)
		// fmt.Println(r)
		// fmt.Println(query)
		return r
	}
}
