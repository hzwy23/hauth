<div class="row subsystem-header">
    <div class="pull-left">
        <span style="font-size: 16px;">系统资源管理</span>
    </div>
</div>

<div class="row subsystem-toolbar">
    <div class="col-sm-12 col-md-6 col-lg-4" >
        <div class="pull-left">
            <button onclick="ResObj.add()" class="btn btn-success btn-sm">
                <i class="icon-plus"> 新增</i>
            </button>
            <button onclick="ResObj.delete()" class="btn btn-danger btn-sm" title="删除机构信息">
                <span class="icon-trash"> 删除</span>
            </button>
        </div>
    </div>
    <div class="col-sm-12 col-md-6 col-lg-4" style="padding-left: 0px;">
        <div class="pull-left">
            <button style="margin-top: 9px;" onclick="ResObj.edit()" class="btn btn-success btn-sm" title="删除机构信息">
                <span class="icon-trash"> 编辑</span>
            </button>
        </div>
    </div>
    <div class="col-sm-12 col-md-6 col-lg-4" style="padding-left: 0px;">
        <div class="pull-left">
            <button style="margin-top: 9px;" onclick="ResObj.configTheme()" class="btn btn-success btn-sm" title="编辑机构信息">
                <span class="icon-edit"> 配置</span>
            </button>
        </div>
    </div>
</div>

<div class="row">
    <div class="col-sm-12 col-md-6 col-lg-4">
        <div id="h-resource-tree-info" style="border: #598f56 solid 1px;height: 300px;">
            <div class="col-ms-12 col-md-12 col-lg-12">
                <div style="border-bottom: #598f56 solid 2px;height: 44px; line-height: 44px;">
                    <div class="pull-left">
                        <span><i class="icon-sitemap"> </i>菜单资源信息</span>
                    </div>
                    <div class="pull-right">
                    <span>
                        <i class=" icon-search" style="margin-top: 15px;"></i>&nbsp;
                    </span>
                    </div>
                </div>
            </div>
            <div id="h-resource-list-info" class="col-sm-12 col-md-12 col-lg-12"
                 style="padding:15px 5px;overflow: auto">
            </div>
        </div>
    </div>
    <div class="col-sm-12 col-md-6 col-lg-4" style="padding-left: 0px;">
        <div id="h-resource-details-info" style="border: #598f56 solid 1px;height: 300px;">
            <div class="col-ms-12 col-md-12 col-lg-12">
                <div style="border-bottom: #598f56 solid 2px;height: 44px; line-height: 44px;">
                    <div class="pull-left">
                        <span>资源详细信息</span>
                    </div>
                </div>
            </div>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <table class="table table-bordered table-condensed" style="margin-top: 20px;">
                    <tr style="background-color: #009966;color: white;"><th style="text-align: center">字段</th><th style="text-align: center">值</th></tr>
                    <tr style="height: 36px; line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">资源编码</td>
                        <td id="h-resource-show-id" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                    <tr style="height: 36px; line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">资源描述</td>
                        <td id="h-resource-show-name" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                    <tr style="height: 36px; line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">上级资源编码</td>
                        <td id="h-resource-show-up-id" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                    <tr style="height: 36px; line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">资源属性</td>
                        <td id="h-resource-show-attr-desc" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                    <tr style="height: 36px; line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">资源类别</td>
                        <td id="h-resource-show-type-desc" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td>
                        <td id="h-resource-show-type-id" style="display: none;"></td></tr>
                </table>
            </div>
        </div>
    </div>
    <div class="col-sm-12 col-md-6 col-lg-4" style="padding-left: 0px;">
        <div id="h-resource-theme-info" style="border: #598f56 solid 1px;height: 300px;">
            <div class="col-ms-12 col-md-12 col-lg-12">
                <div style="border-bottom: #598f56 solid 2px;height: 44px; line-height: 44px;">
                    <div class="pull-left">
                        <span>主题风格信息</span>
                    </div>
                    <div class="pull-right" style="height: 44px; line-height: 44px; width: 260px;">
                        <span style="text-align:right;width:80px;height: 30px; line-height: 30px; margin-top: 7px;display: inline" class="pull-left">&nbsp;&nbsp;主题风格：</span>
                        <select onchange="ResObj.updateTheme()" id="h-resource-theme-style-code" class="form-control pull-right" style="width: 180px;height: 24px; line-height: 24px; margin-top: 10px;padding: 0px;">
                            <option value="1001">绿色主题</option>
                            <option value="1002">蓝色主题</option>
                            <option value="1003">粉丝主题</option>
                            <option value="1004">青色主题</option>
                        </select>
                    </div>
                </div>
            </div>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <table class="table table-bordered table-condensed" style="margin-top: 20px;">
                    <tr style="background-color: #009966;color: white;"><th style="text-align: center">字段</th><th style="text-align: center">值</th></tr>
                    <tr style="height: 36px;line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">主题编码</td>
                        <td id="h-resource-show-theme-id" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                    <tr style="height: 36px;line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">主题名称</td>
                        <td id="h-resource-show-theme-desc" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                    <tr style="height: 36px;line-height: 36px; display: none;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">打开方式</td>
                        <td id="h-resource-show-res-res-type" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;"></td></tr>
                    <tr style="height: 36px;line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">打开方式</td>
                        <td id="h-resource-show-res-res-type-desc" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                    <tr style="height: 36px;line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">资源色彩</td>
                        <td id="h-resource-show-res-bg-color" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                    <tr style="height: 36px;line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">资源样式</td>
                        <td id="h-resource-show-res-class" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                    <tr style="height: 36px;line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">图标</td>
                        <td id="h-resource-show-res-img" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                    <tr style="height: 36px;line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">路由</td>
                        <td id="h-resource-show-res-url" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                    <tr style="height: 36px;line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">分组号</td>
                        <td id="h-resource-show-res-group-id" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                    <tr style="height: 36px;line-height: 36px;"><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;vertical-align: middle;">组内排序号</td>
                        <td id="h-resource-show-res-sort-id" class="col-sm-8 col-md-8 col-lg-8" style="vertical-align: middle;padding-left: 15px;">-</td></tr>
                </table>
            </div>
        </div>
    </div>
