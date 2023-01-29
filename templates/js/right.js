let icon_rmb = document.querySelector('.icon-rmb');
let rmb = document.querySelector('.rmb');
let history = document.querySelector('.icon-history');
let qiandao = document.querySelector('.icon-xianxing-15');
let hongbao = document.querySelector('.icon-hongbao');
   
icon_rmb.addEventListener('mouseenter',()=>{
    console.log('111')
    rmb.classList.add('moved')
})