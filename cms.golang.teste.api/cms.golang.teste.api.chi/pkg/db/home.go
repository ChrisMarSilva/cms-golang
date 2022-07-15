package db

import (
	"github.com/go-pg/pg/v10"
)

type Home struct {
	ID          int64  `json:"id"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	AgentID     int64  `json:"agent_id"`
	Agent       *Agent `pg:"rel:has-one" json:"agent"`
}

func CreateHome(db *pg.DB, req *Home) (*Home, error) {
	_, err := db.Model(req).Insert()
	if err != nil {
		return nil, err
	}

	home := &Home{}
	err = db.Model(home).Relation("Agent").Where("home.id = ?", req.ID).Select()
	if err != nil {
		return nil, err
	}

	return home, nil
}
