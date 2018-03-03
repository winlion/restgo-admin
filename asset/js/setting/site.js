(function($){
var siteConfig={
		
}
	
function buildArgs(){
    var data={}
    
    if(!!$("#kword").val()){
        data.kword =$("#kword").val()
    }

	return data;
}
$("#btn-addconfig").click(function(){
    $("#mod-config").modal("show")
})

$("#submitconfig").click(function(){
    var data={}
    data.name = $("[v-model='dres.name']").val();
    data.label = $("[v-model='dres.label']").val();
    data.value = $("[v-model='dres.value']").val();
    data.format = $("[v-model='dres.format']").val();
    if(!/^[a-zA-Z][a-zA-Z0-9]+$/.test(data.name)){
        alert("参数名称必须是英文")
        return ;
    }

    if(!data.label){
        alert("请说明一下这个参数是干什么的")
        return ;
    }

    if(!data.value){
        alert("参数初始化值为空")
        
    }
    restgo.post("config/create",data).then(function(resp){
      
        if (resp.code==200){
            alert("配置添加成功")
            $("#btn-search").trigger("click")
            $("#mod-config").modal("hide")
            
        }else{
            alert("服务器繁忙请稍后")
        }
    })    


})
var  alldata=[]
$(".picker").datetimepicker({
        format: 'yyyy-mm-dd hh:ii:ss',
		autoclose:true,
		startView:2,
		language:"zh-CN",
		todayBtn:true,
		todayHighlight:true
 });
$("#table").on("click",".update",function(){

    var id = $(this).attr("data-id");
    $("#config-value").val("")
    $("#config-value").attr("data-format","")
    $("#config-label").html()
    $("#config-value").attr("data-name","")
    $("#mod-editconfig").modal("show")
    for(var i in alldata){
        if(alldata[i].name==id){
            $("#config-value").val(alldata[i].value)
            $("#config-value").attr("data-format",alldata[i].format)
            $("#config-value").attr("data-name",alldata[i].name)
            $("#config-label").html("正则规则:"+alldata[i].format||"无")
        }
    }
    
})

$("#btn-submit-config").click(function(){
    var data = $("#config-value").data();
    data.value=$("#config-value").val();
    if (!data.name){
        alert("缺少参数编号");
        return ;
    }

    if (!data.value){
        alert("缺少参数值");
        return ;
    }

    if (!!data.format){
        var reg = new RegExp(data.format);
        if (!reg.test(data.value)){
            alert("参数格式不正确");
            return ;
        }
    }
   
    delete data.format
    restgo.post("config/update",data).then(function(resp){
      
        if (resp.code==200){
            alert("配置修改成功")
            $("#btn-search").trigger("click")
            $("#mod-editconfig").modal("hide")
            
        }else{
            alert("服务器繁忙请稍后")
        }
    })    

})
 $('#table').bootstrapTable({
     queryParamsType: '',              //默认值为 'limit' ,在默认情况下 传给服务端的参数为：offset,limit,sort
     queryParams: buildArgs,//传递参数（*）
     method: "post",
     onLoadSuccess:function(d){
        alldata = d.rows||[]
     },
     contentType:"application/x-www-form-urlencoded",
     columns: [{
        field: 'name',
        title: '参数ID'
        
    }, {
        field: 'label',
        title: '参数名称'
    }, {
         field: 'format',
         title: '参数格式',
         visible:false
     },  {
        field: 'value',
        title: '参数值'
      
    },{
         field: 'opt',
         title: '操作',
         formatter:function(value, row, index, field){
            
            return '<input type="button" class="btn btn-xs btn-success update " data-id="'+row.name+'"  value="修改">'
             

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