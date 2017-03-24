{{define "header"}}
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>数据管理系统</title>
<!-- 新 Bootstrap 核心 CSS 文件 -->
	<link rel="stylesheet" href="/static/css/bootstrap.min.css" />
	<link rel="stylesheet" href="/static/css/font-awesome.min.css" />
	<link rel="stylesheet" href="/static/css/common.css" />
	<link rel="stylesheet" href="/static/css/index.css"/>
	<link href="/static/css/default/dark.theme.css" type="text/css" rel="stylesheet"/>

	<link rel="stylesheet" href="/static/css/jquery.mCustomScrollbar.min.css"/>

	<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
	<script src="/static/js/jquery-1.12.3.min.js"></script>
	<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
	<script src="/static/js/bootstrap.min.js"></script>
	<script src="/static/js/angular.min.js"></script>
	<script src="/static/js/jquery.validate.min.js"></script>
	<script src="/static/js/jquery.mCustomScrollbar.concat.min.js"></script>


	<script src="/static/js/modal.js"></script>
	<script src="/static/js/hzwy23.js"></script>


</head>
{{end}}
{{define "body"}}
<body ng-app="" style="background-color: #1a2a57; background-size:cover">
<div id="masFirstPageHeader" class="col-sm-12 col-sm-md-12 col-lg-12">
	<div id="logo" class="col-sm-2 col-md-2 col-lg-2" style="background-image: none !important;">
		<ul><li><strong style="color:#FFF">大数据分析平台</strong></li></ul>
	</div>
	<div id="main-menu-info" class="col-sm-10 col-md-10 col-lg-10">
	</div>
	<div class="mas-main-user-hand">
		<ul style="height: 66px; line-height: 66px;">
			<li style="cursor: pointer; float: left; margin: auto 6px;"><i class="icon-user" title="用户信息"> </i>
				<ul class="btn-group">
					<li type="button" class="dropdown-toggle"
						data-toggle="dropdown">
						{{.UserId}}<span class="caret"></span>
					</li>
					<ul class="dropdown-menu" role="menu" style="margin-left: -55px; background-color: transparent; background-image: url(/static/img/index_user_info_bg.png); border-radius: 15px; height: 140px;">
						<li style="color: #0f0f0f; text-align: center; font-weight: normal; font-size: 10px; height: 30px; line-height: 30px;">用户信息</li>
						<li style="color: #0f0f0f;text-align: center; font-weight: normal;font-size: 10px; height: 30px; line-height: 30px;" onclick="changePasswd()">密码修改</li>
						<li style="color: #0f0f0f;text-align: center; font-weight: normal;font-size: 10px; height: 30px; line-height: 30px;">风格设计</li>

					</ul>
				</ul>
			</li>
			<li style="cursor: pointer; float: left; margin: auto 6px;"><i class="icon-question-sign"> </i>帮助</li>
			<li onclick="LogOut()" style="float: left; cursor: pointer;margin: auto 6px;" title="退出系统" ><i class="icon-off"> </i>退出</li>
		</ul>
	</div>

</div>

<div id="mas-content-list" class="container" style=" padding-top:100px;">
	<div id="mas-menu-list-info" class="row col-sm-12 col-md-12 col-lg-12" style="text-align: center;">

	</div>
</div>
<script>
	/*
	 * 当浏览器不支持jquery时，通过dom操作调整内容大小。
	 * */
	window.onload = function(){
		if ((typeof $) == "undefined"){
			var hh = document.documentElement.clientHeight ;
			var mh = document.getElementById("masHeader").clientHeight;
			var w =  document.documentElement.clientWidth ;
			document.getElementById("mas-content").style.height = hh-mh-3 ;
			document.getElementById("left-content").style.height = hh - mh - 3;
			document.getElementById("right-content").style.height = hh-mh-3;
			document.getElementById("mas-show").style.height = hh-mh-3;
			document.getElementById("ajust-area").style.height = hh-mh-3;
			document.getElementById("right-content").style.width = w-198;
			document.getElementById("mas-content-footer").style.width = w-198;
		}
	};

	/*
	 * 调整页面初始高度，填充整个浏览器
	 *
	 * 缺陷：目前对于IE8之前的浏览器布局存在错误情况
	 *
	 */
	$(document).ready(function(e) {
		/*
		 * 调整页面高度与宽度
		 * #mas-content 表示显示内容
		 * #left-content 表示菜单栏目
		 * #right-content 表示右边显示区域大小
		 * #mas-show 表示每一个菜单内容
		 * #mas-content-footer 表示分区页面显示信息
		 * #ajust-area 表示调整页面大小的竖线条
		 * */
		var h = $(window).innerHeight()

		$("#mas-content-list").css("height",h)



		$("#mas-content-list").mCustomScrollbar({
			axis:"y",
			theme:"dark-thin",
			scrollbarPosition:"relative",
		});


	});
	/*
	 * 根据浏览器变化为调整内容大小
	 * */
	window.onresize=function(){
		/*
		 * 当浏览器窗口大小调整时，动态调整浏览器中各元素框的宽度与高度
		 * */
		var hh = $(window).innerHeight();
		var mh = $("#masHeader").height();
		var w = $(document).innerWidth();
		$("#masHeader").width(w);
		$("#mas-content").width(w);
		$("#mas-content").height(hh - mh - 3);
	}

	/*
	 * 加载菜单信息
	 * */
