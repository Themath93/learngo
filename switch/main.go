package main

import "fmt"

func canIDrink(age int) bool {
	// 너무많은 if 에 대한 else if 를 대신하기위한 switch 문 java와 거의 비슷
	switch koreanAge := age + 2; koreanAge {

	case 10:
		return false
	case 18:
		return true
	}
	return true

}

func main() {
	fmt.Println(canIDrink(16))
}
