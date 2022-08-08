package Database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB
var lock = &sync.Mutex{}

func Connection() (*sql.DB, error) {

	if conn == nil {
		return nil, ErrNoOpenConnection
	}

	if err := conn.Ping(); err != nil {
		return nil, ErrConnectionLost
	}

	return conn, nil
}

func OpenMySql(cnf MySqlSettings) (*sql.DB, error) {

	if conn != nil {
		return conn, nil
	}

	lock.Lock()
	defer lock.Unlock()

	setMySqlDefaults(&cnf)

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?tls=skip-verify&autocommit=%v&multiStatements=true&parseTime=true&maxAllowedPacket=%v",
		cnf.Username,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Database,
		cnf.AutoCommit,
		cnf.MaxAllowedPacket,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(cnf.ConnMaxLifetime)
	db.SetMaxOpenConns(cnf.MaxOpenConns)
	db.SetMaxIdleConns(cnf.MaxIdleConns)

	conn = db

	return conn, nil
}

func Close() {
	lock.Lock()
	defer lock.Unlock()
	conn.Close()
	conn = nil
}
