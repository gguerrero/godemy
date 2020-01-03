// Package dog allows us to more fully understand dogs.
package godog

// ToDogYears converts human years to dog years.
func ToDogYears(n int) int {
	return n * 7
}

// ToDogYears2 converts human years to dog years.
func ToDogYears2(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}
