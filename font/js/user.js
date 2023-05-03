var u = {}


var has_login  = Boolean(0)
var image = ""
var username = ""

var career = ""
var email = ""
var phone = ""
var address = ""
var age = ""
var id =0
// 判断用户是否已登录
function setUser() {
    //login("jackjdeng","1")
    if (u==={}){
        return
    }

    // 创建一个 XMLHttpRequest 对象
    let xhr = new XMLHttpRequest();
    xhr.withCredentials = true;
    // 设置请求的方式和 URL
    xhr.open('GET','/webapp/user/info', false);
    // 监听请求的状态变化
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
                // 解析响应数据
                let resp = JSON.parse(xhr.responseText)
                if (resp.code === 200) {
                    if (resp.data) {
                        has_login = Boolean(1);
                        image = resp.data.avatar
                        username = resp.data.username
                        career = resp.data.career
                        phone = resp.data.phone
                        email = resp.data.email
                        address = resp.data.addr
                        age = resp.data.age
                        id = resp.data.id
                        //保存到u
                        u=resp
                    }
                } else {
                    image = "https://cdn.redrock.team/web_app_default.jpeg"
                    username = "未登录"
                }
            } else {
                console.error('请求出错，状态码为：' + xhr.status);
            }
        }
    };
    // 发送请求
    xhr.send();
}

function login(username, password) {
    let xhr = new XMLHttpRequest();

    let data = {
        username: username,
        password: password
    };

    let json = JSON.stringify(data);
    xhr.withCredentials = true;
    xhr.open("POST",  '/webapp/user/login', false);
    xhr.setRequestHeader("Content-type", "application/json;charset=UTF-8");

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            let responseData = JSON.parse(xhr.responseText);
            if (responseData.code===200){
                alert('登陆成功');
                window.location.href = 'home.html';
            }else {

                alert('账号或密码错误');
            }
        }else {
            alert(xhr.status);
            alert('后台服务故障');
        }
    };
    xhr.send(json);
}

function logout() {
    let xhr = new XMLHttpRequest();
    xhr.withCredentials = true;
    xhr.open("POST",  '/webapp/user/logout', true);
    xhr.send();
    has_login=Boolean(0);
}

function register(username, password) {
    let xhr = new XMLHttpRequest();

    let data = {
        username: username,
        password: password
    };

    let json = JSON.stringify(data);
    xhr.withCredentials = true;
    xhr.open("POST",  '/webapp/user/register', false);
    xhr.setRequestHeader("Content-type", "application/json;charset=UTF-8");

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var responseData = JSON.parse(xhr.responseText);
            if (responseData.code===200){
                alert('注册成功,请前往登陆');
                window.location.href = 'home.html';
            }else {
                alert(responseData.data);
            }
        }else {
            alert('后台服务故障');
        }
    };
    xhr.send(json);
}
