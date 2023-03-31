package reader

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type Storage interface {
	CreateRow(data Reader) (dest string, err error)
	GetRowByID(id string) (dest Reader, err error)
	SelectRows() (dest []Reader, err error)
	UpdateRow(data Reader) (err error)
	DeleteRow(id string) (err error)
}

type storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) Storage {
	return &storage{
		db: db,
	}
}
func (s *storage) CreateRow(data Reader) (dest string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		INSERT INTO Readers (full_name,book_list)
		VALUES ($1, $2, $3)
		RETURNING id`

	args := []any{data.FullName, data.BookList}

	err = s.db.QueryRowContext(ctx, query, args...).Scan(&dest)

	return
}

func (s *storage) GetRowByID(id string) (dest Reader, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT created_at, updated_at,full_name,book_list 
		FROM readers
		WHERE id=$1`

	args := []any{id}

	err = s.db.GetContext(ctx, &dest, query, args...)

	return
}

func (s *storage) SelectRows() (dest []Reader, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT created_at, updated_at, id, full_name,book_list
		FROM readers`

	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *storage) UpdateRow(data Reader) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sets, args := s.prepareArgs(data)
	if len(args) > 0 {

		args = append(args, data.ID)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")

		query := fmt.Sprintf("UPDATE readers SET %s WHERE id=$%d", strings.Join(sets, ", "), len(args))
		_, err = s.db.ExecContext(ctx, query, args...)
	}
	return
}

func (s *storage) prepareArgs(data Reader) (sets []string, args []any) {
	if data.BookList != nil {
		args = append(args, data.BookList)
		sets = append(sets, fmt.Sprintf("bookList=$%d", len(args)))
	}
	return
}

func (s *storage) DeleteRow(id string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		DELETE 
		FROM readers
		WHERE id=$1`

	args := []any{id}

	_, err = s.db.ExecContext(ctx, query, args...)

	return
}
