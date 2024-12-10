package storage

import (
	"context"
	"testing"

	"github.com/HsiaoCz/monster-clone/luccia/st"
)

func TestCreateUser(t *testing.T) {
	if err := Init(); err != nil {
		t.Fatal(err)
	}
	userColl := client.Database(test_dbname).Collection("user_test")
	us := UserStoreInit(client, userColl)
	user := st.NewUserFromReq(st.CreateUserParam{
		Username: "zhangsan",
		Email:    "lisi@gmai.com",
		Password: "122334asd",
		IsAdmin:  false,
	})
	us.CreateUser(context.Background(),user)
}

func TestGetUserByID(t *testing.T) {

}
