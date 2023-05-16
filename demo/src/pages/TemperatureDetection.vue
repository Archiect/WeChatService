<template>
  <div class="lsm-container" style="height: 100%;">
    <el-scrollbar class="scrollbar-wrapper">
      <!-- 内部文本-->
        <el-table :data="message" stripe style="width: 100%; margin-top: 20px; margin-left: 0px;" >
          <el-table-column prop="Sid" label="学号" width="180"> </el-table-column>
          <el-table-column prop="Content" label="描述" width="180"> </el-table-column>
          <el-table-column prop="Label" label="类型" width="180"> </el-table-column>
        
          <el-table-column label="图片" prop="img" width="180">
            <template scope="scope">
              <!-- http://localhost:8080/img/ -->
              <el-image :src="('http://localhost:8080/img/'+scope.row.Imgpath)"></el-image>
            </template>
          </el-table-column>
        
          <el-table-column align="center" label="操作" width="180">
            <template scope="scope">
              <el-row>
                <el-button-group>
                  <el-tooltip effect="light" content="通过" placement="top-start">
                    <!-- showView可以获取当前表单的信息 -->
                    <el-button icon="el-icon-check" type="success" @click="Allow(scope.row)" />
                  </el-tooltip>
                  <el-tooltip effect="light" content="不通过" placement="top-start">
                    <el-button icon="el-icon-close" type="danger" @click="Refuse(scope.row)" />
                  </el-tooltip>
                </el-button-group>
              </el-row>
            </template>
          </el-table-column>
        </el-table>
    </el-scrollbar>
  </div>
</template>
<script>
  // import Axios from 'axios';
export default {
  name: "TemperatureDetection",
  data() {
    return {
      message: [],
    }
  },
  created() {
    this.Loading();
  },
  methods: {
    Loading() {
      this.$axios.get('/Minfo', {}).then(res => {
        // console.log(res.data.Minfo.message)
        this.message = res.data.Minfo.message
      })
    },
    Allow(res) {
      // console.log(res)
      // 发送请求
      this.$axios.post('/update', {
        content: res.Content,
        imgpath: res.Imgpath,
        info: res.Info,
        infoid: res.Infoid,
        label: res.Label,
        sid:res.Sid,
        type:"allow"
      }).then(res => {
        if (res.data.msg) {
          this.$message({
            duration: 900,
            message: "更新成功",
            type: 'success',
            onClose: () => {
              //此处写提示关闭后需要执行的函数
              location.reload()
            }
          })
        } else {
          this.$message({
            duration: 900,
            message: "更新失败",
            type: 'error',
            onClose: () => {
              //此处写提示关闭后需要执行的函数
              location.reload()
            }
          })
          }
      })
    },
    Refuse(res) {
      // console.log(res)
      this.$axios.post('/update', {
        content: res.Content,
        imgpath: res.Imgpath,
        info: res.Info,
        infoid: res.Infoid,
        label: res.Label,
        sid: res.Sid,
        type: "refuse"
      }).then(res => {
        if (res.data.msg) {
          this.$message({
            duration: 900,
            message: "更新成功",
            type: 'success',
            onClose: () => {
              //此处写提示关闭后需要执行的函数
              location.reload()
            }
          })
        } else {
          this.$message({
            duration: 900,
            message: "更新失败",
            type: 'error',
            onClose: () => {
              //此处写提示关闭后需要执行的函数
              location.reload()
            }
          })
        }
      })
    }
  }
}
</script>

<style>
.scrollbar-wrapper{
  width: 100%;
  height: 100%;
  margin: top 100px; ;
}
</style>