package main

import (
	"fmt"
	"strings"
)

// 곱셈 func
func multiply(a int, b int) int { // a, b 를 int로 argu 로받고 int type 으로 return 한다.
	// func multiply(a, b int) a, b 모두 int 라면 이렇게 사용도 가능하다.
	return a * b
}

// 2 개의 value 를 return 하는 func
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 위에서는 받는 argument 가 하나로 정해져 있는데 만약 정해져있지 않다면
// 받는 type 문 옆에 ... 점을 3개 찍어주면 연속된 arguments 를 받는다
// return type 은 적지 않는다 적으면 error
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Printf("곱셈 함수 2 * 2 : ")
	fmt.Println(multiply(2, 2))

	totalLenght, upperName := lenAndUpper("themath")
	fmt.Println("themath 의 길이는 : ", totalLenght)
	fmt.Println("themath 를 대문자로... : ", upperName)

	// 만약 lenAndUpper("themath") 에서 totalLenght 만 사용하고 싶다면
	// compiler 가 무시할 수 있게 사용하지 않는 리턴값에는 _(언더바)를 명시해주면 선언되지 않는다.
	totalLenght, _ = lenAndUpper("upperName")
	fmt.Println("upperName 의 길이는 : ", totalLenght)

	repeatMe("Amumu", "Aatrox", "Alistar", "Annie", "Anivia", "Ashe")

}
