<template>
    <div class="mainBody">
<!--        左侧边栏-->
        <div class="leftBar">
            <LeftNavBar/>
        </div>
<!--        数据可视化部分-->
        <div class="dataView">
            <transition
                    name="animate__animated animate__bounce router_view"
                    enter-active-class="animate__fadeIn"
                    leave-active-class="animate__fadeOut"
                    appear
            >
            <router-view style="transition: all 0.1s" />
            </transition>
        </div>


    </div>
</template>
<script>
// import 'animate.css';
    import LeftNavBar from "../components/LeftNavBar";
    // import PersonalView from "../components/PersonalView";
    export default {
        name: "Manage",
        components: { LeftNavBar }, //,PersonalView
        data(){
            return{
                personalViewFlag:true
            }
          },
        mounted() {}
    }
</script>

<style scoped>
.mainBody{
    height: 100vh;
    background-color: #257B5E;
    display: flex;
    justify-content: center;
    align-items: center;
    display: flex;
    background-image: url("../assets/images/background.png");
    background-size: 100%;
    position: relative;
}
    .dataView{
        width: 900px;
        height: 650px;
        background-color:#FDFFFD ;
        border-bottom-right-radius: 15px;
        border-top-right-radius: 15px;
        overflow: hidden;
        position: relative;
    }
    .personalCenter{
        position: absolute;
        right: 0px;
        top: 0px;
    }
    .router_view-leave-active {
        opacity: 0;
    }
</style>