package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"user-app/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	// ①-1
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("fail: load .env, %v\n", err)
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	// ①-2
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// ①-3
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	switch r.Method {
	case http.MethodGet:
		usecase.RegisterUser()
		//type UserResGet struct {
		//	Id   string `json:"id"`
		//	Name string `json:"name"`
		//	Age  int    `json:"age"`
		//}
		//

		// ②-2
		//rows, err := db.Query("SELECT id,  name, age FROM frontUser")
		//if err != nil {
		//	log.Printf("fail: db.Query, %v\n", err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}

		// ②-3
		//users := make([]UserResGet, 0)
		//for rows.Next() {
		//	var u UserResGet
		//	if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
		//		log.Printf("fail: rows.Scan, %v\n", err)
		//
		//		if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
		//			log.Printf("fail: rows.Close(), %v\n", err)
		//		}
		//		w.WriteHeader(http.StatusInternalServerError)
		//		return
		//	}
		//	users = append(users, u)

		// ②-4
		//bytes, err := json.Marshal(users)
		//if err != nil {
		//	log.Printf("fail: json.Marshal, %v\n", err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//w.Header().Set("Content-Type", "application/json")
		//_, err = w.Write(bytes)
		//if err != nil {
		//	log.Printf("fail: w.Write(bytes), %v\n", err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
	//
	case http.MethodPost:
		usecase.SerchUser()
		//type UserResPOST struct {
		//	Id   string `json:"id"`
		//	Name string `json:"name"`
		//	Age  int    `json:"age"`
		//}
		//
		//var u UserResPOST

		//t, _ := io.ReadAll(r.Body)
		//if err := json.Unmarshal([]byte(t), &u); err != nil {
		//	log.Printf("fail: json.Unmarshal, %v\n", err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//if u.Name == "" {
		//	log.Println("fail: name is empty")
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//} else if utf8.RuneCountInString(u.Name) >= 50 {
		//	log.Println("fail: name is too long")
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}

		//if u.Age < 20 {
		//	log.Println("fail: age is too young")
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//} else if u.Age > 80 {
		//	log.Println("fail: age is too old")
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}
		//fmt.Println(u.Name)
		//fmt.Println(u.Age)
		//u.Id = strconv.FormatUint(ulid.Timestamp(time.Now()), 10)
		//fmt.Println(u.Id)
		//tx, err := db.Begin()
		//if err != nil {

		//log.Printf("fail: db.Begin, %v\n", err)
		//w.WriteHeader(http.StatusInternalServerError)

		//}
		//fmt.Println(tx)
		//query := fmt.Sprintf("INSERT INTO user (id, name, age) VALUES (?,?,?)")
		//result, err := tx.Exec("INSERT INTO frontUser (id, name, age) VALUES (?, ?, ?);", u.Id, u.Name, u.Age)
		//fmt.Println(result)
		//if err != nil {
		//	err := tx.Rollback()
		//	if err != nil {
		//		log.Printf("fail: tx.Exec, %v\n", err)
		//		w.WriteHeader(http.StatusInternalServerError)
		//		return
		//	}
		//}
		//if err := tx.Commit(); err != nil {
		//	log.Printf("fail: tx.Commit, %v/n", err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}

		//type responseMessage struct {
		//	Name string `json:"name"`
		//}
		//
		//bytes, err := json.Marshal(responseMessage{u.Name})
		//if err != nil {
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//
		//w.WriteHeader(http.StatusOK)
		//_, err = w.Write(bytes)
		//if err != nil {
		//	log.Printf("fail: w.Write, %v\n", err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", handler)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
