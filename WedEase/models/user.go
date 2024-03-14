package models

// ------------------------------ Constants ------------------------------ //

// ------------------------------ Definitions ------------------------------ //

// ------------------------------ Structures ------------------------------ //

type User struct {
	UserName    string `json:"userName" bson:"userName"`
	Password    string `json:"password" bson:"password"`
	Email       string `json:"email" bson:"email"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
}

type UserLogin struct {
	UserName string `json:"userName" bson:"userName"`
	Password string `json:"password" bson:"password"`
}

type CreateUserResponse struct {
	Token  string `json:"token" bson:"token"`
	Status string `json:"status" bson:"status"`
}

type GenericRequest struct {
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"`
	CreatedBy string `json:"createdBy" bson:"createdBy"`
	UpdatedBy string `json:"updatedBy" bson:"updatedBy"`
}
