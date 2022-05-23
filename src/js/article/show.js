'use strict'

// DOMツリーの構築が完了次第処理を開始
document.addEventListener('DOMContentLoaded', function () {

  // DOM APIを使用してHTML要素を取得
  const elm = document.getElementById('article-body');

  // カスタムデータ属性からMarkdown形式のテキストを取得
  // RemarkableでHTMLに変換して要素を追加
  elm.innerHTML = md.render(elm.dataset.markdown);
})