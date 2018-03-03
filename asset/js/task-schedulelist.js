(function($){
var TaskList={
		
}
	
function buildArgs(){
    var data={}
    var ttype = $("#ttype").val()
    if(!!ttype){
        if(ttype=="createat"){
            data.ttype = "createat"
        }else {
            data.stat = ttype
        }
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
$("#ttype").change(function () {
    var v = $(this).val();
    if(v=="createat"){
        $(".picker").show()
    }else{
        $(".picker").hide()
    }
})
 $('#table').bootstrapTable({
     queryParamsType: '',              //默认值为 'limit' ,在默认情况下 传给服务端的参数为：offset,limit,sort
     queryParams: buildArgs,//传递参数（*）
     method: "post",
     contentType:"application/x-www-form-urlencoded",
     columns: [{
        field: 'id',
        title: '编号'
    }, {
        field: 'name',
        title: '任务描述'
    }, {
        field: 'createat',
        title: '发布时间'
    },{
         field: 'stat',
         title: '当前状态'
     },{
         field: 'id',
         title: '操作'
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