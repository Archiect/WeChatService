import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter)
    // 创建并暴露一个路由器
export default new VueRouter({
    routes: [
        // 一级路由
        {
            path: '/',
            redirect: 'login'
        },
        {
            path: '/login',
            name: 'Login',
            component: () =>
                import ("../pages/Login.vue")
        },
        {
            path: '/manage',
            name: 'Manage',
            redirect: 'manage/home',
            component: () =>
                import ("../pages/Manage.vue"),
            children: [
                {
                    path:'weather',
                    name:'Weather',
                    component: ()=> import("../pages/Weather"),
                },
                {
                    path:'home',
                    name:'Home',
                    component: ()=> import("../pages/Home"),
                },
                {
                    path:'personnelDetection',
                    name:'PersonnelDetection',
                    component: ()=> import("../pages/PersonnelDetection"),
              },
              {
                path: 'Stufix',
                name: 'Stufix',
                component: ()=> import("../pages/Stufix"),
              },
                // {
                //     path: 'MyPublic',
                //     name: 'MyPublic',
                //     component:()=>import("../pages/MyPublic.vue")
                // },
                // // {
                //     path:'weather',
                //     name:'Weather',
                //     component: ()=> import("../pages/Weather"),
                // },

                {
                    path:'temperatureDetection',
                    name:'TemperatureDetection',
                    component: ()=> import("../pages/TemperatureDetection"),
                },
            ]
        }

        // // 一级路由
        // {
        //     name: 'guanyu',
        //     path: '/about',
        //     component: About
        // },
        // {
        //     path: '/home',
        //     component: Home,
        //     children: [
        //         // 子路由
        //         {
        //             path: 'news',
        //             component: News
        //         },
        //         {
        //             path: 'Message',
        //             component: Message,
        //             children: [{
        //                 name: 'xiangqing',
        //                 path: 'detail',
        //                 component: Detail,
        //                 // 第一种写法 值为对象  该对象中的值都以props的形式传递个该组件
        //                 // props: { a: 1, b: 'hello' }

        //                 // props的第二种写法 值为布尔值 若布尔值为真 将所有的params属性传递给组件
        //                 // props: true

        //                 // 第三种：函数
        //                 props($route) {
        //                     return { id: $route.query.id, title: $route.query.title }
        //                 }
        //             }]
        //         }
        //     ]
        // },

  ]
})