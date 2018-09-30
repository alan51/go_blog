<div style="margin: 15px">
    <form class="layui-form" action="">
        <input type="hidden" name="_xsrf" value="<<< .xsrf_token >>>">
        <div class="layui-form-item">
            <label class="layui-form-label">分类名称</label>
            <div class="layui-input-block">
                <input type="text" name="Name" lay-verify="required" placeholder="请输入分类名称" autocomplete="off"
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
            <label class="layui-form-label">关键字</label>
            <div class="layui-input-block">
                <input type="text" name="Keywords" lay-verify="required" placeholder="请输入关键字" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">图标</label>
            <div class="layui-input-block">
                <input type="text" name="Icon" lay-verify="required" placeholder="请输入图标库标识" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">自定义链接</label>
            <div class="layui-input-block">
                <input type="text" name="Url" lay-verify="" placeholder="请输入自定义链接" autocomplete="off"
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
            <label class="layui-form-label">父分类</label>
            <div class="layui-input-block">
                <select name="ParentId" class="form-control select">
                    <option value="">顶级</option>
                    <<< range $category := .category>>>
                    <option value="<<< $category.Id >>>"><<<$category.Name>>></option>
                    <<< end >>>
                </select>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">顶部菜单</label>
            <div class="layui-input-block">
                <input type="radio" name="IsTop" value="1" title="是" checked="">
                <input type="radio" name="IsTop" value="2" title="否">
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
