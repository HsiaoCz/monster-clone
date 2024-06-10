package store

type Store struct {
	User    UserStorer
	Tag     TagStorer
	Post    PostStorer
	Comment CommentStorer
}
