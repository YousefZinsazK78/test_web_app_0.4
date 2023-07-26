package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/yousefzinsazk78/test_web_app_0.4/server/database"
	"github.com/yousefzinsazk78/test_web_app_0.4/server/handler"
)

type Server struct {
	db *sql.DB
}

func New() *Server {
	return &Server{
		db: database.RunningDataBase(),
	}
}

func (s *Server) Serve(port string) {

	handler.SetDb(s.db)
	//handle incoming request
	http.HandleFunc("/createPost", handler.CreatePost)
	http.HandleFunc("/createAuthor", handler.CreateAuthor)
	http.HandleFunc("/readAuthor", handler.ReadAuthor)
	http.HandleFunc("/updateAuthor", handler.UpdateAuthor)
	http.HandleFunc("/deleteAuthor", handler.DeleteAuthor)

	fmt.Printf("server running🏃‍♂️ and listening👂 on %s \n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
