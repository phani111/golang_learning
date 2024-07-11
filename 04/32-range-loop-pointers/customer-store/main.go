package main

import "fmt"

type Customer struct {
	ID      string
	Balance float64
}

type Store struct {
	// m is a map that stores Customer pointers, keyed by the Customer ID.
	m map[string]*Customer
}

// main creates a new Store, adds some Customers to it, and then prints the contents of the Store.
func main() {
	s := Store{
		m: make(map[string]*Customer),
	}

	s.storeCustomers([]Customer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})

	printStore(s.m)
}

// storeCustomers adds the provided customers to the Store's internal map.
// For each customer, it prints the address of the customer object and stores a pointer to it in the map,
// keyed by the customer's ID.
// After adding all the customers, it prints a newline.
func (s *Store) storeCustomers(customers []Customer) {
	for _, customer := range customers {
		fmt.Printf("%p\n", &customer)
		s.m[customer.ID] = &customer
	}
	fmt.Println() // Add a newline after printing addresses
}

func printStore(store map[string]*Customer) {
	for id, customer := range store {
		fmt.Printf("ID: %s, Balance: %.2f\n", id, customer.Balance)
	}
}
