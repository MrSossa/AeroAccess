package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/MrSossa/AeroAccess/cmd/api/routes"
)

func main() {
	db, err := sql.Open("mysql", "admin:admin123@tcp(usuarios.cvyyymg2wll4.us-east-1.rds.amazonaws.com:3306)/usuarios")
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
