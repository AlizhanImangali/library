package author

import (
	"crud/author"
	"database/sql"
)

func InsertDB(db *sql.DB) {
	insertStmt := `insert into author ("id", "fullname", "pseudo", "specialty") values(1, 'John Priest','Cory','Roman')`
	_, i := db.Exec(insertStmt)
	author.CheckError(i)
}

/*import (
	"crud/book"
	"database/sql"
)

func InsertDB(db *sql.DB) {
	insertStmt := `insert into "books"("id", "name", "genre", "codeisbn") values(3, '3','Comedy2','8985665-4554654-594')`
	_, i := db.Exec(insertStmt)
	book.CheckError(i)
}*/
