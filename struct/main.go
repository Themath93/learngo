package main

import "fmt"

type person struct {
	name    string
	age     int
	sex     string
	favFood []string
}

func main() {
	favFood := []string{"ramen", "burger", "pizza"}
	mike := person{name: "mike",
		age:     18,
		sex:     "male",
		favFood: favFood}
	fmt.Println(mike)
	// {mike 18 male [ramen burger pizza]}
}
