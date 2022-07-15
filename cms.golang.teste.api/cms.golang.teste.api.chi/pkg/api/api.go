package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ChrisMarSilva/cms.golang.teste.api.chi/pkg/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
)

func NewAPI(pgdb *pg.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/lista", getLista)

	r.Route("/homes", func(r chi.Router) {
		r.Post("/", createHome)
		r.Get("/", getHome)
		r.Get("/{homeID}", getHomeById)
		r.Put("/{homeID}", updateHomeById)
		r.Delete("/{homeID}", deleteHomeById)
	})

	return r
}

func getLista(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("lista"))
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------

type CreateHomeRequest struct {
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	AgentID     int64  `json:"agent_id"`
}

type CreateHomeResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error"`
	Home    *db.Home `json:"home"`
}

func createHome(w http.ResponseWriter, r *http.Request) {
	// parse in the request body
	req := &CreateHomeRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &CreateHomeResponse{Success: false, Error: err.Error(), Home: nil}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get the database somehow
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &CreateHomeResponse{Success: false, Error: "could not get database from context", Home: nil}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// insert our home
	h := &db.Home{Price: req.Price, Description: req.Description, Address: req.Address, AgentID: req.AgentID}
	home, err := db.CreateHome(pgdb, h)
	if err != nil {
		res := &CreateHomeResponse{Success: false, Error: err.Error(), Home: nil}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//  return a response
	res := &CreateHomeResponse{Success: true, Error: "", Home: home}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------

type GetHomeByResponse struct {
	Homes []db.Home `json:"homes"`
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all home"))
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------

func getHomeById(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	w.Write([]byte(fmt.Sprintf("get home by id: %s", homeID)))
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------

func updateHomeById(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	w.Write([]byte(fmt.Sprintf("update home by id: %s", homeID)))
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------

type DeleteHomeByResponse struct {
	Success string `json:"success"`
}

func deleteHomeById(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	w.Write([]byte(fmt.Sprintf("delete home by id: %s", homeID)))
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------
