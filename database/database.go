package database

var connection string

func init() {
	connection = "sql"
}

func GetDatabase() string {
	return connection
}