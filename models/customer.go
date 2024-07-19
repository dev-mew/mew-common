package models

type NewCustomer struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Dob       string `json:"dob"`
}

type Customer struct {
	JWT  string `json:"jwt"`
	User User   `json:"user"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Provider  string `json:"provider"`
	Confirmed bool   `json:"confirmed"`
	Blocked   bool   `json:"blocked"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Dob       string `json:"dob"`
}

type SigninReq struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
