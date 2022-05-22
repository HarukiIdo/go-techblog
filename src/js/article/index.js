'use strict'

// DOMツリーの構築が完了したら処理を開始
document.addEventListener('DOMContentLoaded', () => {

  const deleteBtns = document.querySelectorAll('.articles__item-delete')
  const csrfToken = document.getElementsByName('csrf')[0].content;

  // 記事を削除する関数
  const deleteArticle = id => {
    let statusCode;

    // Fetch APIを使って削除リクエストを送信
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

        // 削除に成功したら、記事のHTML要素を削除
        if (statusCode == 200) {
          document.querySelector(`.articles__item-${id}`).remove();
        }
      })
      .catch(err => console.error(err))
  };

  // 削除ボタンをそれぞれに対してイベントリスナーを設定
  for (let elm of deleteBtns) {
    elm.addEventListener('click', event => {
      event.preventDefault();

      // 削除ボタンのカスタムデータ属性からIDを取得し引数に渡す
      deleteArticle(elm.getAttribute('data-id'));
    });
  }
});