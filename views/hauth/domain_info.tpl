<div class="row subsystem-header">
    <span style="font-size: 14px;">域信息管理</span>
</div>
<div class="row subsystem-toolbar">
    <div class="pull-right">
        <button onclick="DomainObj.domainInsertRow()" class="btn btn btn-info btn-sm">
            <i class="icon-plus"> 新增</i>
        </button>
        <button onclick="DomainObj.domainEditRow()" class="btn btn btn-info btn-sm">
            <i class="icon-edit"> 编辑</i>
        </button>
        <button  onclick="DomainObj.domainDeleteRow()" class="btn btn btn-danger btn-sm">
            <i class="icon-trash"> 删除</i>
        </button>
    </div>
</div>
<div class="row">
    <div id="h-domain-info" class="col-sm-12 col-md-12 col-lg-12">
        <table id="HdomainInfoTable"></table>
    </div>
</div>

<script type="text/javascript">

    var DomainObj = {
        getSharePage:function(id,name){
            var name = name+"共享配置"
            Hutils.openTab({
                url:"/v1/auth/domain/share/page?domain_id="+id,
                id:"sharedomain9999888",
                title:name,
                error:function (m) {
                    $.Notify({
                        title:"温馨提示：",
                        message:"您没有被授权访问这个域。",
                        type:"info",
                    })
                }
            })
        },
        getDomainInfo:function(){
            $("#HdomainInfoTable").bootstrapTable({
                url:'/v1/auth/domain/get',
                height:document.documentElement.clientHeight-118,
                uniqueId:'domain_id',
                striped: true,
                pagination: true,
                pageList:[80,160,400,800,3000,"All"],
                showRefresh:true,
                pageSize: 40,
                showExport:false,
                search:false,
                sidePagination: "client",
                showColumns: false,
                columns:[{
                    checkbox:true,
                }, {

                    field: 'domain_id',

                    title: '域编码',

                    align: 'left',

                    valign: 'middle',

                    sortable: false

                }, {

                    field: 'domain_desc',

                    title: '域名称',

                    align: 'left',

                    valign: 'middle',

                    sortable: false,

                }, {

                    field: 'domain_status',

                    title: '域状态',

                    align: 'center',

                    valign: 'middle',

                    sortable: false,

                },{

                    field: 'maintance_date',

                    title: '创建日期',

                    align: 'center',

                    valign: 'top',

                    sortable: false

                }, {

                    field: 'create_user_id',

                    title: '创建人',

                    align: 'center',

                    valign: 'middle',

                    sortable: false

                }, {

                    field: 'domain_modify_date',

                    title: '修改日期',

                    align: 'center',

                    valign: 'middle',

                    sortable: false

                }, {

                    field: 'domain_modify_user',

                    title: '修改人',

                    align: 'center',

                    valign: 'middle',

                    sortable: false,
                }, {

                    title: '操作',

                    align: 'center',

                    valign: 'middle',

                    sortable: false,

                    formatter:function(value,rows,index){

                        return '<span class="h-td-btn btn-success btn-xs" onclick="DomainObj.getSharePage(\''+rows.domain_id+'\',\''+ rows.domain_desc+'\')">共享管理</span>'

                    }
                }]
            }).closest(".bootstrap-table").find(".columns-right").hide();
        },

        domainInsertRow:function(e){
            var submitMenu = function(hmode){
                $.HAjaxRequest({
                    type:"Post",
                    url:"/v1/auth/domain/post",
                    cache:false,
                    data:$("#h-domain-add-tpl").serialize(),
                    async:true,
                    dataType:"text",
                    success: function(data){
                        $.Notify({
                            title:"操作成功",
                            message:"创建域信息成功",
                            type:"success",
                        });
                        $(hmode).remove();
                        $("#HdomainInfoTable").bootstrapTable('refresh');
                    }
                });
            };

            $.Hmodal({
                callback:submitMenu,
                header:"新增域信息",
                body:$("#domain_input_form").html(),
                height:"360px",
                preprocess:function () {
                    $("#h-domain-add-status").Hselect({
                        height:"30px",
                    })
                },
            })
        },

        domainDeleteRow:function(){
            var rst = $("#HdomainInfoTable").bootstrapTable("getAllSelections")
            if (rst.length > 0){
                $.Hconfirm({
                    callback:function () {
                        $.HAjaxRequest({
                            type:'post',
                            url:"/v1/auth/domain/delete",
                            data:{JSON:JSON.stringify(rst)},
                            success: function(data){
                                $.Notify({
                                    title:"操作成功",
                                    message:"删除域信息成功",
                                    type:"success",
                                });
                                $(rst).each(function(index,element){
                                    $("#HdomainInfoTable").bootstrapTable('removeByUniqueId', element.domain_id);
                                })
                            }
                        });
                    },
                    body:"确定删除这个域，删除后将无法恢复",
                });
            }else{
                $.Notify({
                    title:"温馨提示",
                    message:"您没有勾选任何需要删除的域",
                    type:"info",
                });
            }
        },
        domainEditRow:function(){

            var rst = $("#HdomainInfoTable").bootstrapTable('getSelections');
            if (rst.length == 0){
                $.Notify({
                    title:"温馨提示",
                    message:"请在下方表格中选中需要修改的行",
                    type:"info",
                })
                return
            }else if (rst.length > 1){
                $.Notify({
                    title:"温馨提示",
                    message:"您选中了多行，不知道您具体需要编辑哪一行",
                    type:"warning",
                })
                return
            }else if (rst.length == 1){
                var tr = rst[0];
                var editMenu = function(hmode){
                    var domainid = tr.domain_id;
                    var domainDesc = $("#h-domain-add-tpl").find("input[name='domainDesc']").val();
                    var domainStatus = $("#h-domain-add-tpl").find("select[name='domainStatus'] option:selected").val();

                    $.HAjaxRequest({
                        type:"Put",
                        url:"/v1/auth/domain/update",
                        cache:false,
                        data:{
                            domainId: domainid,
                            domainDesc: domainDesc,
                            domainStatus: domainStatus,
                        },
                        async:false,
                        dataType:"text",
                        success: function(data){
                            $.Notify({
                                title:"温馨提示:",
                                message:"更新域信息成功",
                                type:"success",
                            });

                            $(hmode).remove();
                            $("#HdomainInfoTable").bootstrapTable('refresh');
                        }
                    });
                };

                var preHand = function(){
                    var domainid = tr.domain_id;
                    var domainDesc = tr.domain_desc;
                    var domainStatus = tr.domain_status=="正常"?0:1;
                    $("#h-domain-add-tpl").find("input[name='domainId']").val(domainid).attr("readonly","readonly");
                    $("#h-domain-add-tpl").find("input[name='domainDesc']").val(domainDesc);

                    $("#h-domain-add-status").Hselect({
                        height:"30px",
                    });
                    $("#h-domain-add-status").val(domainStatus).trigger("change");

                };
                $.Hmodal({
                    callback:editMenu,
                    preprocess:preHand,
                    header:"修改域信息",
                    body:$("#domain_input_form").html(),
                    height:"360px",
                })
            }
        },
    };
    $(document).ready(function(){
        DomainObj.getDomainInfo()
    });

