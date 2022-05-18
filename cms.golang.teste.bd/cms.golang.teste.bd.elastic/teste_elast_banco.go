package main

// ├── cmd
// │   └── elastic
// │       └── main.go
// ├── docker
// │   ├── docker-compose.yaml
// └── internal
//     ├── pkg
//     │   ├── domain
//     │   │   └── error.go
//     │   └── storage
//     │       ├── elasticsearch
//     │       │   ├── elasticsearch.go
//     │       │   └── post_storage.go
//     │       └── post_storer.go
//     └── post
//         ├── handler.go
//         ├── request.go
//         ├── response.go
//         └── service.go

// main.go

// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/you/elastic/internal/pkg/storage/elasticsearch"
// 	"github.com/you/elastic/internal/post"

// 	"github.com/julienschmidt/httprouter"
// )

// func main() {
// 	// Bootstrap elasticsearch.
// 	elastic, err := elasticsearch.New([]string{"http://0.0.0.0:9200"})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	if err := elastic.CreateIndex("post"); err != nil {
// 		log.Fatalln(err)
// 	}

// 	// Bootstrap storage.
// 	storage, err := elasticsearch.NewPostStorage(*elastic)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	postAPI := post.New(storage)
// 	router := httprouter.New()
// 	router.HandlerFunc("POST", "/api/v1/posts", postAPI.Create)
// 	router.HandlerFunc("PATCH", "/api/v1/posts/:id", postAPI.Update)
// 	router.HandlerFunc("DELETE", "/api/v1/posts/:id", postAPI.Delete)
// 	router.HandlerFunc("GET", "/api/v1/posts/:id", postAPI.Find)
// 	log.Fatalln(http.ListenAndServe(":3000", router))
// }

// error.go

// package domain

// import "errors"

// var (
// 	ErrNotFound = errors.New("not found")
// 	ErrConflict = errors.New("conflict")
// )

// elasticsearch.go

// package elasticsearch

// import (
// 	"fmt"

// 	"github.com/elastic/go-elasticsearch/v7"
// )

// type ElasticSearch struct {
// 	client *elasticsearch.Client
// 	index  string
// 	alias  string
// }

// func New(addresses []string) (*ElasticSearch, error) {
// 	cfg := elasticsearch.Config{
// 		Addresses: addresses,
// 	}

// 	client, err := elasticsearch.NewClient(cfg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ElasticSearch{
// 		client: client,
// 	}, nil
// }

// func (e *ElasticSearch) CreateIndex(index string) error {
// 	e.index = index
// 	e.alias = index + "_alias"

// 	res, err := e.client.Indices.Exists([]string{e.index})
// 	if err != nil {
// 		return fmt.Errorf("cannot check index existence: %w", err)
// 	}
// 	if res.StatusCode == 200 {
// 		return nil
// 	}
// 	if res.StatusCode != 404 {
// 		return fmt.Errorf("error in index existence response: %s", res.String())
// 	}

// 	res, err = e.client.Indices.Create(e.index)
// 	if err != nil {
// 		return fmt.Errorf("cannot create index: %w", err)
// 	}
// 	if res.IsError() {
// 		return fmt.Errorf("error in index creation response: %s", res.String())
// 	}

// 	res, err = e.client.Indices.PutAlias([]string{e.index}, e.alias)
// 	if err != nil {
// 		return fmt.Errorf("cannot create index alias: %w", err)
// 	}
// 	if res.IsError() {
// 		return fmt.Errorf("error in index alias creation response: %s", res.String())
// 	}

// 	return nil
// }

// // document represents a single document in Get API response body.
// type document struct {
// 	Source interface{} `json:"_source"`
// }

// post_storage.go

// package elasticsearch

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"time"

// 	"github.com/you/elastic/internal/pkg/domain"
// 	"github.com/you/elastic/internal/pkg/storage"

// 	"github.com/elastic/go-elasticsearch/v7/esapi"
// )

// var _ storage.PostStorer = PostStorage{}

// type PostStorage struct {
// 	elastic ElasticSearch
// 	timeout time.Duration
// }

// func NewPostStorage(elastic ElasticSearch) (PostStorage, error) {
// 	return PostStorage{
// 		elastic: elastic,
// 		timeout: time.Second * 10,
// 	}, nil
// }

// func (p PostStorage) Insert(ctx context.Context, post storage.Post) error {
// 	bdy, err := json.Marshal(post)
// 	if err != nil {
// 		return fmt.Errorf("insert: marshall: %w", err)
// 	}

