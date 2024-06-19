package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"time"

	"github.com/HsiaoCz/monster-clone/leaf/conf"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username   string             `bson:"username" json:"username"`
	Password   string             `bson:"password" json:"-"`
	Level      int                `bson:"level" json:"level"`
	Content    string             `bson:"content" json:"content"`
	Email      string             `bson:"email" json:"email"`
	Job        string             `bson:"job" json:"job"`
	Company    string             `bson:"company" json:"company"`
	Birthday   string             `bson:"birthday" json:"birthday"`
	Age        int                `bson:"age" json:"age"`
	Gender     string             `bson:"gender" json:"gender"`
	Avatar     string             `bson:"avatar" json:"avatar"`
	Tags       []string           `bson:"tags" json:"tags"`
	Likes      string             `bson:"likes" json:"likes"`
	Describe   string             `bson:"describe" json:"describe"`
	Collection string             `bson:"collection" json:"collection"`
	IsAdmin    bool               `bson:"isAdmin" json:"isAdmin"`
}

type CreateUserParams struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Birthday string   `json:"birthday"`
	Gender   string   `json:"gender"`
	Tags     []string `json:"tags"`
	IsAdmin  bool     `json:"isAdmin"`
}

var (
	minUsernameLen = 2
	maxUsernameLen = 12
	minPasswordLen = 8
	maxPasswordLen = 16
	minTagsLen     = 4
	genderMap      = map[string]struct{}{"male": {}, "female": {}}
)

func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.Username) < minUsernameLen || len(params.Username) > maxUsernameLen {
		errors["username"] = fmt.Sprintf("the username shouldn't less then %d and more then %d", minUsernameLen, maxUsernameLen)
	}
	if len(params.Password) < minPasswordLen || len(params.Password) > maxPasswordLen {
		errors["password"] = fmt.Sprintf("the password shouldn't less then %d and more then %d", minPasswordLen, maxPasswordLen)
	}
	if len(params.Tags) < minTagsLen {
		errors["tags"] = fmt.Sprintf("the tags shouldn't less then %d", minTagsLen)
	}
	if !isValidGender(params.Gender) {
		errors["gender"] = fmt.Sprintf("the gender should use (%s) or (%s)", "male", "female")
	}
	if !isEmailValidate(params.Email) {
		errors["email"] = "invalid email"
	}
	if err := isBirthdayValidate(params.Birthday); err != nil {
		errors["birthday"] = err.Error()
	}
	return errors
}

func isValidGender(gender string) bool {
	_, ok := genderMap[gender]
	return ok
}

func isEmailValidate(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func isBirthdayValidate(b string) error {
	_, err := time.Parse("2006/01/02", b)
	return err
}

func NewUserFromParams(params CreateUserParams) *User {
	return &User{
		Username:   params.Username,
		Email:      params.Email,
		Password:   encryptPassword(params.Password),
		Level:      1,
		Content:    "",
		Job:        "",
		Company:    "",
		Birthday:   params.Birthday,
		Age:        getUserAge(params.Birthday),
		Gender:     params.Gender,
		Avatar:     "./static/user/avatar/1211.jpg",
		Tags:       params.Tags,
		Likes:      "0",
		Describe:   "0",
		Collection: "0",
		IsAdmin:    params.IsAdmin,
	}
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(conf.Conf.App.MD5Secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func getUserAge(birthday string) int {
	t, _ := time.Parse("2006/01/02", birthday)
	age := time.Now().Year() - t.Year()
	return age
}

type UpdateUserParmas struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	Job      string `json:"job"`
	Company  string `json:"company"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
}

func (param UpdateUserParmas) Validate() map[string]string {
	errors := map[string]string{}
	if !isValidGender(param.Gender) {
		errors["gender"] = fmt.Sprintf("the gender should use (%s) or (%s)", "male", "female")
	}
	if err := isBirthdayValidate(param.Birthday); err != nil {
		errors["birthday"] = err.Error()
	}
	return errors
}

type UserInfo struct {
	UserID  primitive.ObjectID
	Email   string
	IsAdmin bool
}
