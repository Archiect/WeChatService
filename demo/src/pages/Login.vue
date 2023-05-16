<template>
  <div class="base">
    <!-- 注册登录界面 -->
   <div  class="loginAndRegist">
        <!--登录表单-->
        <div  class="loginArea">
            <transition
                    name="animate__animated animate__bounce"
                    enter-active-class="animate__fadeInUp"
                    leave-active-class="animate__zoomOut"
                    appear
            >
            <!-- 标语 -->
            <div v-show="isShow" class="title">
                LOGIN
            </div>
            </transition>
            <transition
                    name="animate__animated animate__bounce"
                    enter-active-class="animate__fadeInUp"
                    leave-active-class="animate__zoomOut"
                    appear
            >
            <!-- 密码框和用户名框 -->
            <div v-show="isShow" class="pwdArea">
              <div style="flex: 1;justify-content: center;display: flex;align-items: center">
                  <el-input class="username" v-model="loginUser.name" style="width: 165px"  placeholder="用户名"></el-input>
              </div>
                <div style="flex: 1;justify-content: center;display: flex;align-items: center">
                    <el-input placeholder="密码"  v-model="loginUser.password" style="width: 165px" show-password></el-input>
                </div>
            </div>
            </transition>
            <transition
                    name="animate__animated animate__bounce"
                    enter-active-class="animate__fadeInUp"
                    leave-active-class="animate__zoomOut"
                    appear
            >
            <!-- 登录注册按钮 -->
            <div v-show="isShow" class="btnArea">
                <el-button type="success" round style="background-color: #257B5E;letter-spacing: 5px" @click="UserLogin"  >登录</el-button>
            </div>
            </transition>
        </div>
        <!-- 注册表单 -->
        <div class="registArea">
            <transition
                    name="animate__animated animate__bounce"
                    enter-active-class="animate__fadeInUp"
                    leave-active-class="animate__zoomOut"
                    appear
            >
            </transition>
        </div>
        <!-- 信息展示界面 -->
        <div id="aaa" class="showInfo"
              :style="{
            borderTopRightRadius:styleObj.bordertoprightradius,
            borderBottomRightRadius:styleObj.borderbottomrightradius,
            borderTopLeftRadius:styleObj.bordertopleftradius,
            borderBottomLeftRadius:styleObj.borderbottomleftradius,
            right:styleObj.rightDis
            }"
              ref="showInfoView">

            <transition
                    name="animate__animated animate__bounce"
                    enter-active-class="animate__fadeInUp"
                    leave-active-class="animate__zoomOut"
                    appear
            >
            <!-- 没有用户输入用户名或者找不到用户名的时候 -->
            <div v-show="isShow" style="display: flex;flex-direction: column;align-items: center;justify-content: center;width: 100%;height: 100%">
            <!-- 欢迎语 -->
                <div style="flex: 2;display: flex;align-items: center;font-size: 22px;color: #FFFFFF;font-weight: bold">
                 Demo_
                 </div>
            <!-- 欢迎图片 -->
                 <!-- <div style="flex: 2">
                     <el-button type="success"  style="background-color:#257B5E;border: 1px solid #ffffff;" round @click="changeToRegiest">还没有账户？点击注册</el-button>
                 </div> -->
           </div>
             </transition>
        </div>
   </div>
   
  </div>
</template>
<script >
    import 'animate.css';
    // eslint-disable-next-line no-unused-vars
    import axios from "axios";
