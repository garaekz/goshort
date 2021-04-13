package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/config"
	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/internal/healthcheck"
	"github.com/garaekz/goshort/internal/link"
	"github.com/garaekz/goshort/internal/page"
	"github.com/garaekz/goshort/internal/user"
	"github.com/garaekz/goshort/pkg/accesslog"
	"github.com/garaekz/goshort/pkg/dbcontext"
	"github.com/garaekz/goshort/pkg/log"
	dbx "github.com/go-ozzo/ozzo-dbx"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	"github.com/go-ozzo/ozzo-routing/v2/file"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// Version indicates the current version of the application.
var Version = "1.0.0"

var flagConfig = flag.String("config", "./config/local.yml", "path to the config file")

func main() {
	flag.Parse()
	// create root logger tagged with server version
	logger := log.New().With(context.TODO(), "version", Version)

	// load application configurations
	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	// connect to the database
	db, err := dbx.MustOpen(cfg.DBType, cfg.DSN)
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	db.QueryLogFunc = logDBQuery(logger)
	db.ExecLogFunc = logDBExec(logger)
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error(err)
		}
	}()

	// build HTTP server
	address := fmt.Sprintf(":%v", cfg.ServerPort)
	hs := &http.Server{
		Addr:    address,
		Handler: buildHandler(logger, dbcontext.New(db), cfg),
	}

	// start the HTTP server with graceful shutdown
	go routing.GracefulShutdown(hs, 10*time.Second, logger.Infof)
	logger.Infof("server %v is running at %v", Version, address)
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(err)
		os.Exit(-1)
	}
}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(logger log.Logger, db *dbcontext.DB, cfg *config.Config) http.Handler {
	router := routing.New()

	router.Use(
		accesslog.Handler(logger),
		errors.Handler(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.AllowAll),
	)

	healthcheck.RegisterHandlers(router, Version)

	rg := router.Group("/v1")
	rg2 := router.Group("")

	authHandler := auth.Handler(cfg.JWTSigningKey)
	customAuthHandler := auth.CustomHandler(cfg.JWTSigningKey)

	auth.RegisterHandlers(rg.Group(""),
		auth.NewService(db, cfg.JWTSigningKey, cfg.JWTExpiration, logger),
		logger,
	)

	user.RegisterHandlers(rg.Group(""),
		user.NewService(user.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	link.RegisterHandlers(rg.Group(""),
		link.NewService(link.NewRepository(db, logger), logger),
		authHandler, customAuthHandler, logger,
	)

	// serve index file
	router.Get("/", file.Content("./view/dist/index.html"))
	// serve assets folder
	router.Get("/assets/*", file.Server(file.PathMap{
		"/": "./view/dist/",
	}))

	page.RegisterHandlers(rg2.Group(""),
		page.NewService(page.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	return router
}

// logDBQuery returns a logging function that can be used to log SQL queries.
func logDBQuery(logger log.Logger) dbx.QueryLogFunc {
	return func(ctx context.Context, t time.Duration, sql string, rows *sql.Rows, err error) {
		if err == nil {
			logger.With(ctx, "duration", t.Milliseconds(), "sql", sql).Info("DB query successful")
		} else {
			logger.With(ctx, "sql", sql).Errorf("DB query error: %v", err)
		}
	}
}

// logDBExec returns a logging function that can be used to log SQL executions.
func logDBExec(logger log.Logger) dbx.ExecLogFunc {
	return func(ctx context.Context, t time.Duration, sql string, result sql.Result, err error) {
		if err == nil {
			logger.With(ctx, "duration", t.Milliseconds(), "sql", sql).Info("DB execution successful")
		} else {
			logger.With(ctx, "sql", sql).Errorf("DB execution error: %v", err)
		}
	}
}
