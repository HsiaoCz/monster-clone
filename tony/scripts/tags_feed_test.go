package scripts

import (
	"context"
	"os"
	"testing"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"github.com/joho/godotenv"
)

func TestCreateTags(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
	params := []types.CreateTagParams{
		// {Content: "运动"},
		// {Content: "学习"},
		// {Content: "穿搭"},
		// {Content: "教育"},
		// {Content: "美食"},
		// {Content: "影视"},
		// {Content: "职场"},
		// {Content: "家居"},
		// {Content: "游戏"},
		// {Content: "旅行"},
		// {Content: "健身"},
		// {Content: "搞笑"},
		// {Content: "考试"},
		{Content: "彩妆"},
	}
	ctx := context.Background()
	tagTestStore, err := newTagTestStore(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, param := range params {
		t.Log(len(param.Content))
		msg := param.Validate()
		if len(msg) != 0 {
			t.Fatal(msg)
		}
		tag := types.TagFromParams(param)
		resp, err := tagTestStore.CreateTags(ctx, tag)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v\n", resp)
	}
}

func TestGetTags(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
	t.Log(os.Getenv("TAGSCOLL"))
	ctx := context.Background()
	tagTestStore, err := newTagTestStore(ctx)
	if err != nil {
		t.Fatal(err)
	}
	tags, err := tagTestStore.GetTags(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", tags)
}
