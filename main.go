package main

import "fmt"

var RomanToDecimal = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func Decode(roman string) int {
	return SumRoman([]rune(roman), 0)
}

func SumRoman(romanNumbers []rune, currRoman int) int {
	currRomanNum := string(romanNumbers[currRoman])
	currDecimalNum := RomanToDecimal[currRomanNum]
	if currRoman == len(romanNumbers)-1 {
		return currDecimalNum
	}
	nextRomanNum := string(romanNumbers[currRoman+1])
	nextDecimalNum := RomanToDecimal[nextRomanNum]
	if currDecimalNum >= nextDecimalNum {
		return SumRoman(romanNumbers, currRoman+1) + currDecimalNum
	}
	return SumRoman(romanNumbers, currRoman+1) - currDecimalNum
}

func main() {
	fmt.Println(Decode("XXI"))
	fmt.Println(Decode("XXV"))
}
