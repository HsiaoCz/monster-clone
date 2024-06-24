package store

type Store struct {
	US UserStorer
	CS CommentStorer
	PS PostStorer
}
