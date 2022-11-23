package model

import (
  "time"

  "gopkg.in/go-playground/validator.v9"
)

// Article ...
type Article struct {
  ID      int       `db:"id" form:"id" json:"id"`
  Title   string    `db:"title" form:"title" validate:"required,max=50" json:"title"`
  Body    string    `db:"body" form:"body" validate:"required" json:"body"`
  Created time.Time `db:"created" json:"created"`
  Updated time.Time `db:"updated" json:"updated"`
}

// ValidationErrors ...
func (a *Article) ValidationErrors(err error) []string {
  // メッセージを格納するスライスを宣言
  var errMessages []string

  // 複数のエラーが発生する場合があるのでループ処理を行う
  for _, err := range err.(validator.ValidationErrors) {
    // メッセージを格納する変数を宣言
    var message string

    // エラーになったフィールドを特定する
    switch err.Field() {
    case "Title":
      // エラーになったバリデーションルールを特定する
      switch err.Tag() {
      case "required":
        message = "タイトルは必須です。"
      case "max":
        message = "タイトルは最大50文字です。"
      }
    case "Body":
      message = "本文は必須です。"
    }

    // メッセージをスライスに追加する
    if message != "" {
      errMessages = append(errMessages, message)
    }
  }

  return errMessages
}
