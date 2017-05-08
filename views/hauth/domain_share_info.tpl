<div class="row subsystem-header">
    <div class="pull-left">
        <span style="font-size: 16px;">域信息共享配置管理</span>
    </div>
</div>

<div class="row subsystem-toolbar">
    <div class="col-sm-12 col-md-5 col-lg-3" >
        <div class="pull-left">
            <span>域信息</span>
        </div>
    </div>
    <div class="col-sm-12 col-md-7 col-lg-9" style="padding-left: 0px;">
        <div class="pull-left">
            共享域信息与权限
        </div>
        <div class="pull-right" style="height: 44px; line-height: 44px; width: 260px;">
            <span style="text-align:right;width:80px;height: 30px; line-height: 30px; margin-top: 7px;display: inline" class="pull-left">&nbsp;&nbsp;所属域：</span>
            <select id="h-domain-share-domain-list" class="form-control pull-right" style="width: 180px;height: 24px; line-height: 24px; margin-top: 10px;padding: 0px;">
            </select>
        </div>
    </div>
</div>

<div class="row">
    <div class="col-sm-12 col-md-5 col-md-3">
        <div id="h-domain-info-shareid" style="border: #598f56 solid 1px;">
            <div class="col-sm-12 col-md-12 col-lg-12">
                <table class="table table-bordered table-condensed" style="margin-top: 30px;">
                    <tr style="background-color: #009966;color: white;"><th style="text-align: center">字段</th><th style="text-align: center">值</th></tr>
                    <tr style="height: 36px; line-height: 36px;"><td style="text-align: right;padding-right: 15px;vertical-align: middle;">域编码</td>
                        <td id="h-domain-share-did" style="vertical-align: middle;">{{.Project_id}}</td></tr>
                    <tr style="height: 36px; line-height: 36px;"><td  style="text-align: right;padding-right: 15px; vertical-align: middle;">域描述</td>
                        <td id="h-domain-share-did-name" style="vertical-align: middle;">{{.Project_name}}</td></tr>
                    <tr style="height: 36px; line-height: 36px;"><td style="text-align: right;padding-right: 15px; vertical-align: middle;">域状态</td>
                        <td id="h-domain-share-did-status" style="vertical-align: middle;">{{.Project_status}}</td></tr>
                    <tr style="height: 36px; line-height: 36px;"><td style="text-align: right;padding-right: 15px; vertical-align: middle;">创建日期</td>
                        <td id="h-domain-share-did-create-date" style="vertical-align: middle;">{{.Maintance_date}}</td></tr>
                    <tr style="height: 36px; line-height: 36px;"><td style="text-align: right;padding-right: 15px; vertical-align: middle;">创建人</td>
                        <td id="h-domain-share-did-create-user" style="vertical-align: middle;">{{.User_id}}</td></tr>
                    <tr style="height: 36px; line-height: 36px;"><td style="text-align: right;padding-right: 15px; vertical-align: middle;">修改日期</td>
                        <td id="h-domain-share-did-modify-date" style="vertical-align: middle;">{{.Domain_maintance_date}}</td></tr>
                    <tr style="height: 36px; line-height: 36px;"><td style="text-align: right;padding-right: 15px; vertical-align: middle;">修改人</td>
                        <td id="h-domain-share-did-modify-user" style="vertical-align: middle;">{{.Domain_maintance_user}}</td></tr>
                </table>
            </div>
        </div>
    </div>
    <div class="col-sm-12 col-md-7 col-lg-9" style="padding-left: 0px;">
        <div id="h-domain-share-info-details" style="border: #598f56 solid 1px;">
            <div id="h-domain-share-toolbar-info" style="height: 24px; line-height: 24px;">
                <div class="pull-left">
                    <button onclick="DomainShareObj.add()" class="btn btn btn-success btn-sm">
                        <i class="icon-plus"> 新增</i>
                    </button>
                    <button onclick="DomainShareObj.edit()" class="btn btn btn-success btn-sm">
                        <i class="icon-edit"> 编辑</i>
                    </button>
                    <button  onclick="DomainShareObj.delete()" class="btn btn btn-danger btn-sm">
                        <i class="icon-trash"> 删除</i>
                    </button>
                </div>
            </div>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <table id="h-domain-share-info-table"
                       data-toggle="table"
                       data-side-pagination="client"
                       data-pagination="true"
                       data-striped="true"
                       data-url="/v1/auth/domain/share/get"
                       data-page-list="[20, 50, 100, 200]"
                       data-toolbar="#h-domain-share-toolbar-info"
                       data-search="true">
                    <thead>
                    <tr>
                        <th data-field="state" data-checkbox="true"></th>
                        <th data-field="target_domain_id">域编码</th>
                        <th data-field="domain_name">域描述</th>
                        <th data-field="create_date">创建日期</th>
                        <th data-field="create_user">创建人</th>
                        <th data-field="modify_date">修改日期</th>
                        <th data-field="modify_user">修改人</th>
                        <th data-field="auth_level" data-formatter="DomainShareObj.authformatter">共享模式</th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>