</div>

<script type="text/javascript">

    var ResObj = {
        delete:function () {
            $.Hconfirm({
                body:"点击确定删除菜单资源信息",
                callback:function () {
                    $.HAjaxRequest({
                        url:"/v1/auth/resource/delete",
                        type:"post",
                        data:{res_id:$("#h-resource-show-id").html()},
                        success:function () {
                            $.Notify({
                                title:"操作成功",
                                message:"删除菜单资源成功",
                                type:"success",
                            });
                            ResObj.tree()
                        },
                    })
                }
            })
        },
        add:function () {
            $.Hmodal({
                header:"新增资源",
                body:$("#res_input_form").html(),
                height:"460px",
                width:"744px",
                callback:function (hmode) {
                    $.HAjaxRequest({
                        url:"/v1/auth/resource/post",
                        data:$("#h-res-add-info").serialize(),
                        type:"post",
                        success:function () {
                            $.Notify({
                                title:"操作成功",
                                message:"新增菜单资源成功",
                                type:"success",
                            })
                            $(hmode).remove()
                            ResObj.tree()
                        },
                    })
                }
            })
        },
        selectType:function (obj) {
            var type_id = $(obj).val()
            switch (type_id){
                case "0":
                    // 主菜单系统
                    $("#h-res-add-up-res-id").parent().parent().hide();
                    $("#h-res-add-sort-id").parent().parent().show();
                    $("#h-res-add-group-id").parent().parent().show();
                    $("#h-res-add-res-bg-color").parent().parent().show();
                    $("#h-res-add-res-img").parent().parent().show();
                    $("#h-res-add-res-class").parent().parent().show();
                    $("#h-res-add-res-url").parent().parent().show();
                    $("#h-res-modify-res-open-type").parent().parent().show();
                    break;
                case "1":
                    // 子页面系统
                    $("#h-res-add-up-res-id").parent().parent().show();
                    $.getJSON("/v1/auth/resource/get",function (data) {
                        var arr = new Array()
                        $(data).each(function (index, element) {
                            if (element.res_type == "0" || element.res_type == "4"){
                                var ijs = {}
                                ijs.id=element.res_id
                                ijs.text = element.res_name
                                ijs.upId=element.res_up_id;
                                arr.push(ijs)
                            }
                        });

                        $("#h-res-add-up-res-id").Hselect({
                            data: arr,
                            height:"30px",
                        })
                    });
                    $("#h-res-add-sort-id").parent().parent().show();
                    $("#h-res-add-group-id").parent().parent().show();
                    $("#h-res-add-res-bg-color").parent().parent().show();
                    $("#h-res-add-res-img").parent().parent().show();
                    $("#h-res-add-res-class").parent().parent().show();
                    $("#h-res-add-res-url").parent().parent().show();
                    $("#h-res-modify-res-open-type").parent().parent().show();
                    break;
                case "2":
                    // 功能按钮
                    $("#h-res-add-up-res-id").parent().parent().show();
                    $.getJSON("/v1/auth/resource/get",function (data) {
                        var arr = new Array()
                        $(data).each(function (index, element) {
                            if (element.res_type != "3"){
                                var ijs = {}
                                ijs.id=element.res_id
                                ijs.text = element.res_name
                                ijs.upId=element.res_up_id;
                                arr.push(ijs)
                            }
                        });

                        $("#h-res-add-up-res-id").Hselect({
                            data: arr,
                            height:"30px",
                        })
                    });
                    $("#h-res-add-sort-id").parent().parent().hide();
                    $("#h-res-add-group-id").parent().parent().hide();
                    $("#h-res-add-res-bg-color").parent().parent().hide();
                    $("#h-res-add-res-img").parent().parent().hide();
                    $("#h-res-add-res-class").parent().parent().hide();
                    $("#h-res-add-res-url").parent().parent().show();
                    $("#h-res-modify-res-open-type").parent().parent().hide();
                    break;
                case "4":
                    // 虚拟节点
                    $("#h-res-add-up-res-id").parent().parent().show();
                    $.getJSON("/v1/auth/resource/get",function (data) {
                        var arr = new Array()
                        $(data).each(function (index, element) {
                            if (element.res_type == "0" || element.res_type == "4"){
                                var ijs = {}
                                ijs.id=element.res_id
                                ijs.text = element.res_name
                                ijs.upId=element.res_up_id;
                                arr.push(ijs)
                            }
                        });
                        $("#h-res-add-up-res-id").Hselect({
                            data: arr,
                            height:"30px",
                        })
                    });
                    $("#h-res-add-sort-id").parent().parent().hide();
                    $("#h-res-add-group-id").parent().parent().hide();
                    $("#h-res-add-res-bg-color").parent().parent().hide();
                    $("#h-res-add-res-img").parent().parent().hide();
                    $("#h-res-add-res-class").parent().parent().hide();
                    $("#h-res-add-res-url").parent().parent().hide();
                    $("#h-res-modify-res-open-type").parent().parent().hide();
                    break;
            }
        },
        updateTheme:function () {
            var id = $("#h-resource-show-id").html()
            var theme_id = $("#h-resource-theme-style-code").val()
            $.getJSON("/v1/auth/resource/queryTheme",{res_id:id,theme_id:theme_id},function (e) {
                if (e.length==0){
                    $("#h-resource-show-theme-id").html("-")
                    $("#h-resource-show-theme-desc").html("-")
                    $("#h-resource-show-res-bg-color").html("-")
                    $("#h-resource-show-res-class").html("-")
                    $("#h-resource-show-res-img").html("-")
                    $("#h-resource-show-res-url").html("-")
                    $("#h-resource-show-res-group-id").html("-")
                    $("#h-resource-show-res-sort-id").html("-")
                    $("#h-resource-show-res-res-type").html("-")
                    $("#h-resource-show-res-res-type-desc").html("-")
                } else {
                    $(e).each(function(index,element){
                        $("#h-resource-show-theme-id").html(element.theme_id)
                        $("#h-resource-show-theme-desc").html(element.theme_desc)
                        $("#h-resource-show-res-bg-color").html(element.res_bg_color)
                        $("#h-resource-show-res-class").html(element.res_class)
                        $("#h-resource-show-res-img").html(element.res_img)
                        $("#h-resource-show-res-url").html(element.res_url)
                        $("#h-resource-show-res-group-id").html(element.group_id)
                        $("#h-resource-show-res-sort-id").html(element.sort_id)

                        $("#h-resource-show-res-res-type").html(element.res_type)
                        if (element.res_type == "0") {
                            $("#h-resource-show-res-res-type-desc").html("内嵌页面")
                        } else if (element.res_type == "1") {
                            $("#h-resource-show-res-res-type-desc").html("新建选项卡")
                        } else {
                            $("#h-resource-show-res-res-type-desc").html("")
                        }
                    })
                }
            })
        },
        tree:function () {
            $.getJSON("/v1/auth/resource/get",function (data) {
                if (data.length==0){
                    $.Notify({
                        title:"温馨提示",
                        message:"查询结果为空",
                        type:"info",
                    });
                    $("#h-resource-list-info").Htree({
                        data:[],
                    })
                } else {
                    var arr = new Array()
                    $(data).each(function(index,element){
                        var ijs = {};
                        ijs.id = element.res_id
                        ijs.text = element.res_name
                        ijs.upId = element.res_up_id
                        arr.push(ijs)
                    });
                    $("#h-resource-list-info").Htree({
                        data:arr,
                        onChange:function(obj){
                            var id = $(obj).attr("data-id")
                            $.getJSON("/v1/auth/resource/query",{res_id:id},function (e) {
                                $(e).each(function (index, element) {
                                    $("#h-resource-show-id").html(element.res_id)
                                    $("#h-resource-show-name").html(element.res_name)
                                    $("#h-resource-show-up-id").html(element.res_up_id)
                                    $("#h-resource-show-attr-desc").html(element.res_attr_desc)
                                    $("#h-resource-show-type-desc").html(element.res_type_desc)
                                    $("#h-resource-show-type-id").html(element.res_type)
                                })
                            })
                            var theme_id = $("#h-resource-theme-style-code").val();
                            $.getJSON("/v1/auth/resource/queryTheme",{res_id:id,theme_id:theme_id},function (e) {
                                if (e.length==0){
                                    $("#h-resource-show-theme-id").html("-");
                                    $("#h-resource-show-theme-desc").html("-");
                                    $("#h-resource-show-res-bg-color").html("-");
                                    $("#h-resource-show-res-class").html("-");
                                    $("#h-resource-show-res-img").html("-");
                                    $("#h-resource-show-res-url").html("-");
                                    $("#h-resource-show-res-group-id").html("-");
                                    $("#h-resource-show-res-sort-id").html("-");
                                    $("#h-resource-show-res-res-type").html("-");
                                    $("#h-resource-show-res-res-type-desc").html("-");
                                } else {
                                    $(e).each(function(index,element){
                                        $("#h-resource-show-theme-id").html(element.theme_id)
                                        $("#h-resource-show-theme-desc").html(element.theme_desc)
                                        $("#h-resource-show-res-bg-color").html(element.res_bg_color)
                                        $("#h-resource-show-res-class").html(element.res_class)
                                        $("#h-resource-show-res-img").html(element.res_img)
                                        $("#h-resource-show-res-url").html(element.res_url)
                                        $("#h-resource-show-res-group-id").html(element.group_id)
                                        $("#h-resource-show-res-sort-id").html(element.sort_id)
                                        $("#h-resource-show-res-res-type").html(element.res_type)
                                        if (element.res_type == "0") {
                                            $("#h-resource-show-res-res-type-desc").html("内嵌页面")
                                        } else if (element.res_type == "1") {
                                            $("#h-resource-show-res-res-type-desc").html("新建选项卡")
                                        } else {
                                            $("#h-resource-show-res-res-type-desc").html("")
                                        }
                                    })
                                }
                            })
                        }
                    });
                }
            })
        },
        edit:function(){
            var res_id = $("#h-resource-show-id").html();
            if (res_id == "-" || res_id == "") {
                $.Notify({
                    message:"请在菜单资源中选择需要编辑的菜单",
                    type:"warning",
                });
                return
            }
            var res_name = $("#h-resource-show-name").html();
            if ( res_id == "" ){
                $.Notify({
                    title:"温馨提示：",
                    message:"请选择需要编辑的菜单资源",
                    type:"info",
                })
                return
            }
            $.Hmodal({
                header:"编辑资源信息",
                body:$("#res_input_form_modify").html(),
                height:"360px",
                callback:function (hmode) {
                    $.HAjaxRequest({
                        url:"/v1/auth/resource/update",
                        type:"put",
                        data:$("#h-res-modify-info").serialize(),
                        success:function () {
                            $.Notify({
                                title:"温馨提示：",
                                message:"修改菜单资源名称成功",
                                type:"success",
                            })
                            $(hmode).remove()
                            ResObj.tree()
                        }
                    })
                },
                preprocess:function () {
                    $("#h-res-modify-res-id").val(res_id);
                    $("#h-res-modify-res-name").val(res_name);
                },
            })
        },
        configTheme:function(){

            var res_type = $("#h-resource-show-type-id").html();
            if (res_type == "4") {
                $.Notify({
                    message:"虚拟节点不允许编辑",
                    type:"warning",
                });
                return
            } else if (res_type == "") {
                $.Notify({
                    message:"请在菜单资源树中,选择需要设置主题的菜单",
                    type:"warning",
                });
                return
            }

            $.Hmodal({
                header:"配置主题信息",
                body:$("#res_modify_theme_form").html(),
                height:"420px",
                callback:function (hmode) {

                    var res_id = $("#h-resource-show-id").html()
                    var theme_id = $("#h-res-modify-theme-id").val()
                    var res_url = $("#h-res-modify-res-url").val()
                    var res_class = $("#h-res-modify-res-class").val()
                    var res_img = $("#h-res-modify-res-img").val()
                    var res_by_color = $("#h-res-modify-res-bg-color").val()
                    var res_group_id = $("#h-res-modify-group-id").val()
                    var res_sort_id = $("#h-res-modify-sort-id").val()
                    var res_type = $("#h-res-modify-res-type").val()

                    $.HAjaxRequest({
                        url:"/v1/auth/resource/config/theme",
                        type:"Put",
                        data:{
                            res_id:res_id,
                            theme_id:theme_id,
                            res_url:res_url,
                            res_class:res_class,
                            res_img:res_img,
                            res_by_color:res_by_color,
                            res_group_id:res_group_id,
                            res_sort_id:res_sort_id,
                            res_openType:res_type,
                        },
                        success:function(){
                            $(hmode).remove()
                            $.Notify({
                                title:"温馨提示：",
                                message:"配置主题信息成功",
                                type:"success",
                            })
                        }
                    })
                },
                preprocess:function(){
                    var theme_id = $("#h-resource-show-theme-id").html()
                    var res_by_color = $("#h-resource-show-res-bg-color").html()
                    var res_class = $("#h-resource-show-res-class").html()
                    var res_img = $("#h-resource-show-res-img").html()
                    var res_url = $("#h-resource-show-res-url").html()
                    var res_group_id = $("#h-resource-show-res-group-id").html()
                    var res_sort_id = $("#h-resource-show-res-sort-id").html()
                    var open_type = $("#h-resource-show-res-res-type").html();

                    if (res_type == "2") {
                        $("#h-res-modify-res-class").parent().parent().hide()
                        $("#h-res-modify-res-img").parent().parent().hide()
                        $("#h-res-modify-res-bg-color").parent().parent().hide()
                        $("#h-res-modify-group-id").parent().parent().hide()
                        $("#h-res-modify-sort-id").parent().parent().hide()
                        $("#h-res-modify-res-type").parent().parent().hide()
                    } else {
                        $("#h-res-modify-res-class").Hselect({
                            height:"30px",
                            value:res_class,
                        });
                        $("#h-res-modify-res-type").Hselect({
                            height: "30px",
                            value: open_type,
                        });
                        $("#h-res-modify-res-img").val(res_img)
                        $("#h-res-modify-res-bg-color").val(res_by_color)
                        $("#h-res-modify-group-id").val(res_group_id)
                        $("#h-res-modify-sort-id").val(res_sort_id)
                    }

                    $("#h-res-modify-theme-id").Hselect({
                        height:"30px",
                        value:theme_id,
                    });
                    $("#h-res-modify-res-url").val(res_url)
                },
            })
        },
    }

    $(document).ready(function() {
        /*
        * 调整属性信息显示框大小
        * 高度填充全屏高度
        * */
        var hwindow = document.documentElement.clientHeight;
        $("#h-resource-tree-info").height(hwindow - 130);
        $("#h-resource-details-info").height(hwindow - 130);
        $("#h-resource-theme-info").height(hwindow - 130);
        $("#h-resource-list-info").height(hwindow - 204);
        ResObj.tree()
    });
