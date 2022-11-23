'use strict';

// DOM Tree の構築が完了したら処理を開始
document.addEventListener('DOMContentLoaded', function() {
  // DOM API を利用して HTML 要素を取得
  const elm = document.getElementById('article-body');

  // カスタムデータ属性から Markdown 形式のテキストを取得、Remarkable で HTML に変換して要素に追加
  elm.innerHTML = md.render(elm.dataset.markdown);
});