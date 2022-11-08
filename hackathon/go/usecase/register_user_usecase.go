package usecase

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid/v2"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"
	"user-app/model"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	//type UserResPOST struct {
	//	Id   string `json:"id"`
	//	Name string `json:"name"`
	//	Age  int    `json:"age"`
	//}
	var db *sql.DB
	var u model.User

	t, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal([]byte(t), &u); err != nil {
		log.Printf("fail: json.Unmarshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if u.Name == "" {
		log.Println("fail: name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if utf8.RuneCountInString(u.Name) >= 50 {
		log.Println("fail: name is too long")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//intAge, _ := strconv.Atoi(u.Age)
	if u.Age < 20 {
		log.Println("fail: age is too young")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if u.Age > 80 {
		log.Println("fail: age is too old")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(u.Name)
	fmt.Println(u.Age)
	//var stmt *sql.Stmt
	u.Id = strconv.FormatUint(ulid.Timestamp(time.Now()), 10)
	fmt.Println(u.Id)
	tx, err := db.Begin()
	if err != nil {
		//tx.Rollback()
		log.Printf("fail: db.Begin, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(tx)
	//query := fmt.Sprintf("INSERT INTO user (id, name, age) VALUES (?,?,?)")
	result, err := tx.Exec("INSERT INTO frontUser (id, name, age) VALUES (?, ?, ?);", u.Id, u.Name, u.Age)
	fmt.Println(result)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			log.Printf("fail: tx.Exec, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v/n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type responseMessage struct {
		Name string `json:"name"`
	}

	bytes, err := json.Marshal(responseMessage{u.Name})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(bytes)
	if err != nil {
		log.Printf("fail: w.Write, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
