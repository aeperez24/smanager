package user

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Password string
	Enabled  bool
}
