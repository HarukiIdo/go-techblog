'use strict';

document.addEventListener('DOMContentLoaded', () => {

  const createBtn = document.querySelector('.page__new');
  const deleteBtn = document.querySelectorAll('.articles__item-delete');


  if (window.location.pathname.endsWith('/')) {
    createBtn.style.display = 'none';
    deleteBtn.forEach(elem => {
      elem.style.display = 'none';
    });
  }
});
