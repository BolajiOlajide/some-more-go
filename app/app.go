package main

import (
	"fmt"
)

func main() {
	customers := GetCustomers()

	for _, customer := range customers {
		fmt.Println(customer, len(customers))
	}
}

// func getData() (customers []string) {
// 	// create one record
// 	customers = []string{
// 		"Marcel Dempers",
// 		"Bob Smith",
// 	}

// 	customers = append(customers, "Jaime Farnan")

// 	// for {
// 	// 	fmt.Println("Infinite Loop 1")
// 	// 	time.Sleep(time.Second)
// 	// 	break
// 	// }

// 	// for x := 0; x < len(customers); x++ {
// 	// 	fmt.Println("Infinite Loop 1", customers[x])
// 	// 	time.Sleep(time.Second)
// 	// }

// 	for x, customer := range customers {
// 		fmt.Println("Infinite Loop ", x+1, " ", customer)
// 		time.Sleep(time.Second)
// 	}

// 	return customers
// }
