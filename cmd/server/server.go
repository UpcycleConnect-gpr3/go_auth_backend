package server

import (
	"authentication_backend/app/handlers/auth_handlers"
	"authentication_backend/app/handlers/metric_handlers"
	"authentication_backend/app/handlers/totp_handlers"
	"authentication_backend/app/handlers/user_handlers"
	"authentication_backend/app/middleware/auth_middleware"
	"authentication_backend/app/middleware/ratelimit_middleware"
	"authentication_backend/app/middleware/source_middleware"
	"authentication_backend/config"
	"authentication_backend/var/database"
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
	config.InitEmail()
	config.InitFilesystem()

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

	isAuth := auth_middleware.IsAuth

	containerTest := source_middleware.Container("test")

	logger := log.NewLoggerBuilder().WithLogLevel(zerolog.DebugLevel).WithBufferSize(10000).WithRateLimit(1000).WithGroupWindow(2 * time.Second).WithOutput(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).Build()
	defer logger.Close()

	http.HandleFunc("GET /health/{$}", limiterLow.RateLimit(containerTest(metric_handlers.Health)))

	http.HandleFunc("POST /auth/login/{$}", limiterLow.RateLimit(auth_handlers.LoginHandler))
	http.HandleFunc("POST /auth/login-totp/{$}", limiterLow.RateLimit(auth_handlers.ValidateTOTPHandler))
	http.HandleFunc("POST /auth/register/{$}", limiterLow.RateLimit(auth_handlers.RegisterHandler))

	http.HandleFunc("GET /user/me/{$}", limiterLow.RateLimit(isAuth(auth_handlers.ShowMeHandler)))
	http.HandleFunc("GET /user/{$}", limiterLow.RateLimit(isAuth(user_handlers.IndexUserHandler)))
	http.HandleFunc("GET /user/{id}/{$}", limiterLow.RateLimit(isAuth(user_handlers.ShowUserHandler)))
	http.HandleFunc("PATCH /user/{$}", limiterLow.RateLimit(isAuth(user_handlers.UpdateUserHandler)))
	http.HandleFunc("DELETE /user/{$}", limiterLow.RateLimit(isAuth(user_handlers.DeleteUserHandler)))

	http.HandleFunc("POST /auth/totp/{$}", limiterLow.RateLimit(isAuth(totp_handlers.PostTOTP)))
	http.HandleFunc("GET /auth/totp/{$}", limiterLow.RateLimit(isAuth(totp_handlers.GetTOTP)))

	logger.Info().Msg("Listening at http://localhost:" + os.Getenv("APP_PORT"))
	err := http.ListenAndServe(":"+os.Getenv("APP_PORT"), corsMiddleware(http.DefaultServeMux))
	if err != nil {
		return
	}
}

// corsMiddleware enables cross-origin requests from the local dev frontends
// (Vite on a different port). Reflects the request Origin and answers the
// preflight OPTIONS so browser login/register calls are not blocked.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
