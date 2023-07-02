package main

import "fmt"

func main() {
	var employee map[string]string = make(map[string]string)

	//assigning values to the map
	employee["name"] = "Khadga Bahadur Shrestha"
	employee["address"] = "Madhumalla Morang"
	employee["role"] = "Software Developer"
	employee["password"] = "password"

	// get the value from the map
	name, isOk := employee["name"]
	if isOk {
		fmt.Println("name:", name)
	}

	// delete value from the map
	delete(employee, "password")

	// iterate over the map
	for key, value := range employee {
		fmt.Println(key, value)
	}

	// get total number of key-value pairs
	fmt.Println("total pairs:", len(employee))

	fmt.Println(employee)

	balances := map[string]float32{
		"NP": 12.0,
		"EN": 8.0,
		"ES": 4.0,
	}

	fmt.Println(balances)

	balances["USD"] = 15.0
	balances["NED"] = 53.0

	for k, v := range balances {
		fmt.Println(k, v)
	}
}
