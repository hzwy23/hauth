<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>大数据应用分析平台</title>

    <link rel="stylesheet" href="/static/bootstrap-3.3.7-dist/css/bootstrap.min.css" />

    <link rel="stylesheet" href="/static/Font-Awesome-3.2.1/css/font-awesome.min.css" />

    <link rel="stylesheet" href="/static/css/animate.css"/>

    <link rel="stylesheet" href="/static/theme/default/index.css"/>

    <link rel="stylesheet" href="/static/nprogress/nprogress.css"/>

    <!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
    <script src="/static/js/jquery-3.1.1.min.js"></script>

    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>

    <script src="/static/nprogress/nprogress.js"></script>

    <style type="text/css">
        .h-github-source {
            cursor: pointer;
            border: #3caff3 solid 1px;
            height: 40px;
            line-height: 40px;
            color: #FFFFFF;
            font-size: 14px;
            font-weight: 600;
            width: 240px;
        }
        .h-github-source:hover{
            background-color: #cccccc;
            color: #0f0f0f;
        }
    </style>
</head>
<body style="background-size: cover;background-image: url('/static/images/index_bg.jpg')">
<div class="container-fluid">
    <nav class="navbar navbar-fixed-top" role="navigation">
        <div class="pull-left">
            <h3 style="height: 50px; line-height: 50px; margin-left: 15px; color: white;">H 集成开发平台</h3>
        </div>
        <div  class="pull-right">
            <div id="h-login-form" class="pull-left" style="display: none; height: 50px; line-height: 50px;">
                <form id='inputZoon' class="navbar-form navbar-right" style="height: 35px;line-height: 35px;">
                    <div class="form-group input-group-sm">
                        <input name="username" value="caadmin" type="text" class="form-control"
                               placeholder="账号">
                    </div>
                    &nbsp;
                    &nbsp;
                    <div class="form-group input-group-sm">
                        <input name="password" value="123456" type="password" class="form-control"
                               placeholder="密码">
                    </div>
                    &nbsp;
                    <div class="btn-group btn-group-sm">
                        <button onclick="LoginSubmit(this)" type="button" class="btn btn-info">登&nbsp;录</button>
                    </div>
                </form>
            </div>
            <div class="pull-right" style="color: white;">
                <div class="btn btn-link-sm" style="font-weight: 600;margin-top: 11px;font-size: 12px;" onclick="Hlogin(this)"> 登录</div>
                <div class="btn btn-link-sm" style="font-weight: 600;margin-top: 11px;font-size: 12px;" onclick="Hregister()"> 注册</div>
            </div>
        </div>
    </nav>
    <div class="row" style=" margin-top: 120px; text-align: center;height: 320px;padding-top: 60px;">
        <div class="col-sm-4 col-md-4 col-lg-4">
            <div class="col-sm-8 col-md-8 col-lg-8 col-sm-offset-4" style="color: #FFFFFF;text-align: center;">
                <span class="icon-4x icon-magnet"></span>
            </div>
            &nbsp;
            <div class="col-sm-8 col-md-8 col-lg-8 col-sm-offset-4">
                <span style="color: #FFFFFF">
                    <p>成熟的解决方案</p>
                    <p>完全开发源代码</p>
                </span>
            </div>

        </div>
        <div class="col-sm-4 col-md-4 col-lg-4">
            <div class="col-sm-12 col-md-12 col-lg-12" style="color: #FFFFFF;text-align: center;">
                <span class="icon-4x icon-signal"></span>
            </div>
            &nbsp;
            <div class="col-sm-12 col-md-12 col-lg-12">
                <span style="color: #FFFFFF">
                    <p>去session化</p>
                    <p>Metro风格</p>
                    <p>多主题切换</p>
                </span>
            </div>
        </div>
        <div class="col-sm-4 col-md-4 col-lg-4">
            <div class="col-sm-8 col-md-8 col-lg-8" style="color: #FFFFFF;text-align: center;">
                <span class="icon-4x icon-group"></span>
            </div>

            <div class="col-sm-8 col-md-8 col-lg-8">
                &nbsp;
                <span style="color: #FFFFFF">

                    <p>beego基础开发</p>
                    <p>丰富的API接口</p>
                </span>
            </div>

        </div>
    </div>
    <div class="row" style="text-align: center; height: 120px;">
        <div class="col-sm-4 col-md-4 col-lg-4">
        </div>
        <div class="col-sm-4 col-md-4 col-lg-4" align="center">
            <div onclick="OpenGithub()" class="h-github-source">
                源码下载
            </div>
        </div>
        <div class="col-sm-4 col-md-4 col-lg-4">
        </div>
    </div>
