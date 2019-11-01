package main

//contact Info can be called as contactInfo
func main() {

	alex := person{
		firstName: "alex",
		lastName:  "sarf",
		contact:   contactInfo{email: "alex@gmail.com", zip: "1222111"},
	}
	alex.updateName("arjun")

	alex.printPerson()
}
