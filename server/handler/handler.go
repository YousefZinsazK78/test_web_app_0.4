package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/yousefzinsazk78/test_web_app_0.4/server/types"
)

var db *sql.DB

func SetDb(oldDB *sql.DB) {
	db = oldDB
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	log.Print("/CreatePost")
	switch r.Method {
	case http.MethodPost:
		var postBlog types.Post
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&postBlog); err != nil {
			log.Fatal(err.Error())
		}
		result, err := db.Exec("INSERT INTO post_tbl (Title,Description,AuthorID) VALUES (? ,? ,?)", postBlog.Title, postBlog.Description, postBlog.AuthorID)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Print(result.LastInsertId())
		fmt.Fprint(w, "post blog inserted successfully....")
	default:
		log.Fatal("invalid method 😑")
	}

}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	log.Print("/UpdatePost")
	switch r.Method {
	case http.MethodPut:
		var postBlog types.Post
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&postBlog); err != nil {
			log.Fatal(err.Error())
		}
		result, err := db.Exec("UPDATE post_tbl SET Title=? WHERE ID=?", postBlog.Title, postBlog.ID)
		if err != nil {
			log.Fatal("error in update request handling...", err)
			return
		} else {
			res, err := result.RowsAffected()
			if err != nil {
				log.Fatal("error in update request handling...", err)
				return
			}
			fmt.Fprint(w, "update successful...")
			fmt.Fprint(w, res)
		}
	default:
		log.Fatal("invalid method 😑")
	}

}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	log.Print("/CreateAuthor")
	switch r.Method {
	case http.MethodPost:
		var author types.Author
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
			log.Fatal(err.Error())
		}
		result, err := db.Exec("INSERT INTO author_tbl (FullName, Email) VALUES (? ,?)", author.FullName, author.Email)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Print(result.LastInsertId())
		fmt.Fprint(w, "author inserted successfully....")
	default:
		log.Fatal("invalid method!😑")
	}

}

func ReadAuthor(w http.ResponseWriter, r *http.Request) {
	log.Print("/ReadAuthor")
	switch r.Method {
	case http.MethodGet:
		var authors []types.Author

		rows, err := db.Query("SELECT * FROM author_tbl")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var authorItem types.Author
			if err := rows.Scan(&authorItem.ID, &authorItem.FullName, &authorItem.Email); err != nil {
				log.Fatal("put author to author item has err", err)
				return
			}
			authors = append(authors, authorItem)
		}

		if err := rows.Err(); err != nil {
			log.Fatal("err in rows.Err", err)
			return
		}
		//return json encoded data (authors data)
		json.NewEncoder(w).Encode(authors)
	default:
		log.Fatal("invalid method!😑")
	}

}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	log.Print("/UpdateAuthor")
	switch r.Method {
	case http.MethodPut:
		var author types.Author
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
			log.Fatal(err.Error())
		}
		result, err := db.Exec("UPDATE author_tbl SET FullName=? WHERE authorID=?", author.FullName, author.ID)
		if err != nil {
			log.Fatal("error in update request handling...", err)
			return
		} else {
			res, err := result.RowsAffected()
			if err != nil {
				log.Fatal("error in update request handling...", err)
				return
			}
			fmt.Fprint(w, "update successful...")
			fmt.Fprint(w, res)
		}
	default:
		log.Fatal("invalid method!😑")
	}

}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	log.Print("/DeleteAuthor")
	switch r.Method {
	case http.MethodDelete:
		var author types.Author
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
			log.Fatal(err.Error())
		}
		result, err := db.Exec("DELETE FROM author_tbl WHERE authorID=?", author.ID)
		if err != nil {
			log.Fatal("error in delete request handling...", err)
			return
		} else {
			res, err := result.RowsAffected()
			if err != nil {
				log.Fatal("error in delete request handling...", err)
				return
			}
			fmt.Fprint(w, "delete successfully affected...")
			fmt.Fprint(w, res)
		}
	default:
		log.Fatal("invalid method!😑")
	}

}
