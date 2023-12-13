package conexion

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

var Session *sql.DB

var server = "192.100.10.116"
var port = 1433
var user = "sa"
var password = "FutL@mC0AC23"
var database = "FuturoLamanense"

//192.100.10.116 FutL@mC0AC23

func init() {
	Connect()
}

func Connect() {
	var err error

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)

	Session, err = sql.Open("sqlserver", connString)
	if err != nil {
		panic(err)
	}

	if err != nil {
		time.Sleep(10000 * time.Millisecond)
		Connect()
	}
}

func CloseConexion() {
	Session.Close()
}
