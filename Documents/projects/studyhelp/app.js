const button = document.querySelector('#sidebar_logo');
button.addEventListener('click', ()=>{
    document.querySelector('#sidebar').classList.toggle('show');
})