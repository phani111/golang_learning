/*
The first loop demonstrates a common pitfall when using range with struct values.
The second and third loops show correct ways to modify slice elements.
The fourth loop shows how using pointers can allow modifications even when using range.
The createAccounts() function returns a slice of account values.
The createAccountsPtr() function returns a slice of account pointers.
The printAccountsPtr() function is a custom way to print the slice of account pointers.
This code effectively illustrates the differences in behavior when working with value types vs pointer types in slices, and different looping techniques in Go
*/
package main

import (
	"fmt"
	"strings"
)

type account struct {
	balance float32
}

func main() {
	//This loop doesn't modify the original slice because a is a copy of each account. The changes are made to the copy, not the original.

	accounts := createAccounts()
	for _, a := range accounts {
		a.balance += 1000
	}
	fmt.Println(accounts)
	//This loop correctly modifies the original slice by accessing elements through their index.

	accounts = createAccounts()
	for i := range accounts {
		accounts[i].balance += 1000
	}
	fmt.Println(accounts)

	// This is a traditional for loop that also correctly modifies the original slice.

	accounts = createAccounts()
	for i := 0; i < len(accounts); i++ {
		accounts[i].balance += 1000
	}
	fmt.Println(accounts)

	// This loop works correctly because accountsPtr is a slice of pointers. Even though a is a copy of the pointer, it still points to the original account, so modifications affect the original data.

	accountsPtr := createAccountsPtr()
	for _, a := range accountsPtr {
		a.balance += 1000
	}
	printAccountsPtr(accountsPtr)

}

func createAccounts() []account {
	return []account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}
}

func createAccountsPtr() []*account {
	return []*account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}
}

func printAccountsPtr(accounts []*account) {
	sb := strings.Builder{}
	sb.WriteString("[")
	s := make([]string, len(accounts))
	for i, account := range accounts {
		s[i] = fmt.Sprintf("{%.0f}", account.balance)
	}

	sb.WriteString(strings.Join(s, " "))
	sb.WriteString("]")
	fmt.Println(sb.String())
}
