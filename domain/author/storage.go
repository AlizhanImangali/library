package author

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

type Storage interface {
	CreateRow(data Author) (dest string, err error)
	GetRowByID(id string) (dest Author, err error)
	SelectRows() (dest []Author, err error)
	UpdateRow(data Author) (err error)
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

func (s *storage) CreateRow(data Author) (dest string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		INSERT INTO authors (full_name, pseudonym, specialty)
		VALUES ($1, $2, $3)
		RETURNING id`
	args := []any{data.FullName, data.Pseudonym, data.Specialty}
	err = s.db.QueryRowContext(ctx, query, args...).Scan(&dest)
	return
}

func (s *storage) GetRowByID(id string) (dest Author, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := `
		SELECT id, full_name, pseudonym, specialty
		FROM authors
		WHERE id=$1`

	args := []any{id}

	err = s.db.GetContext(ctx, &dest, query, args...)

	return
}

func (s *storage) SelectRows() (dest []Author, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT id, full_name, pseudonym, specialty
		FROM authors
		ORDER BY id`

	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *storage) UpdateRow(data Author) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sets, args := s.prepareArgs(data)
	if len(args) > 0 {

		args = append(args, data.ID)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")

		query := fmt.Sprintf("UPDATE authors SET %s WHERE id=$%d", strings.Join(sets, ", "), len(args))
		_, err = s.db.ExecContext(ctx, query, args...)
	}

	return
}

func (s *storage) prepareArgs(data Author) (sets []string, args []any) {
	if data.Pseudonym != nil {
		args = append(args, data.Pseudonym)
		sets = append(sets, fmt.Sprintf("pseudonym=$%d", len(args)))
	}
	if data.FullName != nil {
		args = append(args, data.FullName)
		sets = append(sets, fmt.Sprintf("full_name=$%d", len(args)))
	}
	if data.Specialty != nil {
		args = append(args, data.Specialty)
		sets = append(sets, fmt.Sprintf("specialty=$%d", len(args)))
	}

	return
}

func (s *storage) DeleteRow(id string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		DELETE 
		FROM authors
		WHERE id=$1`

	args := []any{id}

	_, err = s.db.ExecContext(ctx, query, args...)

	return
}
