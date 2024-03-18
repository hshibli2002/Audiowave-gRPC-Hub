package Models

type User struct {
	ID         int64  `json:"user_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Likes      int32  `json:"likes_given"`
	Followings int32  `json:"follows_given"`
}
