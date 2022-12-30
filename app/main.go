package main

import (
  "log"
  "os"

  "my-tech-blog/app/handler"
  "my-tech-blog/app/repository"

  _ "github.com/go-sql-driver/mysql" // Using MySQL driver
  "github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
  "gopkg.in/go-playground/validator.v9"
)

var db *sqlx.DB
var e = createMux()

func main() {
  db = connectDB()
  repository.SetDB(db)

  // TOP ページに記事の一覧を表示
  e.GET("/", handler.ArticleIndex)

  // 記事に関するページは "/articles" で開始する
  // 記事一覧画面には "/" と "/articles" の両方でアクセスできるようにする
  e.GET("/articles", handler.ArticleIndex)         // 一覧画面
  e.GET("/articles/new", handler.ArticleNew)       // 新規作成画面
  e.GET("/articles/:articleID", handler.ArticleShow)      // 詳細画面
  e.GET("/articles/:articleID/edit", handler.ArticleEdit) // 編集画面

  // HTML ではなく JSON を返却する処理は "/api" で開始する
  // 記事に関する処理なので "/articles"
  e.GET("/api/articles", handler.ArticleList)          // 一覧
  e.POST("/api/articles", handler.ArticleCreate)       // 作成
  e.DELETE("/api/articles/:articleID", handler.ArticleDelete) // 削除
  e.PATCH("/api/articles/:articleID", handler.ArticleUpdate)  // 更新

  e.Logger.Fatal(e.Start(":80"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
  e.Use(middleware.CSRF())

  e.Static("/css", "css")
  e.Static("/js", "js")

  e.Validator = &CustomValidator{validator: validator.New()}

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

// CustomValidator ...
type CustomValidator struct {
  validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
  return cv.validator.Struct(i)
}