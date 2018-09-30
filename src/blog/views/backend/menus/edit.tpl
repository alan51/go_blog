<div style="margin: 15px">
    <form class="layui-form" action="">
        <input type="hidden" name="_xsrf" value="<<< .xsrf_token >>>">
        <input type="hidden" name="Id" value="<<< .menu_info.Id >>>">
        <div class="layui-form-item">
            <label class="layui-form-label">菜单名称</label>
            <div class="layui-input-block">
                <input type="text" name="Name" value="<<< .menu_info.Name >>>" lay-verify="required|name" placeholder="请输入菜单名称" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">描述</label>
            <div class="layui-input-block">
                <textarea placeholder="请输入描述" name="Desc" class="layui-textarea"><<< .menu_info.Desc >>></textarea>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">路径</label>
            <div class="layui-input-block">
                <input type="text" name="Url" value="<<< .menu_info.Url >>>" placeholder="请输入菜单路径" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">图标(font awesome)</label>
            <div class="layui-input-block">
                <input type="text" name="Icon" value="<<< .menu_info.Icon >>>" placeholder="请输入菜单图标" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">排序</label>
            <div class="layui-input-block">
                <input type="text" name="Sort" lay-verify="number" value="<<< .menu_info.Sort >>>" autocomplete="off" placeholder="请输入排序"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">父类</label>
            <div class="layui-input-inline">
                <select name="Pid">
                    <<< $menu_info := .menu_info.Pid >>>
                    <option value="">顶级</option>
                    <<< range $ind,$menu:=.menus >>>
                    <option  <<< if eq $menu_info $menu.Id>>> selected <<< end >>> value="<<< $menu.Id >>>"><<< $menu.Title >>></option>
                    <<< end >>>
                </select>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">关联分类</label>
            <div class="layui-input-inline">
                <select name="CategoryId">
                    <<< if .no_cate>>>
                    <option value="">请选择</option>
                    <<< range  .category >>>
                    <option value="<<< .Id >>>"><<< .Name >>></option>
                    <<< end >>>
                    <<< else>>>
                    <<< $cid := .menu_info.Category.Id >>>
                    <option value="">请选择</option>
                    <<< range  .category >>>
                    <option <<< if eq $cid .Id>>> selected <<< end >>> value="<<< .Id >>>"><<< .Name >>></option>
                    <<< end >>>
                    <<< end>>>
                </select>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">菜单类型</label>
            <div class="layui-input-block">
                <input type="radio" name="IsFront" value="1" title="前台菜单" <<< if eq .menu_info.IsFront 1>>> checked="" <<< end >>> >
                <input type="radio" name="IsFront" value="2" title="后台菜单" <<< if eq .menu_info.IsFront 2>>> checked="" <<< end >>>>
            </div>
        </div>
        <button lay-filter="edit" lay-submit style="display: none;"></button>
    </form>
</div>
