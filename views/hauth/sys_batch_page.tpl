<div class="row subsystem-header">
    <div class="pull-left">
        <span style="font-size: 16px;">授权管理</span>
    </div>
</div>
<div class="row subsystem-toolbar">
    <div class="col-sm-12 col-md-12 col-lg-12">
        <div class="pull-left">
            <button style="margin-top: 8px;" onclick="AuthObj.auth()" class="btn btn-success btn-sm">
                <i class="icon-search"> 授权</i>
            </button>
            <!--<button onclick="window.open('/v1/auth/handle/logs/download.xlsx')" class="btn btn-danger btn-sm" title="下载操作记录">-->
                <!--<span class="icon-wrench"> 撤销</span>-->
            <!--</button>-->
        </div>
        <div class="pull-right" style="height: 44px; line-height: 44px; width: 260px;">
            <span style="text-align:right;width:80px;height: 30px; line-height: 30px; margin-top: 7px;display: inline" class="pull-left">&nbsp;&nbsp;所属域：</span>
            <select id="h-auth-domain-list" class="form-control pull-right" style="width: 180px;height: 24px; line-height: 24px; margin-top: 10px;padding: 0px;">
            </select>
        </div>
    </div>
</div>

<div class="row">
    <div class="col-sm-12 col-md-6 col-lg-7">
        <div id="h-grant-user-table-show" style="border: #598f56 solid 1px;">
            <div id="h-grant-toolbar-list" style="height: 24px; line-height: 24px;">
                <span style="font-size: 10px;font-weight: 600;display: inline">机构:</span>
                <select id="h-auth-org-list" class="form-control" style="height: 24px; line-height: 24px;padding: 0px; display: inline">
                </select>
                <!--<button onclick="AuthObj.search()" class="btn btn-success btn-xs" style="margin-left: 8px;"><i class="icon-search"> </i>查询</button>-->
            </div>
            <div id="h-grant-table-info" class="col-sm-12 col-md-12 col-lg-12">
                <table id="h-grant-info-table-details"
                       data-toggle="table"
                       data-striped="true"
                       data-url="/v1/auth/user/get"
                       data-toolbar="#h-grant-toolbar-list"
                       data-side-pagination="client"
                       data-pagination="true"
                       data-page-list="[20, 50, 100, 200]"
                       data-click-to-select="true"
                       data-search="true">
                    <thead>
                    <tr>
                        <th data-field="state" data-radio="true"></th>
                        <th data-field="user_id" data-sortable="true">账户</th>
                        <th data-field="user_name">用户名称</th>
                        <th data-field="org_unit_desc" data-sortable="true">机构</th>
                        <th data-field="modify_user" data-sortable="true">修改人</th>
                        <th data-field="modify_date" data-sortable="true">修改时间</th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>
    <div class="col-sm-12 col-md-6 col-lg-5" style="padding-left: 0px;">
        <div id="h-grant-user-show-box" style="border: #598f56 solid 1px;">
            <div class="col-sm-12 col-md-12 col-lg-12">
                <div class="row">
                    <div class="col-ms-12 col-md-12 col-lg-12">
                        <div style="height: 44px; line-height: 44px;">
                            <div class="pull-left">
                                <span style="font-weight: 600;">已被授予角色</span>
                            </div>
                        </div>
                    </div>
                </div>
                <table id="h-grant-user-role-table-details"
                       data-toggle="table"
                       data-striped="true"
                       data-url="/v1/auth/user/roles/get"
                       data-side-pagination="client"
                       data-pagination="false"
                       data-click-to-select="true"
                       data-search="false">
                    <thead>
                    <tr>
                        <th data-field="code_number" data-align="center" data-valign="middle">角色编码</th>
                        <th data-field="role_name" data-align="center" data-valign="middle">角色名称</th>
                        <th data-formatter="AuthObj.formatter" data-align="center" data-valign="middle">操作</th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>
</div>


