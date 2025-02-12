package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	mooc "GourseAPI/internal"
	"GourseAPI/internal/creating"
	"GourseAPI/internal/increasing"
	"GourseAPI/internal/platform/bus/inmemory"
	"GourseAPI/internal/platform/server"
	"GourseAPI/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

	dbUser    = "root"
	dbPass    = "yebrai"
	dbHost    = "localhost"
	dbPort    = "3306"
	dbName    = "gourseApi"
	dbTimeout = 5 * time.Second
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	courseRepository := mysql.NewCourseRepository(db, dbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseCounterService := increasing.NewCourseCounterService()

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService),
	)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}
