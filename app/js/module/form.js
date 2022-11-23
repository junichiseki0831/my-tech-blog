// DOM Tree の構築が完了したら処理を開始
document.addEventListener('DOMContentLoaded', () => {
  // HTML 要素を取得
  const inputs = document.getElementsByTagName('input');
  const form = document.forms.namedItem('article-form');
  const saveBtn = document.querySelector('.article-form__save');
  const cancelBtn = document.querySelector('.article-form__cancel');
  const previewOpenBtn = document.querySelector('.article-form__open-preview');
  const previewCloseBtn = document.querySelector('.article-form__close-preview');
  const articleFormBody = document.querySelector('.article-form__body');
  const articleFormPreview = document.querySelector('.article-form__preview');
  const articleFormBodyTextArea = document.querySelector('.article-form__input--body');
  const articleFormPreviewTextArea = document.querySelector('.article-form__preview-body-contents');

  // 新規作成画面か編集画面かを URL から判定
  const mode = { method: '', url: '' };
  if (window.location.pathname.endsWith('new')) {
    // 新規作成時の HTTP メソッドは POST
    mode.method = 'POST';
    // 作成リクエスト、および戻るボタンの遷移先のパスは "/"
    mode.url = '/';
  } else if (window.location.pathname.endsWith('edit')) {
    // 更新時の HTTP メソッドは PATCH を利用
    mode.method = 'PATCH';
    // 更新リクエスト、および戻るボタンの遷移先のパスは "/:articleID"
    mode.url = `/${window.location.pathname.split('/')[1]}`;
  }
  const { method, url } = mode;

  // input 要素にフォーカスが合った状態で Enter が押されると form が送信
  // Enter キーで form が送信されないように挙動を制御
  for (let elm of inputs) {
    elm.addEventListener('keydown', event => {
      if (event.keyCode && event.keyCode === 13) {
        // キーが押された際のデフォルトの挙動をキャンセル
        event.preventDefault();

        // 何もせず処理を終了
        return false;
      }
    });
  }

  // プレビューを開くイベントを設定
  previewOpenBtn.addEventListener('click', event => {
    // form の「本文」に入力された Markdown を HTML に変換してプレビューに埋め込み
    articleFormPreviewTextArea.innerHTML = md.render(articleFormBodyTextArea.value);

    // 入力フォームを非表示
    articleFormBody.style.display = 'none';

    // プレビューを表示
    articleFormPreview.style.display = 'grid';
  });

  // プレビューを閉じるイベントを設定
  previewCloseBtn.addEventListener('click', event => {
    // 入力フォームを表示
    articleFormBody.style.display = 'grid';

    // プレビューを非表示
    articleFormPreview.style.display = 'none';
  });

  // 前のページに戻るイベントを設定
  cancelBtn.addEventListener('click', event => {
    // <button> 要素クリック時のデフォルトの挙動をキャンセル
    event.preventDefault();

    // URL を指定して画面を遷移
    window.location.href = url;
  });
});