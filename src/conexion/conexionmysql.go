package conexion

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var SessionMysql *sql.DB

var usuario = "jordan"
var pass = "9595K9595k." //
var host = "tcp(192.100.10.49:3306)"
var nombreBaseDeDatos = "control_asistencias"

/*
var usuario = "jordan"
var pass = "9595K9595k." //
var host = "tcp(192.100.10.49:3306)"
var nombreBaseDeDatos = "control_asistencias"


var usuario = "root"
var pass = "" //
var host = "tcp(127.0.0.1:3306)"
var nombreBaseDeDatos = "control_asistencias"
*/

func init() {
	ConnectMysql()
}

func ConnectMysql() {
	var err error

	SessionMysql, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		panic(err)
	}

	if err != nil {
		time.Sleep(10000 * time.Millisecond)
		ConnectMysql()
	}
}

func CloseConexionMysql() {
	SessionMysql.Close()
}
