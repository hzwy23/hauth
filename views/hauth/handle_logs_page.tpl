<div class="row subsystem-header">
    <div class="pull-left">
        <span style="font-size: 16px;">操作记录</span>
    </div>
</div>
<div class="row subsystem-toolbar">
    <div class="col-sm-12 col-md-12 col-lg-12">
        <div id="hUserLogsTableTools" class="pull-left">
            <button onclick="LogsHandle.search()" class="btn btn-success btn-sm">
                <i class="icon-search"> 搜索</i>
            </button>
            <button onclick="LogsHandle.download()" class="btn btn-success btn-sm" title="下载操作记录">
                <span class="icon-wrench"> 下载</span>
            </button>
        </div>
        <div class="pull-right" style="height: 44px; line-height: 44px; width: 180px;">
            <input type="text" class="form-control" placeholder="页面内容搜索框" style="height: 30px; line-height: 30px; margin-top: 7px;">
        </div>
    </div>
</div>

<div class="row">
    <div id="h-handle-logs" class="col-sm-12 col-md-12 col-lg-12">
        <table id="HandleLogsPageTable"></table>
    </div>
</div>

<script type="text/javascript">
    var LogsHandle = {
        download:function(){
            var x=new XMLHttpRequest();
            x.open("GET", "/v1/auth/handle/logs/download", true);
            x.responseType = 'blob';
            x.onload=function(e){
                download(x.response, "操作记录.xlsx", "application/vnd.ms-excel" );
            }
            x.send();
        },
        showHandleLogDetails:function(val){
            var optHtml = '<div class="panel panel-default"><table class="table table-striped table-bordered table-condensed">'
            optHtml += "<tr><th class='col-sm-2 col-md-2 col-lg-2'>key</th><th>value</th></tr>"
            var st = val.split("&")

            for(var i = 0 ; i < st.length;i++){
                var keyval = st[i].split("=")
                if (keyval.length==2){
                    if (keyval[0]=="JSON"){
                        console.log(JSON.stringify(decodeURI(keyval[1])))
                    }
                    optHtml += "<tr><td>"+keyval[0]+"</td><td>"+decodeURI(keyval[1])+"</td></tr>"
                }
            }
            optHtml += "</table></div>"

            $.Hmodal({
                header:"客户端发送到服务器的参数信息",
                body:optHtml,
                height:"360px",
            })
        },
        search:function(){

            var hSubmit = function(hmode){
                var userId = $("#h-logs-search-form").find("input[name='UserId']").val();
                var startDate = $("#h-logs-search-form").find("input[name='StartDate']").val();
                var endDate = $("#h-logs-search-form").find("input[name='EndDate']").val();
                $("#HandleLogsPageTable").bootstrapTable('destroy');
                var hwindow = document.documentElement.clientHeight - 120;
                $("#HandleLogsPageTable").bootstrapTable({
                    url:'/v1/auth/handle/logs/search',
                    height:hwindow,
                    uniqueId:'uuid',
                    striped: true,
                    pagination: true,
                    search:false,
                    showRefresh:false,
                    queryParams:function(params){
                        return {
                            UserId:userId,
                            StartDate:startDate,
                            EndDate:endDate,
                        }
                    },
                    pageSize: 20,
                    showExport:true,
                    sidePagination: "client",
                    showColumns: false,
                    minimunCountColumns: 2,
                    columns:[{

                        field: 'uuid',

                        title: '账号',

                        align: 'left',

                        visible:false,

                        valign: 'middle',

                        sortable: true

                    },{

                        field: 'user_id',

                        title: '账号',

                        align: 'left',

                        valign: 'middle',
                    }, {

                        field: 'handle_time',

                        title: '操作日期',

                        align: 'left',

                        valign: 'middle',

                        sortable: false,

                    }, {

                        field: 'client_ip',

                        title: '客户端IP',

                        align: 'left',

                        valign: 'middle',

                        sortable: false,

                    },{

                        field: 'method',

                        title: '请求方式',

                        align: 'left',

                        valign: 'top',

                        sortable: false

                    }, {

                        field: 'url',

                        title: '路由信息',

                        align: 'left',

                        valign: 'middle',

                        sortable: false

                    }, {

                        field: 'status_code',

                        title: '返回状态',

                        align: 'left',

                        valign: 'middle',

                        sortable: false

                    }, {

                        field: 'data',

                        title: '请求发送',

                        align: 'left',

                        valign: 'middle',

                        sortable: false,

                        formatter:function(value,rows,index){
                            if (value.length>30){
                                return '<span ondblclick=LogsHandle.showHandleLogDetails("'+value+'") >'+value.substring(0,30)+'......'+'</span>'
                            }else{
                                return '<span ondblclick=LogsHandle.showHandleLogDetails("'+value+'") >'+value.substring(0,30)+'</span>'
                            }
                        }
                    }]
                }).closest(".bootstrap-table").find(".columns-right").hide();
                $(hmode).remove();
            };
            $.Hmodal({
                header:"高级搜索",
                body:$("#handle-logs-search").html(),
                callback:hSubmit,
                height:"320px",
            })
        },
        getLogs:function(args){

            //$("#h-handle-logs").height(document.documentElement.clientHeight-150)
            $("#HandleLogsPageTable").bootstrapTable({
                url:'/v1/auth/handle/logs',
                height:document.documentElement.clientHeight-120,
                uniqueId:'uuid',
                striped: true,
                pagination: true,
                pageList:[40,80,160,400,800,3000],
                showRefresh:true,
                pageSize: 80,
                showExport:false,
                search:false,
                sidePagination: "server",
                showColumns: true,
                minimunCountColumns: 2,
                columns:[{

                    field: 'uuid',

                    title: '账号',

                    align: 'left',

                    visible:false,

                    valign: 'middle',

                    sortable: false

                },{

                    field: 'user_id',

                    title: '账号',

                    align: 'left',

                    valign: 'middle',

                    sortable: false

                }, {

                    field: 'handle_time',

                    title: '操作日期',

                    align: 'left',

                    valign: 'middle',

                    sortable: false,

                }, {

                    field: 'client_ip',

                    title: '客户端IP',

                    align: 'left',

                    valign: 'middle',

                    sortable: false,

                },{

                    field: 'method',

                    title: '请求方式',

                    align: 'left',

                    valign: 'top',

                    sortable: false

                }, {

                    field: 'url',

                    title: '路由信息',

                    align: 'left',

                    valign: 'middle',

                    sortable: false

                }, {

                    field: 'status_code',

                    title: '返回状态',

                    align: 'left',

                    valign: 'middle',

                    sortable: false

                }, {

                    field: 'data',

                    title: '请求发送',

                    align: 'left',

                    valign: 'middle',

                    sortable: false,

                    formatter:function(value,rows,index){
                        if (value.length>30){
                            return '<span ondblclick=LogsHandle.showHandleLogDetails("'+value+'") >'+value.substring(0,30)+'......'+'</span>'
                        }else{
                            return '<span ondblclick=LogsHandle.showHandleLogDetails("'+value+'") >'+value.substring(0,30)+'</span>'
                        }
                    }
                }]
            }).closest(".bootstrap-table").find(".columns-right").hide();
        }
    };

    $(document).ready(function(){
        LogsHandle.getLogs();
    })

