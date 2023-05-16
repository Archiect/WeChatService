<style>
.box-card{
  height: 200px;
  width: 250px;
  
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 0px;
  margin-top: 10px;
}
.scrollbar-wrapper{
  width: 100%;
  height: 100%;
  margin: top 100px;
  /* overflow-x: hidden; */
}
</style>
<template>
  <el-scrollbar class="scrollbar-wrapper">
    <el-table :data="message" stripe style="width: 100%; margin-top: 20px;">
      <el-table-column prop="Content" label="描述" width="250"> </el-table-column>
      <el-table-column prop="Label" label="类型" width="100"> </el-table-column>
      <el-table-column prop="Enjoy" label="参与人数" width="100"> </el-table-column>
      <el-table-column label="图片" prop="img" width="215">
        <template scope="scope">
          <!-- http://localhost:8080/img/ -->
          <el-image :src="('http://localhost:8080/img/'+scope.row.Imgpath)" style="border-radius: 5%;"></el-image>
        </template>
      </el-table-column>

      <el-table-column align="center" label="操作" width="215">
        <template scope="scope">
          <el-row>
            <el-button-group>
              <el-tooltip effect="light" content="取消发布" placement="top-start">
                <el-button icon="el-icon-close" type="danger" @click="Refuse(scope.row)" />
              </el-tooltip>
            </el-button-group>
          </el-row>
        </template>
      </el-table-column>
    </el-table>
  </el-scrollbar>
</template>

<script>
export default {
  name: "Weather",
  data() {
    return {
      message:[]
    }
  },
  created() {
    this.Loading();
  },
  methods: {
    Loading(){
      this.$axios.get('/Mmessage', {}).then(res => {
        // console.log(res.data.Minfo.message)
        this.message = res.data.Minfo.message
      })
    },
    Refuse(res) {
      console.log(res)
      this.$axios.post('/delet', {
        Content: res.Content,
        Type_: res.Label,
        Imgpath: res.Imgpath,
      }).then(res => {
        console.log(res)
        if (res.status==200) {
          this.$message({
            duration: 900,
            message: "删除成功",
            type: 'success',
            onClose: () => {
              //此处写提示关闭后需要执行的函数
              location.reload()
            }
          })
        } else {
          this.$message({
            duration: 900,
            message: "删除失败",
            type: 'error',
          })
        }
      })
    }
  }
}
</script>
