package conexion

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

var Session *sql.DB

var server = "192.100.10.16"
var port = 1433
var user = "sa"
var password = "FLMAdministrador14"

//192.100.10.16 FLMAdministrador14

func init() {
	Connect()
}

func Connect() {
	var err error

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		server, user, password, port)

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