</script>

<script type="text/html" id="res_input_form">
    <form class="row" id="h-res-add-info">
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 2px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">所属主题：</label>
            <div class="col-sm-12">
                <select id="h-res-add-theme-id" name="theme_id" class="form-control" style="height: 30px; line-height: 30px;">
                    <option value="1001">绿色主题</option>
                    <option value="1002">蓝色主题</option>
                    <option value="1003">粉丝主题</option>
                    <option value="1004">青色主题</option>
                </select>
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 2px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">菜单类别：</label>
            <div class="col-sm-12">
                <select id="h-res-add-type-id" onchange="ResObj.selectType(this)" name="res_type" class="form-control" style="height: 30px; line-height: 30px;">
                    <option value="0">首页系统</option>
                    <option value="1">子页系统</option>
                    <option value="2">功能按钮</option>
                    <option value="4">虚拟节点</option>
                </select>
            </div>
        </div>

        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">打开方式：</label>
            <div class="col-sm-12">
                <select id="h-res-modify-res-open-type" name="res_open_type" class="form-control" style="height: 30px; line-height: 30px;">
                    <option value="0">内嵌页面</option>
                    <option value="1">新建选项卡</option>
                </select>
            </div>
        </div>

        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">资源编码：</label>
            <div class="col-sm-12">
                <input id="h-res-add-res-id" placeholder="1-30位字母、数字组成" name="res_id" type="text" class="form-control" style="height: 30px; line-height: 30px;">
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">资源名称：</label>
            <div class="col-sm-12">
                <input id="h-res-add-res-name" placeholder="1-30位汉字、字母组成" type="text" class="form-control" name="res_name" style="height: 30px; line-height: 30px;">
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="display: none;margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">上级资源编码：</label>
            <div class="col-sm-12">
                <select id="h-res-add-up-res-id" name="res_up_id" type="text" class="form-control" style="height: 30px; line-height: 30px;padding: 0px; display: block">
                </select>
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">路由信息：</label>
            <div class="col-sm-12">
                <input id="h-res-add-res-url" placeholder="如：/v1/auth/help" type="url" class="form-control" name="res_url" style="height: 30px; line-height: 30px;">
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">样式属性：</label>
            <div class="col-sm-12">
                <select id="h-res-add-res-class" name="res_class" class="form-control" style="height: 30px; line-height: 30px;">
                    <option value="tile">小方块图形</option>
                    <option value="tile tile-wide">长方形图形</option>
                    <option value="tile tile-large">大方块图形</option>
                </select>
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">图标路径：</label>
            <div class="col-sm-12">
                <input id="h-res-add-res-img" placeholder="如：/static/images/example.png" name="res_img" type="text" class="form-control" style="height: 30px; line-height: 30px;">
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">图标色彩：</label>
            <div class="col-sm-12">
                <input id="h-res-add-res-bg-color" placeholder="#339999" name="res_bg_color" type="text" class="form-control" style="height: 30px; line-height: 30px;">
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">分组编号：</label>
            <div class="col-sm-12">
                <input id="h-res-add-group-id" placeholder="菜单所属分组，请用数字表示" type="number" class="form-control" name="group_id" style="height: 30px; line-height: 30px;">
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">排序号：</label>
            <div class="col-sm-12">
                <input id="h-res-add-sort-id" placeholder="所在分组的排序号，请用数字表示" name="sort_id" type="number" class="form-control" style="height: 30px; line-height: 30px;">
            </div>
        </div>
    </form>
