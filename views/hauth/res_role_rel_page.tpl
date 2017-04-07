<div class="row subsystem-header">
    <div class="pull-left">
        <span style="font-size: 16px;">角色资源信息</span>
    </div>
</div>

<div class="row subsystem-toolbar">
    <div class="col-sm-12 col-md-5 col-lg-3" >
        <div class="pull-left">
            <span>角色信息</span>
        </div>
    </div>
    <div class="col-sm-12 col-md-7 col-lg-9" style="padding-left: 0px;">
        <div class="pull-left">
            角色分配资源情况
        </div>
        <!--<div class="pull-right" style="height: 44px; line-height: 44px; width: 240px;">-->
            <!--<span style="width:60px;height: 30px; line-height: 30px; margin-top: 7px;display: inline" class="pull-left">角色：</span>-->
            <!--<select class="form-control pull-right" style="width: 180px;height: 24px; line-height: 24px; margin-top: 10px;padding: 0px 8px;">-->
                <!--<option>演示环境</option>-->
                <!--<option>生产环境</option>-->
            <!--</select>-->
        <!--</div>-->
    </div>
</div>

<div class="row">
    <div class="col-sm-12 col-md-5 col-md-3">
        <div id="h-domain-info-shareid" style="border: #598f56 solid 1px;">
            <div class="col-sm-12 col-md-12 col-lg-12">
                <table class="table table-bordered table-condensed" style="margin-top: 30px;">
                    <tr style="background-color: #009966;color: white;"><th style="text-align: center">字段</th><th style="text-align: center">值</th></tr>
                    <tr><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;height: 36px; line-height: 36px;vertical-align: middle;">角色编码</td>
                        <td id="h-role-resource-rel-role-id" class="col-sm-8 col-md-8 col-lg-8" style="height: 36px; line-height: 36px;vertical-align: middle;">{{.Role_id}}</td></tr>
                    <tr><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;height: 36px; line-height: 36px;vertical-align: middle;">角色描述</td>
                        <td class="col-sm-8 col-md-8 col-lg-8" style="height: 36px; line-height: 36px;vertical-align: middle;">{{.Role_name}}</td></tr>
                    <tr><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;height: 36px; line-height: 36px;vertical-align: middle;">状态</td>
                        <td class="col-sm-8 col-md-8 col-lg-8" style="height: 36px; line-height: 36px;vertical-align: middle;">{{.Role_status_desc}}</td></tr>
                    <tr><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;height: 36px; line-height: 36px;vertical-align: middle;">所属域</td>
                        <td class="col-sm-8 col-md-8 col-lg-8" style="height: 36px; line-height: 36px;vertical-align: middle;">{{.Domain_desc}}</td></tr>
                    <tr><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;height: 36px; line-height: 36px;vertical-align: middle;">创建日期</td>
                        <td class="col-sm-8 col-md-8 col-lg-8" style="height: 36px; line-height: 36px;vertical-align: middle;">{{.Role_create_date}}</td></tr>
                    <tr><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;height: 36px; line-height: 36px;vertical-align: middle;">创建人</td>
                        <td class="col-sm-8 col-md-8 col-lg-8" style="height: 36px; line-height: 36px;vertical-align: middle;">{{.Role_owner}}</td></tr>
                    <tr><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;height: 36px; line-height: 36px;vertical-align: middle;">修改日期</td>
                        <td class="col-sm-8 col-md-8 col-lg-8" style="height: 36px; line-height: 36px;vertical-align: middle;">{{.Role_maintance_date}}</td></tr>
                    <tr><td class="col-sm-4 col-md-4 col-lg-4" style="text-align: right;padding-right: 15px;height: 36px; line-height: 36px;vertical-align: middle;">修改人</td>
                        <td class="col-sm-8 col-md-8 col-lg-8" style="height: 36px; line-height: 36px;vertical-align: middle;">{{.Role_maintance_user}}</td></tr>
                </table>
            </div>
        </div>
    </div>
    <div class="col-sm-12 col-md-7 col-lg-9" style="padding-left: 0px;">
        <div id="h-domain-share-info-details" style="border: #598f56 solid 1px;">
            <div class="col-sm-6 col-md-6 col-lg-6">
                <div id="h-role-getted-resource-info" style="border: #598f56 solid 1px;margin-top: 15px;">
                    <div class="col-ms-12 col-md-12 col-lg-12">
                        <div style="border-bottom: #598f56 solid 2px;height: 44px; line-height: 44px;">
                            <div class="pull-left">
                                <span><i class="icon-sitemap"> </i>已经获取资源信息</span>
                            </div>
                            <div class="pull-right">
                                <button onclick="RoleResObj.revoke()" class="btn btn btn-danger btn-sm" style="margin-top: 8px;"><i class="icon-remove-circle"></i>&nbsp撤销
                                </button>
                            </div>
                        </div>
                    </div>
                    <div id="h-role-res-owner-resource" class="col-sm-12 col-md-12 col-lg-12"
                            style="overflow: auto;padding: 0px 15px;">
                    </div>
                </div>
            </div>
            <div class="col-sm-6 col-md-6 col-lg-6">
                <div id="h-role-ungetted-resource-info" style="border: #598f56 solid 1px;margin-top: 15px;">
                    <div class="col-ms-12 col-md-12 col-lg-12">
                        <div style="border-bottom: #598f56 solid 2px;height: 44px; line-height: 44px;">
                            <div class="pull-left">
                                <span><i class="icon-sitemap"> </i>未获取资源信息</span>
                            </div>
                            <div class="pull-right">
                                <button onclick="RoleResObj.auth()" class="btn btn btn-success btn-sm" style="margin-top: 8px;"><i class="icon-plus-sign"></i>&nbsp授权
                                </button>
                            </div>
                        </div>
                    </div>
                    <div id="h-role-res-other-resource" class="col-sm-12 col-md-12 col-lg-12"
                         style="overflow: auto;padding: 0px 15px;">
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<script>
    var RoleResObj = {
        resource_self:function () {
            var role_id = $("#h-role-resource-rel-role-id").html()
            $.HAjaxRequest({
                url:"/v1/auth/role/resource/get",
                type:"get",
                data:{type_id:0,role_id:role_id},
                success:function (data) {
                    var arr = new Array()
                    $(data).each(function (index, element) {
                        var ijs = {}
                        ijs.id=element.res_id
                        ijs.text = element.res_name
                        ijs.upId = element.res_up_id
                        arr.push(ijs)
                    });
                    $("#h-role-res-owner-resource").Htree({
                        data:arr,
                    })
                },
            })
        },
        resource_other:function () {
            var role_id = $("#h-role-resource-rel-role-id").html()
            $.HAjaxRequest({
                url:"/v1/auth/role/resource/get",
                type:"get",
                data:{type_id:1,role_id:role_id},
                success:function (data) {
                    var arr = new Array()
                    $(data).each(function (index, element) {
                        var ijs = {}
                        ijs.id=element.res_id
                        ijs.text = element.res_name
                        ijs.upId = element.res_up_id
                        arr.push(ijs)
                    });
                    $("#h-role-res-other-resource").Htree({
                        data:arr,
                    })
                },
            })
        },
        revoke:function () {
            var res_id = $("#h-role-res-owner-resource").attr("data-selected");
            var role_id = $("#h-role-resource-rel-role-id").html();

            if (res_id ==undefined){
                $.Notify({
                    title:"温馨提示：",
                    message:"请在下列树形结构中选择需要撤销的资源",
                    type:"info",
                })
                return
            }
            $.Hconfirm({
                body:"点击确定删除角色拥有的资源",
                callback:function () {
                    $.HAjaxRequest({
                        url:"/v1/auth/role/resource/rights",
                        type:"post",
                        dataType:"json",
                        data:{role_id:role_id,res_id:res_id,type_id:"0"},
                        success:function () {
                            $.Notify({
                                title:"操作成功",
                                message:"撤销资源权限成功",
                                type:"success",
                            })
                            RoleResObj.resource_self()
                            RoleResObj.resource_other()
                        },
                    })
                }
            })

        },
        auth:function () {
            var res_id = $("#h-role-res-other-resource").attr("data-selected");
            var role_id = $("#h-role-resource-rel-role-id").html();

            if (res_id ==undefined){
                $.Notify({
                    title:"温馨提示：",
                    message:"请在下列树形结构中选择需要撤销的资源",
                    type:"info",
                })
                return
            }
            $.Hconfirm({
                body:"点击确定给角色授予菜单资源",
                callback:function () {
                    $.HAjaxRequest({
                        url:"/v1/auth/role/resource/rights",
                        type:"post",
                        data:{role_id:role_id,res_id:res_id,type_id:"1"},
                        success:function () {
                            $.Notify({
                                title:"操作成功",
                                message:"授权资源成功",
                                type:"success",
                            });
                            RoleResObj.resource_self()
                            RoleResObj.resource_other()
                        },
                    });
                }
            })
        },
    }

    $(document).ready(function(){
        var hwindow = document.documentElement.clientHeight;
        $("#h-domain-info-shareid").height(hwindow-130)
        $("#h-domain-share-info-details").height(hwindow-130)
        $("#h-role-getted-resource-info").height(hwindow-160)
        $("#h-role-ungetted-resource-info").height(hwindow-160)
        $("#h-role-res-owner-resource").height(hwindow-210)
        $("#h-role-res-other-resource").height(hwindow-210)
        RoleResObj.resource_self()
        RoleResObj.resource_other()
    });

</script>