<script type="text/javascript">


    var AuthObj = {
        search:function(){
            $.notifyClose();
            var org_id = $("#h-auth-org-list").val();
            var did = $("#h-auth-domain-list").val();
            $.HAjaxRequest({
                url:"/v1/auth/user/search",
                type:"get",
                data:{org_id:org_id,domain_id:did},
                success:function (data) {
                    $.Notify({
                        title:"查询成功",
                        message:"查询机构下用户信息成功",
                        type:"success",
                    })
                    $("#h-grant-info-table-details").bootstrapTable('load',data)
                    $("#h-grant-user-role-table-details").bootstrapTable('load',[])
                },
            })
        },
        auth:function () {
            var row = $("#h-grant-info-table-details").bootstrapTable('getSelections');
            if (row.length==0){
                $.Notify({
                    title:"温馨提示",
                    message:"请选择一个用户进行授权",
                    type:"info",
                })
                return
            }else{
                var user_id = row[0].user_id
                $.getJSON("/v1/auth/user/roles/other",{user_id:user_id},function (data) {
                    $.Hmodal({
                        header:"授权管理",
                        body:$("#h-other-user-role-table-details-html").html(),
                        height:"360px",
                        submitDesc:"授权",
                        cancelDesc:"关闭",
                        preprocess:function () {
                            var $table =  $("#h-other-user-role-table-details");
                            $table.bootstrapTable({
                                paginationLoop:false,
                                height:222,
                                striped:true,
                            });
                            $table.bootstrapTable('load',data)
                        },
                        callback:function (hmode) {
                            var $table =  $("#h-other-user-role-table-details");
                            var sect = $table.bootstrapTable('getSelections');
                            var arr = new Array();
                            $(sect).each(function (index,element) {
                                element.user_id = user_id
                                arr.push(element)
                            });
                            $.HAjaxRequest({
                                url:"/v1/auth/user/roles/auth",
                                type:"post",
                                data:{JSON:JSON.stringify(arr)},
                                success:function () {
                                    $.Notify({
                                        title:"操作成功",
                                        message:"授权用户角色信息成功",
                                        type:"success",
                                    })
                                    $(hmode).remove();
                                    $.getJSON("/v1/auth/user/roles/get",{user_id:user_id},function (dt) {
                                        $("#h-grant-user-role-table-details").bootstrapTable('load',dt)
                                    })
                                },
                            })
                        }
                    })
                });
            }
        },
        formatter:function (value,row,index) {
            var user_id = row.user_id;
            var role_id = row.role_id;
            return '<button onclick="AuthObj.revoke(\''+user_id+'\',\''+role_id+'\')" class="btn btn-danger btn-xs">删除</button>'
        },
        revoke:function (user_id,role_id) {
            $.Hconfirm({
                body:"点击确定移除权限",
                callback:function () {
                    $.HAjaxRequest({
                        url:"/v1/auth/user/roles/revoke",
                        type:"post",
                        data:{user_id:user_id,role_id:role_id},
                        success:function () {
                            $.Notify({
                                title:"温馨提示：",
                                message:"撤销用户角色成功",
                                type:"success",
                            });
                            $.getJSON("/v1/auth/user/roles/get",{user_id:user_id},function (dt) {
                                $("#h-grant-user-role-table-details").bootstrapTable('load',dt)
                            })
                        },
                    })
                }
            })
        },
    };

    $(document).ready(function(){
        var hwindow = document.documentElement.clientHeight;
        $("#h-grant-user-table-show").height(hwindow-130);
        $("#h-grant-user-show-box").height(hwindow-130);

        //初始化域信息
        $.getJSON("/v1/auth/domain/self/owner",function(data){
            var arr = new Array()
            $(data.owner_list).each(function(index,element){
                var ijs = {}
                ijs.id=element.domain_id
                ijs.text=element.domain_desc
                ijs.upId="####hzwy23###"
                arr.push(ijs)
            });

            $("#h-grant-info-table-details").bootstrapTable({
                height:hwindow-130,
                onClickRow:function (row, $element) {
                    var user_id = row.user_id
                    $.getJSON("/v1/auth/user/roles/get",{user_id:user_id},function (dt) {
                        $("#h-grant-user-role-table-details").bootstrapTable('load',dt)
                    })
                },
                queryParams:function (params) {
                    params.domain_id = $("#h-auth-domain-list").val();
                    return params
                },
            });

            $("#h-auth-domain-list").Hselect({
                data:arr,
                height:"24px",
                width:"180px",
                value:data.domain_id,
                onclick:function () {
                    var did = $("#h-auth-domain-list").val();
                    $.getJSON("/v1/auth/resource/org/get",{domain_id:did},function(data){
                        var arr = new Array()
                        $(data).each(function(index,element){
                            var ijs = {}
                            ijs.id=element.org_id;
                            ijs.text=element.org_desc;
                            ijs.upId=element.up_org_id;
                            arr.push(ijs)
                        });

                        $("#h-auth-org-list").Hselect({
                            data:arr,
                            height:"24px",
                            width:"180px",
                            onclick:function () {
                                AuthObj.search()
                            },
                        });
                        $("#h-grant-info-table-details").bootstrapTable('refresh');
                    });
                    $("#h-grant-user-role-table-details").bootstrapTable('load',[])
                }
            });

            $.getJSON("/v1/auth/resource/org/get",{domain_id:data.domain_id},function(data){
                var arr = new Array()
                $(data).each(function(index,element){
                    var ijs = {}
                    ijs.id=element.org_id;
                    ijs.text=element.org_desc;
                    ijs.upId=element.up_org_id;
                    arr.push(ijs)
                });

                $("#h-auth-org-list").Hselect({
                    data:arr,
                    height:"24px",
                    width:"180px",
                    onclick:function () {
                        AuthObj.search()
                    },
                });
            });
        });

        $("#h-grant-user-role-table-details").bootstrapTable({
            height:hwindow-180,
        })
    });
</script>

<script type="text/html" id="h-other-user-role-table-details-html">
    <table id="h-other-user-role-table-details"
           data-toggle="table"
           data-side-pagination="client"
           data-pagination="false"
           data-page-list="[20, 50, 100, 200]"
           data-click-to-select="true"
           data-search="false">
        <thead>
        <tr>
            <th data-field="state" data-checkbox="true">角色编码</th>
            <th data-field="code_number" data-sortable="true">角色编码</th>
            <th data-field="role_name">角色名称</th>
        </tr>
        </thead>
    </table>
</script>
