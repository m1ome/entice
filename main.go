package main

import (
	"log"
	"os"
	"net/http"
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-redis/redis"
)

const (
	Response = `{"db": "%v", "redis": "%v"}`
)

func main() {
	// Database connection
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DATABASE_ADDR"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Database: os.Getenv("DATABASE_DATABASE"),
	})

	var dbVersion string
	if _, err := db.QueryOne(&dbVersion, "SELECT VERSION()"); err != nil {
		log.Fatalf("Error grep database version: %v", err)
	}

	// Redis connection
	re := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	cmd := re.Time()
	res, err := cmd.Result()
	if err != nil {
		log.Fatalf("Error getting time from Redis: %v", err)
	}

	// Listen and serving using data
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, Response, dbVersion, res.Format(time.RFC3339))
	})

	log.Printf("Start listening on: %v", os.Getenv("LISTEN"))
	// Start listening
	if err := http.ListenAndServe(os.Getenv("LISTEN"), nil); err != nil {
		log.Fatalf("Error listening: %v", err)
	}
}
