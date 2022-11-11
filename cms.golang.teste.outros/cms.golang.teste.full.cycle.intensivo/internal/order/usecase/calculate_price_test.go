package usecase

import (
	"database/sql"
	"testing"

	"github.com/chrismarsilva/cms.golang.teste.intensivo/internal/order/entity"
	"github.com/chrismarsilva/cms.golang.teste.intensivo/internal/order/infra/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

// go test ./...
// go test
// go test -v
// go test -run TestCalculateFinalPrice -v

type CalculateFinalPriceUseCaseTestSuite struct {
	suite.Suite
	Db              *sql.DB
	OrderRepository entity.OrderRepositoryInterface
}

// vai rodar antes de inicar os testes
func (suite *CalculateFinalPriceUseCaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS orders (id integer not null primary key, price text, tax text, final_price text)")
	// suite.NoError(err)

	_, err = db.Exec("CREATE TABLE orders (id varchar(255) not null, price float not null, tax float not null, final_price float not null, primary key(id))")
	suite.NoError(err)

	_, err = db.Exec("delete from orders")
	suite.NoError(err)

	suite.Db = db
	suite.OrderRepository = database.NewOrderRepository(db)
}

// vai rodar no final do teste
func (suite *CalculateFinalPriceUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculateFinalPriceUseCaseTestSuite))
}

func (suite *CalculateFinalPriceUseCaseTestSuite) TestCalculateFinalPrice() {
	order, err := entity.NewOrder("123", 100.0, 10.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())

	input := OrderInputDTO{
		ID:    order.ID,
		Price: order.Price,
		Tax:   order.Tax,
	}

	calculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	output, err := calculateFinalPriceUseCase.Execute(input)
	suite.NoError(err)
	suite.Equal(output.ID, order.ID)
	suite.Equal(output.Price, order.Price)
	suite.Equal(output.Tax, order.Tax)
	suite.Equal(output.FinalPrice, order.FinalPrice)

	var calculateFinalPriceResult entity.Order
	sql := "SELECT id, price, tax, final_price FROM orders WHERE id = ?"
	err = suite.Db.QueryRow(sql, order.ID).Scan(&calculateFinalPriceResult.ID, &calculateFinalPriceResult.Price, &calculateFinalPriceResult.Tax, &calculateFinalPriceResult.FinalPrice)
	suite.NoError(err)
	suite.Equal(calculateFinalPriceResult.ID, order.ID)
	suite.Equal(calculateFinalPriceResult.Price, order.Price)
	suite.Equal(calculateFinalPriceResult.Tax, order.Tax)
	suite.Equal(calculateFinalPriceResult.FinalPrice, order.FinalPrice)
}
