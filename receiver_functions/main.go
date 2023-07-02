package main

import "fmt"

func main() {
	p := Person{
		name:  "Khadga bahadur shrestha",
		role:  "Software Developer",
		email: "khadgalovecoding2016@gmail.com",
	}

	p.showInfo()

	c := Counter{
		count: 0,
	}

	fmt.Println("count:", c.count)
	c.increment()

	fmt.Println("count:", c.count)

}

type Person struct {
	name  string
	email string
	role  string
}

func (p Person) showInfo() {
	fmt.Printf("Hello I am %s and i am working as a %s you can contact me through %s\n", p.name, p.role, p.email)
}

type Counter struct {
	count int
}

func (c *Counter) increment() {
	c.count++
}
