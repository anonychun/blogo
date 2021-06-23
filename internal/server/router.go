package server

import (
	"net/http"

	_ "github.com/anonychun/go-blog-api/docs"
	"github.com/anonychun/go-blog-api/internal/app/handler"
	"github.com/anonychun/go-blog-api/internal/app/repository"
	"github.com/anonychun/go-blog-api/internal/app/service"
	"github.com/anonychun/go-blog-api/internal/config"
	"github.com/anonychun/go-blog-api/internal/db"
	"github.com/anonychun/go-blog-api/internal/security/middleware"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go Blog API
// @version 1.0
// @description Implementing back-end services for blog application
// @BasePath /v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
func NewRouter(mysqlClient db.MysqlClient, redisClient db.RedisClient) *chi.Mux {
	router := chi.NewRouter()

	router.Use(httprate.LimitByIP(
		config.Cfg().HttpRateLimitRequest,
		config.Cfg().HttpRateLimitTime,
	))
	router.Use(cors.AllowAll().Handler)
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)

	accountRepository := repository.NewAccountRepository(mysqlClient, redisClient)
	postRepository := repository.NewPostRepository(mysqlClient, redisClient)

	authService := service.NewAuthService(accountRepository)
	accountService := service.NewAccountService(accountRepository)
	postService := service.NewPostService(postRepository)

	authHandler := handler.NewAuthHandler(authService)
	accountHandler := handler.NewAccountHandler(accountService)
	postHandler := handler.NewPostHandler(postService)

	router.Options("/*", func(w http.ResponseWriter, r *http.Request) {})
	api := router.Route("/v1", func(router chi.Router) {})

	api.Route("/accounts", func(r chi.Router) {
		r.Post("/auth", authHandler.Login())

		r.Post("/", accountHandler.Create())
		r.Get("/", accountHandler.List())
		r.Get("/{account_id}", accountHandler.Get())
		r.With(middleware.JWTVerifier).Put("/{account_id}", accountHandler.Update())
		r.With(middleware.JWTVerifier).Put("/{account_id}/password", accountHandler.UpdatePassword())
		r.With(middleware.JWTVerifier).Delete("/{account_id}", accountHandler.Delete())
	})

	api.Route("/posts", func(r chi.Router) {
		r.With(middleware.JWTVerifier).Post("/", postHandler.Create())
		r.Get("/", postHandler.List())
		r.Get("/{post_id}", postHandler.Get())
		r.With(middleware.JWTVerifier).Put("/{post_id}", postHandler.Update())
		r.With(middleware.JWTVerifier).Delete("/{post_id}", postHandler.Delete())
	})

	api.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	))

	return router
}
