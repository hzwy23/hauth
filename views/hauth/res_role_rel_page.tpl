<div class="row subsystem-header">
    <div class="pull-left">
        <span style="font-size: 14px;">角色资源信息</span>
    </div>
</div>
<div class="row subsystem-toolbar">
    <div style="height: 44px; line-height: 44px;display: inline;">
        <span id="h-role-resource-rel-role-id" style="display: none;">{{.Role_id}}</span>
        <span style="height: 30px; line-height: 30px; margin-top: 7px;display: inline"
              class="pull-left">&nbsp;角色编码 = {{.Code_number}}</span>
        <span style="height: 30px; line-height: 30px; margin-top: 7px;display: inline"
              class="pull-left">&nbsp;&nbsp;&nbsp;<i style="border: #0b4059 dotted 0.5px; height: 44px;"></i>&nbsp;&nbsp;&nbsp;角色名称 = {{.Role_name}}</span>
        <span style="height: 30px; line-height: 30px; margin-top: 7px;display: inline"
              class="pull-left">&nbsp;&nbsp;&nbsp;<i style="border: #0b4059 dotted 0.5px;height: 44px;"></i>&nbsp;&nbsp;&nbsp;域编码 = <span id="h-resource-role-rel-domain-id">{{.Domain_id}}</span></span>
        <span style="height: 30px; line-height: 30px; margin-top: 7px;display: inline"
              class="pull-left">&nbsp;&nbsp;&nbsp;<i style="border: #0b4059 dotted 0.5px;height: 44px;"></i>&nbsp;&nbsp;&nbsp;域描述 = {{.Domain_desc}}</span>
    </div>
    <div class="pull-right">
        <button onclick="RoleResObj.auth()" class="btn btn btn-info btn-sm"><i class="icon-plus-sign"></i>&nbsp;授权
        </button>
        <button onclick="RoleResObj.revoke()" class="btn btn btn-danger btn-sm"><i class="icon-remove-circle"></i>&nbsp;撤销
        </button>
    </div>
</div>
<div class="row" style="background-image: url('/static/images/hauth/pure_book.png');filter:'progid:DXImageTransform.Microsoft.AlphaImageLoader(sizingMethod='scale')';-moz-background-size:100% 100%;background-size:100% 100%;">
    <div id="h-domain-share-info-details">
        <div class="col-sm-6 col-md-6 col-lg-6" style="padding-left: 10%; padding-right: 2%;">
            <div class="col-ms-12 col-md-12 col-lg-12" style="margin-top: 3%">
                <div style="border-bottom: #598f56 solid 1px;height: 44px; line-height: 44px;">
                    <div class="pull-left">
                        <span><i class="icon-sitemap"> </i>已经被授权获取菜单资源:</span>
                    </div>
                    <div class="pull-right">
                        <span>
                            <i class=" icon-search" style="margin-top: 15px;"></i>&nbsp;
                        </span>
                    </div>
                </div>
            </div>
            <div id="h-role-res-owner-resource" class="col-sm-12" style="overflow: auto">
            </div>
        </div>
        <div class="col-sm-6 col-md-6 col-lg-6" style="padding-left: 2%;padding-right: 10%;">
            <div class="col-ms-12 col-md-12 col-lg-12" style="margin-top: 3%;">
                <div style="border-bottom: #8f1121 solid 1px;height: 44px; line-height: 44px;">
                    <div class="pull-left">
                        <span><i class="icon-sitemap"> </i>尚未被授权菜单资源:</span>
                    </div>
                    <div class="pull-right">
                        <span>
                            <i class=" icon-search" style="margin-top: 15px;"></i>&nbsp;
                        </span>
                    </div>
                </div>
            </div>
            <div id="h-role-res-other-resource" class="col-sm-12" style="overflow: auto;">
            </div>
        </div>
    </div>
</div>

<!--<div class="row">-->
    <!--<div class="col-sm-12 col-md-7 col-lg-9" style="padding-left: 0px;">-->
        <!--<div id="h-domain-share-info-details" style="border: #598f56 solid 1px;">-->
            <!--<div class="col-sm-6 col-md-6 col-lg-6">-->
                <!--<div id="h-role-getted-resource-info" style="border: #598f56 solid 1px;margin-top: 15px;">-->
                    <!--<div class="col-ms-12 col-md-12 col-lg-12">-->
                        <!--<div style="border-bottom: #598f56 solid 2px;height: 44px; line-height: 44px;">-->
                            <!--<div class="pull-left">-->
                                <!--<span><i class="icon-sitemap"> </i>已经获取资源信息</span>-->
                            <!--</div>-->
                            <!--<div class="pull-right">-->
                                <!--<button onclick="RoleResObj.revoke()" class="btn btn btn-danger btn-sm" style="margin-top: 8px;"><i class="icon-remove-circle"></i>&nbsp撤销-->
                                <!--</button>-->
                            <!--</div>-->
                        <!--</div>-->
                    <!--</div>-->
                    <!--<div id="h-role-res-owner-resource" class="col-sm-12 col-md-12 col-lg-12"-->
                            <!--style="overflow: auto;padding: 0px 15px;">-->
                    <!--</div>-->
                <!--</div>-->
            <!--</div>-->
            <!--<div class="col-sm-6 col-md-6 col-lg-6">-->
                <!--<div id="h-role-ungetted-resource-info" style="border: #598f56 solid 1px;margin-top: 15px;">-->
                    <!--<div class="col-ms-12 col-md-12 col-lg-12">-->
                        <!--<div style="border-bottom: #598f56 solid 2px;height: 44px; line-height: 44px;">-->
                            <!--<div class="pull-left">-->
                                <!--<span><i class="icon-sitemap"> </i>未获取资源信息</span>-->
                            <!--</div>-->
                            <!--<div class="pull-right">-->
                                <!--<button onclick="RoleResObj.auth()" class="btn btn btn-success btn-sm" style="margin-top: 8px;"><i class="icon-plus-sign"></i>&nbsp授权-->
                                <!--</button>-->
                            <!--</div>-->
                        <!--</div>-->
                    <!--</div>-->
                    <!--<div id="h-role-res-other-resource" class="col-sm-12 col-md-12 col-lg-12"-->
                         <!--style="overflow: auto;padding: 0px 15px;">-->
                    <!--</div>-->
                <!--</div>-->
            <!--</div>-->
        <!--</div>-->
    <!--</div>-->
<!--</div>-->
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