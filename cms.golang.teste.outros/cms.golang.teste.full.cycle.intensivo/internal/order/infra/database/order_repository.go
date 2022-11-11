package database

import (
	"database/sql"
	"github.com/chrismarsilva/cms.golang.teste.intensivo/internal/order/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	sql := "INSERT INTO orders (id, price, tax, final_price) values (?, ?, ?, ?) "

	stmt, err := r.Db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int = 0

	err := r.Db.QueryRow("Select Count(*) From orders").Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
