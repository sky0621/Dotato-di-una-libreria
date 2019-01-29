package form

import "Dotato-di-una-libreria/backend/model"

// User ...
type User struct {
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

// ParseToDto ...
func (u *User) ParseToDto() *model.User {
	return &model.User{
		Name:     u.Name,
		Mail:     u.Mail,
		Password: u.Password,
	}
}
