package data

import (
	"errors"
)

var ErrNoRecord = errors.New("no record found")

// Models that match the data store would go below

// Data example model only from MySQL
// mysql> describe data;
// +-------+--------------+------+-----+---------+----------------+
// | Field | Type         | Null | Key | Default | Extra          |
// +-------+--------------+------+-----+---------+----------------+
// | id    | int(11)      | NO   | PRI | NULL    | auto_increment |
// | name  | varchar(255) | NO   | MUL | NULL    |                |
// | value | text         | NO   |     | NULL    |                |
// +-------+--------------+------+-----+---------+----------------+
type Data struct {
	Id    int64
	Name  string
	Value string
}
