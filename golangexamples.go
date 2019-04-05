//Package golangexamples has three functions ConcatSlice, Encrypt and EZGreetings
package golangexamples

import (
	"github.com/ehteshamz/greetings"
)

//ConcatSlice is a global function
func ConcatSlice(sliceToConcat []byte) string {
	str_orignal := ""
	for i := range sliceToConcat {
		str_orignal += string(sliceToConcat[i])
		if i < len(sliceToConcat)-1 {
			str_orignal += "-"
		}
	}
	return str_orignal
}

//Encrypt is a global function
func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for i := range sliceToEncrypt {
		sliceToEncrypt[i] = string(int(sliceToEncrypt[i]) + ceaserCount)[0]
	}
}

//EZGreetings is a global function
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
