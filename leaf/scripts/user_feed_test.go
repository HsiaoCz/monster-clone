package scripts

import (
	"context"
	"testing"
	"time"

	"github.com/HsiaoCz/monster-clone/leaf/models"
)

func TestCreateUser(t *testing.T) {
	createUserParams := []models.CreateUserParams{
		{Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		{Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		{Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		{Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		{Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		{Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		{Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		{Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		{Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		{Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	testUserStore, err := newTestUserStore(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, params := range createUserParams {
		msg := params.Validate()
		if len(msg) != 0 {
			t.Fatal(msg)
		}
		user := models.NewUserFromParams(params)
		userResp, err := testUserStore.CreateUser(ctx, user)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", userResp)
	}
}
