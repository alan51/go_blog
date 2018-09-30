layui.define(function(exports){ //提示：模块也可以依赖其它模块，如：layui.define('layer', callback);
    "use strict";
    var $ = layui.jquery;
    var ajax = $.ajax;
    var addBoxIndex = -1;
    $.extend({
        ajax: function(url, options) {
            if (typeof url === 'object') {
                options = url;
                url = undefined;
            }
            options = options || {};
            url = options.url;
            var xsrftoken = $('meta[name=_xsrf]').attr('content');
            console.log(xsrftoken)
            if (xsrftoken == undefined || !xsrftoken) {
                xsrftoken = $('meta[name=_xsrf_main]').attr('content');
            }
            var headers = options.headers || {};
            var domain = document.domain.replace(/\./ig, '\\.');
            if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
                headers = $.extend(headers, {'X-Xsrftoken':xsrftoken});
            }
            options.headers = headers;
            return ajax(url, options);
        }
    });
    var obj = {
        ajax: function(url, data, type, async,callback) {
            $.ajax({
                type: type?type:"POST",
                url: url,
                data: data,
                async: async?async:false,
                success:function (res) {
                    callback(res)
                },
                error:function (res) {
                    layer.msg("请求错误")
                }
            })
        },
        ajaxGet: function (url, data, callback) {
            this.ajax(url, data, "GET", false, callback)
        },
        ajaxPost: function (url, data, callback) {
            this.ajax(url, data, "POST", false, callback);
        },

        btnAdd:function (obj, url, title) {
            if (addBoxIndex !== -1)
                return;
            $.get(url, null, function (form) {
                addBoxIndex = layer.open({
                    type: 1,
                    title: title?title:"添加",
                    content: form,
                    btn: ['保存', '取消'],
                    shade: false,
                    offset: ['100px', '30%'],
                    area: ['700px', '450px'],
                    zIndex: 19950924,
                    maxmin: true,
                    yes: function (index) {
                        //触发表单的提交事件
                        $('form.layui-form').find('button[lay-filter=edit]').click();
                    },
                    full: function (elem) {
                        var win = window.top === window.self ? window : parent.window;
                        $(win).on('resize', function () {
                            var $this = $(this);
                            elem.width($this.width()).height($this.height()).css({
                                top: 0,
                                left: 0
                            });
                            elem.children('div.layui-layer-content').height($this.height() - 95);
                        });
                    },
                    success: function (layero, index) {
                        //弹出窗口成功后渲染表单
                        var form = layui.form;
                        form.render();
                        form.on('submit(edit)', function (data) {
                            console.log(data.elem) //被执行事件的元素DOM对象，一般为button对象
                            console.log(data.form) //被执行提交的form对象，一般在存在form标签时才会返回
                            console.log(data.field) //当前容器的全部表单字段，名值对形式：{name: value}
                            //调用父窗口的layer对象
                            layerTips.open({
                                title: '这里面是表单的信息',
                                type: 1,
                                content: JSON.stringify(data.field),
                                area: ['500px', '300px'],
                                btn: ['关闭并刷新', '关闭'],
                                yes: function (index, layero) {
                                    layerTips.msg('你点击了关闭并刷新');
                                    layerTips.close(index);
                                    location.reload(); //刷新
                                }

                            });
                            //这里可以写ajax方法提交表单
                            return false; //阻止表单跳转。如果需要表单跳转，去掉这段即可。
                        });
                        //console.log(layero, index);
                    },
                    end: function () {
                        addBoxIndex = -1;
                    }
                });
            });
        }
    };

    //输出test接口
    exports('ajaxJson', obj);
});