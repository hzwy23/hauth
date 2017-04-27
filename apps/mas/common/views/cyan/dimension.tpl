<div id="wrapper" class="navbar-static-side" style="margin: 0px;background-color: #3b565d">
    <div class="H-logo-area" style="margin: 0px; padding: 0 0 0 30px;">
        公共维度信息
    </div>
    <div class="col-sm-12 col-md-12 col-lg-12" id="H-main-menu" style="margin-bottom: 60px;">
        <div id="h-system-service" class="col-sm-12 col-md-6 col-lg-4">
        </div>
        <div id="h-mas-service" class="col-sm-12 col-md-6 col-lg-4">
        </div>
        <div id="h-other-service"  class="col-sm-12 col-md-6 col-lg-4">
        </div>
    </div>
</div>

<div id="page-wrapper" class="gray-bg col-sm-12 col-md-12 col-lg-12"
     style="margin:0px;padding: 0px;">
    <div id="h-main-content"
         style="padding: 0px 15px; margin: 0px;position: relative; overflow: auto;">
    </div>
</div>

<script type="text/javascript">
    $(document).ready(function(){
        var succ = function(data){
//            var h1 = '<div class="tile-group"><span class="tile-group-title">资源管理</span><div class="tile-container">'
//            var h2 = '<div class="tile-group"><span class="tile-group-title">用户与安全管理</span><div class="tile-container">'
//            var h3 = '<div class="tile-group"><span class="tile-group-title">审计</span><div class="tile-container">'
            var h1 = '<div class="tile-group"><div class="tile-container">'
            var h2 = '<div class="tile-group"><div class="tile-container">'
            var h3 = '<div class="tile-group"><div class="tile-container">'
            var h11 = ""
            var h22 = ""
            var h33 = ""

            $(data).each(function(index,element){
                if (element.Group_id == "1"){
                    h11 += '<div data-id="'+element.Res_id+'" data-url="'+element.Res_url+'" onclick="Hutils.goEntrySubSystem(this)" class="'+element.Res_class+' fg-white" data-role="tile" style="background-color: '+element.Res_bg_color+'">' +
                            '<div class="tile-content iconic"><span class="icon"><img src="'+element.Res_img+'"></span></div>' +
                            '<div class="tile-label">'+element.Res_name+'</div></div>'
                } else if (element.Group_id == "2"){
                    h22 += '<div data-id="'+element.Res_id+'" data-url="'+element.Res_url+'" onclick="Hutils.goEntrySubSystem(this)" class="'+element.Res_class+' fg-white" data-role="tile" style="background-color: '+element.Res_bg_color+'">' +
                            '<div class="tile-content iconic"><span class="icon"><img src="'+element.Res_img+'"></span></div>' +
                            '<div class="tile-label">'+element.Res_name+'</div></div>'
                }else{
                    h33 +='<div data-id="'+element.Res_id+'" data-url="'+element.Res_url+'" onclick="Hutils.goEntrySubSystem(this)" class="'+element.Res_class+' fg-white" data-role="tile" style="background-color: '+element.Res_bg_color+'">' +
                            '<div class="tile-content iconic"><span class="icon"><img src="'+element.Res_img+'"></span></div>' +
                            '<div class="tile-label">'+element.Res_name+'</div></div>'
                }
            })

            if (h11 != ""){
                $("#h-system-service").html(h1+h11+'</div></div>')
            }else{
                $("#h-system-service").remove()
            }
            if (h22 !=""){
                $("#h-mas-service").html(h2+h22+'</div></div>')
            }else{
                $("#h-mas-service").remove()
            }
            if (h33 !=""){
                $("#h-other-service").html(h3+h33+'</div></div>')
            }else{
                $("#h-other-service").remove()
            }

            $(function() {
                //取消水平滑动的插件
                //$.StartScreen();

                var tiles = $(".tile, .tile-small, .tile-sqaure, .tile-wide, .tile-large, .tile-big, .tile-super");

                $.each(tiles, function() {
                    var tile = $(this);
                    setTimeout(function() {
                        tile.css({
                            opacity: 1,
                            "-webkit-transform": "scale(1)",
                            "transform": "scale(1)",
                            "-webkit-transition": ".3s",
                            "transition": ".3s"
                        });
                    }, Math.floor(Math.random() * 500));
                });

                $(".tile-group").animate({
                    left: 0
                });
            });
        }

        $.HAjaxRequest({
            url:'/v1/auth/main/menu',
            data:{TypeId:1,Id:'0400000000'},
            success:succ,
        })
    })

    /*
    * 调整页面宽度和高度
    * */
    $(document).ready(function(){
        var hwindow = document.documentElement.clientHeight;
        var wwindow = document.documentElement.clientWidth;
        $("#wrapper").height(hwindow);
        $("#wrapper").width(wwindow);
        $("#page-wrapper").height(hwindow);
        $("#page-wrapper").width(wwindow);
        $("#main-menu-bar").height(hwindow);
        $("#h-main-content").height(hwindow);
        $(".H-content-tab").width(wwindow);
        $(".navbar-static-side").width(wwindow);
    });

    $(function() {
        //取消水平滑动的插件
        //$.StartScreen();

        var tiles = $(".tile, .tile-small, .tile-sqaure, .tile-wide, .tile-large, .tile-big, .tile-super");

        $.each(tiles, function() {
            var tile = $(this);
            setTimeout(function() {
                tile.css({
                    opacity: 1,
                    "-webkit-transform": "scale(1)",
                    "transform": "scale(1)",
                    "-webkit-transition": ".3s",
                    "transition": ".3s"
                });
            }, Math.floor(Math.random() * 500));
        });

        $(".tile-group").animate({
            left: 0
        });
    });

    window.onresize = function(){
        var hw = document.documentElement.clientWidth;
        var hh = document.documentElement.clientHeight;
        $("#wrapper").height(hh);
        $("#wrapper").width(hw);
        $("#page-wrapper").height(hh)
        $("#page-wrapper").width(hw)
        $("#H-main-content").height(hh);
    }
</script>