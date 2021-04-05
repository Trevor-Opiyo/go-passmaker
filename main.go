package main

import (
	"fmt"
	"reflect"
	"strings"
	"encoding/binary"
	crypto_rand "crypto/rand"
	math_rand "math/rand"
)

func main() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
	panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	options := ""
	var passLen int
	uppers := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowers := strings.ToLower(uppers)
	nums := "1234567890"
	specs := "!@#$^&*()_-+=<>?"

	fmt.Println("----------------------------------\n        Password Generator\n----------------------------------")
	options = addOptions(options, uppers, "upper-case letters")
	options = addOptions(options, lowers, "lower-case letters")
	options = addOptions(options, nums, "integers")
	options = addOptions(options, specs, "special characters")
	passLen = passLength()
	generate(passLen, options)
}

func addOptions(original string, addition string, characterType string) string {
	var userInput string
	for {
		fmt.Println("Include " + characterType + "? [Y/N]")
		fmt.Scan(&userInput)
		switch userInput {
		case "Y", "y", "yes", "Yes", "YES":
			return original + addition
		case "N", "n", "no", "No", "NO":
			return original
		default:
			fmt.Println("Please enter a valid response.")
		}
	}
}

func passLength() int {
	for {
		var userInput int
		fmt.Println("Enter an integer length for the password.")
		fmt.Scan(&userInput)
		temp := reflect.TypeOf(userInput).Kind()
		if temp == reflect.Int {
			return userInput
		}
		fmt.Println("Please enter a valid response.")
	}
}

func generate(passLen int, options string) {
outer:
	for {
		fmt.Println("The password is:")
		for iterator := 0; iterator < passLen; iterator++ {
			opttionsLen := len(options)
			randIndex := math_rand.Intn(opttionsLen)
			fmt.Printf("%c", options[randIndex])
		}
		var userInput string
		fmt.Println("\nWould you like to regenerate another password with the same parameters? [Y/N]")
		fmt.Scan(&userInput)
		switch userInput {
		case "Y", "y", "yes", "Yes", "YES":
			break
		case "N", "n", "no", "No", "NO":
			break outer
		default:
			fmt.Println("Please enter a valid response.")
		}
	}
}