// 	// res, err := p.elastic.client.Create()
// 	req := esapi.CreateRequest{
// 		Index:      p.elastic.alias,
// 		DocumentID: post.ID,
// 		Body:       bytes.NewReader(bdy),
// 	}

// 	ctx, cancel := context.WithTimeout(ctx, p.timeout)
// 	defer cancel()

// 	res, err := req.Do(ctx, p.elastic.client)
// 	if err != nil {
// 		return fmt.Errorf("insert: request: %w", err)
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode == 409 {
// 		return domain.ErrConflict
// 	}

// 	if res.IsError() {
// 		return fmt.Errorf("insert: response: %s", res.String())
// 	}

// 	return nil
// }

// func (p PostStorage) Update(ctx context.Context, post storage.Post) error {
// 	bdy, err := json.Marshal(post)
// 	if err != nil {
// 		return fmt.Errorf("update: marshall: %w", err)
// 	}

// 	// res, err := p.elastic.client.Update()
// 	req := esapi.UpdateRequest{
// 		Index:      p.elastic.alias,
// 		DocumentID: post.ID,
// 		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, bdy))),
// 	}

// 	ctx, cancel := context.WithTimeout(ctx, p.timeout)
// 	defer cancel()

// 	res, err := req.Do(ctx, p.elastic.client)
// 	if err != nil {
// 		return fmt.Errorf("update: request: %w", err)
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode == 404 {
// 		return domain.ErrNotFound
// 	}

// 	if res.IsError() {
// 		return fmt.Errorf("update: response: %s", res.String())
// 	}

// 	return nil
// }

// func (p PostStorage) Delete(ctx context.Context, id string) error {
// 	// res, err := p.elastic.client.Delete()
// 	req := esapi.DeleteRequest{
// 		Index:      p.elastic.alias,
// 		DocumentID: id,
// 	}

// 	ctx, cancel := context.WithTimeout(ctx, p.timeout)
// 	defer cancel()

// 	res, err := req.Do(ctx, p.elastic.client)
// 	if err != nil {
// 		return fmt.Errorf("delete: request: %w", err)
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode == 404 {
// 		return domain.ErrNotFound
// 	}

// 	if res.IsError() {
// 		return fmt.Errorf("delete: response: %s", res.String())
// 	}

// 	return nil
// }

// func (p PostStorage) FindOne(ctx context.Context, id string) (storage.Post, error) {
// 	// res, err := p.elastic.client.Get()
// 	req := esapi.GetRequest{
// 		Index:      p.elastic.alias,
// 		DocumentID: id,
// 	}

// 	ctx, cancel := context.WithTimeout(ctx, p.timeout)
// 	defer cancel()

// 	res, err := req.Do(ctx, p.elastic.client)
// 	if err != nil {
// 		return storage.Post{}, fmt.Errorf("find one: request: %w", err)
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode == 404 {
// 		return storage.Post{}, domain.ErrNotFound
// 	}

// 	if res.IsError() {
// 		return storage.Post{}, fmt.Errorf("find one: response: %s", res.String())
// 	}

// 	var (
// 		post storage.Post
// 		body document
// 	)
// 	body.Source = &post

// 	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
// 		return storage.Post{}, fmt.Errorf("find one: decode: %w", err)
// 	}

// 	return post, nil
// }

// post_storer.go

// package storage

// import (
// 	"context"
// 	"time"
// )

// type PostStorer interface {
// 	Insert(ctx context.Context, post Post) error
// 	Update(ctx context.Context, post Post) error
// 	Delete(ctx context.Context, id string) error
// 	FindOne(ctx context.Context, id string) (Post, error)
// }

// type Post struct {
// 	ID        string     `json:"id"`
// 	Title     string     `json:"title"`
// 	Text      string     `json:"text"`
// 	Tags      []string   `json:"tags"`
// 	CreatedAt *time.Time `json:"created_at,omitempty"`
// }

// handler.go

// package post

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/you/elastic/internal/pkg/domain"
// 	"github.com/you/elastic/internal/pkg/storage"

// 	"github.com/julienschmidt/httprouter"
// )

// type Handler struct {
// 	service service
// }

// func New(storage storage.PostStorer) Handler {
// 	return Handler{
// 		service: service{storage: storage},
// 	}
// }

// // POST /api/v1/posts
// func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
// 	var req createRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Println(err)
// 		return
// 	}

// 	res, err := h.service.create(r.Context(), req)
// 	if err != nil {
// 		switch err {
// 		case domain.ErrConflict:
// 			w.WriteHeader(http.StatusConflict)
// 		default:
// 			w.WriteHeader(http.StatusInternalServerError)
// 			log.Println(err)
// 		}
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(http.StatusCreated)
// 	bdy, _ := json.Marshal(res)
// 	_, _ = w.Write(bdy)
// }

// // PATCH /api/v1/posts/:id
// func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
// 	var req updateRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Println(err)
// 		return
// 	}

// 	req.ID = httprouter.ParamsFromContext(r.Context()).ByName("id")

// 	if err := h.service.update(r.Context(), req); err != nil {
// 		switch err {
// 		case domain.ErrNotFound:
// 			w.WriteHeader(http.StatusNotFound)
// 		default:
// 			w.WriteHeader(http.StatusInternalServerError)
// 			log.Println(err)
// 		}
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

// // DELETE /api/v1/posts/:id
// func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
// 	id := httprouter.ParamsFromContext(r.Context()).ByName("id")

// 	if err := h.service.delete(r.Context(), deleteRequest{ID: id}); err != nil {
// 		switch err {
// 		case domain.ErrNotFound:
// 			w.WriteHeader(http.StatusNotFound)
// 		default:
// 			w.WriteHeader(http.StatusInternalServerError)
// 			log.Println(err)
// 		}
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

// // GET /api/v1/posts/:id
// func (h Handler) Find(w http.ResponseWriter, r *http.Request) {
// 	id := httprouter.ParamsFromContext(r.Context()).ByName("id")

// 	res, err := h.service.find(r.Context(), findRequest{ID: id})
// 	if err != nil {
// 		switch err {
// 		case domain.ErrNotFound:
// 			w.WriteHeader(http.StatusNotFound)
// 		default:
// 			w.WriteHeader(http.StatusInternalServerError)
// 			log.Println(err)
// 		}
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	bdy, _ := json.Marshal(res)
// 	_, _ = w.Write(bdy)
// }

// request.go

// package post

// type createRequest struct {
// 	Title string   `json:"title"`
// 	Text  string   `json:"text"`
// 	Tags  []string `json:"tags"`
// }

// type updateRequest struct {
// 	ID string

// 	Title string   `json:"title"`
// 	Text  string   `json:"text"`
// 	Tags  []string `json:"tags"`
// }

// type deleteRequest struct {
// 	ID string
// }

// type findRequest struct {
// 	ID string
// }

// response.go

// package post

// import "time"

// type createResponse struct {
// 	ID string `json:"id"`
// }

// type findResponse struct {
// 	ID        string    `json:"id"`
// 	Title     string    `json:"title"`
// 	Text      string    `json:"text"`
// 	Tags      []string  `json:"tags"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// service.go

// package post

// import (
// 	"context"
// 	"time"
// 	"github.com/you/elastic/internal/pkg/storage"
// 	"github.com/google/uuid"
// )

// type service struct {
// 	storage storage.PostStorer
// }

// func (s service) create(ctx context.Context, req createRequest) (createResponse, error) {
// 	id := uuid.New().String()
// 	cr := time.Now().UTC()

// 	doc := storage.Post{
// 		ID:        id,
// 		Title:     req.Title,
// 		Text:      req.Text,
// 		Tags:      req.Tags,
// 		CreatedAt: &cr,
// 	}

// 	if err := s.storage.Insert(ctx, doc); err != nil {
// 		return createResponse{}, err
// 	}

// 	return createResponse{ID: id}, nil
// }

// func (s service) update(ctx context.Context, req updateRequest) error {
// 	doc := storage.Post{
// 		ID:    req.ID,
// 		Title: req.Title,
// 		Text:  req.Text,
// 		Tags:  req.Tags,
// 	}

// 	if err := s.storage.Update(ctx, doc); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s service) delete(ctx context.Context, req deleteRequest) error {
// 	if err := s.storage.Delete(ctx, req.ID); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s service) find(ctx context.Context, req findRequest) (findResponse, error) {
// 	post, err := s.storage.FindOne(ctx, req.ID)
// 	if err != nil {
// 		return findResponse{}, err
// 	}

// 	return findResponse{
// 		ID:        post.ID,
// 		Title:     post.Title,
// 		Text:      post.Text,
// 		Tags:      post.Tags,
// 		CreatedAt: *post.CreatedAt,
// 	}, nil
// }
