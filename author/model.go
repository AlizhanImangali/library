package author

import "sync"

type Author struct {
	ID        int    `json:"id"`
	FullName  string `json:"author_fullname"`
	Pseudonym string `json:"author_pseudo"`
	Specialty string `json:"author_specialty"`
}

var (
	authors = map[int]*Author{}
	seq     = 1
	lock    = sync.Mutex{}
)
