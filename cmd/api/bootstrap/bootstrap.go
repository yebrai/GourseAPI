package bootstrap

import (
	"GourseAPI/internal/platform/server"
	"GourseAPI/internal/platform/storage/mysql"
	"database/sql"
	"fmt"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "root"
	dbPass = "yebrai"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "gourseApi"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
