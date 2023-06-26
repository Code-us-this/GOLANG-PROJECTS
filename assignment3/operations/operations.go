// Package operations is used for arithmetic operations of calculator project.
package operations

// Add returns the sum of two numbers.
func Add(numA float32, numbB float32) float32 {
	return numA + numbB
}

// Sub returns the difference between two numbers.
func Sub(numA float32, numbB float32) float32 {
	return numA - numbB
}

// Mul returns multiplication of two numbers.
func Mul(numA float32, numbB float32) float32 {
	return numA * numbB
}

// Div returns division of two numbers.
func Div(numA float32, numbB float32) (int, int) {
	return int(numA) / int(numbB), int(numA) % int(numbB)
}
