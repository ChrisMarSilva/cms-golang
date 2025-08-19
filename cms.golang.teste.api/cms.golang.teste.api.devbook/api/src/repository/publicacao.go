package repository

import (
	"database/sql"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/models"
)

type Publicacao struct {
	db *sql.DB
}

func NovaPublicacao(db *sql.DB) *Publicacao {
	return &Publicacao{
		db: db,
	}
}

func (repo Publicacao) Criar(publicacao models.Publicacao) (uint64, error) {

	stmt, err := repo.db.Prepare("INSERT IGNORE INTO TB_DEVBOOK_PUBLICACAO(TITULO, CONTEUDO, AUTOR_ID) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (repo Publicacao) Buscar(publicacaoID uint64) ([]models.Publicacao, error) {
	
	rows, err := repo.db.Query(`
	SELECT P.ID, P.TITULO, P.CONTEUDO, P.AUTOR_ID, U.NICK AS AUTOR_NICK, P.CURTIDAS, P.CRIADAEM 
	FROM TB_DEVBOOK_PUBLICACAO P 
		JOIN TB_DEVBOOK_USUARIO U ON ( U.ID = P.AUTOR_ID ) 
		JOIN TB_DEVBOOK_SEGUIDORES S ON ( S.USUARIO_ID = P.AUTOR_ID ) 
	WHERE P.AUTOR_ID = ? OR S.AUTOR_IDSEGUIDOR_ID = ?  `,
	publicacaoID, publicacaoID)
	
	// SELECT P.ID, P.TITULO, P.CONTEUDO, P.AUTOR_ID, U.NICK AS AUTOR_NICK, P.CURTIDAS, P.CRIADAEM 
	// FROM TB_DEVBOOK_PUBLICACAO P 
	// 	JOIN TB_DEVBOOK_USUARIO U ON ( U.ID = P.AUTOR_ID ) 
	// WHERE P.AUTOR_ID IN ( 
	// 	SELECT DISTINCT S.USUARIO_ID FROM TB_DEVBOOK_SEGUIDORES S WHERE S.USUARIO_ID = ?
	// 	UNION ALL
	// 	SELECT DISTINCT S.SEGUIDOR_ID FROM TB_DEVBOOK_SEGUIDORES S WHERE S.USUARIO_ID = ?


	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	publicacoes := []models.Publicacao{}

	for rows.Next() {
		var publicacao models.Publicacao
		err = rows.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.AutorNick, &publicacao.Curtidas, &publicacao.CriadaEm)
		if err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repo Publicacao) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {

	rows := repo.db.QueryRow(`
	SELECT P.ID, P.TITULO, P.CONTEUDO, P.AUTOR_ID, U.NICK AS AUTOR_NICK, P.CURTIDAS, P.CRIADAEM 
	FROM TB_DEVBOOK_PUBLICACAO P 
		JOIN TB_DEVBOOK_USUARIO U ON ( U.ID = P.AUTOR_ID ) 
	WHERE P.ID = ? `, publicacaoID)

	var publicacao models.Publicacao

	err := rows.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.AutorNick, &publicacao.Curtidas, &publicacao.CriadaEm)
	if err != nil {
		return publicacao, err
	}

	return publicacao, nil
}

func (repo Publicacao) BuscarPorUsuarioID(usuarioID uint64) ([]models.Publicacao, error) {
	
	rows, err := repo.db.Query(`SELECT P.ID, P.TITULO, P.CONTEUDO, P.AUTOR_ID, U.NICK AS AUTOR_NICK, P.CURTIDAS, P.CRIADAEM FROM TB_DEVBOOK_PUBLICACAO P JOIN TB_DEVBOOK_USUARIO U ON ( U.ID = P.AUTOR_ID ) WHERE P.AUTOR_ID = ?`, usuarioID)

	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	publicacoes := []models.Publicacao{}

	for rows.Next() {
		var publicacao models.Publicacao
		err = rows.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.AutorNick, &publicacao.Curtidas, &publicacao.CriadaEm)
		if err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repo Publicacao) Atualizar(publicacao models.Publicacao) (int64, error) {

	stmt, err := repo.db.Prepare("UPDATE TB_DEVBOOK_PUBLICACAO SET TITULO = ?, CONTEUDO = ? WHERE ID = ? ")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.ID)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (repo Publicacao) Deletar(publicacaoID uint64) error {

	stmt, err := repo.db.Prepare("DELETE FROM TB_DEVBOOK_PUBLICACAO WHERE ID = ? ")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(publicacaoID)
	if err != nil {
		return err
	}

	return nil
}

func (repo Publicacao) Curtir(publicacaoID uint64) error {

	stmt, err := repo.db.Prepare("UPDATE TB_DEVBOOK_PUBLICACAO SET CURTIDAS = CURTIDAS + 1 WHERE ID = ? ")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(publicacaoID)
	if err != nil {
		return err
	}

	return nil
}

func (repo Publicacao) Descurtir(publicacaoID uint64) error {

	stmt, err := repo.db.Prepare("UPDATE TB_DEVBOOK_PUBLICACAO SET CURTIDAS = CURTIDAS - 1 WHERE ID = ? AND CURTIDAS > 0")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(publicacaoID)
	if err != nil {
		return err
	}

	return nil
}

// DROP TABLE TB_DEVBOOK_PUBLICACAO;
// CREATE TABLE TB_DEVBOOK_PUBLICACAO (
//    ID         int          NOT NULL AUTO_INCREMENT PRIMARY KEY,
//    TITULO     varchar(50)  NOT NULL,
//    CONTEUDO   varchar(300) NOT NULL,
//    AUTOR_ID   int          NOT NULL,
//    CURTIDAS   int          NOT NULL DEFAULT 0,
//    CRIADAEM   TIMESTAMP    NOT NULL DEFAULT current_timestamp,
//    FOREIGN KEY (AUTOR_ID)  REFERENCES  TB_DEVBOOK_USUARIO (ID) ON DELETE CASCADE
// ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
