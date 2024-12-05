package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"passm/auth"
	"passm/config"
	"passm/data"

	_ "github.com/lib/pq"
)

type Server struct {
	pgsql_conn *sql.DB
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start() {
	config.LoadConfig()
	s.intiDB()
	auth.SetConnectionForAuth(s.pgsql_conn)
	data.SetConnectionForData(s.pgsql_conn)

	http.HandleFunc("/v1/register", auth.Register)
	http.HandleFunc("/v1/login", auth.Login)
	// http.HandleFunc("/reqpass", auth.RequestPass)
	http.HandleFunc("/v1/addpass", auth.AddPass)
	http.HandleFunc("/v1/delpass", auth.DelPass)

	log.Fatalln(
		http.ListenAndServe(
			fmt.Sprintf("%s:%d", config.GLOBAL_CONFIG.APPLICATION_HOSTNAME, config.GLOBAL_CONFIG.APPLICATION_HOSTPORT),
			nil,
		),
	)
}

func (s *Server) intiDB() {

	pgsql_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.GLOBAL_CONFIG.PGSQL_HOSTNAME,
		config.GLOBAL_CONFIG.PGSQL_HOSTPORT,
		config.GLOBAL_CONFIG.PGSQL_USERNAME,
		config.GLOBAL_CONFIG.PGSQL_PASSWORD,
		config.GLOBAL_CONFIG.PGSQL_DATABASE,
	)
	db, err := sql.Open("postgres", pgsql_info)
	if err != nil {
		log.Fatalln("db conn failed", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("pgsql initializtion failed: ", err.Error())
	}
	s.pgsql_conn = db
}
