<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>文章管理</title>
    <meta name="renderer" content="webkit">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="format-detection" content="telephone=no">

    <link rel="stylesheet" href="/static/plugins/layui/css/layui.css" media="all" />
    <link rel="stylesheet" href="/static/css/btable.css" />
    <link rel="stylesheet" type="text/css" href="/static/plugins/editor/css/editormd.min.css"/>
</head>
<body>
<div style="margin: 15px">
    <form class="layui-form" action="/admin/article/post_add_article" method="post">
        <input type="hidden" name="_xsrf" value="<<< .xsrf_token >>>">
        <div class="layui-form-item">
            <label class="layui-form-label">标题</label>
            <div class="layui-input-block">
                <input type="text" name="Title" lay-verify="required" placeholder="请输入标题" autocomplete="off"
                       class="layui-input">
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">描述</label>
            <div class="layui-input-block">
                <textarea placeholder="请输入描述" name="Desc" class="layui-textarea"></textarea>
            </div>
        </div>
       <!-- <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">关键字</label>
            <div class="layui-input-block">
                <input type="text" name="Keywords" lay-verify="required" placeholder="请输入关键字" autocomplete="off"
                       class="layui-input">
            </div>
        </div>-->
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">分类</label>
            <div class="layui-input-block">
                <select name="CategoryId" class="form-control select">
                    <<< range $category := .category>>>
                    <option value="<<< $category.Id >>>"><<<$category.Name>>></option>
                    <<< end >>>
                </select>
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">标签</label>
            <div class="layui-input-block">
                <textarea placeholder="请输入标签,换行一个" name="Tags" class="layui-textarea"></textarea>
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">首图</label>
            <div class="layui-input-block">
                <a class="layui-btn select_img"><i class="layui-icon"></i> 选择图片</a>
                <div id="hidden_input">
                    <input type="hidden" name="IndexImg" id="img" value="" />
                </div>
                <div class="img_show"></div>
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">内容</label>
            <div class="layui-input-block" id="editormd">
                <textarea class="editormd-html-textarea" name="Content"></textarea>
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
            <label class="layui-form-label">是否推荐</label>
            <div class="layui-input-block">
                <input type="radio" name="IsRecommend" value="1" title="是" checked="">
                <input type="radio" name="IsRecommend" value="2" title="否">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">是否置顶</label>
            <div class="layui-input-block">
                <input type="radio" name="IsTop" value="1" title="是">
                <input type="radio" name="IsTop" value="2" title="否"  checked="">
            </div>
        </div>
        <button lay-filter="edit" lay-submit style="display: none;"></button>
        <!-- <div class="layui-form-item">
            <div class="layui-input-block">
                <button class="layui-btn" lay-submit="" lay-filter="demo1">立即提交</button>
                <button type="reset" class="layui-btn layui-btn-primary">重置</button>
            </div>
        </div>
        -->
    </form>
</div>
<script type="text/javascript" src="/static/js/jquery.min.js"></script>
<script type="text/javascript" src="/static/plugins/editor/js/editormd.min.js"></script>
<!-- <script type="text/javascript" src="/static/front/layui/layui.all.js"></script>-->
<script>
    $(function() {
        var editor = editormd("editormd", {
            width  : "90%",
            height : 640,
            path : "/static/plugins/editor/lib/", // Autoload modules mode, codemirror, marked... dependents libs path
            saveHTMLToTextarea : true,
            syncScrolling : "single",
            editorTheme: "pastel-on-dark",
            theme: "gray",
            previewTheme: "dark",
            emoji: true,
            imageUpload:true,
            imageFormats   : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
            imageUploadURL : "/admin/file/edit_upload"
        });
        editor.getHTML()
        editor.getPreviewedHTML()
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
        /*$("form").on("submit", function(){
            var data = $(this).serializeArray()
            console.log(data)
            return false
        })*/
    });
</script>
</body>
</html>
