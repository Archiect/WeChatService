<template>
  <transition name="animate__animated animate__bounce" enter-active-class="animate__fadeInUp"
    leave-active-class="animate__zoomOut" appear>
  <el-form class="Root" :model="ruleForm"  :rules="rules" ref="ruleForm">
    <!-- 描述 -->
    <el-form-item prop="text">
      <!--  @input="change($event)"强制更新组件  resize="none" 禁止改变大小-->
      <el-input type="textarea" :rows="10" resize="none" placeholder="请输入您的描述" v-model="ruleForm.text" @input="change($event)" class="Input-Box">
      </el-input>
    </el-form-item>
    <!-- 下拉栏 -->
    <el-form-item prop="type">
      <el-select v-model="ruleForm.type" placeholder="类型">
        <el-option v-for="item in selectList" :key="item.id" :label="item.name" :value="item.id"></el-option>
      </el-select>
    </el-form-item>
    <!-- 图片 -->
    <el-form-item prop="imageUrl" ref="uploadElement">
      <el-upload ref="upload"
        class="avatar-uploader"
        list-type="picture-card"
        action="/api/Mupload"
        :auto-upload="false"
        :before-upload="beforeAvatarUpload"
        :show-file-list="false"
        :data="ruleForm"
        :on-change="handleChange"
        :on-success="handle_success"
        >
        <img v-if="ruleForm.imageUrl" :src="ruleForm.imageUrl" class="avatar">
        <i class="el-icon-plus"></i>
      </el-upload>
    </el-form-item>
    <el-button type="success" plain class="Btn-Box" @click="submitForm('ruleForm')">提交</el-button>
    <el-button @click="resetForm('ruleForm')" type="info" class="Btn-Box">重置</el-button>
  </el-form>
</transition>
</template>

<script>

// import { request } from '../utils/axios';
// import 'animate.css';
export default {
  name: "PersonnelDetection",
  data() {
    return {
      selectList: [
        { id: '校园公告', name: '校园公告' },
        { id: '社团活动', name: '社团活动' },
      ],
      // imageUrl:'',
      ruleForm: {
        text: '',
        type: '',
        imageUrl: '',
      },
      uploadistrue:false,
      rules: {
        text: [
          { required: true, message: '请输入您的描述', trigger: 'blur' },
        ],
        type: [
          { required: true, message: '请选择类型', trigger: 'blur' },
        ],
        imageUrl: [
          { required: true, message: '请上传图片', trigger: 'blur' },
        ],
      }
    }
  },
  methods: {
    beforeAvatarUpload() {
      return true
    },
    change() {
      this.$forceUpdate()
    },
    // formName
    submitForm(formName) {
      let vm = this;
      this.$refs[formName].validate((valid) => {
        if (valid) {
          // console.log(vm.ruleForm)
          vm.$refs.upload.submit()
        } else {
          return false;
        }
      });
    },
    handle_success(res) {
      if (res.istrue) {
        this.$message({
          duration: 900,
          message: "上传成功",
          type: 'success',
          onClose: () => {
            //此处写提示关闭后需要执行的函数
            location.reload()
          }
        })
      } else {
        this.$message.error('上传失败')
      }
    },
    handleChange(file) {
      this.ruleForm.imageUrl = URL.createObjectURL(file.raw);
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
      this.ruleForm.imageUrl = '';
    },
  }
}
</script>

<style>
.Root{
  margin-top: 20px;
  width: 900px;
  height: 600px;
  /* border: 3px solid red; */ /* 边框 */
  text-align: center;
}
.Btn-Box{
  width: 100px;
  height: 40px;
}
.Input-Box{
  width: 500px;
}
.avatar-uploader .el-upload {
      border: 1px dashed #d9d9d9;
      border-radius: 6px;
      cursor: pointer;
      position: relative;
      overflow: hidden;
    }
  
    .avatar-uploader .el-upload:hover {
      border-color: #409EFF;
    }
  
    .avatar-uploader-icon {
      font-size: 28px;
      color: #8c939d;
      width: 178px;
      height: 178px;
      line-height: 178px;
      text-align: center;
    }
  
    .avatar {
      width: 178px;
      height: 178px;
      display: block;
    }
</style>