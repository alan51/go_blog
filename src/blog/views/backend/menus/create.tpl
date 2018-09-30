<div style="margin: 15px">
    <form class="layui-form" action="">
        <input type="hidden" name="_xsrf" value="<<< .xsrf_token >>>">
        <div class="layui-form-item">
            <label class="layui-form-label">菜单名称</label>
            <div class="layui-input-block">
                <input type="text" name="Name" lay-verify="required|name" placeholder="请输入菜单名称" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">描述</label>
            <div class="layui-input-block">
                <textarea placeholder="请输入描述" name="Desc" class="layui-textarea"></textarea>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">路径</label>
            <div class="layui-input-block">
                <input type="text" name="Url" placeholder="请输入菜单路径" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">图标(font awesome)</label>
            <div class="layui-input-block">
                <input type="text" name="Icon"  placeholder="请输入菜单图标" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">排序</label>
            <div class="layui-input-block">
                <input type="text" name="Sort" lay-verify="number" autocomplete="off" placeholder="请输入排序"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">父类</label>
            <div class="layui-input-inline">
                <select name="Pid">
                    <option value="">顶级</option>
                    <<< range  .menus >>>
                    <option value="<<< .Id >>>"><<< .Title >>></option>
                    <<< end >>>
                </select>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">关联分类</label>
            <div class="layui-input-inline">
                <select name="CategoryId">
                    <option value="">请选择</option>
                    <<< range  .category >>>
                    <option value="<<< .Id >>>"><<< .Name >>></option>
                    <<< end >>>
                </select>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">菜单类型</label>
            <div class="layui-input-block">
                <input type="radio" name="IsFront" value="1" title="前台菜单" checked="">
                <input type="radio" name="IsFront" value="2" title="后台菜单">
            </div>
        </div>
        <button lay-filter="edit" lay-submit style="display: none;"></button>
    </form>
</div>
