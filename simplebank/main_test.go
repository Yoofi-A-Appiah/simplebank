package simplebank

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect:", err)
	}

	testQueries = New(conn)

	// Run all tests and capture the exit code.
	exitCode := m.Run()

	// Close any resources if needed. For example:
	// conn.Close()

	os.Exit(exitCode)
}
