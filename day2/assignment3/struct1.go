package assignment3

type Details struct {
	Name        string
	Age         int
	PhoneNumber int
	Address
}

type Address struct {
	City    string `json:"city"`
	State   string
	PinCode int
}
