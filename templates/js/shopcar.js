var storage = window.localStorage;

const pic = document.querySelectorAll('.pic');
const name = document.querySelectorAll('.name');
const version = document.querySelectorAll('.version');
const price = document.querySelectorAll('.price');
const num = document.querySelectorAll('#num');
const sum = document.querySelectorAll('.sum');

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
                pic[i].children[0].src = data.data[i].pictureLink;
                name[i].innerHTML = data.data[i].name;
                //version[i].innerHTML = data.data[i].name;
                price[i].innerHTML = Number(data.data[i].price);
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







//全选 ok
let input = document.querySelectorAll('.check');
let yes = document.querySelectorAll('.yes');
let j=0;
for(let i = 0; i < yes.length ; i++){
    yes[i].addEventListener('click',function(){
        j++;
        if (  j % 2 != 0 ){
            for(let i=0;i<input.length;i++){
                input[i].checked = true;
            }
        }else{
            for(let i=0;i<input.length;i++){             
                input[i].checked = false;
            }
        }   
    });
}

//更新总价和总数
let totalPrice = document.querySelector('.totalprice');
let totalNum = document.querySelector('.totalnum');
for(let i = 0 ;i< input.length ;i++){  
    if( input[i].checked = true){
        console.log('999')
        totalNum.innerHTML = Number(totalNum.innerHTML) + Number(num[i]);
        totalPrice.innerHTML = Number(totalPrice.innerHTML) + Number(sum[i]);
    }
}


//删除ok
const del = document.querySelector('.delete');
del.addEventListener('click',async()=>{
    for(let i = 0; i<input.length ; i++){
        if( input[i].checked == true){
            let url = `http://39.101.72.18:9090/user/car/${Number(i)}`
            const res = await fetch(url,{
                method:"delete",
                headers : {"Token": storage.getItem("token")},
            })
            const data = await res.json();
            console.log(data);
            //console.log(url);
        }

    }
})

//去结算
const btn = document.querySelector('.totalbtn');
btn.addEventListener('click',async()=>{
    
})

//数量加减ok
const del1 = document.querySelectorAll('.del1');
const add1 = document.querySelectorAll('.add1');
for(let i = 0;i < input.length ; i++){
    del1[i].addEventListener('click',()=>{ 
        if(num[i].value > 1){
            num[i].value = num[i].value - 1;
        }
        if(num[i].value <= 1){
            console.log('111');
            //del1[i].style.cursor = 'not-allowed';
            del1[i].style.color = '#dedede';
        }
    })
add1[i].addEventListener('click',()=>{
    num[i].value = Number(num[i].value)+ 1;
    if(num[i].value > 1){
        console.log('111')
        //del1[i].style.cursor = 'point';
        del1[i].style.color = '#666666';
    }
})
}
