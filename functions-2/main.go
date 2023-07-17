package main

import (
	"fmt"
	"strings"
)

// naked return 방식의 func
func lenAndUpper_2(name string) (lenght int, uppercase string) {

	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
	// 반환할 타입 앞에 미리 변수를 입력해주면 굳이 return var1, var2 ... 이런식으로 안해도된다.
	// 기본형
	// return lenght, uppercase
}

func main() {
	totalLenght, up := lenAndUpper_2("themath")
	fmt.Println(totalLenght, up)
}
