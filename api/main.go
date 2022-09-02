package main

import (
		"encoding/json"
		"fmt"
		"log"
		"net/http"
		"database/sql"

		"github.com/gorilla/mux"
		"github.com/joho/godotenv"
		_ "github.com/lib/pq"
)

var db *sql.DB

type Article struct {
		ArticleID string `json:"article_id"`
		Title string `json:"title"`
		UserName string `json:"user_name"`
		DatePosted string `json:"date_posted"`
		Tag string `json:"tag"`
		ArticleImage string `json:"article_image"`
		ArticlePreviewText string `json:"article_preview_text"`
		UserProfileUrl string `json:"user_profile_url"`
		ArticleUrl string `json:"article_url"`
		TimeToRead string `json:"time_to_read"`
}

func getAllArticles(w http.ResponseWriter, r*http.Request) {
		w.Header().Set("Content-Type", "application/json")
				
		rows, err := db.Query("select * from article_schema.articles")

		if err != nil {
				checkError(err)
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
		}
		defer rows.Close()

		articles := make([]Article, 0)

		// loop through the values of rows
		for rows.Next() {
				article := Article{}
				err := rows.Scan(
						&article.ArticleID, 
						&article.Title, 
						&article.UserName, 
						&article.DatePosted, 
						&article.Tag,
						&article.ArticleImage,
						&article.ArticlePreviewText,
						&article.UserProfileUrl,
						&article.ArticleUrl,
						&article.TimeToRead,
				)

				if err != nil {
						log.Println(err)
						http.Error(w, http.StatusText(500), 500)
						return
				}

				articles = append(articles, article)
		}

		if err = rows.Err(); err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
		}

		json.NewEncoder(w).Encode(articles)
}

func initializeDatabase() {
		var err error

		dotenvErr := godotenv.Load()
		host := "localhost"
		port := 5432
		user := "postgres"
		password := "postgres"
		dbname := "google_medium"

		if dotenvErr != nil {
				log.Fatal("Error loading .env file")
		}
		
		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		db, err = sql.Open("postgres", connStr)
		checkError(err)

		if err = db.Ping(); err != nil {
				panic(err)
		}

		fmt.Println("Connected to postgres database")
}

func checkError(err error) {
		if err != nil {
				panic(err)
		}
}

func initializeRouter() {
		router := mux.NewRouter()

		router.HandleFunc("/api/articles", getAllArticles).Methods("GET")

		fmt.Println("Running on port 8000")

		log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
		initializeDatabase()
		initializeRouter()
}