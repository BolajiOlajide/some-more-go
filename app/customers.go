package main

// Customer definition for customers in the system
type Customer struct {
	FirstName string
	LastName  string
	FullName  string
}

// GetCustomers method used to fetch all customers in the system
func GetCustomers() (customers []Customer) {
	marcel := Customer{
		FirstName: "Marcel",
		LastName:  "Dempers",
	}

	bob := Customer{
		FirstName: "Bob",
		LastName:  "Smith",
	}

	customers = append(customers, marcel, bob)

	return customers
}
