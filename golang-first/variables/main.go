package main

import "fmt"

func main() {
	// golang 은 type language 이다. java와 비슷
	const name = "mike"
	fmt.Println(name)
	// const 는 constant(상수) 이므로 한 번 선언되면 바꿀 수 없다.
	// 그래서 잘 사용하지는 않는다.

	// name = "minsoo" 불가 error
	var name_1 string = "minsoo"
	fmt.Println(name_1)
	name_1 = "misoo"
	fmt.Println(name_1)

	// var name_1 string = "minsoo" 가 너무 길다면 아래와 같이 써도된다.
	name_2 := "John"
	fmt.Println(":= 로선언한 이름은 : ", name_2)
	// 또한 축약형 := 는 func 안에서만 적용할 수 있다.

}
