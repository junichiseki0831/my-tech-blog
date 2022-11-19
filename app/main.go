package main

import (
  "log"
  "os"

  "my-tech-blog/handler"
  "my-tech-blog/repository"

  _ "github.com/go-sql-driver/mysql" // Using MySQL driver
  "github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sqlx.DB
var e = createMux()

func main() {
  db = connectDB()
  repository.SetDB(db)

  e.GET("/", handler.ArticleIndex)
  e.GET("/new", handler.ArticleNew)
  e.GET("/:id", handler.ArticleShow)
  e.GET("/:id/edit", handler.ArticleEdit)

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

  e.Static("/css", "css")
  e.Static("/js", "js")

	return e
}

func connectDB() *sqlx.DB {
  dsn := os.Getenv("DSN")
  db, err := sqlx.Open("mysql", dsn)
  if err != nil {
      e.Logger.Fatal(err)
  }
  if err := db.Ping(); err != nil {
      e.Logger.Fatal(err)
  }
  log.Println("db connection succeeded")
  return db
}
