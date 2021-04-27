package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kaunamohammed/kg_kaunamohammed_2021/algos"
)

func main() {
	var args = os.Args[1:]
	var err error
	argsConvertedToIntegerArray := make([]int, len(args))
	for i := 0; i < len(args); i++ {
		// check for an error in converting an argument from a string to an integer
		if argsConvertedToIntegerArray[i], err = strconv.Atoi(args[i]); err != nil {
			fmt.Println("Please input only numbers")
			panic(err)
		}
	}
	
	algos.PhoneticEquivalent(argsConvertedToIntegerArray)
}

