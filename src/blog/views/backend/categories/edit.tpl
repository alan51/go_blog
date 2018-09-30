<div style="margin: 15px">
    <form class="layui-form" action="">
        <input type="hidden" name="_xsrf" value="<<< .xsrf_token >>>">
        <input type="hidden" name="Id" value="<<< .category_info.Id >>>">
        <div class="layui-form-item">
            <label class="layui-form-label">分类名称</label>
            <div class="layui-input-block">
                <input type="text" name="Name" value="<<< .category_info.Name >>>" lay-verify="required" placeholder="请输入分类名称" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">描述</label>
            <div class="layui-input-block">
                <textarea placeholder="请输入描述" name="Desc" class="layui-textarea"><<< .category_info.Desc >>></textarea>
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">关键字</label>
            <div class="layui-input-block">
                <input type="text" name="Keywords" lay-verify="required" placeholder="请输入关键字" autocomplete="off"
                       class="layui-input" value="<<< .category_info.Keywords >>>">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">图标</label>
            <div class="layui-input-block">
                <input type="text" name="Icon" value="<<< .category_info.Icon >>>" lay-verify="required" placeholder="请输入图标库标识" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">自定义链接</label>
            <div class="layui-input-block">
                <input type="text" name="Url" value="<<< .category_info.Url >>>" lay-verify="" placeholder="请输入自定义链接" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label"> 排序</label>
            <div class="layui-input-block">
                <input type="text" name="Sort" value="<<< .category_info.Sort >>>" lay-verify="required" placeholder="请输入排序" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">父分类</label>
            <div class="layui-input-block">
                <select name="ParentId" class="form-control select">
                    <option value="">顶级</option>
                    <<<$cid := .category_info.ParentId>>>
                    <<< range $category := .category>>>
                    <option <<< if eq $category.Id $cid>>>selected<<< end>>> value="<<< $category.Id >>>"><<<$category.Name>>></option>
                    <<< end >>>
                </select>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">顶部菜单</label>
            <div class="layui-input-block">
                <input type="radio" name="IsTop" <<< if eq .category_info.IsTop 1 >>> checked="" <<< end >>> value="1" title="是">
                <input type="radio" name="IsTop" <<< if eq .category_info.IsTop 2 >>> checked="" <<< end >>> value="2" title="否">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">状态</label>
            <div class="layui-input-block">
                <input type="radio" name="Status" <<< if eq .category_info.Status 1 >>> checked="" <<< end >>> value="1" title="启用">
                <input type="radio" name="Status" <<< if eq .category_info.Status 2 >>> checked="" <<< end >>> value="2" title="禁用">
            </div>
        </div>
        <button lay-filter="edit" lay-submit style="display: none;"></button>
    </form>
</div>
