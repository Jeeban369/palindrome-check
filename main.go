package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("")
	fmt.Println("🛠️  CHECK YOUR PALINDROME HERE")
	fmt.Println("▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️")
	fmt.Println("")
	fmt.Print("🗯️  Type some string:")
	var input string
	fmt.Scanln(&input)
	fmt.Println("")


	normalized := ""
	for _, r := range input {
		if unicode.IsLetter(r) {
			normalized += string(unicode.ToLower(r))
		}
	}

	var isPalindrome bool = true
	length := len(normalized)

	for i:=0; i<length/2; i++{
		if normalized[i] != normalized[length-1-i]{
			isPalindrome = false
			break
		}
	}

	if isPalindrome{
		fmt.Println("It is a Palindrome. ✅ ")
		fmt.Println("")
	}else{
		fmt.Println("it is not a palindrome. ❎")
		fmt.Println("")
	}
}
