package feed

import (
	"sort"

	"morgan.io/internal/post"
)

type Feed struct {
	Posts []*post.Post
}

func NewFeed(posts []*post.Post) *Feed {
	return &Feed{
		Posts: mix(posts),
	}
}

func mix(posts []*post.Post) []*post.Post {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})
	// TODO: sort by likes and comments

	return posts
}
