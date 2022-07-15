package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
		r.Get("/{homeID}", getHomeByID)
		r.Put("/{homeID}", updateHomeByID)
		r.Delete("/{homeID}", deleteHomeByID)
	})

	return r
}

func getLista(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("lista"))
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------

type HomeResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error"`
	Home    *db.Home `json:"home"`
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------

type CreateHomeRequest struct {
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	AgentID     int64  `json:"agent_id"`
}

func createHome(w http.ResponseWriter, r *http.Request) {
	// parse in the request body
	req := &CreateHomeRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &HomeResponse{Success: false, Error: err.Error(), Home: nil}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{Success: false, Error: "could not get database from context", Home: nil}
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
		res := &HomeResponse{Success: false, Error: err.Error(), Home: nil}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//  return a response
	res := &HomeResponse{Success: true, Error: "", Home: home}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------

type HomesResponse struct {
	Success bool       `json:"success"`
	Error   string     `json:"error"`
	Homes   []*db.Home `json:"homes"`
}

func getHome(w http.ResponseWriter, r *http.Request) {

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomesResponse{Success: false, Error: "could not get database from context", Homes: nil}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	homes, err := db.GetHomes(pgdb)
	if err != nil {
		res := &HomesResponse{Success: false, Error: err.Error(), Homes: nil}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//  return a response
	res := &HomesResponse{Success: true, Error: "", Homes: homes}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------

func getHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{Success: false, Error: "could not get database from context", Home: nil}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// query for the home
	home, err := db.GetHome(pgdb, homeID)
	if err != nil {
		res := &HomeResponse{Success: false, Error: err.Error(), Home: nil}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//  return a response
	res := &HomeResponse{Success: true, Error: "", Home: home}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------

type UpdateHomebyIDRequest struct {
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	AgentID     int64  `json:"agent_id"`
}

func updateHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	// parse in the request body
	req := &UpdateHomebyIDRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &HomeResponse{Success: false, Error: err.Error(), Home: nil}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{Success: false, Error: "could not get database from context", Home: nil}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ihomeID, err := strconv.ParseInt(homeID, 10, 64)
	if err != nil {
		res := &HomeResponse{Success: false, Error: err.Error(), Home: nil}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// update the home
	h := &db.Home{ID: ihomeID, Price: req.Price, Description: req.Description, Address: req.Address, AgentID: req.AgentID}
	home, err := db.UpdateHome(pgdb, h)
	if err != nil {
		res := &HomeResponse{Success: false, Error: err.Error(), Home: nil}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//  return a response
	res := &HomeResponse{Success: true, Error: "", Home: home}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------

func deleteHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{Success: false, Error: "could not get database from context", Home: nil}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ihomeID, err := strconv.ParseInt(homeID, 10, 64)
	if err != nil {
		res := &HomeResponse{Success: false, Error: err.Error(), Home: nil}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// delete the home
	err = db.DeleteHome(pgdb, ihomeID)
	if err != nil {
		res := &HomeResponse{Success: false, Error: err.Error(), Home: nil}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("delete home by id: %s", homeID)))
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------
