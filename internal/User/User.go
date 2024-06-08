package User

import (
	"crypto/sha256"
	"math/rand/v2"
)

type usertype int

const (
	StandardUser = iota
	Admin
	SuperAdmin
)

var Usertype = map[usertype]string{
	StandardUser: "Standard User",
	Admin:        "Administrator",
	SuperAdmin:   "Super Administrator",
}

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
	u.password = hashPassword(password)
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

func (u *User) StringifyUsertype() string {
	return Usertype[u.usertype]
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

func CreateUser(fullname string, username string, password string, email string, usertype usertype, enabled bool) *User {
	u := User{
		uid:      rand.IntN(1000),
		fullname: fullname,
		username: username,
		password: hashPassword(password),
		email:    email,
		usertype: usertype,
		enabled:  enabled,
	}

	return &u
}

func hashPassword(password string) []byte {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hash.Sum(nil)
}