export default {

    name:'Login',
    data(){
        return{
            //看看用不用转成用户对象
            loginUser:{
                name:"",
                password:""
            },

            admins:{},
            ////看看用不用转成用户对象
            regUser:{
                regUsername:'',
                regRePwd:'',
                regPwd:'',
                selectValue:"",
            },
            styleObj:{
                bordertoprightradius:'15px',
                borderbottomrightradius: '15px',
                bordertopleftradius:'0px',
                borderbottomleftradius:'0px',
                rightDis:'0px'
            },
            isShow:true
        }
    },
    created() {
        // this.loadInfoOfAdmin();
    },
    methods:{
        changeToRegiest(){
            this.styleObj.bordertoprightradius= '0px'
            this.styleObj.borderbottomrightradius='0px'
            this.styleObj.bordertopleftradius='15px'
            this.styleObj.borderbottomleftradius='15px'
            this.styleObj.rightDis='50%'
            this.isShow = !this.isShow
        },
        changeToLogin(){
            this.styleObj.bordertoprightradius= '15px'
            this.styleObj.borderbottomrightradius='15px'
            this.styleObj.bordertopleftradius='0px'
            this.styleObj.borderbottomleftradius='0px'
            this.styleObj.rightDis='0px'
            this.isShow = !this.isShow
        },
        //用户登录
      UserLogin() {
        // console.log(this.loginUser)
        // 这里有个跨域问题
        /*
          axios是ajax ajax不止axios
        */
        this.$axios.post('/userinfo', {
          name: this.loginUser.name,
          password: this.loginUser.password
        }).then(res => {
          // console.log(res.status==200)
          if (res.status == 200) {
            localStorage.setItem("user", JSON.stringify(res.data))
            this.$message.success("登陆成功！")
            this.$router.push("/manage")
          } if (res.status == 202) {
            this.$message.error("用户名或密码错误！")
          }
        })
      }
    },
}
</script>

<style>
/* @import '../assets/css/Login.css' */
.base{
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background-image: url("../assets/images/background.png");
    background-size: 100%;
}
.loginAndRegist{
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    overflow: hidden;
}
.loginArea{
    background-color: rgba(255,255,255,0.8);
    border-top-left-radius: 15px;
    border-bottom-left-radius: 15px;
    height: 400px;
    width: 350px;
    z-index: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    overflow: hidden;
}
.registArea{
    border-top-right-radius: 15px;
    border-bottom-right-radius: 15px;
    height: 400px;
    width: 350px;
    background-color: rgba(255,255,255,0.8);
    z-index: 1;
    display: flex;
    flex-direction: column;
    justify-content:center ;
    align-items: center;
}
.showInfo{
    border-top-right-radius: 15px;
    border-bottom-right-radius: 15px;
    position: absolute;
    height: 400px;
    width: 350px;
    z-index:2;
    top: 0;
    right: 0;
    background-image: url("../assets/images/welcome.png");
    background-size: 90%;
}
.showInfo:hover{
    background-size: 100%;
    background-position: -50px -50px;
}
.title{
    width: 70%;
    flex:1;
    border-bottom: 1px solid #257B5E;
    display: flex;
    align-items: center;
    color: #257B5E;
    font-weight: bold;
    font-size: 24px;
    display: flex;
    justify-content: center;
}
#aaa{
    transition: 0.3s linear;
}
.pwdArea{
    width: 100%;
    flex:2;
    display: flex;
    flex-direction: column;
    display: flex;
    flex-direction: column;


}
.pwdArea input{
    outline: none;
    height: 30%;
    border-radius:13px ;
    padding-left: 10px;
    font-size: 12px;
    border: 1px solid gray;
}
.pwdArea input:focus{
    border: 2px solid #257B5E;
}
    .btnArea{
        flex:1;
        width: 100%;
        display: flex;
        justify-content: space-around;
        align-items: center;
    }
    .rigestTitle{
        width: 70%;
        flex: 1;
        color: #257B5E;
        font-weight: bold;
        font-size: 24px;
        display: flex;
        justify-content: center;
        align-items: center;
        border-bottom: 1px solid #257B5E;
    }
    .registForm{
        flex: 2;
        display: flex;
        flex-direction: column;
        color: #257B5E;
        font-size: 16px;
    }

    .registForm input{
    outline: none;
    height: 30%;
    border-radius:13px ;
    padding-left: 10px;
    font-size: 12px;
    border: 1px solid gray;
    }
.registForm input:focus{
    border: 2px solid #257B5E;
}
#elselect:focus{
    border: 2px solid #257B5E;
}
    .registBtn{
        flex: 1;
        display: flex;
        justify-content: center;
        align-items: center;
    }


</style>