<!DOCTYPE html>
<html lang="en">
<head>
    <!-- META SECTION -->
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />

    <link rel="icon" href="favicon.ico" type="image/x-icon" />
    <!-- END META SECTION -->

    <!-- CSS INCLUDE -->
    <link rel="stylesheet" type="text/css" id="theme" href="/static/plugins/gallery/css/theme-default.css?v=1.1"/>
    <link rel="stylesheet" type="text/css" href="/static/plugins/layui/css/layui.css"/>
    <!-- EOF CSS INCLUDE -->
</head>
<body>
<!-- START PAGE CONTAINER -->
<div class="page-container">

    <!-- PAGE CONTENT -->
    <div class="page-content">


        <!-- START CONTENT FRAME -->
        <div class="content-frame">

            <!-- START CONTENT FRAME TOP -->
            <div class="content-frame-top">
                <div class="page-title">
                    <h2><span class="fa fa-image"></span>图 册</h2>
                </div>
                <div class="pull-right">
                    <!-- <button class="btn btn-primary"><span class="fa fa-upload"></span>上 传</button>-->
                    <button class="btn btn-default content-frame-right-toggle"><span class="fa fa-bars"></span></button>
                </div>
            </div>

            <!-- START CONTENT FRAME RIGHT -->
            <div class="content-frame-right">
                <div class="block push-up-10">
                    <form action="/admin/file/upload" class="dropzone dropzone-mini" id="my-awesome-dropzone">
                        <input type="hidden" name="_xsrf" value="<<< .xsrf_token >>>">
                        <input type="hidden" name="category_name" value="<<< .category_name >>>">
                    </form>
                </div>
                <h4>分组:</h4>
                <div class="list-group border-bottom push-down-20">
                    <<< range $cate := .file_category >>>
                    <!-- <a href="/admin/file?query=category_name__exact:<<<$cate.CategoryName>>>" class="list-group-item active"><<<$cate.CategoryName>>> -->
                    <a href="/admin/file?category_name=<<<$cate.CategoryName>>>" class="list-group-item active"><<<$cate.CategoryName>>>
                        <!-- <span class="badge badge-primary">12</span>-->
                    </a>
                    <<< end >>>
                </div>
            </div>
            <!-- END CONTENT FRAME RIGHT -->

            <!-- START CONTENT FRAME BODY -->
            <div class="content-frame-body content-frame-body-left">

                <div class="pull-left push-up-10">
                    <button class="btn btn-primary" id="gallery-toggle-items">全选/反选</button>
                </div>
                <div class="pull-right push-up-10">
                    <div class="btn-group">

                        <button class="btn btn-primary gallery-selected-remove"><span class="fa fa-trash-o"></span>删除</button>
                    </div>
                </div>

                <div class="gallery" id="links">
                    <<< range $img := .images >>>
                    <a class="gallery-item" data-id="<<< $img.Id >>>" href="<<< $img.Url >>>" title="Nature Image 1" data-gallery>
                        <div class="image">
                            <img src="<<< $img.Url >>>" alt="<<< $img.Name >>>"/>
                            <ul class="gallery-item-controls">
                                <li><label class="check"><input type="checkbox" class="icheckbox"/></label></li>
                                <li><span class="gallery-item-remove"><i class="fa fa-times"></i></span></li>
                            </ul>
                        </div>
                        <div class="meta">
                            <strong><<< $img.Name >>></strong>
                        </div>
                    </a>
                    <<< end >>>

                </div>
                <div id="pager" class="pagination pagination-sm pull-right push-down-20 push-up-20">

                </div>
                <!-- <ul class="pagination pagination-sm pull-right push-down-20 push-up-20">
                    <<< if .pageIndex >>>
                    <li><a href="<<< .pre_url >>>">上一页</a></li>
                    <<< else >>>
                    <li class="disabled"><a href="#">上一页</a></li>
                    <<< end >>>
                    <li class="active"><a href="<<< .next_url >>>">下一页</a></li>
                </ul> -->
            </div>
            <!-- END CONTENT FRAME BODY -->
        </div>
        <!-- END CONTENT FRAME -->


    </div>
    <!-- END PAGE CONTENT -->
