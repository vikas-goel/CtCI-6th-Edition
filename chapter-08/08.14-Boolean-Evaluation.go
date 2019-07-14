// Given a boolean expression consisting of the symbols 0 (false), 1 (true), &
// (AND), | (OR), and ^ (XOR), and a desired boolean result value result,
// implement a function to count the number of ways of parenthesizing the
// expression such that it evaluates to result. The expression should be fully
// parenthesized (e.g., (0)^(1)) but not extraneously (e.g., (((0))^(1)).
// EXAMPLE
// countEval("1^0|0|1", false) -> 2
// countEval("0&0&0&1^1|0", true) -> 10

package main

import "fmt"

func Evaluate(expression string, result bool) int {
	if len(expression) == 0 {
		return 0
	}

	expressionArray := []byte(expression)
	mapExpResCount := make(map[string]int)

	var findWays func(int, int, byte) int
	findWays = func(start, end int, subres byte) int {
		if start > end {
			return 0
		} else if start == end {
			if expressionArray[start] == subres {
				return 1
			} else {
				return 0
			}
		}

		// If already evaluated this expression, then return the count
		// from the lookup table.
		mapkey := fmt.Sprintf("%v-%v-%v", start, end, subres)
		if count, ok := mapExpResCount[mapkey]; ok {
			return count
		}

		totalSubWays := 0

		// Divide the expression to two subsets and evaluate their
		// counts.
		for i := start+1; i < end; i += 2 {
			totalTrue, totalWays := 0, 0

			// Get all possible results for the two
			// sub-expressions.
			leftFalse := findWays(start, i-1, '0')
			leftTrue := findWays(start, i-1, '1')
			rightFalse := findWays(i+1, end, '0')
			rightTrue := findWays(i+1, end, '1')

			if expressionArray[i] == '^' {
				totalTrue += leftFalse * rightTrue
				totalTrue += leftTrue * rightFalse
			} else if expressionArray[i] == '&' {
				totalTrue += leftTrue * rightTrue
			} else if expressionArray[i] == '|' {
				totalTrue += leftTrue * rightTrue
				totalTrue += leftTrue * rightFalse
				totalTrue += leftFalse * rightTrue
			}

			if subres == '1' {
				totalSubWays += totalTrue
			} else {
				totalWays = (leftTrue+leftFalse)*(rightTrue+rightFalse)
				totalSubWays += totalWays - totalTrue
			}
		}

		// Add the count against the expression to the lookup table.
		mapExpResCount[mapkey] = totalSubWays
		return totalSubWays
	}

	var resultByte byte = '0'
	if result {
		resultByte = '1'
	}

	return findWays(0, len(expressionArray)-1, resultByte)
}

func main() {
	exp1, res1 := "1^0|0|1", false
	exp2, res2 := "0&0&0&1^1|0", true

	fmt.Printf("Ways(%v, %v) = %v\n", exp1, res1, Evaluate(exp1, res1))
	fmt.Printf("Ways(%v, %v) = %v\n", exp2, res2, Evaluate(exp2, res2))
}
