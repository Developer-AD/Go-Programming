package main

import "fmt"

func main() {

	// No Shorthand for const.
	// Const - only two way
	// way-1: const varname dtype = value
	// way-2: const varname = value

	const company_name = "Aceable Cyber Solutions"

	fmt.Println("----------- Constants -----------")
	// company_name = "New Company" // cannot assign to company_name
	fmt.Printf("Company - %s", company_name)

	// const new_company = "Cybler solutions."
	// const new_company string = "Microsoft"
	const new_company string = "Microsoft Pvt. Ltd."
	// new_company = "Google.com"

	fmt.Printf("\nNew Company - %s", new_company)

	// Multiple elements
	const (
		port = 8000
		host = "localhost"
		db   = "postgres"
	)
	fmt.Printf("\nport - %d", port)
	fmt.Printf("\nhost - %s", host)
	fmt.Printf("\ndb - %s", db)
}
