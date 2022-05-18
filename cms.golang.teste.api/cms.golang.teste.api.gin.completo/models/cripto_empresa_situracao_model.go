package models

import (
	"gorm.io/gorm"
	//"database/sql"
	"cms.golang.tnb.api/entities"
)

type CriptoEmpresaSituacaoModel struct {
}

func (CriptoEmpresaSituacaoModel) GetSituacoes(db *gorm.DB, situacao *[]entities.CriptoEmpresaSituacao) (err error) {
	err = db.Find(situacao).Error

	// result := db.Find(&situacao)
	// result.RowsAffected // returns found records count, equals `len(users)`
	// result.Error
	// errors.Is(result.Error, gorm.ErrRecordNotFound)

	// m := make(map[string]interface{})
	// m["id"] = 10
	// m["name"] = "chetan"
	// db.Where(m).Find(&users)

	// result := map[string]interface{}{}
	// db.Model(&User{}).First(&result)

	// db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)

	if err != nil {
		return err
	}
	return nil
}

func (CriptoEmpresaSituacaoModel) GetSituacao(db *gorm.DB, Situacao *entities.CriptoEmpresaSituacao, codigo string) (err error) {
	//err = db.Where("CODIGO = ?", codigo).First(Situacao).Error
	// err = db.First(Situacao, "CODIGO = ?", codigo).Error
	db.Raw("SELECT CODIGO, DESCRICAO FROM TBCRIPTO_EMPRESA_ST WHERE CODIGO = ?", codigo).Scan(&Situacao)
	if err != nil {
		return err
	}
	return nil
}

func (CriptoEmpresaSituacaoModel) CreateSituacao(db *gorm.DB, Situacao *entities.CriptoEmpresaSituacao) (err error) {
	err = db.Create(Situacao).Error
	if err != nil {
		return err
	}
	return nil
}

func (CriptoEmpresaSituacaoModel) UpdateSituacao(db *gorm.DB, Situacao *entities.CriptoEmpresaSituacao) (err error) {
	db.Save(Situacao)
	return nil
}

func (CriptoEmpresaSituacaoModel) DeleteSituacao(db *gorm.DB, Situacao *entities.CriptoEmpresaSituacao, codigo string) (err error) {
	db.Where("CODIGO = ?", codigo).Delete(Situacao)
	return nil
}

// type CriptoEmpresaSituacaoModel struct {
// 	Db *sql.DB
// }

// func (model CriptoEmpresaSituacaoModel) FindAll() ([]entities.CriptoEmpresaSituacao, error) {

// 	rows, errQuery := model.Db.Query("SELECT CODIGO, DESCRICAO FROM TBCRIPTO_EMPRESA_ST")
// 	if errQuery != nil {
// 		return nil, errQuery
// 	}

// 	defer rows.Close()

// 	itens := []entities.CriptoEmpresaSituacao{}
// 	for rows.Next() {
// 		item := entities.CriptoEmpresaSituacao{}
// 		errScan := rows.Scan(&item.Codigo, &item.Descricao)
// 		if errScan != nil {
// 			return nil, errScan
// 		}
// 		itens = append(itens, item)
// 	}

// 	return itens, nil
// }

// func (model CriptoEmpresaSituacaoModel) Find(codigo string) (entities.CriptoEmpresaSituacao, error) {

// 	item := entities.CriptoEmpresaSituacao{}

// 	rows, errQuery := model.Db.Query("SELECT CODIGO, DESCRICAO FROM TBCRIPTO_EMPRESA_ST WHERE CODIGO = ?", codigo)
// 	if errQuery != nil {
// 		return item, errQuery
// 	}

// 	defer rows.Close()

// 	// row := bd.QueryRow("SELECT id, name, genre, year FROM video_games WHERE id = ?", id)
// 	// err = row.Scan(&videoGame.Id, &videoGame.Name, &videoGame.Genre, &videoGame.Year)

// 	for rows.Next() {
// 		errScan := rows.Scan(&item.Codigo, &item.Descricao)
// 		if errScan != nil {
// 			return item, errQuery
// 		}
// 	}

// 	return item, nil
// }

// // fnc (userModel UsuarioModel) Create(user *entities.Usuario) error {
// //   result, err := userModel.Db.Exec("inser into dddd () values (?,?) ", user.id, user.name)
// //   if err != nil{
// //     return nil, err
// //   }
// //   // result.RowsAffected()
// //   user.id, _ = result.LastInsertId()
// //   return nil
// // }
