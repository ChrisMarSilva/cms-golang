package repositories

import (
	"context"
	"database/sql"
	"log"
	"time"

	data "github.com/chrismarsilva/cms.golang.tnb.cripo.database"
	"github.com/chrismarsilva/cms.golang.tnb.cripo.domain/models"
	"github.com/google/uuid"
)

type IUserRepository interface {
	GetById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*models.UserModel, error)
	GetByEmail(ctx context.Context, tx *sql.Tx, email string) (*models.UserModel, error)
	GetAll(ctx context.Context, tx *sql.Tx) (*[]models.UserModel, error)

	Create(ctx context.Context, tx *sql.Tx, user *models.UserModel) error
	Update(ctx context.Context, tx *sql.Tx, user *models.UserModel) error
	Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID)
}

type UserRepository struct {
	db *data.Database
}

func NewUserRepository(db *data.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (this UserRepository) GetById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*models.UserModel, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE id = ?`

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(timeoutCtx, query, id)
	} else {
		row = this.db.QueryRowContext(timeoutCtx, query, id)
	}

	user := &models.UserModel{} // var user models.User
	err := row.Scan(&user.ID, &user.Nome, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	log.Error("Erro no ErrNoRows:", err.Error())
		// 	return nil, fmt.Errorf("No user found with Email '%s'", email)
		// }
		log.Println("Erro no Scan:", err.Error())
		return nil, err
	}

	//log.Info("ID:", user.ID, "Nome:", user.Nome, "Email:", user.Email, "Password:", user.Password, "IsActive:", user.IsActive, "Created_at:", user.Created_at)
	return user, nil
}

func (this UserRepository) GetByEmail(ctx context.Context, tx *sql.Tx, email string) (*models.UserModel, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE email = ?`

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(timeoutCtx, query, email)
	} else {
		row = this.db.QueryRowContext(timeoutCtx, query, email)
	}

	user := &models.UserModel{}
	err := row.Scan(&user.ID, &user.Nome, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	log.Error("Erro no ErrNoRows:", err.Error())
		// 	return nil, fmt.Errorf("No user found with Email '%s'", email)
		// }
		log.Println("Erro no Scan:", err.Error())
		return nil, err
	}

	//log.Info("ID:", user.ID, "Nome:", user.Nome, "Email:", user.Email, "Password:", user.Password, "IsActive:", user.IsActive, "Created_at:", user.Created_at)
	return user, nil
}

func (this UserRepository) GetAll(ctx context.Context, tx *sql.Tx) (map[int]models.UserModel, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var rows *sql.Rows
	var err error
	query := `SELECT * FROM users`

	if tx != nil {
		rows, err = tx.QueryContext(timeoutCtx, query)
	} else {
		rows, err = this.db.QueryContext(timeoutCtx, query)
	}
	if err != nil {
		return nil, err
	}

	//var users = make([]models.UserModel, 0)
	users := make(map[int]models.UserModel)

	for rows.Next() {
		var user models.UserModel

		err := rows.Scan(&user.ID, &user.Nome, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt)
		if err != nil {
			log.Println("Erro no Scan:", err.Error())
			return nil, err
		}

		//users = append(users, user)
		users[len(users)] = user
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	//log.Info("ID:", user.ID, "Nome:", user.Nome, "Email:", user.Email, "Password:", user.Password, "IsActive:", user.IsActive, "Created_at:", user.Created_at)
	return users, nil
}

func (this UserRepository) Create(ctx context.Context, tx *sql.Tx, user *models.UserModel) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var stmt *sql.Stmt
	var err error
	query := `INSERT INTO users (id, nome, email, password, is_active, created_at) VALUES (?, ?, ?, ?, ?, ?)`

	if tx != nil {
		stmt, err = tx.Prepare(query)
	} else {
		stmt, err = this.db.Prepare(query)
	}

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(timeoutCtx, user.ID, user.Nome, user.Email, user.Password, user.IsActive, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (this UserRepository) Update(ctx context.Context, tx *sql.Tx, user *models.UserModel) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var stmt *sql.Stmt
	var err error
	query := `UPDATE users SET nome = ?, password = ?, is_active = ? WHERE id = ?`

	if tx != nil {
		stmt, err = tx.Prepare(query)
	} else {
		stmt, err = this.db.Prepare(query)
	}

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(timeoutCtx, user.Nome, user.Password, user.IsActive, user.ID.String())
	if err != nil {
		return err
	}

	return nil
}

func (this UserRepository) Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var stmt *sql.Stmt
	var err error
	query := `DELETE FROM users WHERE id = ?`

	if tx != nil {
		stmt, err = tx.Prepare(query)
	} else {
		stmt, err = this.db.Prepare(query)
	}

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(timeoutCtx, id) // .String()
	if err != nil {
		return err
	}

	return nil
}

/*
import (
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
    "github.com/koddr/tutorial-go-fiber-rest-api/app/models"
)

// BookQueries struct for queries from Book model.
type BookQueries struct {
    *sqlx.DB
}

// GetBooks method for getting all books.
func (q *BookQueries) GetBooks() ([]models.Book, error) {
    // Define books variable.
    books := []models.Book{}

    // Define query string.
    query := `SELECT * FROM books`

    // Send query to database.
    err := q.Get(&books, query)
    if err != nil {
        // Return empty object and error.
        return books, err
    }

    // Return query result.
    return books, nil
}

// GetBook method for getting one book by given ID.
func (q *BookQueries) GetBook(id uuid.UUID) (models.Book, error) {
    // Define book variable.
    book := models.Book{}

    // Define query string.
    query := `SELECT * FROM books WHERE id = $1`

    // Send query to database.
    err := q.Get(&book, query, id)
    if err != nil {
        // Return empty object and error.
        return book, err
    }

    // Return query result.
    return book, nil
}

// CreateBook method for creating book by given Book object.
func (q *BookQueries) CreateBook(b *models.Book) error {
    // Define query string.
    query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

    // Send query to database.
    _, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.UserID, b.Title, b.Author, b.BookStatus, b.BookAttrs)
    if err != nil {
        // Return only error.
        return err
    }

    // This query returns nothing.
    return nil
}

// UpdateBook method for updating book by given Book object.
func (q *BookQueries) UpdateBook(id uuid.UUID, b *models.Book) error {
    // Define query string.
    query := `UPDATE books SET updated_at = $2, title = $3, author = $4, book_status = $5, book_attrs = $6 WHERE id = $1`

    // Send query to database.
    _, err := q.Exec(query, id, b.UpdatedAt, b.Title, b.Author, b.BookStatus, b.BookAttrs)
    if err != nil {
        // Return only error.
        return err
    }

    // This query returns nothing.
    return nil
}

// DeleteBook method for delete book by given ID.
func (q *BookQueries) DeleteBook(id uuid.UUID) error {
    // Define query string.
    query := `DELETE FROM books WHERE id = $1`

    // Send query to database.
    _, err := q.Exec(query, id)
    if err != nil {
        // Return only error.
        return err
    }

    // This query returns nothing.
    return nil
}
*/
