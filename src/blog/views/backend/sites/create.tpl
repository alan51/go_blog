<div style="margin: 15px">
    <form class="layui-form" action="">
        <input type="hidden" name="_xsrf" value="<<< .xsrf_token >>>">
        <div class="layui-form-item">
            <label class="layui-form-label">网站标识</label>
            <div class="layui-input-block">
                <input type="text" name="Key" lay-verify="required" placeholder="请输入网站标识" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">标识值</label>
            <div class="layui-input-block">
                <textarea placeholder="请输入标识值" name="Value" class="layui-textarea"></textarea>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">状态</label>
            <div class="layui-input-block">
                <input type="radio" name="Status" value="1" title="启用" checked="">
                <input type="radio" name="Status" value="2" title="禁用">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">类型</label>
            <div class="layui-input-block">
                <input type="radio" name="Type" value="1" title="前台" checked="">
                <input type="radio" name="Type" value="2" title="后台">
            </div>
        </div>
        <button lay-filter="edit" lay-submit style="display: none;"></button>
    </form>
</div>
