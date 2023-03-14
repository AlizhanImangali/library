package book

import "sync"

type (
	Book struct {
		ID       int    `json:"id"`
		Name     string `json:"book_name"`
		Genre    string `json:"book_genre"`
		CodeISBN string `json:"book_code"`
	}
)

var (
	books = map[int]*Book{}
	seq   = 1
	lock  = sync.Mutex{}
)
