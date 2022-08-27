package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	// newrelic "github.com/newrelic/go-agent"
	// "github.com/newrelic/go-agent/internal"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// cd "C:\Users\chris\Desktop\CMS GoLang\cms.golang.teste.outros\cms.golang.teste.trace.new.relic"

// go mod init github.com/chrismarsilva/cms.golang.teste.newrelic
// go get github.com/gorilla/mux
// go get github.com/newrelic/go-agent
// go get github.com/newrelic/go-agent/v3/newrelic
// go get github.com/newrelic/go-agent/v3/integrations/nrgorilla
// go get github.com/newrelic/go-agent/v3/integrations/nrlogrus
// go get gorm.io/gorm/logger
// go get github.com/newrelic/go-agent/internal
// go mod tidy

// go run main.go

// func init() {
// 	internal.TrackUsage("integration", "framework", "gorilla", "v1")
// }

func main() {
	log.Println("Ini")

	// app, _ := newrelic.NewApplication(
	// 	newrelic.ConfigAppName("Cms.Golang.Teste.New.Relic"),
	// 	newrelic.ConfigLicense("<TOKEN>"),
	// )

	app, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("Cms.Golang.Teste.New.Relic"),
		newrelic.ConfigLicense("<TOKEN>"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	logger := logrus.New()
	//logger.SetFormatter(nrlogrusplugin.ContextFormatter{})
	logger.Info("Hello New Relic!")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Use(nrgorilla.Middleware(app))
	myRouter.Use(customMiddleware)

	//myRouter.HandleFunc("/", homeHandler)
	myRouter.HandleFunc(newrelic.WrapHandleFunc(app, "/", homeHandler))
	myRouter.HandleFunc(newrelic.WrapHandleFunc(app, "/users", usersHandler))
	myRouter.HandleFunc(newrelic.WrapHandleFunc(app, "/home", homeHandler))
	_, myRouter.NotFoundHandler = newrelic.WrapHandle(app, "NotFoundHandler", makeHandler("not found"))
	_, myRouter.MethodNotAllowedHandler = newrelic.WrapHandle(app, "MethodNotAllowedHandler", makeHandler("method not allowed"))

	// log.Fatal(http.ListenAndServe(":3000", myRouter))
	log.Fatal(http.ListenAndServe(":3000", nrgorilla.InstrumentRoutes(myRouter, app)))

	log.Println("Fim")
}

func makeHandler(text string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(text))
	})
}

func customMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txn := newrelic.FromContext(r.Context())
		segment := txn.StartSegment("customMiddleware")
		time.Sleep(50 * time.Millisecond)
		segment.End()
		// name := routeName(r)
		// txn := app.StartTransaction(name)
		// defer txn.End()
		// txn.SetWebRequestHTTP(r)
		// w = txn.SetWebResponse(w)
		// r = newrelic.RequestWithTransactionContext(r, txn)
		next.ServeHTTP(w, r)
	})
}

func newrelicMiddleware(app *newrelic.Application) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			txn := app.StartTransaction("ddddd")
			w = txn.SetWebResponse(w)
			txn.SetWebRequestHTTP(r)
			defer txn.End()
			r = newrelic.RequestWithTransactionContext(r, txn)
			next.ServeHTTP(w, r)
		})
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//logger.Info("homeHandler.homeHandler")
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//logger.Info("usersHandler.usersHandler")
}

// type instrumentedHandler struct {
// 	name string
// 	app  newrelic.Application
// 	orig http.Handler
// }

// func (h instrumentedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	txn := h.app.StartTransaction(h.name, w, r)
// 	defer txn.End()

// 	r = newrelic.RequestWithTransactionContext(r, txn)

// 	h.orig.ServeHTTP(txn, r)
// }

// func instrumentRoute(h http.Handler, app newrelic.Application, name string) http.Handler {
// 	if _, ok := h.(instrumentedHandler); ok {
// 		return h
// 	}
// 	return instrumentedHandler{
// 		name: name,
// 		orig: h,
// 		app:  app,
// 	}
// }

// func routeName(route *mux.Route) string {
// 	if nil == route {
// 		return ""
// 	}
// 	if n := route.GetName(); n != "" {
// 		return n
// 	}
// 	if n, _ := route.GetPathTemplate(); n != "" {
// 		return n
// 	}
// 	n, _ := route.GetHostTemplate()
// 	return n
// }

// // InstrumentRoutes instruments requests through the provided mux.Router.  Use
// // this after the routes have been added to the router.
// func InstrumentRoutes(r *mux.Router, app newrelic.Application) *mux.Router {
// 	if app != nil {
// 		r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
// 			h := instrumentRoute(route.GetHandler(), app, routeName(route))
// 			route.Handler(h)
// 			return nil
// 		})
// 		if nil != r.NotFoundHandler {
// 			r.NotFoundHandler = instrumentRoute(r.NotFoundHandler, app, "NotFoundHandler")
// 		}
// 	}
// 	return r
// }
