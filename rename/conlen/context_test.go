package conlen

import (
	"context"
	"log"
	"testing"
)

type UserInfoKey string

const (
	CtxUserInfo UserInfoKey = "userInfo"
	CtxUserName UserInfoKey = "username"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxUserInfo, "11234567")

	func1(ctx)
}

func func1(ctx context.Context) {
	user_id, ok := ctx.Value(CtxUserInfo).(string)
	if !ok {
		log.Fatal("no context value")
	}
	ctx = context.WithValue(ctx, CtxUserName, "admin")

	func2(ctx)

	log.Println(user_id)
}

func func2(ctx context.Context) {
	username, ok := ctx.Value(CtxUserName).(string)
	if !ok {
		log.Fatal("no context value")
	}
	user_id, ok := ctx.Value(CtxUserInfo).(string)
	if !ok {
		log.Fatal("no context value")
	}
	log.Println("username : ", username, "user_id : ", user_id)
}
