package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type UserRepo interface {
	CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error)
	GetAUser(ctx context.Context, tx *sql.Tx, id int) (models.User, error)
	GetAllUser(ctx context.Context, tx *sql.Tx) ([]models.User, error)
	UpdateAUsersName(ctx context.Context, tx *sql.Tx, id int, firstName, lastName string)(error)
	DeleteUserByID(ctx context.Context, tx *sql.Tx, id int) error
}

type DBRepo interface {
	Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error 
}



type Repository interface {
	Create(ctx context.Context, trip *Trip) error
	Delete(ctx context.Context, id int) error
	FindByFilter(ctx context.Context, trip *Filter) ([]Trip, error)
	FindByTripID(ctx context.Context, tripID int) (*Trip, error)
	GetSoldTicketNumber(ctx context.Context, tripID int) (int, error)
	UpdateAvailableSeat(ctx context.Context, tripID int, ticketNum int) error

	Save(user *models.User) error
	Update(user *models.User) error
	GetById(id string) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (users []*models.User, err error)
	Delete(id string) error
}

type UserRepository struct { // defaultRepository
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo UserRepository) GetByEmail(ctx context.Context, email string) (*UserModel, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "SELECT * FROM users WHERE email = ?"
	row := repo.db.QueryRowContext(timeoutCtx, query, email)

	user := &UserModel{}
	err := row.Scan(&user.ID, &user.Nome, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	log.Error("Erro no ErrNoRows:", err.Error())
		// 	return nil, fmt.Errorf("No user found with Email '%s'", email)
		// }
		log.Error("Erro no Scan:", err.Error())
		return nil, err
	}

	//log.Info("ID:", user.ID, "Nome:", user.Nome, "Email:", user.Email, "Password:", user.Password, "IsActive:", user.IsActive, "Created_at:", user.Created_at)
	return user, nil
}

func (repo UserRepository) Create(ctx context.Context, user *UserModel) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "INSERT INTO users (id, nome, email, password, is_active, created_at) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := repo.db.Prepare(query)
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

func (repo UserRepository) Update(ctx context.Context, user *UserModel) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "UPDATE users SET nome = ?, password = ?, is_active = ? WHERE id = ?"

	stmt, err := repo.db.Prepare(query)
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

func (repo UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "DELETE FROM users WHERE id = ?"
	stmt, err := repo.db.Prepare(query)
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




package repository

import (
    "context"
    "database/sql"
)

type dbRepo struct {
    DB *sql.DB
}

func NewDBRepo(conn *sql.DB) DBRepo {
    return &dbRepo{
        DB: conn,
    }
}

func (m *dbRepo) Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error {
    tx, err := m.DB.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    defer func() error{
        if err != nil {
            tx.Rollback()
            return err
        }

        if err := tx.Commit(); err != nil {
            return err
        }

        return nil
    }()

    if err := operation(ctx, tx); err != nil {
        return err
    }

    return nil
}


type UserRepo interface {
	-   CreateAUser(user models.User) (int, error)
	-   GetAUser(id int) (models.User, error)
	-   GetAllUser() ([]models.User, error)
	-   UpdateAUsersName(id int, firstName, lastName string)(error)
	-   DeleteUserByID(id int) error
	+   CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error)
	+   GetAUser(ctx context.Context, tx *sql.Tx, id int) (models.User, error)
	+   GetAllUser(ctx context.Context, tx *sql.Tx) ([]models.User, error)
	+   UpdateAUsersName(ctx context.Context, tx *sql.Tx, id int, firstName, lastName string)(error)
	+   DeleteUserByID(ctx context.Context, tx *sql.Tx, id int) error
	}
	
	type DBRepo interface {
		Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error 
	}

	