</div>
<!-- END PAGE CONTAINER -->

<!-- BLUEIMP GALLERY -->
<div id="blueimp-gallery" class="blueimp-gallery blueimp-gallery-controls">
    <div class="slides"></div>
    <h3 class="title"></h3>
    <a class="prev">‹</a>
    <a class="next">›</a>
    <a class="close">×</a>
    <a class="play-pause"></a>
    <ol class="indicator"></ol>
</div>
<!-- END BLUEIMP GALLERY -->

<script type="text/javascript" src="/static/plugins/gallery/js/plugins/jquery/jquery.min.js"></script>
<script type="text/javascript" src="/static/plugins/gallery/js/plugins/jquery/jquery-ui.min.js"></script>
<script type="text/javascript" src="/static/plugins/gallery/js/plugins/bootstrap/bootstrap.min.js"></script>
<!-- END PLUGINS -->

<!-- START THIS PAGE PLUGINS-->
<script type="text/javascript" src="/static/plugins/gallery/js/plugins/mcustomscrollbar/jquery.mCustomScrollbar.min.js"></script>

<script type="text/javascript" src="/static/plugins/gallery/js/plugins/blueimp/jquery.blueimp-gallery.min.js"></script>
<script type="text/javascript" src="/static/plugins/gallery/js/plugins/dropzone/dropzone.min.js"></script>
<script type="text/javascript" src="/static/plugins/gallery/js/plugins/icheck/icheck.min.js"></script>
<!-- END THIS PAGE PLUGINS-->

<!-- START TEMPLATE -->

<script type="text/javascript" src="/static/plugins/gallery/js/plugins.js"></script>
<script type="text/javascript" src="/static/plugins/gallery/js/actions.js"></script>
<script type="text/javascript" src="/static/front/layui/layui.all.js"></script>
<!-- END TEMPLATE -->

<script>
    document.getElementById('links').onclick = function (event) {
        event = event || window.event;
        var target = event.target || event.srcElement;
        var link = target.src ? target.parentNode : target;
        var options = {index: link, event: event,onclosed: function(){
        setTimeout(function(){
        $("body").css("overflow","");
        },200);
    }};
    var links = this.getElementsByTagName('a');
    blueimp.Gallery(links, options);
    };
    Dropzone.options.myAwesomeDropzone = {
        paramName: "image", // The name that will be used to transfer the file
        maxFilesize: 10, // MB
        accept: function(file, done) {
            console.log(file)
            done();
        },
        init:function(){
            this.on("success", function(file) {
                /*layui.layer.open({
                    shade:false,
                    content: '上传成功!'
                    ,btn: ['确定']
                    ,yes: function(index, layero){
                        //按钮【按钮一】的回调
                        window.location.reload()
                    }
                    ,cancel: function(){
                        //右上角关闭回调
                        return false
                    }
                });*/

            });
        }
    };
    var category_name = '<<<.category_name>>>'
    layui.use('laypage', function(){
        var laypage = layui.laypage;

        //执行一个laypage实例
        laypage.render({
            elem: 'pager' //注意，这里的 test1 是 ID，不用加 # 号
            ,count: '<<<.images_num>>>' //数据总数，从服务端得到
            ,limit:12
            //,curr: location.hash.replace('#!fenye=', '') //获取起始页
            //,hash: 'pageIndex=' //自定义hash值
            ,curr: location.hash.replace('#!fenye=', '') //获取hash值为fenye的当前页
            ,hash: 'fenye' //自定义hash值
            ,jump: function(obj, first){
                //obj包含了当前分页的所有参数，比如：
                //console.log(obj.curr); //得到当前页，以便向服务端请求对应页的数据。
                //console.log(obj.limit); //得到每页显示的条数

                //首次不执行
                if(!first){
                    //do something
                    var biao = ''
                    if (category_name) {
                        biao = "&"
                    }else {
                        biao = "?"
                    }
                    location.href = '/admin/file/upload' + biao+"pageIndex=" + obj.curr + "#!fenye=" + obj.curr
                }
            }
        });
    });

</script>

<!-- END SCRIPTS -->
</body>
</html>






