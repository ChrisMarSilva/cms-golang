package repository

import (
	"database/sql"
	"fmt"
	//"errors"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/models"
)

type Usuario struct {
	db *sql.DB
}

func NovoUsuario(db *sql.DB) *Usuario {
	return &Usuario{
		db: db,
	}
}

func (repo Usuario) Criar(usuario models.Usuario) (uint64, error) {

	stmt, err := repo.db.Prepare("INSERT INTO TB_DEVBOOK_USUARIO(NOME, NICK, EMAIL, SENHA) VALUES(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (repo Usuario) Buscar(nomeOuNick string) ([]models.Usuario, error) {

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	
	rows, err := repo.db.Query("SELECT ID, NOME, NICK, EMAIL, SENHA, CRIADOEM FROM TB_DEVBOOK_USUARIO WHERE NOME LIKE ? OR NICK LIKE ? ", nomeOuNick, nomeOuNick)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	usuarios := []models.Usuario{}

	for rows.Next() {
		var usuario models.Usuario
		err = rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.Senha, &usuario.CriadoEm)
		if err != nil {
			return nil, err
		}
		usuario.Senha = ""
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repo Usuario) BuscarPorID(usuarioID uint64) (models.Usuario, error) {

	rows := repo.db.QueryRow("SELECT ID, NOME, NICK, EMAIL, SENHA, CRIADOEM FROM TB_DEVBOOK_USUARIO WHERE ID = ? ", usuarioID)
	// if err != nil {
	// 	return usuario, err
	// }
	//defer rows.Close()
	
	// if rows.Next(){

	// }
	
	var usuario models.Usuario

	err := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.Senha, &usuario.CriadoEm)
	if err != nil {
		return usuario, err
	}

	return usuario, nil
}

func (repo Usuario) BuscarPorEmail(usuarioEmail string) (models.Usuario, error) {

	rows := repo.db.QueryRow("SELECT ID, NOME, NICK, EMAIL, SENHA, CRIADOEM FROM TB_DEVBOOK_USUARIO WHERE EMAIL = ? ", usuarioEmail)

	var usuario models.Usuario
	
	err := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.Senha, &usuario.CriadoEm)
	if err != nil {
		return usuario, err
	}

	return usuario, nil
}

func (repo Usuario) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error) {
	
	rows, err := repo.db.Query(`SELECT U.ID, U.NOME, U.NICK, U.EMAIL, U.SENHA, U.CRIADOEM FROM TB_DEVBOOK_USUARIO U JOIN TB_DEVBOOK_SEGUIDORES S ON ( S.SEGUIDOR_ID = U.ID ) WHERE S.USUARIO_ID = ? `, usuarioID)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	usuarios := []models.Usuario{}

	for rows.Next() {
		var usuario models.Usuario
		err = rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.Senha, &usuario.CriadoEm)
		if err != nil {
			return nil, err
		}
		usuario.Senha = ""
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repo Usuario) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error) {
	
	rows, err := repo.db.Query(`SELECT U.ID, U.NOME, U.NICK, U.EMAIL, U.SENHA, U.CRIADOEM FROM TB_DEVBOOK_USUARIO U JOIN TB_DEVBOOK_SEGUIDORES S ON ( S.USUARIO_ID = U.ID ) WHERE S.SEGUIDOR_ID = ? `, usuarioID)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	usuarios := []models.Usuario{}

	for rows.Next() {
		var usuario models.Usuario
		err = rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.Senha, &usuario.CriadoEm)
		if err != nil {
			return nil, err
		}
		usuario.Senha = ""
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repo Usuario) Atualizar(usuario models.Usuario) (int64, error) {

	stmt, err := repo.db.Prepare("UPDATE TB_DEVBOOK_USUARIO SET NOME = ?, NICK = ?, EMAIL = ? WHERE ID = ? ")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.ID)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	// println("usuario.ID", usuario.ID)
	// println("rowsAffected", rowsAffected)

	// if rowsAffected != 1 {
	// 	return 0, errors.New("não encontrado")
	// }

	return rowsAffected, nil
}

func (repo Usuario) AtualizarSenha(usuario models.Usuario) error {

	stmt, err := repo.db.Prepare("UPDATE TB_DEVBOOK_USUARIO SET SENHA = ? WHERE ID = ? ")
	if err != nil {
		return err 
	}

	defer stmt.Close()

	_, err = stmt.Exec(usuario.Senha, usuario.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo Usuario) Deletar(usuarioID uint64) error {

	stmt, err := repo.db.Prepare("DELETE FROM TB_DEVBOOK_USUARIO WHERE ID = ? ")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(usuarioID)
	if err != nil {
		return err
	}

	return nil
}

// DROP TABLE TB_DEVBOOK_USUARIO;
// CREATE TABLE TB_DEVBOOK_USUARIO (
//   ID       int          NOT NULL AUTO_INCREMENT PRIMARY KEY,
//   NOME     varchar(50)  NOT NULL,
//   NICK     varchar(50)  NOT NULL UNIQUE,
//   EMAIL    varchar(50)  NOT NULL UNIQUE,
//   SENHA    varchar(250) NOT NULL,
//   CRIADOEM TIMESTAMP    NOT NULL DEFAULT current_timestamp
// ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
