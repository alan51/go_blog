(function() {  //加载
    var obj =  {};
    obj.loadScript = function(url,callback){
        var doc = document;
        var script = doc.createElement("script");
        script.type = "text/javascript";
        if(script.readyState){  //IE
            script.onreadystatechange = function(){
                if(script.readyState=="load"||script.readyState=="complete"){
                    script.onreadystatechange = null;
                    callback();
                }
            };
        }else{
            script.onload = function(){
                callback();
            };
        }
        script.src = url;
        doc.getElementsByTagName("head")[0].appendChild(script);
    };
    var jsList = [
        "/static/plugins/editor/lib/marked.min.js",
        "/static/plugins/editor/lib/prettify.min.js",
        "/static/plugins/editor/js/editormd.min.js",
    ];
    function callback(){
        jsList.length?obj.loadScript(jsList.shift(),callback)
            :(function(){time = null})();
        var testEditormdView;
        testEditormdView = editormd.markdownToHTML("test-editormd-view", {
            htmlDecode      : "style,script,iframe",  // you can filter tags decode
            emoji           : true,
            taskList        : true,
            tex             : true,  // 默认不解析
            flowChart       : true,  // 默认不解析
            sequenceDiagram : true,  // 默认不解析
            tocm            : true,    // Using [TOCM]
            tocContainer    : "#custom-toc-container", // 自定义 ToC 容器层
        });
    }
    var time = setTimeout(function(){obj.loadScript(jsList.shift(),callback)},25);

})();
