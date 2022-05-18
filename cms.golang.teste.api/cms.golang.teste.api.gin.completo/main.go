package main

import (
	"cms.golang.tnb.api/database"
	"cms.golang.tnb.api/server"
	//"gorm-test/controllers"
	// "database/sql"
	// "encoding/json"
	// "fmt"
	// "log"
	// "net/http"
	// "time"
	//"strconv"
	// "encoding/json"
	// "io/ioutil"
	// "github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	// "cms.golang.tnb.api/config"
	// "cms.golang.tnb.api/models"
	// "github.com/gorilla/mux"
)



/*
redirect.CreatedAt = time.Now().UTC().Unix()
	db := models.SetupModels()
	r.Use(func(c *gin.Context){
		c.Set("db", db)
		c.Next()
	})

	//Tampil data mahasiswa
func MahasiswaTampil (c *gin.Context) {
	db:= c.MustGet("db").(*gorm.DB)

	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{"data":mhs})
}

*/









func main() {
	database.StartDB()
	s := server.NewServer()
	s.Run()
}

// var db *sql.DB
// var errDb error

// func main() {

// 	db, errDb = config.GetMySQLBD()
// 	if errDb != nil {
// 		// panic(errDb.Error())
// 		log.Printf("Error with database" + errDb.Error())
// 		return
// 	}
// 	log.Println("Connection Established")

// 	defer db.Close()

// 	handleRequests()

// }

// //type App struct { Router   *mux.Router Database *sql.DB}

// func handleRequests() {
// 	port := ":8000"
// 	log.Println("Rest API v2.0 - Mux Routers")
// 	log.Println("Starting development server at http://127.0.0.1" + port + "/")
// 	log.Println("Quit the server with CONTROL-C.")

// 	router := mux.NewRouter().StrictSlash(true)
// 	router.Use(middleware)
// 	router.HandleFunc("/", returnHome)
// 	setupRouterCORS(router)
// 	setupRouterSituacoes(router)

// 	// app := &app.App{Router:   mux.NewRouter().StrictSlash(true),Database: database,}
// 	// app.SetupRouter()

// 	// log.Fatal(http.ListenAndServe(port, router)) // app.Router

// 	server := &http.Server{Handler: router, Addr: port, WriteTimeout: 15 * time.Second, ReadTimeout: 15 * time.Second}
// 	log.Fatal(server.ListenAndServe())
// }

// func setupRouterCORS(router *mux.Router) {
// 	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Content-Type", "application/json") // "text/html; charset=utf-8" // "text/html; charset=ascii"
// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 		// (*w).Header().Set("Access-Control-Max-Age", "3600")
// 		// if e := fn(w, r); e != nil { // e is *appError, not os.Error.
// 		//     http.Error(w, e.Message, e.Code)
// 		// }
// 	}).Methods(http.MethodOptions)
// }

// func setupRouterSituacoes(router *mux.Router) {
// 	// router.HandleFunc("/situacoes", getSituacoes).Methods("GET")
// 	router.Methods("GET").Path("/situacoes").HandlerFunc(getSituacoes)
// 	router.Methods("POST").Path("/situacoes").HandlerFunc(createSituacao)
// 	router.Methods("GET").Path("/situacoes/{codigo}").HandlerFunc(getSituacao)
// 	router.Methods("PUT").Path("/situacoes/{codigo}").HandlerFunc(updateSituacao)
// 	router.Methods("DELETE").Path("/situacoes/{codigo}").HandlerFunc(deleteSituacao)
// }

// // func setupResponse(w *http.ResponseWriter, r *http.Request) {
// // 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// // 	(*w).Header().Set("Content-Type", "application/json") // "text/html; charset=utf-8" // "text/html; charset=ascii"
// // 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// // 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// // }

// var middleware = func(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		requestPath := r.URL.Path
// 		log.Println(requestPath)
// 		next.ServeHTTP(w, r) //proceed in the middleware chain!
// 	})
// }

// // func stringToInt64(s string) (int64, error) {
// // 	numero, err := strconv.ParseInt(s, 0, 64)
// // 	if err != nil {
// // 		return 0, err
// // 	}
// // 	return numero, err
// // }

// // // id, err := stringToInt64(idAsString)

// // func (app *App) SetupRouter() {
// // 	app.Router.Methods("POST").Path("/endpoint").HandlerFunc(app.postFunction)
// // }

// // func (app *App) postFunction(w http.ResponseWriter, r *http.Request) {
// // 	_, err := app.Database.Exec("INSERT INTO `test` (name) VALUES ('myname')")
// // 	if err != nil {
// // 		log.Fatal("Database INSERT failed")
// // 	}
// // 	log.Println("You called a thing!")
// // 	w.WriteHeader(http.StatusOK)
// // }

// // func (app *App) getFunction(w http.ResponseWriter, r *http.Request) {
// // 	vars := mux.Vars(r)
// // 	id, ok := vars["id"]
// // 	if !ok {
// // 		log.Fatal("No ID in the path")
// // 	}
// // 	dbdata := &DbData{}
// // 	err := app.Database.QueryRow("SELECT id, date, name FROM `test` WHERE id = ?", id).Scan(&dbdata.ID, &dbdata.Date, &dbdata.Name)
// // 	if err != nil {
// // 		log.Fatal("Database SELECT failed")
// // 	}
// // 	log.Println("You fetched a thing!")
// // 	w.WriteHeader(http.StatusOK)
// // 	if err := json.NewEncoder(w).Encode(dbdata); err != nil {
// // 		panic(err)
// // 	}
// // }

