package database

import (
	"database/sql"
	"testing"

	"github.com/chrismarsilva/cms.golang.teste.intensivo/internal/order/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

// go test ./...
// go test
// go test -v
// go test -run TestGivenAnOrder_WhenSave_ThenShouldSaveOrder -v

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

// vai rodar antes de inicar os testes
func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS orders (id integer not null primary key, price text, tax text, final_price text)")
	// suite.NoError(err)

	_, err = db.Exec("CREATE TABLE orders (id varchar(255) not null, price float not null, tax float not null, final_price float not null, primary key(id))")
	suite.NoError(err)

	_, err = db.Exec("delete from orders")
	suite.NoError(err)

	suite.Db = db
}

// vai rodar no final do teste
func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 100.0, 10.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())

	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	sql := "SELECT id, price, tax, final_price FROM orders WHERE id = ?"
	err = suite.Db.QueryRow(sql, order.ID).Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)
	suite.NoError(err)
	suite.Equal(orderResult.ID, order.ID)
	suite.Equal(orderResult.Price, order.Price)
	suite.Equal(orderResult.Tax, order.Tax)
	suite.Equal(orderResult.FinalPrice, order.FinalPrice)
}
