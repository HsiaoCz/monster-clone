package types

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"time"

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
	minUsernameLen    = 2
	maxUsernameLen    = 12
	minPasswordLen    = 8
	maxPasswordLen    = 16
	minTagsLen        = 4
	maxUserContentLen = 1000
	genderMap         = map[string]struct{}{"male": {}, "female": {}}
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
		Avatar:     "./static/avatar/1211.jpg",
		Tags:       params.Tags,
		Likes:      "0",
		Describe:   "0",
		Collection: "0",
		IsAdmin:    params.IsAdmin,
	}
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(os.Getenv("MD5SECRET")))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func getUserAge(birthday string) int {
	t, _ := time.Parse("2006/01/02", birthday)
	age := time.Now().Year() - t.Year()
	return age
}

type UpdateUserParmas struct {
	Username string   `json:"username,omitempty"`
	Content  string   `json:"content,omitempty"`
	Job      string   `json:"job,omitempty"`
	Company  string   `json:"company,omitempty"`
	Birthday string   `json:"birthday,omitempty"`
	Gender   string   `json:"gender,omitempty"`
	Avatar   string   `json:"avatar,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

func (params UpdateUserParmas) Validate() map[string]string {
	errors := map[string]string{}
	if params.Username != "" {
		if len(params.Username) < minUsernameLen || len(params.Username) > maxUsernameLen {
			errors["username"] = fmt.Sprintf("the username shouldn't less then %d and more then %d", minUsernameLen, maxUsernameLen)
		}
	}
	if params.Content != "" {
		if len(params.Content) > maxUserContentLen {
			errors["content"] = fmt.Sprintf("the content shouldn't more then %d", maxUserContentLen)
		}
	}
	if err := isBirthdayValidate(params.Birthday); err != nil {
		errors["birthday"] = err.Error()
	}
	return errors
}

type InsertDBUpdateUserParmas struct {
	Username string
	Content  string
	Job      string
	Company  string
	Birthday string
	Age      int
	Gender   string
	Avatar   string
	Tags     []string
}

func NewInstertDBUpdateUserParams(param UpdateUserParmas) *InsertDBUpdateUserParmas {
	age := getUserAge(param.Birthday)
	return &InsertDBUpdateUserParmas{
		Age:      age,
		Content:  param.Content,
		Job:      param.Job,
		Username: param.Username,
		Company:  param.Company,
		Gender:   param.Gender,
		Avatar:   param.Avatar,
		Tags:     param.Tags,
		Birthday: param.Birthday,
	}
}

type UserInfo struct {
	UserID  primitive.ObjectID
	Email   string
	IsAdmin bool
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
