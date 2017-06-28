<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
	<meta name="description" content="Metro, a sleek, intuitive, and powerful framework for faster and easier web development for Windows Metro Style.">
	<meta name="keywords" content="HTML, CSS, JS, JavaScript, framework, metro, front-end, frontend, web development">
	<meta name="author" content="Sergey Pimenov and Metro UI CSS contributors">

	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=8">
	<meta http-equiv="Expires" content="0">
	<meta http-equiv="Pragma" content="no-cache">
	<meta http-equiv="Cache-control" content="no-cache">
	<meta http-equiv="Cache" content="no-cache">


	<title>大数据应用分析平台</title>
	<link rel="stylesheet" href="/static/css/metro.css">
	<link rel="stylesheet" href="/static/bootstrap-3.3.7-dist/css/bootstrap.min.css"/>
	<link rel="stylesheet" href="/static/Font-Awesome-3.2.1/css/font-awesome.min.css"/>

	<link rel="stylesheet" href="/static/theme/common.css"/>
	<link rel="stylesheet" href="/static/theme/blue/index.css" type="text/css" />
	<link rel="stylesheet" href="/static/css/animate.css"/>
	<link rel="stylesheet" href="/static/nprogress/nprogress.css"/>

	<script type="text/javascript" src="/static/js/jquery-3.1.1.min.js"></script>
	<script src="/static/nprogress/nprogress.js"></script>

	<script type="text/javascript" src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>
	<script type="text/javascript" src="/static/jquery-i18n-properties/jquery.i18n.properties.min.js"></script>
	<script type="text/javascript" src="/static/js/utils.min.js"></script>


	<!--bootstrap-table表格-->
	<!-- Latest compiled and minified CSS -->
	<link rel="stylesheet" href="/static/bootstrap-table/dist/bootstrap-table.min.css">
	<!-- Latest compiled and minified JavaScript -->
	<script src="/static/bootstrap-table/dist/bootstrap-table.min.js"></script>
	<!-- Latest compiled and minified Locales -->
	<script src="/static/bootstrap-table/dist/locale/bootstrap-table-zh-CN.min.js"></script>


	<!--bootstrap switch-->
	<link rel="stylesheet" href="/static/bootstrap-switch-master/dist/css/bootstrap3/bootstrap-switch.min.css"/>
	<script src="/static/bootstrap-switch-master/dist/js/bootstrap-switch.min.js"></script>

	<!--webupload-->
	<link rel="stylesheet" href="/static/webuploader/dist/webuploader.css"/>
	<script src="/static/webuploader/dist/webuploader.min.js"></script>

</head>

<body style="overflow: hidden" class="hzwy23-theme-background">
<div id="bigdata-platform-subsystem"
	 style="margin-right:0px; background-size: cover; overflow: hidden;">
	<div style="position: relative; height: 60px; text-align: left;">
		<div class="col-sm-6 col-md-6 col-lg-6" style="padding-left: 30px;">
			<h3 id="huuid" style="color: #ffffff; font-size: 23px; font-weight: 700; height: 60px; line-height: 60px;">大数据应用分析平台</h3>
		</div>
		<div class="col-sm-6 col-md-6 col-lg-6" style="padding-left: 30px; text-align: right">
			<span class="label label-primary" style="color: #ffffff; font-size: 14px; font-weight: 700; height: 60px; line-height: 60px;"><i class="icon-user"></i>&nbsp;{{.}}</span>
		</div>
	</div>
	<div id="wrap" class="col-sm-12 col-md-12 col-lg-12" style="overflow: auto;">
		<div id="h-system-service" class="col-sm-12 col-md-6 col-lg-4"></div>
		<div id="h-mas-service" class="col-sm-12 col-md-6 col-lg-4"></div>
		<div id="h-other-service"  class="col-sm-12 col-md-6 col-lg-4"></div>
	</div>
</div>
<!--导航栏,标签切换栏, 修改为隐藏-->
<div class="H-content-tab theme-tab-color">
	<div class="H-tab-bar pull-left" id="H-tab-left">
		<button class="H-left-tab theme-tab-color" onclick="Hutils.H_HomePage()"><i style="color: white" class="icon-th-large"></i></button>
		<nav class="H-tabs-index"></nav>
	</div>
	<div class="H-tab-bar pull-right" id="H-tab-right">
		<button data-toggle="tooltip" title="显示菜单栏" class="H-right-tab theme-tab-color" onclick="Hutils.HchangeWrapper()"><i style="color: white" class="icon-columns"></i></button>
		<button data-toggle="tooltip" title="安全退出" class="H-right-tab theme-tab-color" onclick="Hutils.HLogOut()"><i style="color: white" class="icon-off"></i></button>
		<button data-toggle="tooltip" title="用户信息" class="H-right-tab theme-tab-color" onclick="Hutils.UserMgrInfo()"><i style="color: white" class="icon-user"></i></button>
	</div>
