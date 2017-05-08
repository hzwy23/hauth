<div class="row subsystem-header">
    <div class="pull-left">
        <span style="font-size: 16px;">角色管理</span>
    </div>
</div>
<div class="row subsystem-toolbar">
    <div class="col-sm-12 col-md-12 col-lg-12">
        <div class="pull-left">
            <button onclick="RoleObj.add()" class="btn btn btn-success btn-sm">
                <i class="icon-plus"> 新增</i>
            </button>
            <button onclick="RoleObj.edit()" class="btn btn btn-success btn-sm">
                <i class="icon-edit"> 编辑</i>
            </button>
            <button  onclick="RoleObj.delete()" class="btn btn btn-danger btn-sm">
                <i class="icon-trash"> 删除</i>
            </button>
        </div>
        <div class="pull-right" style="height: 44px; line-height: 44px; width: 260px;">
            <span style="text-align:right;width:80px;height: 30px; line-height: 30px; margin-top: 7px;display: inline" class="pull-left">&nbsp;&nbsp;所属域：</span>
            <select id="h-role-domain-list" class="form-control pull-right" style="width: 180px;height: 24px; line-height: 24px; margin-top: 10px;padding: 0px;">
            </select>
        </div>
    </div>
</div>

<div id="h-role-info" class="row" class="row" style="margin: 0px 0px; border: #598f56 solid 1px;">
    <div id="h-role-details-toolbar" class="pull-left" style="height: 24px; line-height: 24px;">
        <span>角色信息列表</span>
    </div>
    <div id="h-role-table-info" class="col-sm-12 col-md-12 col-lg-12">
        <table id="h-role-info-table-details"
               data-toggle="table"
               data-striped="true"
               data-toolbar="#h-role-details-toolbar"
               data-side-pagination="client"
               data-pagination="true"
               data-page-list="[20, 50, 100, 200]"
               data-search="true">
            <thead>
            <tr>
                <th data-field="state" data-checkbox="true"></th>
                <th data-field="code_number">角色编码</th>
                <th data-field="role_name">角色名称</th>
                <th data-align="center"
                    data-field="role_status_desc">状态</th>
                <th data-field="domain_desc">所属域</th>
                <th data-align="center"
                    data-field="create_user">创建人</th>
                <th data-align="center"
                    data-field="create_date">创建时间</th>
                <th data-align="center"
                    data-field="modify_user">修改人</th>
                <th data-align="center"
                    data-field="modify_date">修改时间</th>
                <th data-field="state-handle"
                    data-align="center"
                    data-formatter="RoleObj.formatter">资源操作</th>
            </tr>
            </thead>
        </table>
    </div>
</div>

<script>

    var RoleObj={
        getResourcePage:function(id,name){
            var name = name+"资源信息"
            Hutils.openTab({
                url:"/v1/auth/role/resource/details?role_id="+id,
                id:"resourcedetails999988899",
                title:name,
                error:function (m) {
                    $.Notify({
                        title:"温馨提示：",
                        message:"权限不足",
                        type:"danger",
                    })
                }
            })
        },
        formatter:function(value,rows,index){
            return '<span class="h-td-btn btn-primary btn-xs" onclick="RoleObj.getResourcePage(\''+rows.role_id+'\',\''+ rows.role_name+'\')">资源管理</span>'
        },
        add:function () {
            var domain_id = $("#h-role-domain-list").val();

            $.Hmodal({
                header:"新增角色",
                body:$("#role_input_form").html(),
                height:"300px",
                preprocess:function () {
                    $.getJSON("/v1/auth/domain/owner",function(data) {
                        var arr = new Array()
                        $(data).each(function (index, element) {
                            var ijs = {}
                            ijs.id = element.domain_id
                            ijs.text = element.domain_desc
                            ijs.upId = "####hzwy23###"
                            arr.push(ijs)
                        });
                        $("#h-role-domain-id").Hselect({
                            data: arr,
                            height: "30px",
                            value:domain_id,
                        });
                    });
                    $("#h-role-add-status").Hselect({height:"30px"})
                },
                callback:function (hmode) {
                    $.HAjaxRequest({
                        url:"/v1/auth/role/post",
                        type:"post",
                        data:$("#h-role-add-info").serialize(),
                        success:function () {
                            $.Notify({
                                title:"操作成功",
                                message:"新增角色信息成功",
                                type:"success",
                            });
                            $(hmode).remove()
                            $("#h-role-info-table-details").bootstrapTable('refresh')
                        },
                    })
                }
            })
        },
        edit:function () {

            var rows = $("#h-role-info-table-details").bootstrapTable('getSelections');
            if (rows.length==0){
                $.Notify({
                    title:"温馨提示",
                    message:"您没有选择需要编辑的角色信息",
                    type:"info",
                });
            }else if (rows.length==1){
                var code_number = rows[0].code_number;
                var role_id = rows[0].role_id;
                var role_name = rows[0].role_name;
                var status_id = rows[0].role_status_code;

                console.log(code_number,role_id,role_name,status_id)
                $.Hmodal({
                    header:"编辑角色信息",
                    body:$("#role_modify_form").html(),
                    height:"320px",
                    preprocess:function () {
                        $("#h-role-modify-role-code-number").val(code_number)
                        $("#h-role-modify-role-name").val(role_name)
                        $("#h-role-modify-role-status-cd").Hselect({height:"30px"})
                        $("#h-role-modify-role-status-cd").val(status_id).trigger("change")
                    },
                    callback:function (hmode) {
                        var new_name = $("#h-role-modify-role-name").val()
                        var new_status = $("#h-role-modify-role-status-cd").val()
                        $.HAjaxRequest({
                            url:"/v1/auth/role/update",
                            type:"put",
                            data:{Role_id:role_id,Role_name:new_name,Role_status:new_status},
                            success:function () {
                                $.Notify({
                                    title:"操作成功",
                                    message:"编辑角色信息成功",
                                    type:"success",
                                });
                                $(hmode).remove();
                                $("#h-role-info-table-details").bootstrapTable('refresh')
                            },
                        })
                    },
                })

            }else {
                $.Notify({
                    title:"温馨提示",
                    message:"您选择了多行角色信息，不知道确定要编辑哪一行",
                    type:"info",
                });
            }
        },
        delete:function () {
           var rows = $("#h-role-info-table-details").bootstrapTable('getSelections');
           if (rows.length==0){
               $.Notify({
                   title:"温馨提示",
                   message:"您没有选择需要删除的角色信息",
                   type:"info",
               })
               return
           }else{
               $.Hconfirm({
                   callback:function () {
                       $.HAjaxRequest({
                           url:"/v1/auth/role/delete",
                           type:"post",
                           data:{JSON:JSON.stringify(rows)},
                           success:function () {
                               $.Notify({
                                   title:"操作成功",
                                   message:"删除角色信息成功",
                                   type:"success",
                               })
                               $("#h-role-info-table-details").bootstrapTable('refresh')
                           },
                       })
                   },
                   body:"确认要删除选中的角色吗"
               })
           }
        },
    };

    $(document).ready(function(obj){

        var hwindow=document.documentElement.clientHeight;

        // 初始化右上角域选择框
        $("#h-role-info-table-details").bootstrapTable({
            height:hwindow-130,
            url:"/v1/auth/role/get",
            queryParams:function (params) {
                params.domain_id = $("#h-role-domain-list").val();
                return params
            }
        });

        $("#h-role-info").height(hwindow-130);

        $.getJSON("/v1/auth/domain/self/owner",function(data) {

            var arr = new Array()
            $(data.owner_list).each(function (index, element) {
                var ijs = {}
                ijs.id = element.domain_id
                ijs.text = element.domain_desc
                ijs.upId = "####hzwy23###"
                arr.push(ijs)
            });

            $("#h-role-domain-list").Hselect({
                data: arr,
                height: "24px",
                width: "180px",
                value:data.domain_id,
                onclick: function () {
                    $("#h-role-info-table-details").bootstrapTable('refresh');
                },
            });
        });
    });
