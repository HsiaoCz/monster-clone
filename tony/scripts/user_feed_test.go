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
		{Username: "wangsu", Password: "wangsu1ssa5", Email: "wangsuw@gmail.com", Birthday: "2003/12/21", Gender: "male", Tags: []string{"美妆", "美食", "明星", "影视", "音乐", "旅行", "家居"}, IsAdmin: false},
		{Username: "songshenz", Password: "songshenzs5", Email: "songshenz@gmail.com", Birthday: "1988/03/21", Gender: "female", Tags: []string{"运动", "学习", "穿搭", "健身", "音乐", "电影", "阅读", "旅行"}, IsAdmin: false},
		{Username: "songhuizong", Password: "songhuiz12345", Email: "songhuizong@gmail.com", Birthday: "1978/01/01", Gender: "female", Tags: []string{"运动", "学习", "教育", "职场", "电影", "考试", "阅读", "健身"}, IsAdmin: false},
		{Username: "zhaoliu", Password: "zhaoliu123", Email: "zhaoliu@gmail.com", Birthday: "1996/05/21", Gender: "female", Tags: []string{"科技", "明星", "运动", "音乐", "电影"}, IsAdmin: false},
		{Username: "qianer", Password: "qianer123", Email: "qianers@gmail.com", Birthday: "1993/06/11", Gender: "female", Tags: []string{"科技", "明星", "运动", "音乐", "电影", "职场", "游戏"}, IsAdmin: false},
		{Username: "tanghe", Password: "tanghe12345", Email: "tanghe@gmail.com", Birthday: "1997/05/27", Gender: "female", Tags: []string{"职场", "科技", "明星", "运动", "音乐", "电影", "阅读", "教育"}, IsAdmin: false},
		{Username: "sandezi", Password: "sandezi1235", Email: "sandezi@gmail.com", Birthday: "1994/02/23", Gender: "male", Tags: []string{"学习", "科技", "明星", "运动", "音乐", "电影", "阅读", "游戏"}, IsAdmin: false},
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
