package usecase

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"user-app/model"
)

func SerchUser(w http.ResponseWriter, r *http.Request) {
	//type UserResGet struct {
	//	Id   string `json:"id"`
	//	Name string `json:"name"`
	//	Age  int    `json:"age"`
	//}

	// ②-2
	var db *sql.DB
	rows, err := db.Query("SELECT id,  name, age FROM frontUser")
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ②-3
	users := make([]model.User, 0)
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}
	// ②-4
	bytes, err := json.Marshal(users)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bytes)
	if err != nil {
		log.Printf("fail: w.Write(bytes), %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
