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
}
