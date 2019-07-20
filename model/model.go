package model

type User struct {
	Username  string `json:"username"`
	Email 	  string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Savings   []Saving `json:"savings"`
}

type Saving struct {
	ID        int	 `json:"ID"`
	Category  string `json:"category"`
	Amount    string `json:"amount"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
	Token  string `json:"token"`
}
