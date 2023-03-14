package reader

import "sync"

type (
	Reader struct {
		ID       int      `json:"id"`
		FullName string   `json:"reader_fullname"`
		BookList []string `json:"reader_booklist"`
	}
)

var (
	readers = map[int]*Reader{}
	seq     = 1
	lock    = sync.Mutex{}
)
