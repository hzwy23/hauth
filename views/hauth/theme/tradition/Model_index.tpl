<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>数据管理系统</title>
    <link href="/static/css/bootstrap-table.min.css" type="text/css" rel="stylesheet"/>
    <link rel="stylesheet" href="/static/css/index.css"/>
    <link href="/static/css/default/dark.theme.css" type="text/css" rel="stylesheet"/>

    <script type="text/javascript" src="/static/js/bootstrap-table.min.js" />
    <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"/>

    <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"/>
    <script type="text/javascript" src="/static/js/bootstrap-table-tableExport.js"/>

    <script type="text/javascript" src="/static/js/hzwy23table-from-bootstrap-table.js"/>
    <script src="/static/js/hzwy23.js"></script>

</head>
<body>
<div class="container-fluid">
    <div class="row">
        <nav id="wrapper" class="navbar-static-side">
            <div class="row H-logo-area">
                <div class="dropdown col-sm-12 col-md-12 col-lg-12">
                    <span  style="text-align: center; display: block; height: 64px; line-height: 64px; margin-top: 20px;"><img alt="image" class="img-circle" src="/static/img/profile_small.jpg"></span>
                    <a data-toggle="dropdown" class="dropdown-toggle" href="#" aria-expanded="false">
                        <span class="clear" style="display: block;height: 46px; line-height: 46px;">
                            <span class=" m-t-xs" style="display: block; height: 23px; line-height: 23px; font-size: 12px; text-align: center; color: #f3f3f4;"><strong class="font-bold">Hzwy23</strong></span>
                            <span class="text-muted text-xs " style="display: block; height: 23px; line-height: 23px; font-size: 12px; text-align: center; color: #a7b1c2">超级管理员<b class="caret"></b></span>
                        </span>
                    </a>
                    <ul class="dropdown-menu" style="margin-left: 30px; text-align: center;">
                        <li><a data-index="0">修改头像</a>
                        </li>
                        <li><a data-index="1">个人资料</a>
                        </li>
                        <li><a data-index="2">联系我们</a>
                        </li>
                        <li class="divider"></li>
                        <li><a onclick="LogOut()" style="cursor: pointer">安全退出</a>
                        </li>
                    </ul>
                </div>
            </div>
            <div class="row" id="H-main-menu">
                <ul class="col-sm-12 col-md-12 col-lg-12"></ul>
            </div>
        </nav>
        <div id="page-wrapper" class="container-fluid gray-bg">
            <div class="row H-main-content">
                <div class="col-md-12 col-sm-12 col-lg-12" id="h-main-content">
                    <div data-id="homepage" class="active">
                        <div id="toolbar">
                            <button onclick="insertRow()" class="btn btn-default"><i class="icon-plus"></i></button>
                            <button onclick="deleteRow()" class="btn btn-default"><i class="icon-trash"></i></button>
                            <button onclick="editRow()" class="btn btn-default"><i class="icon-edit"></i></button>
                        </div>
                        <table id="table"></table>
                    </div>
                </div>
            </div>
            <!--<div class="row H-footer">-->
                <!--<div class="pull-right">© 2014-2015 &nbsp;&nbsp;<a  target="_blank" href="http://www.asofdate.com">www.asofdate.com</a>-->
                <!--</div>-->
            <!--</div>-->
            <div class="row H-content-tab">
                <div class="H-tab-bar pull-left" id="H-tab-left">
                    <button class="H-left-tab" onclick="leftTabShow()"><i class="glyphicon glyphicon-backward"></i></button>
                    <button class="H-left-tab active-tab" data-id="homepage" onclick="changetab(this)">首页</button>
                    <nav class="H-tabs-index"></nav>
                </div>
                <div class="H-tab-bar pull-right" id="H-tab-right">
                    <button class="H-right-tab H-right-tab-padding" onclick="backIndex()"><i class="icon-reply"></i>&nbsp;返回</button>
                    <div class="H-right-tab">
                        <button class="dropdown" type="button" id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
                            关闭操作<span class="caret"></span>
                        </button>
                        <ul role="menu" class="dropdown-menu dropdown-menu-right" style="margin-right: 36px; padding:auto 15px; text-align: center;">
                            <li onclick="lockCurrentTab()"><a>锁定当前选项卡</a></li>
                            <li class="divider"></li>
                            <li onclick="closeOtherTab()"><a>关闭其他选项卡</a></li>
                            <li onclick="closeAllTab()"><a>关闭全部选项卡</a></li>
                        </ul>
                    </div>
                    <button class="H-right-tab H-right-tab-padding" onclick="rightTabShow()"><i class="glyphicon glyphicon-forward"></i></button>
                </div>
            </div>
        </div>
    </div>
