package author

import (
	"crud/author"
	"database/sql"
)

func DeleteDB(db *sql.DB) {
	deleteState := `Delete from author where id = $1 `
	_, err := db.Exec(deleteState, 1)
	author.CheckError(err)
}
func DeleteAllRec(db *sql.DB) {
	deleteState := `Delete from author`
	_, err := db.Exec(deleteState)
	author.CheckError(err)
}
