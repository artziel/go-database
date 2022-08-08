package Database

import "time"

type MySqlSettings struct {
	Username         string
	Password         string
	Host             string
	Port             string
	Database         string
	MaxAllowedPacket int
	ConnMaxLifetime  time.Duration
	MaxOpenConns     int
	MaxIdleConns     int
	AutoCommit       bool
}

func setMySqlDefaults(cnf *MySqlSettings) {
	if cnf.MaxIdleConns == 0 {
		cnf.MaxIdleConns = 10
	}
	if cnf.MaxOpenConns == 0 {
		cnf.MaxOpenConns = 10
	}
	if cnf.ConnMaxLifetime == 0 {
		cnf.ConnMaxLifetime = time.Minute * 3
	}
	if cnf.MaxAllowedPacket == 0 {
		cnf.MaxAllowedPacket = 12194304
	}
	if cnf.Port == "" {
		cnf.Port = "3306"
	}
}
