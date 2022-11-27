package dto

import "WebProject/ginEssential/mapping"

type UserDto struct {
	Id        string
	Name      string
	Telephone string
	Age       int
}

func ToUserDto(user mapping.User) UserDto {

	return UserDto{
		Id:        user.Id,
		Name:      user.Name,
		Telephone: user.Telephone,
		Age:       user.Age,
	}
}
