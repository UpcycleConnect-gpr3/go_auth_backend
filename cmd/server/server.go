package server

import (
	"authentication_backend/app/handlers/auth_handlers"
	"authentication_backend/app/handlers/metric_handlers"
	"authentication_backend/app/middleware/ratelimit_middleware"
	"authentication_backend/app/middleware/source_middleware"
	"authentication_backend/config"
	"authentication_backend/database"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	log "github.com/thedataflows/go-lib-log"
)

func initialize() {

	logger := log.NewLoggerBuilder().WithLogLevel(zerolog.DebugLevel).WithBufferSize(10000).WithRateLimit(1000).WithGroupWindow(2 * time.Second).WithLogFormat(log.LOG_FORMAT_JSON).WithOutput(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).Build()
	defer logger.Close()

	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal().Err(err).Msg("Error loading .env file")
	}

	// Config Initialization
	config.InitDatabase()

	err = database.Auth.Ping()

	if err != nil {
		logger.Fatal().Err(err).Msg("(DATABASE)")
	}
}

func Start() {

	initialize()

	limiterLow := ratelimit_middleware.NewRateLimiter(10, 1*time.Minute)
	//limiter_medium := ratelimit_middleware.NewRateLimiter(10, 1*time.Minute)
	//limiter_hight := ratelimit_middleware.NewRateLimiter(10, 1*time.Minute)

	containerTest := source_middleware.Container("test")

	logger := log.NewLoggerBuilder().WithLogLevel(zerolog.DebugLevel).WithBufferSize(10000).WithRateLimit(1000).WithGroupWindow(2 * time.Second).WithOutput(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).Build()
	defer logger.Close()

	http.HandleFunc("GET /health/{$}", limiterLow.RateLimit(containerTest(metric_handlers.Health)))

	http.HandleFunc("POST /auth/login/{$}", limiterLow.RateLimit(auth_handlers.LoginHandler))
	http.HandleFunc("POST /auth/register/{$}", limiterLow.RateLimit(auth_handlers.RegisterHandler))

	logger.Info().Msg("Listening at http://localhost:" + os.Getenv("APP_PORT"))
	err := http.ListenAndServe(":"+os.Getenv("APP_PORT"), nil)
	if err != nil {
		return
	}
}
