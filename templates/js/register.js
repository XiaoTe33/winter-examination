alert("注册协议[审慎阅读]您在申请注册流程中点击同意前，应当认真阅读以下协议。请您务必审慎阅读、充分理解协议中相关条款内容,其中包括:1.与您约定免除或限制责任的条款:2.与您约定法律适用和管辖的条款:3.其他以粗体下划线标识的重要条款。如您对协议有任何疑问，可向平台客服咨询。[特别提示]当您按照注册页面提示填写信息阅读并同意协议且完成全部注册程序后，即表示您已充分阅读、理解并接受协议的全部内容。阅读协议的过程中,如果您不同意相关协议或其中任何条款约定,您应立即停止注册程序。《苏宁会员章程》《易付宝协议》《苏宁隐私政策》《易付宝隐私政策》《苏宁联盟在线协议》");


const btn=document.querySelector('#submit');
btn.addEventListener('click',function(){
    const xhr = new XMLHttpRequest();
xhr.open("post", "http://39.101.72.18:9090/user/register",true);
//xhr.setRequestHeader("", "");// 设置请求头
xhr.send();//发送请求，可携带参数
xhr.onreadystatechange = () => {
  if (xhr.readyState == 4) {
    if (xhr.status == 200) {//返回状态码
      var data = xhr.responseTEXT;
      return data;
    }
  }
}
})