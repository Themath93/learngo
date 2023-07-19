package main

import "fmt"

func main() {
	mike := map[string]string{"name": "mike", "age": "12"}
	// fmt.Print(mike)
	// map[age:12 name:mike]
	// map 은 map[type of key]type of value{ key : value , key : value} 로 선언된다.
	// 해당 객체는 파이썬의 dict 구조와 비슷한데 선언된 type 만 사용할 수 있다.

	for key, value := range mike {
		fmt.Println(key, value)
	}
	// name mike
	// age 12

	// map 객체는 key, value 를 for loop 할 수 있다.
	// python 의 for k,v in dict.items(): 와 비슷하다

}