</div>

<script type="text/javascript">
    NProgress.start();
    var indexObj = {
		/*
		 * 调整元素位置
		 * 使其铺满全屏
		 * */
        adjustLocation:function(){
            var hh = document.documentElement.clientHeight;
            $("#wrap").height(hh-96);
        },
		/*
		 * 绑定一系列事件
		 * */
        bindEvents:function(){
			/*绑定鼠标指向事件
			 * 鼠标指过去时,背景变成白色,前景色成黑色
			 * */
            $(".H-right-tab").on("mouseover",function(){
                $(this).find("i").css("color","black");
            }).on("mouseout",function(){
                $(this).find("i").css("color","white");
            })
            $(".H-left-tab").on("mouseover",function(){
                $(this).find("i").css("color","black");
            }).on("mouseout",function(){
                $(this).find("i").css("color","white");
            })
			/*
			 * 开启bootstrap的title提示特效
			 * */
            $("[data-toggle='tooltip']").tooltip();
        }
    };

	/*
	 * 禁用浏览器后退按钮
	 * */
    window.onload = function () {
        if (typeof history.pushState === "function") {
            history.pushState("jibberish", null, null);
            window.onpopstate = function () {
                history.pushState('newjibberish', null, null);
                Hutils.H_HomePage()
            };
        }
        else {
            var ignoreHashChange = true;
            window.onhashchange = function () {
                if (!ignoreHashChange) {
                    ignoreHashChange = true;
                    window.location.hash = Math.random();
                }
                else {
                    ignoreHashChange = false;
                }
            };
        }
    };

    var changeTheme = function (id) {
		$.HAjaxRequest({
		    url:"/v1/auth/theme/update",
			type:'post',
			dataType:'json',
			data:{theme_id:id},
			success:function () {
				window.location.href="/HomePage"
            },
		})
    };

    var changemodifypassword = function(){
        $.Hmodal({
            header:"密码修改",
            body:$("#h-user-modify-password").html(),
            height:"420px",
            width:"720px",
            preprocess:function () {
                var user_id = $("#h-user-details-user-id").html()
                $("#h-modify-user-id").val(user_id)
            },
            callback:function(hmode){
                var newpd = $("#plat-change-passwd").find('input[name="newpasswd"]').val()
                var orapd = $("#plat-change-passwd").find('input[name="orapasswd"]').val()
                var surpd = $("#plat-change-passwd").find('input[name="surepasswd"]').val()
                if ($.trim(newpd) =="" || $.trim(orapd) == "" || $.trim(surpd)  == "" ){
                    $.Notify({
                        title:"温馨提示",
                        message:"不能将密码设置成空格",
                        type:"danger",
                    })
                    return
                }else if(newpd != surpd){
                    $.Notify({
                        title:"温馨提示",
                        message:"两次输入的新密码不一致，请确认是否存在多余的空格",
                        type:"danger",
                    })
                    return
                }
                $.HAjaxRequest({
                    type:"post",
                    url:"/v1/auth/passwd/update",
                    data:$("#plat-change-passwd").serialize(),
                    dataType:"json",
                    success:function(){
                        $(hmode).remove();
                        $.Notify({
                            title:"执行成功",
                            message:"修改密码成功",
                        })
                    },
                });
            }
        })
    };

    //调整主菜单的长度和宽度
    $(document).ready(function(){
        $.i18n.properties({
            name:'Messages',
            path:'/static/jquery-i18n-properties/bundle/',
            mode:'both',
            language :(navigator.language || navigator.browserLanguage).toLowerCase(),
            async: true,
        });

        Hutils.initMenu(0,-1,"系统服务","管理会计","公共信息")
        indexObj.adjustLocation()
        indexObj.bindEvents()
        NProgress.done();
    });

    window.onresize = function(){
        var hh = document.documentElement.clientHeight;
        $("#wrap").height(hh-96);
    }
</script>

