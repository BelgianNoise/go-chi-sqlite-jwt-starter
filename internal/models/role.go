package models

type Role string

const (
	Admin       Role = "Admin"
	RegularUser Role = "RegularUser"
)

func (r Role) String() string {
	return string(r)
}
