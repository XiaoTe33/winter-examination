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
const phonenum = document.querySelector('.phonenum');
const example = document.querySelectorAll('.more');
const header = document.querySelector('.header2');
var storage = window.localStorage;
window.addEventListener('load',async()=>{
    //console.log(storage.getItem("token"));
    let url = `http://39.101.72.18:9090/user/info`
    const res = await fetch(url, {
      method : 'get',
      headers : {"Token": storage.getItem("token")}
    })
     let data = await res.json();
     console.log(data);
     if( data.status == 200){
       phonenum.innerText = data.data.phone;
       phonenum.classList.remove('none');
       header.children[0].classList.add('none');
       header.children[1].classList.add('none');
     }
     
})



setInterval(async()=>{
  console.log('111')
  const fd = new FormData()
  fd.append("token",storage.getItem("token"))
  const res = await fetch('http://39.101.72.18:9090/user/login/token', {
      method : 'post',
      body: fd,
  })
  let data = await res.json();
  console.log(data);
  if(data.status == 200){
      storage.setItem("token",data.token)
  }
},2000000)

// console.log(example);
// example[0].id = 3055483356794;
// console.log(example[0].id)

// const imgs = document.querySelectorAll('.img');
// const left = document.querySelector('.zuo');
// const right = document.querySelector('.you');
// const dots = document.querySelector('.dots').querySelector('li');

 
//   // // 鼠标移动到左右箭头的位置更换图片 有灰色背景的图片
//   // leftArrow.addEventListener('mouseenter', function() {
//   //     this.style.backgroundPosition = '0 0';
//   // });

//   // leftArrow.addEventListener('mouseleave', function() {
//   //     this.style.backgroundPosition = '-83px 0';
//   // });

//   // rightArrow.addEventListener('mouseenter', function() {
//   //     this.style.backgroundPosition = '-42px 0';
//   // });

//   // rightArrow.addEventListener('mouseleave', function() {
//   //     this.style.backgroundPosition = '-123px 0';
//   // });


//   // 给图片设置index 属性，好判断当前的图片是哪一张
//   for (let i = 0; i < imgs.length; i++) {
//       imgs[i].setAttribute('data-index', i);
//   }

//   // 获取当前图片 和 图片的index（数组下标）
//   var current = this.document.querySelector('.current');
//   var currentIndex = current.getAttribute('data-index');

//   // 左箭头的点击事件，点击了就返回前一张图片 
//   // 如果当前图片为第一张那么需要更换到最后一张图片也就是第四张
//   left.addEventListener('click', function() {
//       if (currentIndex > 0) {
//           imgs[currentIndex].classList.remove('current');
//           dots[currentIndex].classList.remove('active');
//           imgs[--currentIndex].classList.add('current');
//           dots[currentIndex].classList.add('active');
//       } else {
//           imgs[currentIndex].classList.remove('current');
//           dots[currentIndex].classList.remove('active');
//           currentIndex = 4;
//           imgs[currentIndex].classList.add('current');
//           dots[currentIndex].classList.add('active');
//       }
//   });

//   // 点击右箭头下一张图片切换
//   // 如果当前为第五张图片，那么切换回第一张图片
//   right.addEventListener('click', changeImage);

//   var timer = this.setInterval(changeImage, 8000);

//   function changeImage() {
//       if (currentIndex < 4) {
//           imgs[currentIndex].classList.remove('current');
//           dots[currentIndex].classList.remove('active');
//           imgs[++currentIndex].classList.add('current');
//           dots[currentIndex].classList.add('active');
//       } else {
//           imgs[currentIndex].classList.remove('current');
//           dots[currentIndex].classList.remove('active');
//           currentIndex = 0;
//           imgs[currentIndex].classList.add('current');
//           dots[currentIndex].classList.add('active');
//       }
//   };

//   // 小圆点的点击事件
//   for (let k = 0; k < dots.length; k++) {
//       dots[k].setAttribute('data-index', k);
//       dots[k].addEventListener('click', function() {
//           var index = this.getAttribute('data-index');
//           if (index != currentIndex) {
//               imgs[currentIndex].classList.remove('current');
//               dots[currentIndex].classList.remove('active');
//               imgs[index].classList.add('current');
//               dots[index].classList.add('active');
//               currentIndex = index;
//           }

//       })
//   }





// //注册事件

for( let i = 0; i < example.length; i++){
    example[i].id = '3055483356794';
    example[i].addEventListener('click',()=>{
        let storage = window.localStorage;
        storage.setItem('id',example[i].id)
    })
}

//退出登录
const signOut = document.querySelector('.sign-out');
signOut.addEventListener('click',async()=>{
  const res = await fetch('http://39.101.72.18:9090/user/info', {
    method : 'get',
    headers : {"Token": "123456789"}
  })
   let data = await res.json();
   console.log(data);
   if( data.status != 200){
    console.log('333')
     //phonenum.innerText = data.data.phone;
     phonenum.classList.add('none');
     header.children[0].classList.remove('none');
     header.children[1].classList.remove('none');
     storage.setItem("token","123456789");
   }
})
