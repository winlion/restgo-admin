
var roleapp = new Vue({
    el:"#pagecontent",
    data:{
        roles:[],
      
        drole:{
            "name":""
        },
        allauth:[],
        dmod:{
            "name":""
        },
        icon:{
            "mod":asseturi+"/images/mod.png",
            "api":asseturi+"/images/api.png",
            "page":asseturi+"/images/page.png"
        },
        dres:{
            "name":"",
            "restype":"",
            "pid":0,
            "patern":"",
        },
    },
    methods:{
        loadrole:function(){
            var that = this
            restgo.post("role/search",{}).then(function(res){
                that.roles = res.rows
            })
        },
      
       
        faddrole:function(){
            $("#mod-role").modal("show")
        },
        faddresource:function(){
            $("#mod-res").modal("show")
        },
        faddmodule:function(){
            $("#mod-mod").modal("show")
        },
        confirmmod:function(){

            if (!this.dmod.name){
                alert("请输入名称")
                return
            }
            this.dmod.pid=0;
            this.dmod.patern="^\\s+$";
            restgo.post("resource/addmod",this.dmod).then(function(res){
                alert(res.msg)
                $("#mod-role").modal("hide")
            })
        },
        confirmres:function(){

            if (!this.dres.name){
                alert("请输入名称")
                return
            }
            if (!this.dres.pid){
                alert("请选择模块")
                return
            }

            if (!this.dres.patern){
                alert("请输入格式")
                return
            }
            restgo.post("resource/addres",this.dres).then(function(res){
               alert(res.msg)
            })
        },
        confirmrole:function(){

            if (!this.drole.name){
                alert("请输入名称")
                return
            }
            var that =this
            restgo.post("role/create",this.drole).then(function(res){
                alert(res.msg)
                $("#mod-role").modal("hide")
                that.loadrole()
                
            })
        },
        showauth:function(id){
            var that = this;
            restgo.post("role/loadauth",{"roleid":id}).then(function(res){
                if(res.code!=200){
                    return ;
                }
                for(var i in that.allauth){
                    that.allauth[i].resid = that.allauth[i].id;
                    that.allauth[i].icon = that.icon[that.allauth[i].restype];
                        that.allauth[i].roleid = id;
                        that.allauth[i].checked = false;
                        that.allauth[i].open = true;
                        that.allauth[i].pId = that.allauth[i].pid||0;
                        for(var j in res.rows){
                            if (res.rows[j].id==that.allauth[i].id){
                                that.allauth[i].checked = true
                            }
                        }
                }
                that.setupztree(that.allauth)
            })
            
        },
        loadallauth:function(){
            var that = this;
            restgo.post("resource/search",{}).then(function(res){
                    that.allauth = res.rows;
            })
        },
        zTreeOnCheck:function(event, treeId, treeNode){
               
            this.grantorrevopke(treeNode)
            if(typeof treeNode.children!="object"){
                return ;
            }
            for(var i in treeNode.children){
                this.grantorrevopke(treeNode.children[i])
            }
           
        },
        grantorrevopke:function(node){
                if(node.checked==true){
                    restgo.post("role/grantauth",node).then(function(res){
                        console.log(res)
                    })
                }else{
                    restgo.post("role/revokeauth",node).then(function(res){
                        console.log(res)
                    })
                }
        },
        setupztree:function(zNodes){
            var that = this;
            var setting = {
                check: {
                    enable: true
                },
                data: {
                    simpleData: {
                        enable: true
                    }
                },
                callback: {
                    onCheck:that.zTreeOnCheck
                }
            };
            $("#authtree").html("");     
            $.fn.zTree.init($("#authtree"), setting, zNodes);

        }

    },
    mounted:function(){
      
        this.loadrole()
        this.loadallauth()
    }
})
