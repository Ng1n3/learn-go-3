package main

import (
	"composite-types/exercise"
	"fmt"
	"maps"
)

func main() {
	lecture()
	exercise.Exercise1()
  exercise.Exercise2()
  exercise.Exercise3()
}

func lecture() {
	// [lecture] -> Maps
	totalWins := map[string]int{}

	teams := map[string][]string{
		"orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	fmt.Println(totalWins)
	fmt.Println(teams["Lions"])

	// [lecture] -> map with default size
	ages := make(map[int][]string, 10)
	fmt.Println("ages: ", ages)

	// ! Maps are not comparble. You can check if they are equal to nil, but you cannot check if two maps have identical keys and values using == or differ using !=

	// ! you cannot use := to assign a value to a map key.
	// ? -> When you try to read the value assigned to a map key that was never set, the map returns the zero value for the map’s value type. In this case, the value type is an int, so you get back a 0.

	totalWin := map[string]int{}
	totalWin["Orcas"] = 1
	totalWin["Lions"] = 2
	fmt.Println(totalWin["Orcas"])
	fmt.Println(totalWin["Kittens"]) // 0
	totalWin["Kittens"]++
	fmt.Println(totalWin["Kittens"]) // 1
	totalWin["Lions"] = 3
	fmt.Println(totalWin["Lions"]) // 3

	// [lecture] -> The comma ok idiom
	// ? -> THe comma ok idiom is used in Go when you want to differentiate between reading a value and getting back the zero value.
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	// [Lecture] -> Deleting from Maps
	g := map[string]int{
		"hello": 5,
		"world": 10,
	}
	delete(g, "hello")
	fmt.Println(g)

	// [Lecture] -> Empty a map
	s := map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println(s, len(s))
	clear(s)
	fmt.Println(s, len(s))

	// [lecture] -> comparing Maps
	k := map[string]int{
		"hello": 5,
		"world": 10,
	}
	l := map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println(maps.Equal(k, l))

	// [lecture] -> Using maps as sets
	// ? -> Go doesn't include a set, but ou can use a map to simulate some of its features, one being that you cannot duplicate keys in a map.

	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	// [lecture] -> Structs

	type person struct {
		name string
		age  int
		pet  string
	}

	emmanuel := person{
		"Emmanuel",
		25,
		"dog",
	}

	fmt.Println(emmanuel.name)

	// [lecture] -> Anonymous structs
	/*
	  ! -> You cannot convert an instance of one struct to another if the fields don't match
	  ! -> You cannot convert an instance of one struct to another if threre is an additinal field
	  !-> You cannot convert an instance of one struct to another if fields are not in the same order
	  ! -> You cannot use == to compare an instance of one struct to another struct because they are of different types

	  ? -> Anonymous structs add a small twist: if two struct variables are being compared and at least one has a type that’s an anonymous struct, you can compare them
	*/
	pet := struct {
		name string
		kind string
	}{
		name: "Dragon",
		kind: "dog",
	}

	fmt.Println(pet.kind)

	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "bob",
		age:  50,
	}
	var n struct {
		name string
		age  int
	}
	n = f
	fmt.Println(f == n)
}
