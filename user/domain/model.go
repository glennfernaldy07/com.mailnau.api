package domain

type User struct {
	ID        int64  `gorm:"column:id;primaryKey"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	CreatedAt string `gorm:"column:created_at"`
	CreatedBy string `gorm:"column:created_by"`
	UpdatedAt string `gorm:"column:updated_at"`
	UpdatedBy string `gorm:"column:updated_by"`
}

func (User) TableName() string {
	return "users"
}

type UserResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}
