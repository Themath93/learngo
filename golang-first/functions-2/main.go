package main

import (
	"fmt"
	"strings"
)

// naked return 방식의 func
// defer 를 통해 함수가 return 이후에 작동할 코드를 func 안에 남길 수 있다.
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
	// 여기서의 출력은
	// I'm done
	// 7 THEMATH
	// 이라고 출력이 되는데 이는
	// 1. lenAndUpper_2("themath") 라는 함수 호출을 하여 함수가 역할을 수행한다.
	// 2. totalLenght, up 라는 변수에 값이 저장된다.
	// 3. 그러면 이제 lenAndUpper_2() 함수는 종료가된다.
	// lenAndUpper_2() 안의 defer 가 함수종료을 인식하면 defer 이후의 코드가 작동한다.
	// 그 이후 fmt.Println(totalLenght, up) 이 실행된다.

}
