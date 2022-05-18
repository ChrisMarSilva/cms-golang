package redigo

import (
	"fmt"
	//"context"
	"bytes"
	"encoding/gob"
	"encoding/json"
	redigo "github.com/gomodule/redigo/redis"
	"log"
	"reflect"
	"strconv"
	"time"
)

func newPool() *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     80,
		MaxActive:   12000,
		IdleTimeout: 1 * time.Minute,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial(
				"tcp",
				"localhost:6379",           // "localhost:6379" // redis-12888.c267.us-east-1-4.ec2.cloud.redislabs.com:12888
				redigo.DialPassword("123"), // "123" // "KC6xF4LijcE8AErIDD2KZOzN6rnimQCI"
				redigo.DialConnectTimeout(1*time.Minute),
				redigo.DialReadTimeout(1*time.Minute),
				redigo.DialWriteTimeout(1*time.Minute),
				redigo.DialDatabase(0),
			)
			if err != nil {
				return nil, err
				//panic(err.Error())
			}
			return c, nil
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			//_, err := c.Do("PING")
			s, err := redigo.String(c.Do("PING"))
			log.Println("PING Response =", s)
			return err
		},
	}
}

type Album struct {
	Title  string  `redis:"title"`
	Artist string  `redis:"artist"`
	Price  float64 `redis:"price"`
	Likes  int     `redis:"likes"`
}
