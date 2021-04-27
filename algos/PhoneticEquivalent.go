package algos

import (
	"fmt"
)

// loops over every number in nums and converts the digits to a phonetic equivalent
func PhoneticEquivalent(nums []int) {

	for i := 0; i < len(nums); i++ {
		number := nums[i]
		findDigitString(number)
		// prevents the extra comma
		if i < len(nums) - 1 {
			fmt.Print(",")
		}
	}

}

// finds the phonetic equivalent of a number
func findDigitString(number int) {
	numberMap := make(map[int]string)

	numberMap[0] = "Zero"
	numberMap[1] = "One"
	numberMap[2] = "Two"
	numberMap[3] = "Three"
	numberMap[4] = "Four"
	numberMap[5] = "Five"
	numberMap[6] = "Six"
	numberMap[7] = "Seven"
	numberMap[8] = "Eight"
	numberMap[9] = "Nine"

    if number > 0 {
		findDigitString(number/10)
		fmt.Print(numberMap[(number % 10)])
    }
	
}

