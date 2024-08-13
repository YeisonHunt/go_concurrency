package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "time"
    _ "github.com/lib/pq"
)

type App struct {
    DB *sql.DB
}

type Response struct {
    Data string `json:"data"`
}

func main() {
    connStr := "user=youruser dbname=yourdb sslmode=disable password=yourpassword"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    app := &App{DB: db}

    http.HandleFunc("/query", app.queryHandler)
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func (a *App) queryHandler(w http.ResponseWriter, r *http.Request) {
    start := time.Now()

    var data string
    err := a.DB.QueryRow("SELECT data FROM your_table WHERE id = $1", 1).Scan(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    elapsed := time.Since(start)
    log.Printf("Query took %s", elapsed)

    response := Response{Data: data}
    json.NewEncoder(w).Encode(response)
}