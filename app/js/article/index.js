'use strict';

// DOM Tree の構築が完了したら処理を開始
document.addEventListener('DOMContentLoaded', () => {
  // DOM API を利用して HTML 要素を取得
  const deleteBtns = document.querySelectorAll('.articles__item-delete');

  // CSRF トークンを取得
  const csrfToken = document.getElementsByName('csrf')[0].content;

  // 記事を削除する関数を定義
  const deleteArticle = id => {
    let statusCode;

    // Fetch API を利用して削除リクエストを送信
    fetch(`/${id}`, {
      method: 'DELETE',
      headers: { 'X-CSRF-Token': csrfToken }
    })
      .then(res => {
        statusCode = res.status;
        return res.json();
      })
      .then(data => {
        console.log(JSON.stringify(data));
        if (statusCode == 200) {
          // 削除に成功したら画面から記事の HTML 要素を削除
          document.querySelector(`.articles__item-${id}`).remove();
        }
      })
      .catch(err => console.error(err));
  };

  // 削除ボタンそれぞれに対してイベントリスナーを設定
  for (let elm of deleteBtns) {
    elm.addEventListener('click', event => {
      event.preventDefault();

      // 削除ボタンのカスタムデータ属性からIDを取得して引数に渡す
      deleteArticle(elm.dataset.id);
    });
  }
});