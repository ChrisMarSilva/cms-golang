package goredis

import (
	"bytes"
	"encoding/gob"
	"github.com/go-redis/redis"
	"time"
)

func NewRedis(db int) (*Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379", // "localhost:6379" // redis-12888.c267.us-east-1-4.ec2.cloud.redislabs.com:12888
		Password:     "123",            // "123" // "KC6xF4LijcE8AErIDD2KZOzN6rnimQCI"
		DB:           db,
		DialTimeout:  1 * time.Hour, // Minute
		WriteTimeout: 1 * time.Hour, // Minute,
		ReadTimeout:  1 * time.Hour, // Minute,
		PoolTimeout:  1 * time.Hour, // Minute,
	})

	//_, err := client.Ping(context.Background()).Result()
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

type Message struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	Target  string `json:"target"`
	Sender  string `json:"sender"`
}

type Pessoa struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Client struct {
	client *redis.Client
}

func (c *Client) SetPessoa(pessoa Pessoa) error {
	var pessoaBytes bytes.Buffer
	err := gob.NewEncoder(&pessoaBytes).Encode(pessoa)
	if err != nil {
		return err
	}
	// return c.client.Set(pessoa.ID, pessoaBytes.Bytes(), 25*time.Second).Err()
	return c.client.Set(pessoa.ID, pessoaBytes.Bytes(), 0).Err()

	// json, err := json.Marshal(pessoa)
	// c.client.Set(pessoa.ID, json, 0)

}

func (c *Client) GetPessoa(idPessoa string) (Pessoa, error) {
	value := c.client.Get(idPessoa)

	pessoaBytes, err := value.Bytes()
	if err != nil {
		return Pessoa{}, err
	}

	b := bytes.NewReader(pessoaBytes)

	var pessoa Pessoa
	err = gob.NewDecoder(b).Decode(&pessoa)
	if err != nil {
		return Pessoa{}, err
	}

	return pessoa, nil

	// value, err := c.client.Get(idPessoa).Result()
	// var pessoa Pessoa // pessoa := Pessoa{}
	// err = json.Unmarshal([]byte(value), &pessoa)
	// return &pessoa, nil

}
