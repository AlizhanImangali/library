package book

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

type Storage interface {
	CreateRow(data Book) (dest string, err error)
	GetRowByID(id string) (dest Book, err error)
	SelectRows() (dest []Book, err error)
	UpdateRow(data Book) (err error)
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
func (s *storage) CreateRow(data Book) (dest string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		INSERT INTO books (name,genre,codeisbn)
		VALUES ($1, $2, $3)
		RETURNING id`

	args := []any{data.Name, data.Genre, data.CodeISBN}
	err = s.db.QueryRowContext(ctx, query, args...).Scan(&dest)

	return
}

func (s *storage) GetRowByID(id string) (dest Book, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := `
		SELECT id, name, genre ,codeisbn
		FROM books
		WHERE id=$1`

	args := []any{id}

	err = s.db.GetContext(ctx, &dest, query, args...)

	return
}

func (s *storage) SelectRows() (dest []Book, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT  id, name, genre, codeisbn
		FROM books
		ORDER BY id`

	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *storage) UpdateRow(data Book) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sets, args := s.prepareArgs(data)
	if len(args) > 0 {

		args = append(args, data.ID)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")

		query := fmt.Sprintf("UPDATE books SET %s WHERE id=$%d", strings.Join(sets, ", "), len(args))
		_, err = s.db.ExecContext(ctx, query, args...)
	}
	return
}

func (s *storage) prepareArgs(data Book) (sets []string, args []any) {
	if data.Name != nil {
		args = append(args, data.Name)
		sets = append(sets, fmt.Sprintf("name=$%d", len(args)))
	}
	if data.Genre != nil {
		args = append(args, data.Genre)
		sets = append(sets, fmt.Sprintf("genre=$%d", len(args)))
	}

	if data.CodeISBN != nil {
		args = append(args, data.CodeISBN)
		sets = append(sets, fmt.Sprintf("codeISBN=$%d", len(args)))
	}
	return
}

func (s *storage) DeleteRow(id string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		DELETE 
		FROM books
		WHERE id=$1`

	args := []any{id}

	_, err = s.db.ExecContext(ctx, query, args...)

	return
}
