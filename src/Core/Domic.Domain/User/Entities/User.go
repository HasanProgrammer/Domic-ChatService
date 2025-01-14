package UserEntity

import (
	"Domic.Domain/Chat/Entities"
	"errors"
	"regexp"
	"time"
)

type User struct {
	id        string
	fullName  string
	email     string
	createdBy string
	createdAt time.Time
	version   string
	isDeleted bool
	isActive  bool
	chats     *[]ChatEntity.Chat
}

func (user *User) GetFullName() string {
	return user.fullName
}

func (user *User) GetEmail() string {
	return user.email
}

func NewUser(fullName string, email string) (*User, error) {
	if len(fullName) <= 3 && len(fullName) >= 100 {
		return nil, errors.New("نام و نام خانوادگی باید بیشتر از 3 و کمتر از 100 عبارت داشته باشد!")
	}

	rc := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !rc.MatchString(email) {
		return nil, errors.New("فرمت ایمیل ارسالی صحیح نمی باشد!")
	}

	return &User{}, nil
}