// func returnHome(w http.ResponseWriter, r *http.Request) {
// 	//fmt.Println("Endpoint Hit: homePage")
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// }

// func respondWithSuccess(data interface{}, w http.ResponseWriter) {
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(data)
// }

// func respondWithError(err error, w http.ResponseWriter) {
// 	w.WriteHeader(http.StatusInternalServerError)
// 	json.NewEncoder(w).Encode(err.Error())
// }

// func getSituacoes(w http.ResponseWriter, r *http.Request) {
// 	//setupResponse(&w, r)
// 	situacoes, err := models.CriptoEmpresaSituacaoModel{Db: db}.FindAll()
// 	if err != nil {
// 		// panic(err.Error())
// 		respondWithError(err, w)
// 		return
// 	}
// 	// bookings := []Booking{}
// 	// db.Find(&bookings)

// 	// var orders []Order
// 	// db.Preload("Items").Find(&orders)

// 	//w.WriteHeader(http.StatusOK)
// 	//json.NewEncoder(w).Encode(situacoes)

// 	respondWithSuccess(situacoes, w)
// }

// func getSituacao(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	codigo := params["codigo"]
// 	// idAsString := mux.Vars(r)["id"]
// 	// id, err := stringToInt64(idAsString)
// 	situacao, err := models.CriptoEmpresaSituacaoModel{Db: db}.Find(codigo)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// var person entity.Person
// 	// database.Connector.First(&person, key)

// 	// var order Order
// 	// db.Preload("Items").First(&order, inputOrderID)

// 	// bookings := []Booking{}
// 	// db.Find(&bookings)
// 	// for _, booking := range bookings {
// 	// 	s , err:= strconv.Atoi(key)
// 	// 	if err == nil{
// 	// 	   if booking.Id == s {
// 	// 		  json.NewEncoder(w).Encode(booking)
// 	// 	   }
// 	// 	}
// 	//  }

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(situacao)
// }

// func createSituacao(w http.ResponseWriter, r *http.Request) {

// 	// stmt, err := db.Prepare("INSERT INTO posts(title) VALUES(?)")
// 	// body, err := ioutil.ReadAll(r.Body)
// 	// keyVal := make(map[string]string)
// 	// json.Unmarshal(body, &keyVal)
// 	// title := keyVal["title"]
// 	// _, err = stmt.Exec(title)
// 	// w.WriteHeader(http.StatusCreated)

// 	// body, _ := ioutil.ReadAll(r.Body)
// 	// var person entity.Person
// 	// json.Unmarshal(body, &person)
// 	// database.Connector.Create(person)
// 	// db.Create(&person)
// 	// w.WriteHeader(http.StatusCreated)
// 	// json.NewEncoder(w).Encode(person)

// 	// var order Order
// 	// json.NewDecoder(r.Body).Decode(&order)
// 	// db.Create(&order)

// 	fmt.Fprintf(w, "New post was created")
// }

// func updateSituacao(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	codigo := params["codigo"]
// 	// stmt, err := db.Prepare("UPDATE posts SET title = ? WHERE id = ?")
// 	// body, err := ioutil.ReadAll(r.Body)
// 	// keyVal := make(map[string]string)
// 	// json.Unmarshal(body, &keyVal)
// 	// title := keyVal["title"]
// 	// _, err = stmt.Exec(title, codigo)

// 	// var person entity.Person
// 	// json.Unmarshal(body, &person)
// 	// database.Connector.Save(&person)
// 	// w.WriteHeader(http.StatusOK)
// 	// json.NewEncoder(w).Encode(person)

// 	// var updatedOrder Order
// 	// json.NewDecoder(r.Body).Decode(&updatedOrder)
// 	// db.Save(&updatedOrder)

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Post with ID = %s was updated", codigo)
// }

// func deleteSituacao(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	codigo := params["codigo"]
// 	// 	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?")
// 	// 	_, err = stmt.Exec(codigo)

// 	// var person entity.Person
// 	// id, _ := strconv.ParseInt(key, 10, 64)
// 	// database.Connector.Where("id = ?", id).Delete(&person)
// 	//

// 	// id64, _ := strconv.ParseUint(inputOrderID, 10, 64)
// 	//idToDelete := uint(id64)
// 	//db.Where("order_id = ?", idToDelete).Delete(&Item{})
// 	//db.Where("order_id = ?", idToDelete).Delete(&Order{})

// 	fmt.Fprintf(w, "Post with ID = %s was deleted", codigo)
// 	w.WriteHeader(http.StatusNoContent)
// }

// //---------------------------------------------------------
// //---------------------------------------------------------

// // db, errDb := config.GetMySQLBD()
// // if errDb != nil {
// // 	panic(errDb.Error())
// // }
// // defer db.Close()
// // situacaoModel := models.CriptoEmpresaSituacaoModel{Db: db}
// // situacoes, errModel := situacaoModel.FindAll()
// // if errModel != nil {
// // 	panic(errModel.Error())
// // }
// // fmt.Println("  Situações")
// // for _, situacao := range situacoes {
// // 	fmt.Println("    Codigo:", situacao.Codigo, " - Descricao:", situacao.Descricao)
// // }
// // // situacao := entites.CriptoEmpresaSituacaoModel{Codigo: "ddd", Descricao: "ddd"}
// // // errCreate := situacaoModel.Create(&user)
// // // situacao, errFind := situacaoModel.Find(3)

// //---------------------------------------------------------
// //---------------------------------------------------------
