package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _ "github.com/lib/pq" // Import anonyme du pilote PostgreSQL
)

var db *sql.DB

// Cette fonction établit une connexion à la base de données une seule fois.
func init() {
    var err error

    connStr := "postgres://postgres:postgres@postgres_invites/postgres?sslmode=disable"
    db, err = sql.Open("postgres", connStr)

    if err != nil {
        panic(err)
    }

    if err = db.Ping(); err != nil {
        panic(err)
    }
    // Ceci sera imprimé dans le terminal, confirmant la connexion à la base de données
    fmt.Println("The database is connected")
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
