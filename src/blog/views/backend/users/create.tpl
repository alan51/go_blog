<div style="margin: 15px">
    <form class="layui-form" action="">
        <input type="hidden" name="_xsrf" value="<<< .xsrf_token >>>">
        <div class="layui-form-item">
            <label class="layui-form-label">登录名</label>
            <div class="layui-input-block">
                <input type="text" name="UserName" lay-verify="required" placeholder="请输入登录名" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">密码</label>
            <div class="layui-input-block">
                <input type="text" name="Password" lay-verify="required" placeholder="请输入密码" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">昵称</label>
            <div class="layui-input-block">
                <input type="text" name="NickName" lay-verify="required" placeholder="请输入昵称" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">头像</label>
            <div class="layui-input-block">
                <a class="layui-btn select_img"><i class="layui-icon"></i> 选择图片</a>
                <div id="hidden_input">
                    <input type="hidden" name="Img" id="img" value="">
                </div>
                <div class="img_show"></div>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">状态</label>
            <div class="layui-input-block">
                <input type="radio" name="Status" value="1" title="启用" checked="">
                <input type="radio" name="Status" value="2" title="禁用">
            </div>
        </div>
        <button lay-filter="edit" lay-submit style="display: none;"></button>
    </form>
</div>
<script type="text/javascript" src="/static/js/jquery.min.js"></script>
<script>
    $(function(){
        $('.select_img').on("click", function(){

            layui.layer.open({
                type: 2,
                area: ['700px', '450px'],
                fixed: false, //不固定
                maxmin: true,
                zIndex: 19950925,
                content: '/admin/file/upload'
            });
        })
    })
</script>
