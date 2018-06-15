(function($){
var UserList={
		
}
	
function buildArgs(){
    var data={}
    if(!!$("#ttype").val()){
        data.ttype =$("#ttype").val()
    }
    if(!!$("#datefrom").val()){
        data.datefrom =$("#datefrom").val()
    }
    if(!!$("#dateto").val()){
        data.dateto =$("#dateto").val()
    }
    if(!!$("#kword").val()){
        data.kword =$("#kword").val()
    }

	return data;
}

$(".picker").datetimepicker({
        format: 'yyyy-mm-dd hh:ii:ss',
		autoclose:true,
		startView:2,
		language:"zh-CN",
		todayBtn:true,
		todayHighlight:true
 });
$("#table").on("click",".updatestat",function(){
    var data = $(this).data()
    restgo.post("user/updatestat",data).then(function(r){
        alert(r.msg)
        if(r.code==200){
            $('#table').bootstrapTable("refresh")
        }
    })
})
 $('#table').bootstrapTable({
     queryParamsType: '',              //默认值为 'limit' ,在默认情况下 传给服务端的参数为：offset,limit,sort
     queryParams: buildArgs,//传递参数（*）
     method: "post",
     contentType:"application/x-www-form-urlencoded",
     columns: [{
        field: 'id',
        title: 'Item ID'
    }, {
        field: 'account',
        title: '账号'
    }, {
         field: 'email',
         title: '邮箱'
     },{
         field: 'mobile',
         title: '手机号'
     },{
         field: 'nickname',
         title: '昵称'
     }, {
        field: 'createat',
        title: '注册时间'
    },{
         field: 'stat',
         title: '状态',
         formatter:function(value, row, index, field){
             return value==1?"可用":"禁用"
         }
     },{
         field: 'id',
         title: '操作',
         formatter:function(value, row, index, field){
             if (row.stat==1){
                 return '<input type="button" class="btn btn-xs btn-danger updatestat " data-id="'+value+'"  data-stat="0" value="禁用">'
             }else{
                 return '<input type="button" class="btn btn-xs btn-success updatestat " data-id="'+value+'" data-stat="1" value="启用">'
             }

         }
     }],
	showFooter:false,
	pageNumber:1,
	pageSize:20,
	pageList:[10, 25, 50, 100, 200],
	pagination: true,                   //是否显示分页（*）
	sortName:"id",
	showColumns:true,
	showToggle:true,
	showPaginationSwitch:true,
	minimumCountColumns:10,
    sidePagination: "server",           //分页方式：client客户端分页，server服务端分页（*）
	cardView:false,
	detailView:false,
	singleSelect:true,
	sortable:true,
	silentSort:true,
	toolbar:"#toolbar",
}); 
$("#btn-search").click(function(){
    $('#table').bootstrapTable("refresh")
})
})(jQuery)