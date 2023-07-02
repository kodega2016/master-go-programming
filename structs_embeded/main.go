package main

import "fmt"

func main() {
	e := Employee{
		name:  "Khadga Bahadur Shrestha",
		email: "khadgaloveocoding2016@gmail.com",
		role:  "Software Developer",
		address: Address{
			streetName: "Lainchaur Kathmandu",
			city:       "Kathmandu",
			country:    "Nepal",
		},
	}

	fmt.Println(e.name, e.address.country)
	e.print()
	fmt.Println("Employee email address:", e.email)
	fmt.Println("Employee address:", e.address.city)

}

type Employee struct {
	name    string
	email   string
	role    string
	address Address
}

type Address struct {
	streetName string
	city       string
	country    string
}

func (e Employee) print() {
	fmt.Printf("+%v\n", e)
}
