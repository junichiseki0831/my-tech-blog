package repository

import (
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