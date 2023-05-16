import axios from 'axios';
//load是上传图片时放加载画面的盒子
//let load=document.querySelector(".load");
//config是图片加载画面的一些属性，可以自己定义，也可以在axios官网头像上传那块找到上找到。
export function request(url,method="get",data={},config={}){
    //load.style.display="block";
    return axiosRequest(url, method,data, config)
}

function axiosRequest(url,method,data) {

    if (method.toLowerCase()==='post'){
        if (data instanceof Object){
            let params=new URLSearchParams();
            for (let key in data){
                params.append(key, data[key]);
            }
            data = params;
        }
    }else if (method.toLowerCase()==='file'){
        method="post";
        if (data instanceof Object){
            let params=new FormData();
            for (let key in data){
                params.append(key, data[key]);
            }
            data = params;
        }
    }
    let axiosConfig={
        url:url,
        method:method.toLowerCase(),
        data:data
    };
    //if (config instanceof Object){
        //for (let key in config){
           // axiosConfig[key]=config[key];
        //}
   // }
    return axios(axiosConfig).then(res=>{
        //load.style.display="none";
        return res.data
    });
}
