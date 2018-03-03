var Regsiter={
    data:{
        build:function () {
            Regsiter.data.user.email = $("#account").val();
            Regsiter.data.user.account = $("#account").val();
            Regsiter.data.user.passwd = $("#passwd").val();
            Regsiter.data.user.verify = $("#code").val();
            Regsiter.data.user.nickname = $("#nickname").val();
        },
        user:{
            "email":null,
            "account":null,
            "passwd":null,
            "nickname":null,
            "verify":null
        }
    },
    validate:{
        user:function(){
        if (!restgo.testEmail(Regsiter.data.user.email)){
            return [false,"请输入正确的邮箱"]
        }

        if (!restgo.testEmail(Regsiter.data.user.account)){
            return [false,"请输入正确的邮箱"]
        }

        if (!restgo.testReg("[a-zA-Z]+",Regsiter.data.user.passwd)){
            return [false,"密码必须包含字母"]
        }
        if (!restgo.testReg("[0-9]+",Regsiter.data.user.passwd)){
                return [false,"密码必须包含数字"]
        }
        if (Regsiter.data.user.passwd.length<6){
            return [false,"密码至少长6位"]
        }

        if (!$("#accept").prop("checked")){
            return [false,"只有同意我们的协议才可以继续"]
        }
        if (!Regsiter.data.user.nickname){
            Regsiter.data.user.nickname = "会员"+(new Date().getTime()%100);
        }
        return [true,Regsiter.data.user]
     }
    },
    event:{
            init:function(){
                $("#btn-register").click(Regsiter.service.register)
            }
    },
    service:{
        register:function(){
                Regsiter.data.build();
                var r = Regsiter.validate.user()
                if (!r[0]){
                    restgo.error(r[1])
                    return
                }

                restgo.post("user/register",r[1]).then(function (value) {
                    alert(value.msg)
                    if (value.code==200){
                        location.href="/"
                    }
                })
        }
    },

}
$(function(){
    Regsiter.event.init()
})