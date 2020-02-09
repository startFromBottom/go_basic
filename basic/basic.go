package basic

import (
	"fmt"
	"strings"
)

// function

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // return 후 호출되는 부분
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // same with return length, uppercase
}

// function - multiple inputs

func repeatMe(words ...string) {
	fmt.Println(words)
}

// for 문 (while 문 없음)

func superAdd(numbers ...int) int {
	// var total int = 0
	total := 0
	// // 방법 1
	// for i := 0; i < len(numbers); i++ {
	// 	total += numbers[i]
	// }

	// 방법 2
	for _, number := range numbers {
		total += number
	}

	return total

}

// if, switch

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrink2(age int) bool {

	// switch age {
	// case 10:
	// 	return false
	// case 20:
	// 	return true
	// }
	// return false

	switch koreanAge := age + 2; koreanAge {
	case 20:
		return true
	}
	return false

}

// pointer

func pointer() {

	a := 2
	b := &a
	*b = 20 // a의 값도 20으로 변함

	fmt.Println(a, b)

}

// array

func array() {

	var names = []string{"nico", "lynn", "dal"}
	names = append(names, "aaaa")

}

// maps

func maps() {
	nico := map[string]string{"name":"nico", "age": "12"}
	for key, value := range nico {
		fmt.Println(key, value)
	}
}

// struct

type person struct {
	name string
	age int
	favFood []string
}

var favFood = []string{"a", "b", "c"}
var me = person{name:"name", age:18, favFood:favFood}
