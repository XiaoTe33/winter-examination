// const leftbtn = document.querySelector('.zuo');
// const rightbtn = document.querySelector('.you');
// const images = document.querySelector('.images');
//console.log(images.children[0])

// const timer = setInterval(function () {
//         rightbtn.click();
//     }, 10000);
    
// images.addEventListener('mouseenter', ()=> {
//         leftbtn.style.display = 'block';
//         rightbtn.style.display = 'block';
//         clearInterval(timer);
//         timer = null;//清除计时器
//     });
// images.addEventListener('mouseleave',()=> {
//         leftbtn.style.display = 'none';
//         rightbtn.style.display = 'none';
//         timer = setInterval(function () {
//             rightbtn.click();
//         }, 10000);
//     });


// var num = 1;
// //点击左右按钮
// rightbtn.addEventListener('click', function () {
//     // 最后一张图片
//     if (num == images.children.length) {
//             num = 0;
//     }
//     num++;
//     //排他清除一下   
//     for (let i = 0; i < dots.children.length; i++) {
//         dots.children[i].className = '';
//         images.children[i].classList.remove = 'current';
//     }
//     dots.children[num-1].className = 'active';
//     images.children[num-1].classList.add('current');
// })

// leftbtn.addEventListener('click',()=> {
//     if (num == 0) {
//          num = images.children.length;
//      }
//     num--;
//     for (let i = 0; i < dots.children.length; i++) {
//         dots.children[i].className = '';
//         images.children[i].classList.remove = 'current';
//     }
//     dots.children[num].className = 'active';
//     images.children[num].classList.add('current');
//     })



const example = document.querySelectorAll('.more');
console.log(example);
example[0].id = 3055483356794;
console.log(example[0].id)
for( let i = 0; i < example.length; i++){
    example[i].addEventListener('click',()=>{
        let storage = window.localStorage;
        storage.setItem('id',example[i].id)
    })
}