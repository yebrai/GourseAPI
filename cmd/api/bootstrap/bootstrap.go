package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	mooc "GourseAPI/internal"
	"GourseAPI/internal/creating"
	"GourseAPI/internal/increasing"
	"GourseAPI/internal/platform/bus/inmemory"
	"GourseAPI/internal/platform/server"
	"GourseAPI/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

func Run() error {
	var cfg config
	err := envconfig.Process("MOOC", &cfg)
	if err != nil {
		return err
	}

	log.Printf("Configuración del servidor: Host=%s, Port=%d", cfg.Host, cfg.Port) // Only for depure

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	courseRepository := mysql.NewCourseRepository(db, cfg.DbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseService := increasing.NewCourseCounterService()

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseService),
	)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"admin"`
	DbPass    string        `default:"admin"`
	DbHost    string        `default:"mysql"`
	DbPort    uint          `default:"3306"`
	DbName    string        `default:"gourse_api_db"`
	DbTimeout time.Duration `default:"5s"`
}
