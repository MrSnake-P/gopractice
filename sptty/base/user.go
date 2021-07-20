package base

const (
	ServiceUser = "user"
)

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Email        string `json:"email"`
	VerifyCode   string `json:"verify_code"'`
	VerifyStatus bool `json:"verify_status"`
}

type EmailStatus struct {
	Id       int    `json:"user_id"`
	Email        string `json:"email"`
	VerifyCode   string `json:"verify_code"'`
	VerifyStatus string `json:"verify_status"`
}
