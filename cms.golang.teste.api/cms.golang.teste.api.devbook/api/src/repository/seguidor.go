package repository

import (
	"database/sql"
)

type Seguidor struct {
	db *sql.DB
}

func NovoSeguidor(db *sql.DB) *Seguidor {
	return &Seguidor{
		db: db,
	}
}

func (repo Seguidor) Criar(usuarioID uint64, seguidorID uint64) error {

	stmt, err := repo.db.Prepare("INSERT IGNORE INTO TB_DEVBOOK_SEGUIDORES(USUARIO_ID, SEGUIDOR_ID) VALUES(?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(usuarioID, seguidorID)
	if err != nil {
		return err
	}

	return nil
}

func (repo Seguidor) Deletar(usuarioID uint64, seguidorID uint64) error {

	stmt, err := repo.db.Prepare("DELETE FROM TB_DEVBOOK_SEGUIDORES WHERE USUARIO_ID = ? AND SEGUIDOR_ID = ? ")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(usuarioID, seguidorID)
	if err != nil {
		return err
	}

	return nil
}

// DROP TABLE TB_DEVBOOK_SEGUIDORES;
// CREATE TABLE TB_DEVBOOK_SEGUIDORES (
//   USUARIO_ID  int NOT NULL,
//   SEGUIDOR_ID int NOT NULL,
//    PRIMARY KEY (USUARIO_ID, SEGUIDOR_ID),
//   FOREIGN KEY (USUARIO_ID)  REFERENCES  TB_DEVBOOK_USUARIO (ID) ON DELETE CASCADE,
//   FOREIGN KEY (SEGUIDOR_ID) REFERENCES  TB_DEVBOOK_USUARIO (ID) ON DELETE CASCADE
// ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
