package base

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Email        string `json:"email"`
	VerifyCode   string `json:"verify_code"'`
	VerifyStatus string `json:"verify_status"`
}