func (m *user) CreateAUser(user models.User) (int, error){
    -   ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    +func (m *user) CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error){
    +   ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
        defer cancel()
    
        var newId int
    
        query := `
                INSERT into users 
                    (first_name, last_name, email, password, created_at, updated_at)
                values 
                    ($1, $2, $3, $4, $5, $6)
                returning id`
    
    -   err := m.DB.QueryRowContext(ctx, query, 
    +   var err error;
    +   if tx != nil {
    +       err = tx.QueryRowContext(ctx, query, 
    +           user.FirstName, 
    +           user.LastName, 
    +           user.Email, 
    +           user.Password,
    +           time.Now(),
    +           time.Now(),
    +       ).Scan(&newId)
    +   }else{
            err = m.DB.QueryRowContext(ctx, query, 
                user.FirstName, 
                user.LastName, 
                user.Email, 
                user.Password,
                time.Now(),
                time.Now(),
            ).Scan(&newId)
    +   }


    package repository

import (
    "context"
    "database/sql"
    "time"

    "github.com/orololuwa/crispy-octo-guacamole/models"
)

type user struct {
    DB *sql.DB
}

func NewUserRepo(conn *sql.DB) UserRepo {
    return &user{
        DB: conn,
    }
}

func (m *user) CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error){
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    var newId int

    query := `
            INSERT into users 
                (first_name, last_name, email, password, created_at, updated_at)
            values 
                ($1, $2, $3, $4, $5, $6)
            returning id`

    var err error;
    if tx != nil {
        err = tx.QueryRowContext(ctx, query, 
            user.FirstName, 
            user.LastName, 
            user.Email, 
            user.Password,
            time.Now(),
            time.Now(),
        ).Scan(&newId)
    }else{
        err = m.DB.QueryRowContext(ctx, query, 
            user.FirstName, 
            user.LastName, 
            user.Email, 
            user.Password,
            time.Now(),
            time.Now(),
        ).Scan(&newId)
    }

    if err != nil {
        return 0, err
    }

    return newId, nil
}

func (m *user) GetAUser(ctx context.Context, tx *sql.Tx, id int) (models.User, error){
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    var user models.User

    query := `
            SELECT (id, first_name, last_name, email, password, created_at, updated_at)
            from users
            WHERE
            id=$1
    `

    var err error
    if tx != nil {
        err = tx.QueryRowContext(ctx, query, id).Scan(
            &user.ID,
            &user.FirstName,
            &user.LastName,
            &user.Email,
            &user.Password,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
    }else{
        err = m.DB.QueryRowContext(ctx, query, id).Scan(
            &user.ID,
            &user.FirstName,
            &user.LastName,
            &user.Email,
            &user.Password,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
    }

    if err != nil {
        return user, err
    }

    return user, nil
}

func (m *user) GetAllUser(ctx context.Context, tx *sql.Tx) ([]models.User, error){
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    var users = make([]models.User, 0)

    query := `
        SELECT (id, first_name, last_name, email, password, created_at, updated_at)
        from users
    `

    var rows *sql.Rows
    var err error

    if tx != nil {
        rows, err = tx.QueryContext(ctx, query)
    }else{
        rows, err = m.DB.QueryContext(ctx, query)
    }
    if err != nil {
        return users, err
    }

    for rows.Next(){
        var user models.User
        err := rows.Scan(
            &user.ID,
            &user.FirstName,
            &user.LastName,
            &user.Email,
            &user.Password,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
        if err != nil {
            return users, err
        }
        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        return users, err
    }

    return users, nil
}

func (m *user) UpdateAUsersName(ctx context.Context, tx *sql.Tx, id int, firstName, lastName string)(error){
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    query := `
        UPDATE 
            users set (first_name, last_name) = ($1, $2)
        WHERE
            id = $3
    `

    var err error
    if tx != nil{
        _, err = tx.ExecContext(ctx, query, firstName, lastName, id)
    }else{
        _, err = m.DB.ExecContext(ctx, query, firstName, lastName, id)
    }

    if err != nil{
        return  err
    }

    return nil
}

func (m *user) DeleteUserByID(ctx context.Context, tx *sql.Tx, id int) error {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    query := "DELETE FROM users WHERE id = $1"

    var err error 

    if tx != nil {
        _, err = tx.ExecContext(ctx, query, id)
    }else{
        _, err = m.DB.ExecContext(ctx, query, id)
    }

    if err != nil {
        return err
    }

    return nil
}



func (m *user) CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error){
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var newId int

	query := `
			INSERT into users 
				(first_name, last_name, email, password, created_at, updated_at)
			values 
				($1, $2, $3, $4, $5, $6)
			returning id`

	var err error;
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, 
			user.FirstName, 
			user.LastName, 
			user.Email, 
			user.Password,
			time.Now(),
			time.Now(),
		).Scan(&newId)
	}else{
		err = m.DB.QueryRowContext(ctx, query, 
			user.FirstName, 
			user.LastName, 
			user.Email, 
			user.Password,
			time.Now(),
			time.Now(),
		).Scan(&newId)
	}

	if err != nil {
		return 0, err
	}

	return newId, nil
}


func (m *user) GetAUser(ctx context.Context, tx *sql.Tx, id int) (models.User, error){
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var user models.User

	query := `
			SELECT (id, first_name, last_name, email, password, created_at, updated_at)
			from users
			WHERE
			id=$1
	`

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, id).Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	}else{
		err = m.DB.QueryRowContext(ctx, query, id).Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	}

	if err != nil {
		return user, err
	}

	return user, nil
}