//	$(document).ready(function(){
//		$.ajax({
//			type:"get",
//			url:"/platform/menu",
//			cache:false,
//			async:false,
//			dataType:"text",
//			error: function(){
//				alert("获取菜单信息失败，请重新登陆");
//			},
//			success: function(data){
//				$("#menu-info").html(data)
//			}
//		});
//	});

	function changePasswd(){

		var submitPasswd = function(){
			$.ajax({
				type:"post",
				url:"/platform/passwd",
				data:$("#plat-change-passwd").serialize(),
				cache:false,
				async:false,
				dataType:"text",
				error: function(a,b,c){
					alert("error")
					alert(a.readyState)
					alert(a.responseText)
					alert(a.statusText)
				},
				success: function(data){
					alert("success")
					alert(data)
				}
			});
		}
		modal.newModal(submitPasswd,"密码修改",$("#mas-passwd-prop").html())
	}

	$(document).ready(function(){

		$.ajax({
			type:"get",
			url:"/platform/ResInfo",
			cache:false,
			data:{offset:0,limit:99999},
			async:false,
			dataType:"json",
			error: function(){
			},
			success: function(data){
				var opthtml = "";
				$(data.rows).each(function(index,element){
					if (element.Res_type == "0"){
						opthtml+='<div id='+element.Res_id+' onmouseover="indexMouseOn(this)" onmouseout="indexMouseOut(this)" onclick="go_entry(this)" class="col-sm-2 col-md-2 col-lg-2" style="height: 120px; cursor: pointer; width: 120px; margin: 15px 30px; padding: 15px 15px;"> <img src="'+element.Res_icon+'"  style="height: 70px;line-height: 70px;"/><br/><span style="font-weight: bold; color: #ffffff; height: 30px; line-height: 30px;">'+element.Res_name+'</span></div>'
					}
				});
				$("#mas-menu-list-info").html(opthtml);
			}
		});



	});

	function go_entry(e){
		var id = $(e).attr("id")
		var quit = function(){
			window.location.href="/"
		};
		$.ajax({
			type:"get",
			url:"/platform/select",
			data:{Id:id},
			cache:false,
			async:false,
			dataType:"text",
			error: function(){
				modal.confirm(quit,"连接已断开,请重新登录系统","off");
			},
			success: function(data){
				$("body").html(data)
			}
		});
	}

	function indexMouseOn(e){
		$(e).css("background-image","url(/static/img/index_menu_bg_m.png)")
	}

	function indexMouseOut(e){
		$(e).css("background-image","")
	}

	function redirectLoginPage(){

		var Ext = function(){
			$.ajax({
				type:"Get",
				url:"/logout",
				cache:false,
				async:false,
				dataType:"text",
				error: function(){
					window.location.href="/"
				},
				success: function(data){
					window.location.href="/"
				}
			});
		}
		//modal.confirm(Ext,"连接已断开,请重新登录系统","off");
	}

	function LogOut(){
		var Ext = function(){
			$.ajax({
				type:"Get",
				url:"/logout",
				cache:false,
				async:false,
				dataType:"text",
				error: function(){
					window.location.href="/"
				},
				success: function(data){
					window.location.href="/"
				}
			});
		}
		modal.confirm(Ext);
	}
</script>


<script id="mas-passwd-prop" type="text/html">
	<form id="plat-change-passwd" class="col-sm-12 col-md-12 col-lg-12" style="margin-top: 80px;">
		<table class="table table-condensed table-bordered">
			<tr>
				<td style="vertical-align: middle; width: 100px;"><label>原密码：</label></td>
				<td class="input-group"><input type="password" name="ora_passwd" class="form-control" style="border-right: none;"/><span class="input-group-addon" style="background-color: inherit; border-left:none;"></span></td>
			</tr>
			<tr>
				<td style="vertical-align: middle;width: 100px;"><label>新密码：</label></td>
				<td class="input-group"><input type="password" name="new_passwd" class="form-control" style="border-right: none;" /><span class="input-group-addon" style="background-color: inherit; border-left:none;"></span></td>
			</tr>
			<tr>
				<td style="vertical-align: middle;width: 100px;"><label>确　认：</label></td>
				<td class="input-group"><input type="password" name="sure_passwd" class="form-control" style="border-right: none;" /><span class="input-group-addon" style="background-color: inherit; border-left:none;"></span></td>
			</tr>
		</table>
	</form>
</script>

<script type="text/javascript">

	//处理键盘事件 禁止后退键（Backspace）密码或单行、多行文本框除外
	function banBackSpace(e){
		var ev = e || window.event;//获取event对象
		var obj = ev.target || ev.srcElement;//获取事件源

		var t = obj.type || obj.getAttribute('type');//获取事件源类型

//获取作为判断条件的事件类型
		var vReadOnly = obj.getAttribute('readonly');
		var vEnabled = obj.getAttribute('enabled');
//处理null值情况
		vReadOnly = (vReadOnly == null) ? false : vReadOnly;
		vEnabled = (vEnabled == null) ? true : vEnabled;

//当敲Backspace键时，事件源类型为密码或单行、多行文本的，
//并且readonly属性为true或enabled属性为false的，则退格键失效
		var flag1=(ev.keyCode == 8 && (t=="password" || t=="text" || t=="textarea")
		&& (vReadOnly==true || vEnabled!=true))?true:false;

//当敲Backspace键时，事件源类型非密码或单行、多行文本的，则退格键失效
		var flag2=(ev.keyCode == 8 && t != "password" && t != "text" && t != "textarea")
				?true:false;

//判断
		if(flag2){
			return false;
		}
		if(flag1){
			return false;
		}
	}

	//禁止后退键 作用于Firefox、Opera
	document.onkeypress=banBackSpace;
	//禁止后退键 作用于IE、Chrome
	document.onkeydown=banBackSpace;

</script>
</body>
{{end}}
{{define "footer"}}
</html>
{{end}}

