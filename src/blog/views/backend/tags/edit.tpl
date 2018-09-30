<div style="margin: 15px">
    <form class="layui-form" action="">
        <input type="hidden" name="_xsrf" value="<<< .xsrf_token >>>">
        <input type="hidden" name="Id" value="<<< .tag_info.Id >>>">
        <div class="layui-form-item">
            <label class="layui-form-label">标签名称</label>
            <div class="layui-input-block">
                <input type="text" name="Name" value="<<< .category_info.Name >>>" lay-verify="required" placeholder="请输入标签名称" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <button lay-filter="edit" lay-submit style="display: none;"></button>
    </form>
</div>