</script>
<script id="handle-logs-search" type="text/html">
    <form id="h-logs-search-form" class="form-horizontal col-sm-12 col-md-12 col-lg-12" style="margin-top: 20px;">
        <div class="form-group">
            <label class="col-sm-3 col-md-3 col-lg-3 control-label" style="font-size: 14px; font-weight: 500;">账 号：</label>
            <div class="col-sm-8 col-md-8 col-lg-8">
                <input style="height: 30px;line-height: 30px;" name="UserId"  type="text" class="form-control" placeholder="待查找的账号">
            </div>
        </div>
        <div class="form-group">
            <label class="col-sm-3 col-md-3 col-lg-3 control-label" style="font-size: 14px;font-weight: 500;">开始时间：</label>
            <div class="col-sm-8 col-md-8 col-lg-8">
                <input style="height: 30px;line-height: 30px;" onclick="laydate()" name="StartDate" class="form-control" placeholder="开始时间">
            </div>
        </div>
        <div class="form-group">
            <label class="col-sm-3 col-md-3 col-lg-3 control-label" style="font-size: 14px;font-weight: 500;">结束时间：</label>
            <div class="col-sm-8 col-md-8 col-lg-8">
                <input style="height: 30px;line-height: 30px;" onclick="laydate()" name="EndDate" class="form-control" placeholder="结束时间">
            </div>
        </div>
    </form>
</script>