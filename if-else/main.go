package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	// variable expression : if 문안에서만 사용할 수 있는 변수를 선언가능
	if koreanAge := age + 2; koreanAge > 18 {
		return true
	}

	return false
}

func main() {
	fmt.Println(canIDrink(17))
}
