package dto

type FriendRequest struct {
	ID            string `json:"id"`
	FullName      string `json:"full_name"`
	IDCard        string `json:"id_card"`
	KTP           string `json:"ktp"`
	AccountNumber string `json:"account_number"`
	Salary        string `json:"salary"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Whatsapp      string `json:"whatsapp"`
	PlaceOfBirth  string `json:"place_of_birth"`
	DateOfBirth   string `json:"date_of_birth"`
	Address       string `json:"address"`
	StatusAccount bool   `json:"status_account"`
	UserID        string `json:"user_id"`
	LikeID        string `json:"like_id"`
	PackageID     string `json:"package_id"`
	CompanyID     string `json:"company_id"`
	CreatedBy     string `json:"created_by"`
}

type FriendResponse struct {
	FriendID string `json:"friend_id"`
	FullName string `json:"full_name"`
}
