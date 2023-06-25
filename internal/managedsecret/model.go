package managedsecret

type ManagedSecret struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Value  string
	UserId uint
}
