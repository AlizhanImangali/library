package reader

type Reader struct {
	ID       string  `json:"id" db:"id"`
	FullName *string `json:"fullname" db:"fullname"`
	BookList *string `json:"booklist" db:"booklist"`
}
