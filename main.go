package main

import (
	"fmt"
)

// fmt 는 package 이며 import "" 로 import 할 수 있다.

func main() {
	fmt.Println("Hello world!")
	// fmt 라는 package 안에 Println() 이라는 func() 함수가 있으며 package 안의 함수의 첫 알파벳은 대문자로 온다.
	// something.SayHello()
	// something.sayBye()
	// error undefined: something.sayBye -> func 의 첫알파벳이 소문자라면 Private 로 분류되어 pakage의 func 로 import 가 불가능다.

	// variable

}
