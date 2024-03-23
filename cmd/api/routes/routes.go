package routes

import (
	"database/sql"

	"github.com/MrSossa/AeroAccess/cmd/api/handler"
	"github.com/MrSossa/AeroAccess/internal/user"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{
		r:  r,
		db: db,
	}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildPing()
	r.buildUserRoutes()
}

func (r *router) setGroup() {
	r.r.Use(CORSMiddleware())
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildPing() {
	r.r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	},
	)
}

func (r *router) buildUserRoutes() {
	repo := user.NewUserRepository(r.db)
	service := user.NewUserService(repo)
	handler := handler.NewUser(service)
	r.rg.POST("/login", handler.Login)
	r.rg.POST("/register", handler.SaveUser)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
