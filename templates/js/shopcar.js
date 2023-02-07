var storage = window.localStorage;
let details = document.querySelector('.details');
const del = document.querySelector('.delete');
let yes = document.querySelector('.yes');
let totalPrice = document.querySelector('.totalprice');
let totalNum = document.querySelector('.totalnum');

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
        const res = await fetch('http://39.101.72.18:9090/user/car', {
            method : 'get',
            headers : {"Token": storage.getItem("token")}
          })
           let data = await res.json();
           console.log(data);
           if(data.status == 200){
            for(let i = 0; i<data.data.length; i++){
                let ul = document.createElement('ul');
                ul.classList.add('flex');
                details.appendChild(ul);
                for( let j=1;j<=8;j++){
                    let li = document.createElement('li');
                    var classname = `list${j}`;
                    li.classList.add(classname);
                    ul.appendChild(li);
                    if(j==6){
                        li.classList.add('flex');
                        li.classList.add('num');
                        let btn1 = document.createElement('button');
                        li.appendChild(btn1);
                        btn1.classList.add('del1');
                        btn1.classList.add('btn');
                        btn1.innerHTML = '-'
                        let input = document.createElement('input');
                        input.classList.add('num1');
                        input.setAttribute("type","text")
                        input.value = 1;
                        li.appendChild(input);
                        let btn2 = document.createElement('button');
                        btn2.classList.add('add1');
                        btn2.classList.add('btn');
                        btn2.innerHTML = '+'
                        li.appendChild(btn2);
                        //在生成的同时绑定事件监听
                        btn1.addEventListener('click',()=>{ 
                                if(input.value > 1){
                                    input.value = input.value - 1;
                                }
                                if(input.value <= 1){
                                    console.log('111');
                                    //del1[i].style.cursor = 'not-allowed';
                                    btn1.style.color = '#dedede';
                                }
                            })
                        btn2.addEventListener('click',()=>{
                            input.value = Number(input.value)+ 1;
                            if(input.value > 1){
                                console.log('111')
                                //del1[i].style.cursor = 'point';
                                btn1.style.color = '#666666';
                            }
                        })
                        

                    }
                    if( j==1){
                        let input = document.createElement('input');
                        input.setAttribute("type","checkbox");
                        input.classList.add('check');
                        li.appendChild(input);
                        del.addEventListener('click',async()=>{
                                if( input.checked == true){
                                    let url = `http://39.101.72.18:9090/user/car/${Number(i)}`
                                    const res = await fetch(url,{
                                        method:"delete",
                                        headers : {"Token": storage.getItem("token")},
                                    })
                                    const data = await res.json();
                                    console.log(data);
                                    //console.log(url);
                                }
                        })
                        let m = 0;
                        yes.addEventListener('click',function(){
                            m++;
                            if (  m % 2 != 0 ){                               
                                    input.checked = true;
                            }else{            
                                    input.checked = false;
                            }   
                        });
                        // if( input.checked = true){
                        //     console.log('999')
                        //     totalNum.innerHTML = Number(totalNum.innerHTML) + Number(num[i]);
                        //     totalPrice.innerHTML = Number(totalPrice.innerHTML) + Number(sum[i]);
                        // }
                    }
                    if(j==2){
                        li.classList.add('pic');
                        let img = document.createElement('img');
                        li.appendChild(img);
                        img.src = data.data[i].pictureLink
                    }
                    if(j==3){
                        li.classList.add('name');
                        li.innerHTML = data.data[i].name;
                    }
                    if(j==4){
                        li.classList.add('version');
                    }
                    if(j==5){
                        li.classList.add('price');
                        li.innerHTML = Number(data.data[i].price);
                    }
                    if(j==7){li.classList.add('sum')}
                    if(j==8){
                        li.innerHTML = '移入关注'
                    }
                }
            }
           }
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
        storage.setItem("token",data.token);
    }
},2000000)

const pic = document.querySelectorAll('.pic');
const name = document.querySelectorAll('.name');
const version = document.querySelectorAll('.version');
const price = document.querySelectorAll('.price');
const num = document.querySelectorAll('.num1');
const sum = document.querySelectorAll('.sum');


// //全选 ok
let input = document.querySelectorAll('.check');
// let yes = document.querySelectorAll('.yes');
// let j=0;
// for(let i = 0; i < yes.length ; i++){
//     yes[i].addEventListener('click',function(){
//         j++;
//         if (  j % 2 != 0 ){
//             for(let i=0;i<input.length;i++){
//                 input[i].checked = true;
//             }
//         }else{
//             for(let i=0;i<input.length;i++){             
//                 input[i].checked = false;
//             }
//         }   
//     });
// }

//更新总价和总数

// for(let i = 0 ;i< input.length ;i++){  
//     if( input[i].checked = true){
//         console.log('999')
//         totalNum.innerHTML = Number(totalNum.innerHTML) + Number(num[i]);
//         totalPrice.innerHTML = Number(totalPrice.innerHTML) + Number(sum[i]);
//     }
// }


//删除ok
// del.addEventListener('click',async()=>{
//     for(let i = 0; i<input.length ; i++){
//         if( input[i].checked == true){
//             let url = `http://39.101.72.18:9090/user/car/${Number(i)}`
//             const res = await fetch(url,{
//                 method:"delete",
//                 headers : {"Token": storage.getItem("token")},
//             })
//             const data = await res.json();
//             console.log(data);
//             //console.log(url);
//         }

//     }
// })

//去结算
const btn = document.querySelector('.totalbtn');
btn.addEventListener('click',async()=>{
    
})

// //数量加减ok
// const del1 = document.querySelectorAll('.del1');
// const add1 = document.querySelectorAll('.add1');
// for(let i = 0;i < input.length ; i++){
//     del1[i].addEventListener('click',()=>{ 
//         if(num[i].value > 1){
//             num[i].value = num[i].value - 1;
//         }
//         if(num[i].value <= 1){
//             console.log('111');
//             //del1[i].style.cursor = 'not-allowed';
//             del1[i].style.color = '#dedede';
//         }
//     })
// add1[i].addEventListener('click',()=>{
//     num[i].value = Number(num[i].value)+ 1;
//     if(num[i].value > 1){
//         console.log('111')
//         //del1[i].style.cursor = 'point';
//         del1[i].style.color = '#666666';
//     }
// })
// }
