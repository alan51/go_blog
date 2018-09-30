<!DOCTYPE html>

<html>

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <title>登录</title>
    <link rel="stylesheet" href="/static/plugins/layui/css/layui.css" media="all" />
    <link rel="stylesheet" href="/static/css/login.css" />
</head>

<body class="beg-login-bg">
<div class="beg-login-box">
    <header>
        <h1>后台登录</h1>
    </header>
    <div class="beg-login-main">
        <form action="/admin/login" class="layui-form" method="post">
            <input type="hidden" name="_xsrf" value="<<< .xsrf_token >>>">
            <div class="layui-form-item">
                <label class="beg-login-icon">
                    <i class="layui-icon">&#xe612;</i>
                </label>
                <input type="text" name="username" lay-verify="userName" autocomplete="off" placeholder="这里输入登录名" class="layui-input">
            </div>
            <div class="layui-form-item">
                <label class="beg-login-icon">
                    <i class="layui-icon">&#xe642;</i>
                </label>
                <input type="password" name="password" lay-verify="password" autocomplete="off" placeholder="这里输入密码" class="layui-input">
            </div>
            <div class="layui-form-item">
                <!-- <div class="beg-pull-left beg-login-remember">
                    <label>记住帐号？</label>
                    <input type="checkbox" name="rememberMe" value="true" lay-skin="switch" checked title="记住帐号">
                </div> -->
                <div class="beg-pull-center">
                    <button class="layui-btn layui-btn-primary" lay-submit lay-filter="login">
                        <i class="layui-icon">&#xe650;</i> 登录
                    </button>
                </div>
                <div class="beg-clear"></div>
            </div>
        </form>
    </div>
    <footer>
        <p>后台管理系统 © www.aicyun.com</p>
    </footer>
</div>
<script type="text/javascript" src="/static/plugins/layui/layui.js"></script>
<script>
    layui.config({
        base: '/static/js/',
        v: new Date().getTime()
    }).use(['layer', 'form', 'baajax'], function() {
        var layer = layui.layer,
            $ = layui.jquery,
            form = layui.form(),
            baajax = layui.baajax
        ;

        form.on('submit(login)',function(data){
            baajax.post('/public/login', data.field, function(ret){
                if (ret.code == 200) {
                    layer.alert(ret.msg, {
                        icon: 1,
                        skin: 'layer-ext-moon' //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
                    })
                    window.location.href = '/admin/'
                } else {
                    layer.alert(ret.msg, {
                        icon: 2,
                        skin: 'layer-ext-moon' //该皮肤由layer.seaning.com友情扩展。关于皮肤的扩展规则，去这里查阅
                    })
                }
            })
            //location.href='index.html';
            return false;
        });
    });
</script>
</body>

</html>