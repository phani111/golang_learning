package main

import "fmt"

type Customer struct {
	ID      string
	Balance float64
}

type Store struct {
	m map[string]*Customer
}

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
