package author

import (
	"crud/author"
	"database/sql"
)

func CreateDB(db *sql.DB) {
	createData := `Create table author ( id  int primary key,fullname text,pseudo text,specialty text)`
	_, i := db.Exec(createData)
	author.CheckError(i)
}
