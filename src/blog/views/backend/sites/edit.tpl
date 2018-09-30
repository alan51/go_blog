<div style="margin: 15px">
    <form class="layui-form" action="">
        <input type="hidden" name="_xsrf" value="<<< .xsrf_token >>>">
        <input type="hidden" name="Id" value="<<< .site_info.Id >>>">
        <div class="layui-form-item">
            <label class="layui-form-label">网站标识</label>
            <div class="layui-input-block">
                <input type="text" name="Key" value="<<< .site_info.Key >>>" lay-verify="required" placeholder="请输入网站标识" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">标识值</label>
            <div class="layui-input-block">
                <textarea placeholder="请输入标识值" name="Value" class="layui-textarea"> <<< .site_info.Value >>> </textarea>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">状态</label>
            <div class="layui-input-block">
                <input type="radio" name="Status" value="1"  <<< if eq .site_info.Status 1>>> checked="" <<< end >>> title="启用">
                <input type="radio" name="Status" value="2"  <<< if eq .site_info.Status 2>>> checked="" <<< end >>> title="禁用">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">类型</label>
            <div class="layui-input-block">
                <input type="radio" name="Type" value="1"  <<< if eq .site_info.Type 1>>> checked="" <<< end >>> title="前台">
                <input type="radio" name="Type" value="2"  <<< if eq .site_info.Type 2>>> checked="" <<< end >>> title="后台">
            </div>
        </div>
        <button lay-filter="edit" lay-submit style="display: none;"></button>
    </form>
</div>
