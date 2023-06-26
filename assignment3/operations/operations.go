// Package operations is used for arithmetic operations of calculator project.
package operations

// Add returns the sum of two numbers.
func Add(numA int, numbB int) int {
	return numA + numbB
}

// Sub returns the difference between two numbers.
func Sub(numA int, numbB int) int {
	return numA - numbB
}

// Mul returns multiplication of two numbers.
func Mul(numA int, numbB int) int {
	return numA * numbB
}

// Div returns division of two numbers.
func Div(numA int, numbB int) (int, int) {
	return numA / numbB, numA % numbB   
}