</div>
<script>

    var DomainShareObj = {
        add:function () {
           $.Hmodal({
               header:"共享域信息",
               body:$("#domain-share-input-form").html(),
               height:"320px",
               preprocess:function () {

                   //获取列表中被授权访问的域信息
                   var arr = new Array()
                   var domain_id = $("#h-domain-share-did").html()
                   $.getJSON("/v1/auth/domain/share/unauth",{domain_id:domain_id},function (data) {
                       $(data).each(function (index, element) {
                           if (element.domain_id != domain_id){
                               var ijs = {}
                               ijs.id=element.domain_id
                               ijs.text=element.domain_name
                               ijs.upId="###hzwy23###"
                               arr.push(ijs)
                           }
                       })
                       $("#h-domain-share-target-domain-id").Hselect({
                           height:"30px",
                           data:arr,
                       })
                   });

                   $("#h-domain-share-auth-level").Hselect({
                       height:"30px",
                   })
               },
               callback:function (hmode) {
                   var domain_id = $("#h-domain-share-did").html() ;
                   var target_domain_id = $("#h-domain-share-target-domain-id").val();
                   var auth_level = $("#h-domain-share-auth-level").val();
                   $.HAjaxRequest({
                       url:"/v1/auth/domain/share/post",
                       type:"post",
                       data:{domain_id:domain_id,target_domain_id:target_domain_id,auth_level:auth_level},
                       success:function () {
                           $(hmode).remove()
                           $.Notify({
                               title:"操作成功",
                               message:"共享域信息成功",
                               type:"success",
                           })
                           $('#h-domain-share-info-table').bootstrapTable('refresh')
                       },
                   })
               }
           })
        },
        edit:function () {
            var sel = $('#h-domain-share-info-table').bootstrapTable('getSelections');
            if (sel.length == 0){
                $.Notify({
                    title:"温馨提示",
                    message:"请在下表中选择一项进行编辑",
                    type:"info",
                })
                return
            }else if (sel.length == 1){
                $.Hmodal({
                    header:"修改域分享模式",
                    body:$("#domain-share-modify-form").html(),
                    height:"320px",
                    preprocess:function(){
                        //获取列表中被授权访问的域信息
                        $("#h-domain-share-target-modify-domain-id").val(sel[0].target_domain_id)
                        var $select =  $("#h-domain-share-modify-auth-level");
                        $select.Hselect({
                            height:"30px",
                        });
                        $select.val(sel[0].auth_level).trigger("change");
                    },
                    callback:function (hmode) {
                        var uuid = sel[0].uuid;
                        var auth_level = $("#h-domain-share-modify-auth-level").val();
                        $.HAjaxRequest({
                            url:"/v1/auth/domain/share/put",
                            type:"put",
                            data:{uuid:uuid,auth_level:auth_level,domain_id:$("#h-domain-share-did").html()},
                            success:function () {
                                $.Notify({
                                    title:"操作成功",
                                    message:"修改域共享模式成功",
                                    type:"success",
                                })
                                $(hmode).remove()
                                $('#h-domain-share-info-table').bootstrapTable('refresh');
                            },
                        })
                    },
                })
            }else{
                $.Notify({
                    title:"温馨提示",
                    message:"你选择了多行信息，不知道想要编辑哪一行",
                    type:"info"
                })
                return
            }
        },
        delete:function () {
            $.Hconfirm({
                body:"点击确定删除域共享",
                callback:function () {
                    var sel = $('#h-domain-share-info-table').bootstrapTable('getSelections');
                    $.HAjaxRequest({
                        url:"/v1/auth/domain/share/delete",
                        type:"Post",
                        data:{JSON:JSON.stringify(sel),domain_id:$("#h-domain-share-did").html()},
                        success:function () {
                            $.Notify({
                                title:"操作成功",
                                message:"删除信息成功",
                                type:"success",
                            })
                            $('#h-domain-share-info-table').bootstrapTable('refresh')
                        },
                    })
                }
            })
        },
        authformatter:function (val, row, index) {
            if (val==1){
                return "只读"
            } else if (val==2){
                return "读写"
            } else {
                return "禁止访问"
            }
        }
    };

    $(document).ready(function(){
        var hwindow = document.documentElement.clientHeight;
        $("#h-domain-info-shareid").height(hwindow-130)
        $("#h-domain-share-info-details").height(hwindow-130)

        /*
         * 初始化左边工具栏中，域选择框
         * */
        $.getJSON("/v1/auth/domain/owner",function(data){
            var arr = new Array()
            $(data).each(function(index,element){
                var ijs = {}
                ijs.id=element.domain_id
                ijs.text=element.domain_desc
                ijs.upId="####hzwy23###"
                arr.push(ijs)
            });
            $("#h-domain-share-domain-list").Hselect({
                data:arr,
                height:"24px",
                width:"180px",
                onclick:function(){
                    // 刷新table信息
                    var domain_id =  $("#h-domain-share-domain-list").val();
                    if (domain_id==null){
                        return
                    }
                    //刷新左侧公告栏信息
                    $.getJSON("/v1/auth/domain/row/details",{domain_id:domain_id},function(data){
                        $(data).each(function(index,element){
                            $("#h-domain-share-did").html(element.domain_id)
                            $("#h-domain-share-did-name").html(element.domain_desc)
                            $("#h-domain-share-did-status").html(element.domain_status)
                            $("#h-domain-share-did-create-date").html(element.maintance_date)
                            $("#h-domain-share-did-create-user").html(element.create_user_id)
                            $("#h-domain-share-did-modify-date").html(element.domain_modify_date)
                            $("#h-domain-share-did-modify-user").html(element.domain_modify_user)
                        })
                    });

                    //刷新右侧分享域信息
                    $.getJSON("/v1/auth/domain/share/get",{domain_id:domain_id},function(data){
                        $table.bootstrapTable('load',data)
                    })
                }
            });
            var did = $("#h-domain-share-did").html()
            $("#h-domain-share-domain-list").val(did).trigger("change")
        });
        /*
        * 初始化table中信息
        * */
        var $table = $('#h-domain-share-info-table');
        $table.bootstrapTable({
            height:hwindow-130,
            queryParams:function (params) {
                params.domain_id = $("#h-domain-share-did").html()
                return params
            }
        });
    });
