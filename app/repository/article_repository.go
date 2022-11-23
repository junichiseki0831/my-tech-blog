package repository

import (
    "database/sql"
    "time"

	"my-tech-blog/model"
)

// ArticleList ...
func ArticleList() ([]*model.Article, error) {
  // 実行SQL
	query := `SELECT * FROM articles;`

  // データベースから取得した値を格納する変数を宣言
	var articles []*model.Article
  // Query を実行して、取得した値を変数に格納
	if err := db.Select(&articles, query); err != nil {
		return nil, err
	}

	return articles, nil
}

// ArticleCreate ...
func ArticleCreate(article *model.Article) (sql.Result, error) {
  // 現在日時を取得
  now := time.Now()

  // 構造体に現在日時を設定
  article.Created = now
  article.Updated = now

  // クエリ文字列を生成
  query := `INSERT INTO articles (title, body, created, updated)
  VALUES (:title, :body, :created, :updated);`

  // トランザクションを開始
  tx := db.MustBegin()

  // クエリ文字列と構造体を引数に渡して SQL を実行
  // クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換
  // 構造体タグで指定してあるフィールドが対象（`db:"title"` など）
  res, err := tx.NamedExec(query, article)
  if err != nil {
    // エラーが発生した場合はロールバック
    tx.Rollback()

    // エラー内容を返却
    return nil, err
  }

  // SQL の実行に成功した場合はコミット
  tx.Commit()

  // SQL の実行結果を返却
  return res, nil
}
