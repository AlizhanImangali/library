package reader

type (
	Reader struct {
		ID       string   `json:"id" db:"id"`
		FullName string   `json:"full name" db:"full_name"`
		BookList []string `json:"book list" db:"book_list"`
	}
)
