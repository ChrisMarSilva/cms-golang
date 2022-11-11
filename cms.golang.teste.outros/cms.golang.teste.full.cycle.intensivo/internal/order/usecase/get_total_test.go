package usecase

import (
	"database/sql"
	// "testing"

	"github.com/chrismarsilva/cms.golang.teste.intensivo/internal/order/entity"
	"github.com/chrismarsilva/cms.golang.teste.intensivo/internal/order/infra/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

// go test ./...
// go test
// go test -v
// go test -run TestGetTotalOk -v

type GetTotalUseCaseTestSuite struct {
	suite.Suite
	Db              *sql.DB
	OrderRepository entity.OrderRepositoryInterface
}

// vai rodar antes de inicar os testes
func (suite *GetTotalUseCaseTestSuite) SetupSuite() {
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
func (suite *GetTotalUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

// func TestSuite(t *testing.T) {
// 	suite.Run(t, new(GetTotalUseCaseTestSuite))
// }

// func (suite *GetTotalUseCaseTestSuite) TestGetTotalOk() {
// 	getTotalUseCase := NewGetTotalUseCase(suite.OrderRepository)
// 	output, err := getTotalUseCase.Execute()
// 	suite.NoError(err)
// 	suite.Equal(output.Total, 0)
// }
