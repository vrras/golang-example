package database

import "fmt"

var connection string

func init() {
	fmt.Println("Init dipanggil")
	connection = "MySQL"
}

// GetDatabase is ...
func GetDatabase() string {
	return connection
}

// SetDatabase is ...
func SetDatabase(database string) {
	connection = database
}
