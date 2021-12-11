package serializer

type UserCreateSerializer struct {
	Username string `json:"username" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

type UserUpdateSerializer struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}
