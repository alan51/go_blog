layui.define('layer', function (exports) {
    "use strict";
    var $ = layui.jquery,
        layerTips = parent.layer === undefined ? layui.layer : parent.layer;
    var ajax = $.ajax;
    $.extend({
        ajax: function(url, options) {
            if (typeof url === 'object') {
                options = url;
                url = undefined;
            }
            options = options || {};
            url = options.url;
            var xsrftoken = $('meta[name=_xsrf]').attr('content');
            if (xsrftoken == undefined || !xsrftoken) {
                xsrftoken = $('meta[name=_xsrf_main]').attr('content');
            }
            if (xsrftoken == undefined || !xsrftoken) {
                xsrftoken = $('#xsrf_token').val();
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

    var baajax = {
        post: function (url, params, callback) {
            $.post(url, params, function (res) {
                if (res.statusCode === 101) {
                    layerTips.alert(res.msg, {
                        icon: 2, title: '系统提示', cancel: function (index, layero) {
                            top.location.href = location.origin + '/login';
                        }
                    }, function () {
                        top.location.href = location.origin + '/login';
                    });
                }
                callback(res);
            }, 'json');
        },
        get: function (url, params, callback) {
            $.getJSON(url, params, function (res) {
                if (res.statusCode === 101) {
                    layerTips.alert(res.msg, {
                        icon: 2, title: '系统提示', cancel: function (index, layero) {
                            top.location.href = location.origin + '/login';
                        }
                    }, function () {
                        top.location.href = location.origin + '/login';
                    });
                }
                callback(res);
            });
        },
        v: '1.0.0'
    };


    exports('baajax', baajax);
});