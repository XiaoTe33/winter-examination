const account = document.querySelector('#account');
const password = document.querySelector('#psd');
const btn = document.querySelector('#btn')


btn.addEventListener('click',async()=>{
    if(account.value == ''|| account.value == undefined || account.value == null){
    alert('请输入账户名称')
    }else if(password.value == ''|| password.value == undefined || password.value == null){
    alert('请输入密码')
    }else{
        let fd = new FormData();
        fd.append("username",account.value);
        fd.append("password",password.value);
        const res = await fetch('http://39.101.72.18:9090/user/login', {
            method: "post",
            body: fd,
          })
          const data = await res.json()
          console.log(data);
    }
    
})