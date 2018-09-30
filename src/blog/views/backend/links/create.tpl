<div style="margin: 15px">
    <form class="layui-form" action="">
        <input type="hidden" name="_xsrf" value="<<< .xsrf_token >>>">
        <div class="layui-form-item">
            <label class="layui-form-label">友链名称</label>
            <div class="layui-input-block">
                <input type="text" name="Name" lay-verify="required" placeholder="请输入链接名称" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">描述</label>
            <div class="layui-input-block">
                <textarea placeholder="请输入描述" name="Desc" class="layui-textarea"></textarea>
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">首图</label>
            <div class="layui-input-block">
                <a class="layui-btn select_img"><i class="layui-icon"></i> 选择图片</a>
                <div id="hidden_input">
                    <input type="hidden" name="Img" id="img" value="">
                </div>
                <div class="img_show"></div>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">链接</label>
            <div class="layui-input-block">
                <input type="text" name="Url" lay-verify="" placeholder="请输入链接" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">排序</label>
            <div class="layui-input-block">
                <input type="text" name="Sort" lay-verify="required" placeholder="请输入排序" autocomplete="off"
                       class="layui-input">
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
