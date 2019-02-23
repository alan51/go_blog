<aside class="sidebar">
<<< $config := .config>>>
<<< $qiniu_domain := .qiniu_domain>>>
<<< $qiniu_article := .qiniu_article>>>
<<< if .tui_article_status >>>
<div class="widget widget_meta">
    <h3>精彩推荐</h3>
    <ul>
        <<< range $list :=.tui_article >>>
        <li>
            <a href="/article/<<<$list.Id>>>">
                <!--<span class="thumbnail">
                    <img data-src="<<<$qiniu_domain>>><<< $list.IndexImg>>><<< $qiniu_article>>>" width="220" height="150" alt="<<< $list.Title >>> - <<< $config.app_name>>>" class="thumb">
                </span>
                <span class="text"><<< $list.Title>>></span>
                <span class="muted"><<< $list.CreatedAt>>></span>
                -->
                <<< $list.Title>>>

            </a>
        </li>
        <<< end>>>
    </ul>
</div>
<<< end>>>
<<< if .hot_article_status >>>
 <div class="widget widget_meta">
     <h3>热门文章</h3>
     <ul>
         <<< range $list :=.hot_article >>>
         <li>
             <a href="/article/<<<$list.Id>>>">
                 <!--<span class="thumbnail">
                     <img data-src="<<<$qiniu_domain>>><<< $list.IndexImg>>><<< $qiniu_article>>>" width="220" height="150" alt="<<< $list.Title >>> - <<< $config.app_name>>>" class="thumb">
                 </span>
                 <span class="text"><<< $list.Title>>></span>
                 <span class="muted"><<< $list.CreatedAt>>></span>
                 -->
                 <<< $list.Title>>>
             </a>
         </li>
         <<< end>>>
     </ul>
 </div>
 <<< end>>>
<div class="widget widget_meta affix-top">
    <h3>索引目录</h3>
    <div id="sidebar" style="max-height: 600px;">
        <div class="markdown-body editormd-preview-container" id="custom-toc-container"></div>
    </div>
</div>
</aside>