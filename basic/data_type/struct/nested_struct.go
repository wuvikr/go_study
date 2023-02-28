package main

import "fmt"

type Location struct {
	continent, country, province, city string
}

type Person struct {
	name string
	age  int
	sex  string
	Location
}

func main() {
	zhangsan := Person{
		name: "zhangsan",
		age:  18,
		sex:  "male",
		Location: Location{
			continent: "Asia",
			country:   "China",
			province:  "Shanghai",
			city:      "Shanghai",
		},
	}
	fmt.Printf("zhangsan: %v\n", zhangsan)
	fmt.Printf("zhangsan.address.city: %v\n", zhangsan.city)
}