</div>
<div>

</div>
<script type="text/javascript">
    $(document).ready(function(){
        var hwindow = document.documentElement.clientHeight;
        var wwindow = document.documentElement.clientWidth;
        $("#wrapper").height(hwindow);
        $("#page-wrapper").height(hwindow);
        $("#h-main-content").height(hwindow - 42);
    });

    $(document).ready(function(){
        $.ajax({
            type:"get",
            url:"/platform/menu",
            cache:false,
            async:false,
            data:{Id:{{.Menu_Id}}},
            dataType:"json",
            error:function(){
                setTimeout(redirectLoginPage,300);
            },
            success: function(data){
                var opthtml = ""
                $(data).each(function(index,element){
                    opthtml += '<li data-dept="'+element.Menu_Dept+'" data-url="'+element.Menu_route+'" data-node="'+element.Menu_leaf_flag+'" onclick="platMenuHande(this)" data-id="'+element.Menu_id+'"><span class="'+element.Menu_icon+'"></span>&nbsp;<hzw>'+element.Menu_name+'</hzw><i class="icon-angle-left"></i></li> '
                })
                $("#H-main-menu").find("ul").html(opthtml);
            }
        });

        $("#H-main-menu").find("li").each(function(index,element){
            var dept = $(element).attr("data-dept")
            var pd = parseInt(dept) * 15
            var node = $(element).attr("data-node")

            if (node == "1"){
                $(element).find("i").removeClass("icon-angle-left")
            }
            $(element).css("padding-left",pd)
            if (dept == "1"){
                $(element).addClass("nav-root-level")
                $(element).show();
            }else{
                $(element).addClass("nav-leaf-level")
                $(element).find("span").remove()
                $(element).hide();
            }
        });

        $(".navbar-static-side").mCustomScrollbar({
            axis:"y",
            theme:"dark-thin",
            scrollbarPosition:"outside",
            scrollSpeed:100,
        });
        $(".mCSB_scrollTools").css("width","56px");

    });

    var platMenuHande = function(e){
        var curDept = $(e).attr("data-dept");
        var curNode = $(e).attr("data-node");
        if (curNode == "1"){
            platMenuService(e,curDept)
        }else if (curNode == "0"){
            platMenuDisp(e,curDept)
        }
    };

    var platMenuService = function(e,dept){
        var flag = false;
        var url = $(e).attr("data-url");
        var data_id = $(e).attr("data-id");
        var name = $(e).find("hzw").html();
        var optHtml = '<span data-id="'+data_id+'" class="H-left-tab active-tab" onclick="changetab(this)">'+name+'<i onclick="closetab(this)" class="icon-remove-sign H-gray-close"></i></span>'

        $(".H-tabs-index").find("span").each(function(index,element){
            if (data_id == $(element).attr("data-id")){
                changetab(element)
                flag = true;
                return false;
            }
        });

        if (flag == false){
            $.ajax({
                type:"get",
                url:url,
                cache:false,
                async:true,
                dataType:"text", error: function(){
                    setTimeout(redirectLoginPage,300);
                },
                success: function(data){
                    $(".active-tab").removeClass("active-tab");
                    $(".H-tabs-index").append(optHtml);

                    (function(){
                        do{
                            var maxLT = (function(){
                                var rt = $("#H-tab-right").width();
                                var sl = $("#wrapper").width();
                                var ww = document.documentElement.clientWidth;
                                return ww - sl - rt;
                            })();
                            var lt = $("#H-tab-left").width();
                            if (lt >= maxLT - 20){
                                $(".H-tabs-index").find("span:visible:eq(0)").hide();
                            }
                        }while(lt >= maxLT - 20)
                    })()

                    $("#h-main-content").find("div.active").removeClass("active").addClass("none")
                    var cot = '<div data-type="frame" data-id="'+data_id+'" class="active">'+data+'</div>';
                    $("#h-main-content").append(cot);
                }
            });
        }
    };

    var platMenuDisp = function(e,dept){
        var pDept = parseInt(dept);
        var nextLi = $(e).next("li");
        var nextLiStatus = $(nextLi).css("display");

        if ($(e).attr("data-dept") == 1){
            $(e).closest("ul").find("li").each(function(index, element) {
                if(parseInt($(element).attr("data-dept"))>1){
                    $(element).css("display","none");
                    $(element).find("i.icon-angle-down").removeClass("icon-angle-down").addClass("icon-angle-left");
                }else if(parseInt($(element).attr("data-dept"))==1){
                    $(element).find("i.icon-angle-down").removeClass("icon-angle-down").addClass("icon-angle-left");
                }
            });
        }



        var hzw = $("#H-main-menu").find("hzw").css("display")
        if (hzw == "none"){
            $("#H-main-menu li span").hide();
            $("#wrapper").css("width","220px");
            $(".nav-root-level").css("width","220px");
            $("hzw").show("fast");
            $("#H-main-menu li span").removeClass("icon-2x").show("fast");
            $(".H-logo-area").css("width","220px");
            $("#H-main-menu").find("i").show("slow");
            setTimeout(function(){$(".H-logo-area").find("div").show("fast");},300)
        }


        if (nextLiStatus == "none"){
            $(e).find("i.icon-angle-left").removeClass("icon-angle-left").addClass("icon-angle-down");

            $(e).nextAll("li").each(function(index,element){
                var CurDept = parseInt($(element).attr("data-dept"));
                if (CurDept == pDept + 1) {
                    $(element).show();
                    $(element).find("hzw").show();
                } else if (CurDept > pDept + 1){
                    $(element).hide("fast");
                    $(element).find("hzw").hide("fast");
                } else {
                    return false
                }
            })
        } else {

            $(e).find("i.icon-angle-down").removeClass("icon-angle-down").addClass("icon-angle-left");

            $(e).nextAll("li").each(function(index,element){
                var CurDept = parseInt($(element).attr("data-dept"));
                if (CurDept >= pDept + 1) {
                    $(element).hide("fast");
                    $(element).find("hzw").hide("fast");
                } else {
                    return false
                }
            })
        }
    }

    var menubartoggle = function(){
        event.stopPropagation();
        var hzw = $("#H-main-menu").find("hzw").css("display")
        if (hzw == "none"){
            $("#H-main-menu li span").hide();
            $("#wrapper").css("width","220px");
            $(".nav-root-level").css("width","220px");
            $("hzw").show("fast");
            $("#H-main-menu li span").removeClass("icon-2x").show("fast");
            $(".H-logo-area").css("width","220px");
            $("#H-main-menu").find("i").show("slow");
            setTimeout(function(){$(".H-logo-area").find("div").show("fast");},300)
        } else {
            $("#H-main-menu li span").hide();
            $(".H-logo-area").find("div").hide("slow");
            $(".H-logo-area").css("width","60px");
            $("#H-main-menu").find("i").hide("slow");
            $("#H-main-menu").find("i.icon-angle-down").removeClass("icon-angle-down").addClass("icon-angle-left");
            $(".nav-leaf-level").hide();
            $("#wrapper").css("width","60px");
            $(".nav-root-level").css("width","60px")
            $("hzw").hide("slow");
            setTimeout(function(){$("#H-main-menu li span").addClass("icon-2x").show();},300);
        }
    }

    var backIndex =  function(){
        $.ajax({
            type:"post",
            url:"/platform/IndexPage",
            cache:false,
            async:true,
            dataType:"text", error: function(){
                setTimeout(redirectLoginPage,300);
            },
            success: function(data){
                $("body").html(data);
            }
        });
    };

    var closetab = function(e){
        var id = $(e).parent().attr("data-id");
        if ($(e).parent().hasClass("active-tab")){
            var pobj = $(e).parent().prev("span");
            var pid = $(pobj).attr("data-id");
            var nobj = $(e).parent().next("span");
            var nid = $(nobj).attr("data-id");

            $(e).parent().remove();
            $("#h-main-content").find("div[data-id='"+id+"']").remove();
            if (pid == undefined){
                if (nid == undefined){
                    id = "homepage"
                } else {
                    id = nid
                }
            } else {
                id = pid
            }

            $("#h-main-content").find("div[data-id='"+id+"']").removeClass("none").addClass("active");
            $(".H-left-tab").each(function(index,element){
                if (id == $(element).attr("data-id")){
                    $(element).addClass("active-tab")
                }
            });
            var leftStyle = leftTabShow();
            if (leftStyle == false){
                rightTabShow()
            }
        }else{
            $(e).parent().remove();
            $("#h-main-content").find("div[data-id='"+id+"']").remove();
            var leftStyle = leftTabShow();
            if (leftStyle == false){
                rightTabShow()
            }
        }
        window.event.cancelBubble = true;
    };

    var leftTabShow = function(){
        var firstObj = $(".H-tabs-index").find("span:visible:eq(0)")
        var preOjb = $(firstObj).prev("span").attr("data-id")
        if (preOjb == undefined){
            return false
        }else{
            var lastShowItem = null;
            $(firstObj).nextAll("span").each(function(index,element){
                if ("none" ==  $(element).css("display")){
                    return false;
                }else{
                    lastShowItem = element;
                }
            });

            $(firstObj).prev("span").show();
            do {
                var maxLT = (function(){
                    var rt = $("#H-tab-right").width();
                    var sl = $("#wrapper").width();
                    var ww = document.documentElement.clientWidth;
                    return ww - sl - rt;
                })();
                var lt = $("#H-tab-left").width();
                if (lt >= maxLT - 20){
                    $(lastShowItem).hide();
                    lastShowItem = $(lastShowItem).prev("span");
                }
            }while( lt >= maxLT -20)
            return true;
        }
    };

    var rightTabShow = function(){
        var firstObj = $(".H-tabs-index").find("span:visible:eq(0)")
        $(firstObj).nextAll("span").each(function(index,element){
            if ("none" ==  $(element).css("display")){
                $(element).show();
                do {
                    var maxLT = (function(){
                        var rt = $("#H-tab-right").width();
                        var sl = $("#wrapper").width();
                        var ww = document.documentElement.clientWidth;
                        return ww - sl - rt;
                    })();
                    var lt = $("#H-tab-left").width();
                    if (lt >= maxLT - 20){
                        $(".H-tabs-index").find("span:visible:eq(0)").hide();
                    }
                }while( lt >= maxLT -20)
                return false;
            }
        });
    }

    var changetab = function(e){
        $(".active-tab").removeClass("active-tab")
        $(e).addClass("active-tab")
        var id = $(e).attr("data-id");
        $("#h-main-content").find("div.active").removeClass("active").addClass("none")
        $("#h-main-content").find("div[data-id='"+id+"']").removeClass("none").addClass("active");
    };

    var closeAllTab = function(){
        $(".H-tabs-index").find("span").remove();
        $("#h-main-content").find("div[data-type='frame']").remove();
        $("#h-main-content").find("div[data-id='homepage']").removeClass("none").addClass("active");
        $(".H-left-tab").each(function(index,element){
            if ("homepage" == $(element).attr("data-id")){
                $(element).addClass("active-tab")
            }
        })
    };

    var closeOtherTab = function(){
        var id = new Array();
        var i = 0;
        $(".H-tabs-index").find("span").each(function(index,element){
            if ($(element).hasClass("active-tab") || $(element).hasClass("tab-lock")){
                id[i++] = $(element).attr("data-id");
            } else {
                $(element).remove()
            }
        });

        $("#h-main-content").find("div[data-type='frame']").each(function(index,element){
            if( id.indexOf($(element).attr("data-id")) > -1){

            }else{
                $(element).remove()
            }
        });
    };

    var lockCurrentTab = function(){
        $(".H-tabs-index").find("span.active-tab").addClass("tab-lock").find("i").remove()
    };

    var getHomePage = function(){
        $.ajax({
            type:"get",
            url:"/platform/HomePage",
            cache:false,
            async:true,
            dataType:"text", error: function(){
                setTimeout(redirectLoginPage,300);
            },
            success: function(data){
                $("#h-main-content").html("<div data-id='homepage' class='active'>"+data+"</div>")
            }
        });
    }


    $(document).ready(function(e){
        $("#DomainPageTable").hzwTable({url:'/platform/DomainMgr/page',toolbar:'#DomainPageTools',uniqueId:'Project_id'});

        var insertRow = function(e){
            alert("hello")
            $("#DomainPageTable").hzwInsertRow({index:0,row:{
                Project_id: 'helwro',
                Project_name: 'Item ',
                Project_status: '$',
                Maintance_date: '2016-07-29',
                User_id: 'admin'}
            })
        };
        var deleteRow = function(){
            $("#DomainPageTable").hzwDeleteRow({filed:'Project_id'})
        };
        var editRow = function(){
            $("#DomainPageTable").hzwEditRow({filed:'Project_id',row:{
                Project_name: 'Item ', Project_status: '$'
            }})
        };
    })


    $("#table").hzwTable({
        url: '/platform/DomainMgr',
        toolbar:'#toolbar',
        uniqueId:'Project_id',
        columns:[{

            field: 'state',

            radio: true

        }, {

            field: 'Project_id',

            title: '域名编号',

            align: 'left',

            valign: 'middle',

            sortable: true

        }, {

            field: 'Project_name',

            title: '域名描述',

            align: 'left',

            valign: 'middle',

            sortable: true,

        }, {

            field: 'Project_status',

            title: '域名状态',

            align: 'left',

            valign: 'top',

            sortable: true

        }, {

            field: 'Maintance_date',

            title: '创建日期',

            align: 'left',

            valign: 'center',

            sortable: true

        }, {

            field: 'User_id',

            title: '用户',

            align: 'left',

            valign: 'middle',

            sortable: true

        }]
    });
    var insertRow = function(e){
        $("#table").hzwInsertRow({index:0,row:{
            Project_id: 'helwro',
            Project_name: 'Item ',
            Project_status: '$',
            Maintance_date: '2016-07-29',
            User_id: 'admin'}
        })
    };
    var deleteRow = function(){
        $("#table").hzwDeleteRow({filed:'Project_id'})
    };
    var editRow = function(){
        $("#table").hzwEditRow({filed:'Project_id',row:{
            Project_name: 'Item ', Project_status: '$'
        }})
    };

    $("#page-wrapper").click(function(e){

        var hzw = $("#H-main-menu").find("hzw").css("display")
        if (hzw != "none"){
            //$("#wrapper").css("transform","rotateY(75deg)");
            $("#H-main-menu li span").hide();
            $(".H-logo-area").find("div").hide("slow");
            $(".H-logo-area").css("width","60px");
            $("#H-main-menu").find("i").hide("slow");
            $("#H-main-menu").find("i.icon-angle-down").removeClass("icon-angle-down").addClass("icon-angle-left");
            $(".nav-leaf-level").hide();
            $("#wrapper").css("width","60px");
            $(".nav-root-level").css("width","60px")
            $("hzw").hide("slow");
            setTimeout(function(){$("#H-main-menu li span").addClass("icon-2x").show();},300);
        }
    })

</script>

</body>
</html>
