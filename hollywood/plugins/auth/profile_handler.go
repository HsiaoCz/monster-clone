package auth

import (
	"fmt"
	"hollywood/app/db"

	"github.com/anthdm/superkit/kit"
	v "github.com/anthdm/superkit/validate"
)

var profileSchema = v.Schema{
	"username": v.Rules(v.Min(3), v.Max(50)),
}

type ProfileFormValues struct {
	ID       uint   `form:"id"`
	Username string `form:"username"`
	Email    string
	Success  string
}

func HandleProfileShow(kit *kit.Kit) error {
	auth := kit.Auth().(Auth)

	var user User
	if err := db.Get().First(&user, auth.UserID).Error; err != nil {
		return err
	}

	formValues := ProfileFormValues{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return kit.Render(ProfileShow(formValues))
}

func HandleProfileUpdate(kit *kit.Kit) error {
	var values ProfileFormValues
	errors, ok := v.Request(kit.Request, &values, profileSchema)
	if !ok {
		return kit.Render(ProfileForm(values, errors))
	}

	auth := kit.Auth().(Auth)
	if auth.UserID != values.ID {
		return fmt.Errorf("unauthorized request for profile %d", values.ID)
	}
	err := db.Get().Model(&User{}).
		Where("id = ?", auth.UserID).
		Updates(&User{
			Username: values.Username,
		}).Error
	if err != nil {
		return err
	}

	values.Success = "Profile successfully updated!"
	values.Email = auth.Email

	return kit.Render(ProfileForm(values, v.Errors{}))
}