<script id="mas-passwd-prop" type="text/html">
	<div class="panel panel-default">
		<!-- Default panel contents -->
		<div class="panel-heading">
			<span style="font-size: 12px;font-weight: 600;">主题切换：</span>
			<button onclick="changeTheme(1001)" class="btn btn-sm theme-green-color" style="color: white;">
			</button>
			<button onclick="changeTheme(1004)" class="btn btn-sm theme-cyan-color" style="color: white;">
			</button>
			<button onclick="changeTheme(1002)" class="btn btn-sm theme-blue-color" style="color: white;">
			</button>
			<button onclick="changeTheme(1003)" class="btn btn-sm theme-apple-color" style="color: white;">
			</button>
			<div class="pull-right">
				<button onclick="changemodifypassword()" class="btn btn-success btn-xs">
					<i class="icon-wrench"> 修改密码</i>
				</button>
			</div>
		</div>
		<table class="table table-bordered table-responsive">
			<tr style="height: 36px;line-height: 36px;">
				<td style="text-align: right;">用户id:&nbsp;</td>
				<td id="h-user-details-user-id" style="font-weight: 600">user_id</td>
				<td style="text-align: right;">用户名称:&nbsp;</td>
				<td id="h-user-details-user-name" style="font-weight: 600">user_name</td>
			</tr>
			<tr style="height: 36px;line-height: 36px;">
				<td style="text-align: right;">邮箱:&nbsp;</td>
				<td id="h-user-details-user-email" style="font-weight: 600">user_email</td>
				<td style="text-align: right;">手机号:&nbsp;</td>
				<td id="h-user-details-user-phone" style="font-weight: 600">user_phone</td>
			</tr>
			<tr style="height: 36px;line-height: 36px;">
				<td style="text-align: right;">所属域编码:&nbsp;</td>
				<td id="h-user-details-user-domain" style="font-weight: 600">user_dept</td>
				<td style="text-align: right;">所属域名称:&nbsp;</td>
				<td id="h-user-details-user-domain-name" style="font-weight: 600">user_domain</td>
			</tr>
			<tr style="height: 36px;line-height: 36px;">
				<td style="text-align: right;">组织部门编码:&nbsp;</td>
				<td id="h-user-details-user-org" style="font-weight: 600">user_dept</td>
				<td style="text-align: right;">组织部门描述:&nbsp;</td>
				<td id="h-user-details-user-org-name" style="font-weight: 600">user_domain</td>
			</tr>
			<tr style="height: 36px;line-height: 36px;">
				<td style="text-align: right;">创建人:&nbsp;</td>
				<td id="h-user-details-user-create" style="font-weight: 600">user_create</td>
				<td style="text-align: right;">创建时间:&nbsp;</td>
				<td id="h-user-details-user-create-date" style="font-weight: 600">user_create_date</td>
			</tr>
			<tr style="height: 36px;line-height: 36px;">
				<td style="text-align: right;">修改人:&nbsp;</td>
				<td id="h-user-details-user-modify" style="font-weight: 600">user_create</td>
				<td style="text-align: right;">修改时间:&nbsp;</td>
				<td id="h-user-details-user-modify-date" style="font-weight: 600">user_create_date</td>
			</tr>
		</table>
	</div>

</script>
<script id="h-user-modify-password" type="text/html">
	<form id="plat-change-passwd" class="col-sm-12 col-md-12 col-lg-12">
		<div class="form-group col-sm-12 col-md-12 col-lg-12">
			<label class="h-label" style="width: 100%;">账　号：</label>
			<input id="h-modify-user-id" readonly="readonly" class="form-control" style="width: 100%;height: 30px; line-height: 30px;" type="text" name="userid"/>
		</div>
		<div class="form-group col-sm-12 col-md-12 col-lg-12">
			<label class="h-label" style="width: 100%;">原密码：</label>
			<input placeholder="密码长度必须大于6位，小于30位" class="form-control" style="width:100%;height: 30px; line-height: 30px;" type="password" name="orapasswd"/>
		</div>
		<div class="form-group col-sm-12 col-md-12 col-lg-12">
			<label class="h-label" style="width: 100%;">新密码：</label>
			<input placeholder="密码长度必须大于6位，小于30位" class="form-control" style="width:100%;height: 30px; line-height: 30px;" type="password" name="newpasswd"/>
		</div>
		<div class="form-group col-sm-12 col-md-12 col-lg-12">
			<label class="h-label" style="width: 100%;">确认密码：</label>
			<input placeholder="请确认新密码信息" class="form-control" style="height: 30px; line-height: 30px; width: 100%;" type="password" name="surepasswd"/>
		</div>
	</form>
</script>
<script type="text/javascript" src="/static/laydate/laydate.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-notify.min.js"></script>
<script src="/static/js/download.js"></script>
<script type="text/javascript" src="/static/js/spin.min.js"></script>
</body>
</html>