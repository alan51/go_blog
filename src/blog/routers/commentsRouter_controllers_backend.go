package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["blog/controllers/backend:ArticleController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:ArticleController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:ArticleController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:ArticleController"],
		beego.ControllerComments{
			Method: "AddArticle",
			Router: `/add_article`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:ArticleController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:ArticleController"],
		beego.ControllerComments{
			Method: "DeleteSelect",
			Router: `/delete_select`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:ArticleController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:ArticleController"],
		beego.ControllerComments{
			Method: "EditArticle",
			Router: `/edit_article/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:ArticleController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:ArticleController"],
		beego.ControllerComments{
			Method: "PostAddArticle",
			Router: `/post_add_article`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:ArticleController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:ArticleController"],
		beego.ControllerComments{
			Method: "PostEditArticle",
			Router: `/post_edit_article`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:CategoryController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:CategoryController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:CategoryController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:CategoryController"],
		beego.ControllerComments{
			Method: "AddCategory",
			Router: `/add_category`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:CategoryController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:CategoryController"],
		beego.ControllerComments{
			Method: "DeleteSelect",
			Router: `/delete_select`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:CategoryController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:CategoryController"],
		beego.ControllerComments{
			Method: "EditCategory",
			Router: `/edit_category/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:CategoryController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:CategoryController"],
		beego.ControllerComments{
			Method: "PostAddCategory",
			Router: `/post_add_category`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:CategoryController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:CategoryController"],
		beego.ControllerComments{
			Method: "PostEditCategory",
			Router: `/post_edit_category`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:FileUploadController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:FileUploadController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:FileUploadController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:FileUploadController"],
		beego.ControllerComments{
			Method: "DeleteSelect",
			Router: `/delete_select`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:FileUploadController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:FileUploadController"],
		beego.ControllerComments{
			Method: "UploadEditor",
			Router: `/edit_upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:FileUploadController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:FileUploadController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:FileUploadController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:FileUploadController"],
		beego.ControllerComments{
			Method: "UploadShow",
			Router: `/upload`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:LinkController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:LinkController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:LinkController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:LinkController"],
		beego.ControllerComments{
			Method: "AddLink",
			Router: `/add_link`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:LinkController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:LinkController"],
		beego.ControllerComments{
			Method: "DeleteSelect",
			Router: `/delete_select`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:LinkController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:LinkController"],
		beego.ControllerComments{
			Method: "EditLink",
			Router: `/edit_link/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:LinkController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:LinkController"],
		beego.ControllerComments{
			Method: "PostAddLink",
			Router: `/post_add_link`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:LinkController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:LinkController"],
		beego.ControllerComments{
			Method: "PostEditLink",
			Router: `/post_edit_link`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:MainController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:MainController"],
		beego.ControllerComments{
			Method: "URLMapping",
			Router: `/admin/main`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:MenusController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:MenusController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:MenusController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:MenusController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:MenusController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:MenusController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:MenusController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:MenusController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:MenusController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:MenusController"],
		beego.ControllerComments{
			Method: "AddMenu",
			Router: `/add_menu`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:MenusController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:MenusController"],
		beego.ControllerComments{
			Method: "DeleteSelect",
			Router: `/delete_select`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:MenusController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:MenusController"],
		beego.ControllerComments{
			Method: "EditMenu",
			Router: `/edit_menu/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:MenusController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:MenusController"],
		beego.ControllerComments{
			Method: "PostAddMenu",
			Router: `/post_add_menu`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:MenusController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:MenusController"],
		beego.ControllerComments{
			Method: "PostEditMenu",
			Router: `/post_edit_menu`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:SiteController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:SiteController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:SiteController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:SiteController"],
		beego.ControllerComments{
			Method: "AddSite",
			Router: `/add_site`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:SiteController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:SiteController"],
		beego.ControllerComments{
			Method: "DeleteSelect",
			Router: `/delete_select`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:SiteController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:SiteController"],
		beego.ControllerComments{
			Method: "EditSite",
			Router: `/edit_site/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:SiteController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:SiteController"],
		beego.ControllerComments{
			Method: "PostAddSite",
			Router: `/post_add_site`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:SiteController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:SiteController"],
		beego.ControllerComments{
			Method: "PostEditSite",
			Router: `/post_edit_site`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:TagController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:TagController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:TagController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:TagController"],
		beego.ControllerComments{
			Method: "AddTag",
			Router: `/add_tag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:TagController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:TagController"],
		beego.ControllerComments{
			Method: "DeleteSelect",
			Router: `/delete_select`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:TagController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:TagController"],
		beego.ControllerComments{
			Method: "EditTag",
			Router: `/edit_tag/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:TagController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:TagController"],
		beego.ControllerComments{
			Method: "PostAddTag",
			Router: `/post_add_tag`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:TagController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:TagController"],
		beego.ControllerComments{
			Method: "PostEditTag",
			Router: `/post_edit_tag`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:UserController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:UserController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:UserController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:UserController"],
		beego.ControllerComments{
			Method: "AddUser",
			Router: `/add_user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:UserController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:UserController"],
		beego.ControllerComments{
			Method: "DeleteSelect",
			Router: `/delete_select`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:UserController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:UserController"],
		beego.ControllerComments{
			Method: "EditUser",
			Router: `/edit_user/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:UserController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:UserController"],
		beego.ControllerComments{
			Method: "PostAddUser",
			Router: `/post_add_user`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/backend:UserController"] = append(beego.GlobalControllerRouter["blog/controllers/backend:UserController"],
		beego.ControllerComments{
			Method: "PostEditUser",
			Router: `/post_edit_user`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
