'use strict';

document.addEventListener('DOMContentLoaded', () => {
  // DOM APIを利用してHTML要素を取得する
  const inputs = document.getElementsByTagName('input');
  const form = document.forms.namedItem('article-form');
  const saveBtn = document.querySelector('.article-form__save');
  const cancelBtn = document.querySelector('.article-form__cancel');
  const previewOpenBtn = document.querySelector('.article-form__open-preview');
  const previewCloseBtn = document.querySelector('.article-form_close-preview');
  const articleFormBody = document.querySelector('.article-form__body')
  const articleFormPreview = document.querySelector('.article-form__preview');
  const articleFormBodyTextArea = document.querySelector('.article-form__input-body');
  const articleFormPreviewTextArea = document.querySelector('.article-form__preview-body-contents');
})

// 新規作成画面か編集画面かをURLから判定します
const mode = { method: "", url: "" };
if (window.location.pathname.endsWith("new")) {
  mode.method = 'Post'
  mode.url = '/'
} else if (window.location.pathname.endsWith("edit")) {
  mode.method = 'PATCH'
  mode.url = `/${window.location.pathname.split(' / ')[1]}`
}

const { method, url } = mode;

// input要素にフォーカスがあった状態でEnterが押されるとformが送信される
// ここではEnterキーでformが送信されないように挙動を制御する
for (let elm of inputs) {
  elm.addEventListener('keydown', event => {
    if (event.keyCode && event.keyCode == 13) {
      event.preventDefault();
      return false
    }
  })
}

// プレビューを開くイベントを設定する
previewOpenBtn.addEventListener('click', event => {
  //本文に入力された内容をプレビューにコピー
  articleFormPreviewTextArea.innerHTML = articleFormBody.TextArea.value

  // 入力フォームを非表示に、プレビューを表示する
  articleFormBody.style.display = 'none';
  articleFormPreview.style.display = 'grid';
});


// プレビューを閉じるイベントを設定
previewCloseBtn.addEventListener('click', event => {

  // 入力フォームを表示し、プレビューを非表示にする
  articleFormBody.style.display = 'grid';
  articleFormPreview.style.display = 'none';
});

cancelBtn.addEventListener('click', event => {
  // <button>要素のリンクの遷移機能を動作させず、指定したURLに遷移
  event.preventDefault();
  window.location.href = url;
})