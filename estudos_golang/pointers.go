package main

// & operator returns the address of a variable/function
// * opererator returns data at an address (dereferencing)
func pointer() {
	var x int = 1
	var y int
	var ip *int // ip is pointer to int

	ip = &x //ip now points to x
	y = *ip // y is now 1

}
