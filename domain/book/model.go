package book

type Book struct {
	ID       string  `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	Genre    *string `json:"genre" db:"genre"`
	CodeISBN *string `json:"codeisbn" db:"codeisbn"`
}
