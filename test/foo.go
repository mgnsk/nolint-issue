package main

import "fmt"

func main() {
	pouet() //nolint:errcheck
	fmt.Println("coucou")
}

func pouet() error {
	return nil
}
