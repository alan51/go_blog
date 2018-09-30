<div class="content-wrap">
    <div class="content">
        <<< if .top_article_status >>>
        <<< range $top := .top_article >>>
        <article class="excerpt-minic excerpt-minic-index">
            <h2>
                <a class="red" href="/article_cate/<<< $top.Category.Id>>>"><<< $top.Category.Name >>></a>
                <a href="/article/<<< $top.Id>>>" title="<<< $top.Title >>>"><<< $top.Title >>></a>
            </h2>
            <p class="note">
                <<< $top.Desc >>>
            </p>
        </article>
        <<< end >>>
        <<< end>>>
        <div class="title">
            <h3>最新</h3>
            <!-- <div class="more">

            </div>-->
        </div>
        <<< $config := .config>>>
        <<< $qiniu_domain := .qiniu_domain>>>
        <<< $qiniu_article := .qiniu_article>>>
        <<< range $article := .page.List >>>
        <article class="excerpt excerpt-<<<$article.Id>>>">
            <a class="focus" href="/article/<<<$article.Id>>>">
                <img data-src="<<<$qiniu_domain>>><<< $article.IndexImg >>><<<$qiniu_article>>>" alt="<<< $article.Title >>> - <<< $config.app_name >>>" class="thumb" width="220" height="150">
            </a>
            <header>
                <a class="cat" href="/article_cate?cid=<<<$article.Category.Id>>>"><<< $article.Category.Name >>><i></i></a>
                <h2>
                    <a href="/article/<<<$article.Id>>>" title="<<< $article.Title >>> - <<< $config.app_name >>>"><<< $article.Title >>></a>
                </h2>
            </header>
            <p class="meta">
                <time><i class="fa fa-clock-o"></i><<< $article.CreatedAt>>></time>
                <span class="author"><i class="fa fa-user"></i><<< $article.User.NickName >>> </span><span class="pv"><i class="fa fa-eye"></i>阅读(<<< $article.ViewNum >>>)</span>
            </p>
            <p class="note note_space">
                <<< $article.Desc >>>
            </p>
            <p class="tags">
                <<< $article.Tags | StrToSplit | str2html>>>
            </p>
        </article>
        <<< end>>>
        <div class="pagination">
            <ul id="page">

            </ul>
        </div>
    </div>
</div>
<<< template "layout/front/tree.html" .>>>
<script>
    window.jsui={
        www: 'https://www.aicyun.com',
        uri: '<<<.qiniu_domain>>>/static/dux',
        ver: '',
        roll: ["2","3"],
        ajaxpager: '0',
        url_rp: ''
    };
    var TotalPage = parseInt('<<<.page.TotalPage>>>');
    var currentPage = parseInt('<<<.page.PageNo>>>');
    var title = '<<<.search>>>';
    var qiniu_domain = '<<< .qiniu_domain>>>';
</script>
<script type='text/javascript' src='<<<.qiniu_domain>>>/static/dux/js/home_page.js' async="async"></script>