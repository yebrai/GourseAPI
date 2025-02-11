package bootstrap

import (
	"GourseAPI/internal/creating"
	"GourseAPI/internal/platform/bus/inmemory"
	"GourseAPI/internal/platform/server"
	"GourseAPI/internal/platform/storage/mysql"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

	var (
		commandBus = inmemory.NewCommandBus()
	)

	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	srv := server.New(host, port, commandBus)
	return srv.Run()
}
