const pic = document.querySelector('#pic');
const bigpic = document.querySelector('#bigpic');
const mask = document.querySelector('#mask');
const headline = document.querySelector('#headline')

//读取localstorage，渲染页面
window.addEventListener('load',async()=>{
    var storage = window.localStorage;
    const url = `http://39.101.72.18:9090/goods?goodsId=${storage.getItem("id")}`
    const res = await fetch(url, {
       method:"get"
           })
     const data = await res.json();
     console.log(data);
     console.log(data.data.pictureLink)
     pic.children[0].src = data.data.pictureLink;
     bigpic.children[0].src = data.data.pictureLink;
     headline.innerHTML = data.data.name;

})




//放大镜
pic.addEventListener('mouseenter',()=>{
    bigpic.style.display = 'block';
    mask.style.display = 'block';
})
pic.addEventListener('mouseleave',()=>{
    bigpic.style.display = 'none';
    mask.style.display = 'none';
})
pic.addEventListener('mousemove',(event)=>{
    let clientX=event.clientX;
    let clientY=event.clientY;
    let eleX=pic.getBoundingClientRect().left;
    let eleY=pic.getBoundingClientRect().top;

    let left=clientX-eleX-mask.offsetWidth/2;
    let top=clientY-eleY-mask.offsetHeight/2;

    if(left<=0){
        left=0;
    }else if(left>=pic.clientWidth-mask.offsetWidth){
        left=pic.clientWidth-mask.offsetWidth;
    }

    if(top<=0){
        top=0;
    }else if(top>pic.clientHeight-mask.offsetHeight){
        top=pic.clientHeight-mask.offsetHeight;
    }
    mask.style.left=left+'px';
    mask.style.top=top+'px';
    bigpic.scrollLeft=left*2;
    bigpic.scrollTop=top*2;
})


//加购数量
const num = document.querySelector('#num')
const del = document.querySelector('.del1');
const add = document.querySelector('.add1');
const buy = document.querySelector('.buyitnow');
const addtocar = document.querySelector('.addtocar');
del.addEventListener('click',()=>{ 
    if(num.value > 1){
        num.value = num.value - 1;
    }
    if(num.value <= 1){
        console.log('111');
        //del.style.cursor = 'not-allowed';
        del.style.color = '#dedede';
    }
})
add.addEventListener('click',()=>{
    num.value = Number(num.value)+ 1;
    if(num.value > 1){
        console.log('111')
        //del.style.cursor = 'point';
        del.style.color = '#666666';
    }
})
//立即购买，跳转订单页面
buy.addEventListener('click',async()=>{
    const fd = new FormData();
    fd.append("goodsId","张三");
    fd.append("amount",num.value);
    fd.append("style","23");
    fd.append("address","male");  
    console.log(fd.get('amount'));
    const res = await fetch('http://39.101.72.18:9090/order', {
        method: "post",
        body: fd,
      })
      const data = await res.json();
      console.log(data);
})
//加入购物车
addtocar.addEventListener('click',async()=>{
    
})

//选择款式
