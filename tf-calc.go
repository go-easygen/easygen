package easygen

import (
	"errors"
	"strconv"
)

//==========================================================================
// template function somewhat related to calculations

// Iterate returns a slice whose end (& start) is controlled by the ic
func Iterate(ic ...string) []int {
	// cannot think of a better way to report the Atoi error, other than --
	end, _ := strconv.Atoi(ic[0])
	start := 0
	if len(ic) > 1 {
		start, _ = strconv.Atoi(ic[1])
	}
	var Items []int
	for i := start; i <= (end); i++ {
		Items = append(Items, i)
	}
	return Items
}

// ArgsA creates an array slice from the given arguments
func ArgsA(vs ...interface{}) []interface{} { return vs }

// ArgsM creates a map from the given K/V argument pairs
func ArgsM(kvs ...interface{}) (map[string]interface{}, error) {
	if len(kvs)%2 != 0 {
		return nil, errors.New("argsm requires even number of arguments.")
	}
	m := make(map[string]interface{})
	for i := 0; i < len(kvs); i += 2 {
		s, ok := kvs[i].(string)
		if !ok {
			return nil, errors.New("even args to args must be strings.")
		}
		m[s] = kvs[i+1]
	}
	return m, nil
}

//==========================================================================
// template function for calculations

func Add(b, a int) int   { return a + b }
func Minus(b, a int) int { return a - b }

// By Caleb Spare @gmail.com
// https://groups.google.com/d/msg/golang-nuts/gzAyBLAeUbU/LwgomgxcjQ8J

// Minus1 calculates to input less 1
func Minus1(n int) int { return n - 1 }