</script>


<script type="text/html" id="res_input_form_modify">
    <form class="row form-horizontal" id="h-res-modify-info">
        <div class="col-sm-12 col-md-12 col-lg-12">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">资源编码：</label>
            <div class="col-sm-12">
                <input readonly="readonly" id="h-res-modify-res-id" placeholder="1-30位字母、数字组成" name="res_id" type="text" class="form-control" style="height: 30px; line-height: 30px;">
            </div>
        </div>
        <div class="col-sm-12 col-md-12 col-lg-12" style="margin-top: 15px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">资源名称：</label>
            <div class="col-sm-12">
                <input id="h-res-modify-res-name" placeholder="1-30位汉字、字母组成" type="text" class="form-control" name="res_name" style="height: 30px; line-height: 30px;">
            </div>
        </div>
    </form>
</script>


<script type="text/html" id="res_modify_theme_form">
    <form class="row" id="h-res-modify-theme-info">
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 2px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">所属主题：</label>
            <div class="col-sm-12">
                <select id="h-res-modify-theme-id" name="theme_id" class="form-control" style="height: 30px; line-height: 30px;">
                    <option value="1001">绿色主题</option>
                    <option value="1002">蓝色主题</option>
                    <option value="1003">粉丝主题</option>
                    <option value="1004">青色主题</option>
                </select>
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 2px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">样式属性：</label>
            <div class="col-sm-12">
                <select id="h-res-modify-res-class" name="res_class" class="form-control" style="height: 30px; line-height: 30px;">
                    <option value="tile">小方块图形</option>
                    <option value="tile tile-wide">长方形图形</option>
                    <option value="tile tile-large">大方块图形</option>
                </select>
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">打开方式：</label>
            <div class="col-sm-12">
                <select id="h-res-modify-res-type" name="res_type" class="form-control" style="height: 30px; line-height: 30px;">
                    <option value="0">内嵌页面</option>
                    <option value="1">新建选项卡</option>
                </select>
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">图标路径：</label>
            <div class="col-sm-12">
                <input id="h-res-modify-res-img" placeholder="如：/static/images/example.png" name="res_img" type="text" class="form-control" style="height: 30px; line-height: 30px;">
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">图标色彩：</label>
            <div class="col-sm-12">
                <input id="h-res-modify-res-bg-color" placeholder="#339999" name="res_bg_color" type="text" class="form-control" style="height: 30px; line-height: 30px;">
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">路由信息：</label>
            <div class="col-sm-12">
                <input id="h-res-modify-res-url" placeholder="如：/v1/auth/help" type="url" class="form-control" name="res_url" style="height: 30px; line-height: 30px;">
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">分组编号：</label>
            <div class="col-sm-12">
                <input id="h-res-modify-group-id" placeholder="菜单所属分组，请用数字表示" type="number" class="form-control" name="group_id" style="height: 30px; line-height: 30px;">
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="margin-top: 8px;">
            <label class="col-sm-12 control-label" style="font-size: 14px; font-weight: 500;text-align: left">排序号：</label>
            <div class="col-sm-12">
                <input id="h-res-modify-sort-id" placeholder="所在分组的排序号，请用数字表示" name="sort_id" type="number" class="form-control" style="height: 30px; line-height: 30px;">
            </div>
        </div>
    </form>
</script>
