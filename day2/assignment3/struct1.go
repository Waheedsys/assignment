package assignment3

type Details struct {
	Name         string
	Age          int
	Phone_number int
	Address
}

type Address struct {
	City    string
	State   string
	Pincode int
}