</div>
</body>
<script type="text/javascript" src="/static/jquery-i18n-properties/jquery.i18n.properties.min.js"></script>
<script type="text/javascript" src="/static/js/utils.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-notify.min.js"></script>
<script>
    NProgress.start();
    function LoginSubmit(obj){
        $.HAjaxRequest({
            url:"/login",
            type:"post",
            data:$("#inputZoon").serialize(),
            dataType:'json',
            success:function(){

                window.location.href="/HomePage";
            },
            error:function(msg){
                var imsg = JSON.parse(msg.responseText)
                if (imsg.error_code == 401){
                    $.Notify({
                        title:"温馨提示",
                        message:"用户名不存在,请检查您的账号是否输入错误",
                        type:"warning",
                        placement:{from:"bottom",align:"right"},
                    })
                } else if (imsg.error_code==402){
                    $.Notify({
                        title:"温馨提示",
                        message:"用户名不存在,系统中无法获取用户密码信息",
                        type:"warning",
                        placement:{from:"bottom",align:"right"},
                    })
                } else if (imsg.error_code==403){
                    $.Notify({
                        title:"温馨提示",
                        message:"已经连续6次输错密码,用户已被锁定,请联系管理员",
                        type:"warning",
                        placement:{from:"bottom",align:"right"},
                    })
                } else if (imsg.error_code==404){
                      $.Notify({
                          title:"温馨提示",
                          message:"已经连续6次输错密码,用户已被锁定,请联系管理员",
                          type:"warning",
                          placement:{from:"bottom",align:"right"},
                      })
                  }  else if (imsg.error_code==405){
                      $.Notify({
                          title:"温馨提示",
                          message:"用户密码输入错误,如忘记密码,请联系管理员",
                          type:"warning",
                          placement:{from:"bottom",align:"right"},
                      })
                  } else if (imsg.error_code==406){
                      $.Notify({
                          title:"温馨提示",
                          message:"用户已被锁定,请联系管理员",
                          type:"warning",
                          placement:{from:"bottom",align:"right"},
                      })
                  }
              }
          });

    };

    function Hlogin(obj){
        $("#h-login-form").show();
        $(obj).hide();
    };

    function  OpenGithub() {
    };

    function Hregister(){
        $.Notify({
            message:"演示环境，暂不支持用户注册",
            type:"info"
        })
//        window.location.href="/plat/registerPage";
    };

    $(document).keydown(function(e){
        if (e.keyCode == '13'){
            setTimeout(LoginSubmit,200)
        }
    });

    $(document).ready(function () {
        $.i18n.properties({
            name:'Messages',
            path:'/static/jquery-i18n-properties/bundle/',
            mode:'both',
            language :(navigator.language || navigator.browserLanguage).toLowerCase(),
            async: true,
        });

        var $body = $(".container-fluid");
        $body.height(document.documentElement.clientHeight)
//        $body.css("background-image","url('/static/images/index_bg.png')");
        NProgress.done();
    });
    
    window.onresize = function () {
        $(".container-fluid").height(document.documentElement.clientHeight)
    };


</script>
</html>