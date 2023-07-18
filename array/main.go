package main

import "fmt"

func main() {

	// 기본적인 Array 는 다음과 같이 선언한다.
	// [5] 로 명시된 array 는 길이가 무조건 5로 선언되며 값을 더 넣을 수도없고
	// 값이 덜넣어졌다고 해도 무조건 길이는 5개가 된다. : 3개 넣으면 빈칸 2개로 총 5개길이
	// 또한 [len_of_array]string 같이 string 으로 선언되었으면
	// 해당 배열은 무조건 string 값만 넣을 수 있다.
	names := [5]string{"mike", "john", "young"}
	fmt.Println(names)
	fmt.Println("배열의 길이는 : ", len(names))
	// [mike john young  ]
	// 배열의 길이는 :  5

	names[3] = "cathy"
	names[4] = "halt"
	// names[5] = "july"  배열의 길이는 5개로 되어있기때문에 6번째는 정할 수 없다.
	fmt.Println(names)
	fmt.Println("배열의 길이는 : ", len(names))
	// [mike john young cathy halt]
	// 배열의 길이는 :  5

	// 제한 없는 배열의 선언 : slice
	names_infinity := []string{"mike", "john", "young"}
	names_infinity = append(names_infinity, "cathy")
	names_infinity = append(names_infinity, "june")
	names_infinity = append(names_infinity, "paul")
	fmt.Println(names_infinity)
	fmt.Println("배열의 길이는 : ", len(names_infinity))
	// [mike john young cathy june paul]
	// 배열의 길이는 :  6

	// append() 는 첫번째 argu 로 []type 이라는 slice 객체를 받고
	// 두번째는 배열의 마지막에 삽입할 값을 입력하면 된다.
	// 하지만 파이썬같은 array.append(value) 로 하면 바로 적용되는 것이아니라
	// append([]type, value) 가 하나의 slice 객체가 된다.
	// 그 뒤에 이미 선언된 slice 타입의 array 에다가 넣어주면
	// 이미 선언된 slice 타입의 array 는 값이 추가된 slice array 가된다.

}
