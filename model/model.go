package model

type User struct {
	Username  string `json:"username"`
	Email 	  string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Savings    []Saving
}

type Saving struct {
	Category  string `json:"category"`
	Date      string `json:"date"`
	Amount    string `json:"amount"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
	Token  string `json:"token"`
}
