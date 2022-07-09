package main

import "fmt"

func rot13(str string) string {
	var answer = ""
	for i := 0; i < len(str); i++ {
		asciiNum := str[i]
		if asciiNum >= 65 && asciiNum <= 77 || asciiNum >= 97 && asciiNum <= 109 {
			answer = answer + string(str[i]+13)
		} else if asciiNum >= 78 && asciiNum <= 90 || asciiNum >= 110 && asciiNum <= 122 {
			answer = answer + string(str[i]-13)
		} else {
			answer = answer + string(str[i])
		}
	}
	return answer
}

func main() {
	result := rot13("Pbatenghyngvba, lbh'er ba gur tbbq jnl gb wbva Chyfne NV !")
	fmt.Println(result)
}