</script>

<script type="text/html" id="role_input_form">
    <form id="h-role-add-info">
        <div class="form-group-sm col-sm-6 col-md-6 col-lg-6" style="margin-top: 2px;">
            <label class="h-label" style="width:100%;">角色编码：</label>
            <input placeholder="1..30位数字，字母组成" name="role_id" type="text" class="form-control" style="width: 100%;height: 30px;line-height: 30px;">
        </div>
        <div class="form-group-sm col-sm-6 col-md-6 col-lg-6" style="margin-top: 2px;">
            <label class="h-label" style="width: 100%;">角色名称：</label>
            <input placeholder="1..30位汉字，字母，数字组成" type="text" class="form-control" name="role_name" style="width: 100%;height: 30px;line-height: 30px;">
        </div>
        <div class="form-group-sm col-sm-6 col-md-6 col-lg-6" style="margin-top: 15px;">
            <label class="h-label" style="width: 100%;">状　态：</label>
            <select id="h-role-add-status" name="role_status"  class="form-control" style="width: 100%;height: 30px;line-height: 30px;">
                <option value="0">正常</option>
                <option value="1">失效</option>
            </select>
        </div>
        <div class="form-group-sm col-sm-6 col-md-6 col-lg-6" style="margin-top: 15px;">
            <label class="h-label" style="width: 100%;">所属域：</label>
            <select id="h-role-domain-id" name="domain_id" style="width: 100%;height: 30px;line-height: 30px;">
            </select>
        </div>
    </form>
</script>

<script type="text/html" id="role_modify_form">
    <form novalidate="novalidate" id="h-role-modify-info">
        <div class="col-sm-12 col-md-12 col-lg-12">
            <label class="h-label" style="width:100%;">角色编码：</label>
            <input id="h-role-modify-role-code-number" readonly="readonly" name="code_number" type="text" class="form-control" style="width: 100%;height: 30px;line-height: 30px;">
        </div>
        <div class="col-sm-12 col-md-12 col-lg-12" style="margin-top: 8px;">
            <label class="h-label" style="width: 100%;">角色名称：</label>
            <input id="h-role-modify-role-name" placeholder="" type="text" class="form-control" name="role_name" style="width: 100%;height: 30px;line-height: 30px;">

        </div>
        <div class="col-sm-12 col-md-12 col-lg-12" style="margin-top: 8px;">
            <label class="h-label" style="width: 100%;">状　态：</label>
            <select id="h-role-modify-role-status-cd" name="role_status"  class="form-control" style="width: 100%;height: 30px;line-height: 30px;">
                <option value="0">正常</option>
                <option value="1">失效</option>
            </select>
        </div>
    </form>
</script>