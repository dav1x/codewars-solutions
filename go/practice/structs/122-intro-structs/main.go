package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}
	p3 := person{
		first: "davis",
		last:  "phillips",
		age:   45,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	//fmt.Printf("%T \t %#v\n", p1, p1)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Printf("He is %s %s American Hero!", p1.first, p1.last)
}