</script>

<script id="domain-share-input-form" type="text/html">
    <div class="col-sm-12 col-md-12 col-lg-12">
        <form class="form-horizontal" id="h-domain-share-add-tpl">
            <div class="col-sm-12 col-md-12 col-lg-12">
                <div class="form-group-sm col-sm-12 col-md-12 col-lg-12" style="margin-top: 15px;">
                    <label class="h-label" style="width: 100%;">共享目标：</label>
                    <select id="h-domain-share-target-domain-id" name="target-domain-id"  class="form-control"
                            style="width: 100%;padding: 0px;">
                    </select>
                </div>
                <div class="form-group-sm col-sm-12 col-md-12 col-lg-12" style="margin-top: 15px;">
                    <label class="h-label" style="width: 100%;">共享模式：</label>
                    <select id="h-domain-share-auth-level" name="auth-level"  class="form-control"
                            style="width: 100%; padding:0px;">
                        <option value="1">只读</option>
                        <option value="2">读写</option>
                    </select>
                </div>
            </div>
        </form>
    </div>
</script>

<script id="domain-share-modify-form" type="text/html">
    <div class="col-sm-12 col-md-12 col-lg-12">
        <form class="form-horizontal" id="h-domain-share-modify-tpl">
            <div class="col-sm-12 col-md-12 col-lg-12">
                <div class="form-group-sm col-sm-12 col-md-12 col-lg-12" style="margin-top: 15px;">
                    <label class="h-label" style="width: 100%;">共享目标：</label>
                    <input readonly="readonly" id="h-domain-share-target-modify-domain-id" name="target-domain-id"  class="form-control"
                            style="width: 100%;height:30px;line-height:30px;" />
                </div>
                <div class="form-group-sm col-sm-12 col-md-12 col-lg-12" style="margin-top: 15px;">
                    <label class="h-label" style="width: 100%;">共享模式：</label>
                    <select id="h-domain-share-modify-auth-level" name="auth-level"  class="form-control"
                            style="width: 100%; padding:0px;">
                        <option value="1">只读</option>
                        <option value="2">读写</option>
                    </select>
                </div>
            </div>
        </form>
    </div>
</script>