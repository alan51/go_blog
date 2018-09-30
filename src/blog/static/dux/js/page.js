tbquire(['page'], function(){
    $('#page').jqPaginator({
        totalPages: TotalPage,
        visiblePages:1,
        first: '<li class="first"><a href="javascript:;">首 页</a></li>',
        prev: '<li class="prev"><a href="javascript:;">上一页</a></li>',
        next: '<li class="next"><a href="javascript:;">下一页</a></li>',
        last: '<li class="last"><a href="javascript:;">尾 页</a></li>',
        currentPage: currentPage,
        onPageChange: function (num, type) {
            console.log(type)
            if(type == 'change') {
                window.location.href = '/article_cate?cid='+cid+'&page='+num
            }
        }
    });
})