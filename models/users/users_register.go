package Users

type UserRegister struct {
	FirstName         string `json:"firstname"`
	LastName          string `json:"lastname"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Bio               string `json:"bio"`
	Profile_pic       string `json:"profile_pic"`
	Subscription_type string `json:"subscription_type"`
}
