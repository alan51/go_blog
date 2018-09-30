<div class="content-wrap">
<div class="content">
    <header class="article-header">
        <h1 class="article-title">
            <a href="/article/<<< .article_info.Id>>>"><<< .article_info.Title>>></a>
        </h1>
        <div class="article-meta">
            <span class="item"><<< .article_info.CreatedAt>>></span>
            <span class="item">分类：
                <a href="/article_cate/<<< .article_info.Id>>>" rel="category tag"><<< .article_info.Category.Name>>></a>
            </span>
            <span class="item post-views">阅读(<<< .article_info.ViewNum>>>)</span>
        </div>
    </header>
    <div class="article-tags">
        标签：
        <<< .article_info.Tags | StrToSplit | str2html>>>
    </div>
    <article class="article-content" id="test-editormd-view">
        <textarea id="test-editormd-view2" style="display:none;"><<< .article_info.Content  | str2html>>></textarea>
    </article>
    <div class="post-copyright">
        <b><<<.config.web_state | str2html>>><br></b>
        文章来源:
        <a href="https://www.aicyun.com"><<< .config.app_name>>></a> »
        <a href="/article/<<< .article_info.Id>>>"><<< .article_info.Title>>></a>
    </div>
    <div class="action-share"></div>
    <div class="article-author">
        <img data-src="<<< .article_info.User.Img>>>" class="avatar avatar-100" width="220" height="150" src="<<< .article_info.User.Img>>>"
                style="display: block;">
        <h4>
            <i class="fa fa-user" aria-hidden="true"></i>
            <a title="查看更多文章" href="/"><<< .article_info.User.NickName>>></a>
        </h4>
    </div>
    <div class="text-center">
        <div id="cyReward" role="cylabs" data-use="reward"></div>
    </div>
    <nav class="article-nav">
        <span class="article-nav-prev">
            <<< .pre | str2html>>>
        </span>
        <span class="article-nav-next">
            <<< .next | str2html>>>
        </span>
    </nav>
    <<< if .tui_article_status >>>
    <div class="relates">
        <div class="title"><h3>相关推荐</h3></div>
        <ul>
            <<< range $list :=.tui_article >>>
            <li><a href="/article/<<< $list.Id>>>" target="_blank"> <<< $list.Title >>> </a></li>
            <<< end >>>
        </ul>
    </div>
    <<< end >>>
    <div class="title" id="comments">
        <h3>评论
            <small>抢沙发</small>
        </h3>
    </div>
    <div id="respond" class="no_webshot">
        <div id="SOHUCS" sid="<<< .article_info.Id>>>" ></div>
        <<< .config.comment | str2html>>>
    </div>
</div>
</div>
<<< template "front/toc.tpl" .>>>
<!--<script>
    window.jsui={
        www: 'https://www.aicyun.com',
        uri: '<<<.qiniu_domain>>>/static/dux',
        ver: '',
        roll: ["2","3"],
        ajaxpager: '0',
        url_rp: ''
    };
    var qiniu_domain = "<<<.qiniu_domain>>>";
</script>-->

<script type="text/javascript" src="<<<.qiniu_domain>>>/static/js/jquery.min.js"></script>
<!--<script type='text/javascript' src='<<<.qiniu_domain>>>/static/dux/js/libs/bootstrap.min.js'></script>
<script type='text/javascript' src='<<<.qiniu_domain>>>/static/dux/js/loader.js'></script>-->
<script src="<<<.qiniu_domain>>>/static/plugins/editor/lib/marked.min.js"></script>
<script src="<<<.qiniu_domain>>>/static/plugins/editor/lib/prettify.min.js"></script>
<!--
<script src="<<<.qiniu_domain>>>/static/plugins/editor/lib/raphael.min.js"></script>
<script src="<<<.qiniu_domain>>>/static/plugins/editor/lib/underscore.min.js"></script>
<script src="<<<.qiniu_domain>>>/static/plugins/editor/lib/sequence-diagram.min.js"></script>
<script src="<<<.qiniu_domain>>>/static/plugins/editor/lib/flowchart.min.js"></script>
<script src="<<<.qiniu_domain>>>/static/plugins/editor/lib/jquery.flowchart.min.js"></script>-->
<script src="<<<.qiniu_domain>>>/static/plugins/editor/js/editormd.min.js"></script>
<script>
    var testEditormdView;
    testEditormdView = editormd.markdownToHTML("test-editormd-view", {
        htmlDecode      : "style,script,iframe",  // you can filter tags decode
        emoji           : false,
        taskList        : true,
        tex             : true,  // 默认不解析
        flowChart       : false,  // 默认不解析
        sequenceDiagram : false,  // 默认不解析
        tocm            : true,    // Using [TOCM]
        tocContainer    : "#custom-toc-container", // 自定义 ToC 容器层
    });
</script>
<!--<script src="/static/dux/js/article.js" async="async"></script>-->