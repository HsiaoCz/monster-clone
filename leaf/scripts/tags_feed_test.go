package scripts

import (
	"context"
	"testing"
	"time"

	"github.com/HsiaoCz/monster-clone/leaf/models"
)

func TestCreateTags(t *testing.T) {
	tags := []models.CreateTagParams{
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
		{Content: "美妆"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	tagTestStore, err := newTestTagStore(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, params := range tags {
		msg := params.Validate()
		if len(msg) != 0 {
			t.Fatal(msg)
		}
		tag := models.TagFromParams(params)
		tagResp, err := tagTestStore.CreateTags(ctx, tag)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", tagResp)
	}
}