</script>


<script id="domain_input_form" type="text/html">
    <div class="col-sm-12 col-md-12 col-lg-12" style="margin-top: 15px">
        <form class="form-horizontal" id="h-domain-add-tpl">
            <div class="col-sm-12 col-md-12 col-lg-12">
                <div class="form-group-sm col-sm-12 col-md-12 col-lg-12">
                    <label class="h-label" style="width:100%;">域编码：</label>
                    <input placeholder="1至30个字母，数字组成" name="domainId" type="text" class="form-control" style="width: 100%;height: 30px;line-height: 30px;">
                </div>
                <div class="form-group-sm col-sm-12 col-md-12 col-lg-12" style="margin-top: 15px;">
                    <label class="h-label" style="width: 100%;">域描述：</label>
                    <input placeholder="1至30个汉字、字母、数字组成" type="text" class="form-control" name="domainDesc" style="width: 100%;height: 30px;line-height: 30px;">
                </div>
                <div class="form-group-sm col-sm-12 col-md-12 col-lg-12" style="margin-top: 15px;">
                    <label class="h-label" style="width: 100%;">状　态：</label>
                    <select id="h-domain-add-status" name="domainStatus"  class="form-control" style="width: 100%;height: 30px;line-height: 30px;">
                        <option value="0">正常</option>
                        <option value="1">失效</option>
                    </select>
                </div>
            </div>
        </form>
    </div>
</script>