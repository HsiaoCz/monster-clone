package scripts

import (
	"context"
	"testing"
	"time"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"github.com/joho/godotenv"
)

func TestCreateUser(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
	createUserParams := []types.CreateUserParams{
		{Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		{Username: "zhangsan", Password: "zhangsan123", Email: "zhangsan@gmail.com", Birthday: "1994/12/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐"}, IsAdmin: false},
		{Username: "lisi", Password: "lisi12345", Email: "slisi@gmail.com", Birthday: "2001/05/12", Gender: "female", Tags: []string{"科技", "明星", "运动", "音乐", "电影"}, IsAdmin: false},
		// {Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		// {Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		// {Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		// {Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		// {Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		// {Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
		// {Username: "HsiaoL1", Password: "shaw12345", Email: "shaw123@gmail.com", Birthday: "1998/04/01", Gender: "female", Tags: []string{"美妆", "科技", "明星", "运动", "音乐", "电影", "阅读"}, IsAdmin: false},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	testUserStore, err := newUserfeed(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, params := range createUserParams {
		msg := params.Validate()
		if len(msg) != 0 {
			t.Fatal(msg)
		}
		user := types.NewUserFromParams(params)
		userResp, err := testUserStore.CreateUser(ctx, user)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", userResp)
	}
}
