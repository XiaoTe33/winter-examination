const account = document.querySelector('#account');
const password = document.querySelector('#psd');
const btn = document.querySelector('#btn')
const form = document.querySelector('#form1');

form.addEventListener('submit',async()=>{
    if(account.value == ''|| account.value == undefined || account.value == null){
    alert('请输入账户名称')
    }else if(password.value == ''|| password.value == undefined || password.value == null){
    alert('请输入密码')
    }else{
        e.preventDefault()
        const fd = new FormData()
            formdata.append('username',account.value)
            formdata.append('password',password.value)
        const xhr = new XMLHttpRequest()
        xhr.open('POST','http://39.101.72.18:9090/user/login');
        xhr.setRequestHeader('Content-Type','application/x-www-form-urlencoded')
        xhr.send(fd)
        
        xhr.onreadystatechange = function(){
            if(xhr.readyState === 4){
                if(xhr.status === 200){
                    console.log(JSON.parse(xhr.responseText))
                }
            }
        }
        // fetch('http://39.101.72.18:9090/user/login', {
        //     method: "post",
        //     headers: {
        //         'Content-Type':'application/x-www-form-urlencoded'
        //     },
        //     body: JSON.stringify({
        //       username : account.value,
        //       password : password.value
        //     })
        //   })
        //   const data = await res.json()
    }
})