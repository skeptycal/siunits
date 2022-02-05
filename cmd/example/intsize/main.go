package main

import (
	"fmt"
	"reflect"

	"github.com/skeptycal/siunits"
)

func main() {
	n := 1
	for i := 1; i < siunits.MaxIntLen; i++ {
		fmt.Printf("%v: %v (len: %v)\n", i, n, siunits.IntLen(n))
		// a := siunits.IntLen(n)
		n *= 10
	}

	fmt.Println()
	for i, test := range siunits.IntLenTests() {
		v := reflect.ValueOf(test)
		fmt.Printf("%3d: %v \n", i, test)
		for j := 0; j < v.NumField(); j++ {

			f := v.Field(j)
			fmt.Printf("...field %2d: %20v\n", j, f)

		}

	}

}
