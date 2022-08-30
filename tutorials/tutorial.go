package main

import (
	"fmt"
	"log"
)

func main() {

	logger()
	loop()
	condition()
	array()
	_map()
}

func logger() {
	// The := syntax is shorthand for declaring and initializing a variable
	n := 1
	s := "one"

	// print
	log.Print("Print")
	log.Println("Println")
	log.Printf("Printf %d", n)
	log.Printf("Printf %s", s)

	// os exist
	// log.Fatal("Fatal")
	// log.Fatalln("Fatalln")
	// log.Fatalf("Printf %d", n)
	// log.Fatalf("Printf %s", s)

	// fmt print
	fmt.Println("Println")
	fmt.Printf("Printf %d \n", n)

}

func loop() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func condition() {

	if 1 > 0 {
		fmt.Println("ok")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}
func array() {
	var a [5]int
	fmt.Println("eiei:", a)
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
}

func _map() {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	m["a1"] = 7
	m["a2"] = 13

	fmt.Println("map:", m)

	v1 := m["a1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "a2")
	fmt.Println("map:", m)

	/* The optional second return value when getting a value from a map indicates
	if the key was present in the map. This can be used to disambiguate between
	missing keys and keys with zero values like 0 or "". Here we didn’t
	need the value itself, so we ignored it with the blank identifier _.
	*/
	_, prs := m["a2"]
	fmt.Println("prs:", prs)

	// Note that maps appear in the form map[k:v k:v]
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
