package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/go-pg/pg"
)

var (
	host     string
	port     string
	db       string
	user     string
	addr     string
	password string
)

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "listening host")
	flag.StringVar(&port, "port", "12345", "listening port")
	flag.StringVar(&db, "db", "devops", "database db")
	flag.StringVar(&user, "db-user", "postgres", "database username")
	flag.StringVar(&addr, "db-addr", "127.0.0.1:5432", "database address")
	flag.StringVar(&password, "db-password", "password", "database password")
	flag.Parse()
}

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: db,
	})

	http.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		var users []User
		if err := db.Model(&users).Select(); err != nil {
			log.Printf("error getting users: %v", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		buf, err := json.Marshal(users)
		if err != nil {
			log.Printf("error marshalling users: %v", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusOK)
		if _, err := res.Write(buf); err != nil {
			log.Printf("error writing response: %v", err)
		}
	})

	log.Printf("server startging at %s:%s", host, port)
	if err := http.ListenAndServe(host+":"+port, nil); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
