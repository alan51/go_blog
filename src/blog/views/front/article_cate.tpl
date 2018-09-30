<div class="content-wrap">
    <div class="content">
        <div class="title">
            <h3><<< .category.Name>>></h3>
            <!-- <div class="more">

            </div>-->
        </div>
        <<< $config := .config>>>
        <<< $qiniu_domain := .qiniu_domain>>>
        <<< $qiniu_article := .qiniu_article>>>
        <<< range $article := .page.List >>>
        <article class="excerpt excerpt-1">
            <a class="focus" href="/article/<<<$article.Id>>>">
                <img data-src="<<<$qiniu_domain>>><<< $article.IndexImg >>><<< $qiniu_article>>>" alt="<<< $article.Title >>> - <<< $config.app_name >>>" class="thumb" width="220" height="150">
            </a>
            <header>
                <a class="cat" href="/article_cate?cid=<<<$article.Category.Id>>>"><<< $article.Category.Name >>><i></i></a>
                <h2>
                    <a href="/article/<<<$article.Id>>>" title="<<< $article.Title >>> - <<< $config.app_name >>>"><<< $article.Title >>></a>
                </h2>
            </header>
            <p class="meta">
                <time><i class="fa fa-clock-o"></i><<< $article.CreatedAt>>></time>
                <span class="author"><i class="fa fa-user"></i><<< $article.User.NickName >>></span><span class="pv"><i class="fa fa-eye"></i>阅读(<<< $article.ViewNum >>>)</span>
            </p>
            <p class="note">
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
    var cid = parseInt('<<<.cid>>>');
    var qiniu_domain = '<<< .qiniu_domain>>>';
</script>
<script type='text/javascript' src='<<<.qiniu_domain>>>/static/dux/js/cate_page.js' async="async"></script>