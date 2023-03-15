package author

import (
	"crud/author"
	"database/sql"
)

func UpdateDB(db *sql.DB) {
	updateStmt := `update "author" set "fullname"=$1 where "id"=$2`
	_, i := db.Exec(updateStmt, "Alizhan2", 1)
	author.CheckError(i)
}
