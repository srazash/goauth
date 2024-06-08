package User

import (
	"crypto/sha256"
)

type usertype int

const (
	StandardUser = iota
	Admin
	SuperAdmin
)

type User struct {
	uid      int
	fullname string
	username string
	password []byte
	email    string
	usertype usertype
	enabled  bool
}

func (u *User) GetUid() int {
	return u.uid
}

func (u *User) SetUid(uid int) {
	u.uid = uid
}

func (u *User) GetFullname() string {
	return u.fullname
}

func (u *User) SetFullname(fullname string) {
	u.fullname = fullname
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) GetPassword() []byte {
	return u.password
}

func (u *User) SetPassword(password string) {
	hash := sha256.New()
	hash.Write([]byte(password))
	u.password = hash.Sum(nil)
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) GetUsertype() usertype {
	return u.usertype
}

func (u *User) SetUsertype(usertype usertype) {
	u.usertype = usertype
}

func (u *User) GetEnabled() bool {
	return u.enabled
}

func (u *User) SetEnabled(enabled bool) {
	u.enabled = enabled
}