func (m *user) GetAllUser(ctx context.Context, tx *sql.Tx) ([]models.User, error){
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var users = make([]models.User, 0)

	query := `
		SELECT (id, first_name, last_name, email, password, created_at, updated_at)
		from users
	`

	var rows *sql.Rows
	var err error

	if tx != nil {
		rows, err = tx.QueryContext(ctx, query)
	}else{
		rows, err = m.DB.QueryContext(ctx, query)
	}
	if err != nil {
		return users, err
	}

	for rows.Next(){
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (m *user) UpdateAUsersName(ctx context.Context, tx *sql.Tx, id int, firstName, lastName string)(error){
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
		UPDATE 
			users set (first_name, last_name) = ($1, $2)
		WHERE
			id = $3
	`

	var err error
	if tx != nil{
		_, err = tx.ExecContext(ctx, query, firstName, lastName, id)
	}else{
		_, err = m.DB.ExecContext(ctx, query, firstName, lastName, id)
	}

	if err != nil{
		return  err
	}

	return nil
}

func (m *user) DeleteUserByID(ctx context.Context, tx *sql.Tx, id int) error {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    query := "DELETE FROM users WHERE id = $1"

	var err error 

	if tx != nil {
		_, err = tx.ExecContext(ctx, query, id)
	}else{
		_, err = m.DB.ExecContext(ctx, query, id)
	}

    if err != nil {
        return err
    }

    return nil
}


func (m *dbRepo) Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error {
    tx, err := m.DB.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
	
    defer func() error{
        if err != nil {
            tx.Rollback()
            return err
        }

        if err := tx.Commit(); err != nil {
            return err
        }

        return nil
    }()

    if err := operation(ctx, tx); err != nil {
        return err
    }

    return nil
}

func (m *user) UpdateAUsersName(id int, firstName, lastName string)(error){
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    query := `
        UPDATE 
            users set (first_name, last_name) = ($1, $2)
        WHERE
            id = $3
    `

    _, err := m.DB.ExecContext(ctx, query, firstName, lastName, id)
    if err != nil{
        return  err
    }

    return nil
}




+type DBRepo interface {
    +   Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error 
    +}
    

    
package repository

import (
    "context"
    "database/sql"
)

type dbRepo struct {
    DB *sql.DB
}

func NewDBRepo(conn *sql.DB) DBRepo {
    return &dbRepo{
        DB: conn,
    }
}

func (m *dbRepo) Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error {
    tx, err := m.DB.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    defer func() error{
        if err != nil {
            tx.Rollback()
            return err
        }

        if err := tx.Commit(); err != nil {
            return err
        }

        return nil
    }()

    if err := operation(ctx, tx); err != nil {
        return err
    }

    return nil
}



type UserRepo interface {
    CreateAUser(user models.User) (int, error)
    GetAUser(id int) (models.User, error)
    GetAllUser() ([]models.User, error)
    UpdateAUsersName(id int, firstName, lastName string)(error)
    DeleteUserByID(id int) error
 +   CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error)
 +   GetAUser(ctx context.Context, tx *sql.Tx, id int) (models.User, error)
 +   GetAllUser(ctx context.Context, tx *sql.Tx) ([]models.User, error)
 +   UpdateAUsersName(ctx context.Context, tx *sql.Tx, id int, firstName, lastName string)(error)
 +   DeleteUserByID(ctx context.Context, tx *sql.Tx, id int) error
 }

 

func (m *user) CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error){
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    var newId int

    query := `
            INSERT into users 
                (first_name, last_name, email, password, created_at, updated_at)
            values 
                ($1, $2, $3, $4, $5, $6)
            returning id`

    var err error;
    if tx != nil {
        err = tx.QueryRowContext(ctx, query, 
            user.FirstName, 
            user.LastName, 
            user.Email, 
            user.Password,
            time.Now(),
            time.Now(),
        ).Scan(&newId)
    }else{
        err = m.DB.QueryRowContext(ctx, query, 
            user.FirstName, 
            user.LastName, 
            user.Email, 
            user.Password,
            time.Now(),
            time.Now(),
        ).Scan(&newId)
    }

    if err != nil {
        return 0, err
    }

    return newId, nil
}

func (m *user) GetAUser(ctx context.Context, tx *sql.Tx, id int) (models.User, error){
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    var user models.User

    query := `
            SELECT (id, first_name, last_name, email, password, created_at, updated_at)
            from users
            WHERE
            id=$1
    `

    var err error
    if tx != nil {
        err = tx.QueryRowContext(ctx, query, id).Scan(
            &user.ID,
            &user.FirstName,
            &user.LastName,
            &user.Email,
            &user.Password,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
    }else{
        err = m.DB.QueryRowContext(ctx, query, id).Scan(
            &user.ID,
            &user.FirstName,
            &user.LastName,
            &user.Email,
            &user.Password,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
    }

    if err != nil {
        return user, err
    }

    return user, nil
}

func (m *user) GetAllUser(ctx context.Context, tx *sql.Tx) ([]models.User, error){
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    var users = make([]models.User, 0)

    query := `
        SELECT (id, first_name, last_name, email, password, created_at, updated_at)
        from users
    `

    var rows *sql.Rows
    var err error

    if tx != nil {
        rows, err = tx.QueryContext(ctx, query)
    }else{
        rows, err = m.DB.QueryContext(ctx, query)
    }
    if err != nil {
        return users, err
    }

    for rows.Next(){
        var user models.User
        err := rows.Scan(
            &user.ID,
            &user.FirstName,
            &user.LastName,
            &user.Email,
            &user.Password,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
        if err != nil {
            return users, err
        }
        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        return users, err
    }

    return users, nil
}

func (m *user) UpdateAUsersName(ctx context.Context, tx *sql.Tx, id int, firstName, lastName string)(error){
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    query := `
        UPDATE 
            users set (first_name, last_name) = ($1, $2)
        WHERE
            id = $3
    `

    var err error
    if tx != nil{
        _, err = tx.ExecContext(ctx, query, firstName, lastName, id)
    }else{
        _, err = m.DB.ExecContext(ctx, query, firstName, lastName, id)
    }

    if err != nil{
        return  err
    }

    return nil
}

func (m *user) DeleteUserByID(ctx context.Context, tx *sql.Tx, id int) error {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    query := "DELETE FROM users WHERE id = $1"

    var err error 

    if tx != nil {
        _, err = tx.ExecContext(ctx, query, id)
    }else{
        _, err = m.DB.ExecContext(ctx, query, id)
    }

    if err != nil {
        return err
    }

    return nil
}


