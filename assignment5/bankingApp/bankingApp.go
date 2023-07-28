// Package bankingapp helps with customer interactions with bank
package bankingapp

import "fmt"

type customer struct {
	firstName, lastName, accountType string
	accountNo, balance               int
	accountStatus                    bool
}

var customerData []customer
var customerLogin = make(map[int]int)

func createNewAccount(inCustomerData []customer) {

	inCustomerData = append(inCustomerData)

}

func register(cust customer) {

	//Prompt the customer to enter the firstName
	fmt.Print("Enter the first name: ")
	_, err := fmt.Scan(&cust.firstName)
	if err != nil {
		panic(err)
	}

	//Prompt the customer to enter the lastName
	fmt.Print("Enter the last name: ")
	_, err = fmt.Scan(&cust.lastName)
	if err != nil {
		panic(err)
	}

	//Prompt the customer to enter the account type
	fmt.Print("Please select the account type (Enter 'S' for Savings or 'C' for Checking ): ")

	//Prompt the customer to enter the lastName
	fmt.Print("Enter the last name: ")
	_, err = fmt.Scan(&cust.lastName)
	if err != nil {
		panic(err)
	}

}
