package models

type Account struct {
	Model

	TokenUser string `json:"token_user" gorm:"index"`

	ProfilePicture string `json:"profile_picture" `
	DisplayName    string `json:"display_name"`
	Status         string `json:"status"`
	Username       string `json:"username"`
	FirstNameTh    string `json:"first_name_th"`
	LastNameTh     string `json:"last_name_th"`

	FirstNameEng string `json:"first_name_eng"`
	LastNameEng  string `json:"last_name_eng"`
	NicknameTh   string `json:"nickname_th"`
	NicknameEng  string `json:"nickname_eng"`

	Email string `json:"email"`

	MobileNo string `json:"mobile_no"`

	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
