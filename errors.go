package Database

import "errors"

var ErrNoOpenConnection = errors.New("no open connection found")
var ErrConnectionLost = errors.New("database connection lost")
