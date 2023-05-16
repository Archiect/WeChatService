<style>
.scrollbar-wrapper{
    width: 100%;
    height: 100%;
    /* margin: top 100px; */
    /* overflow-x: hidden; */
}
.card-header {
  display:flex;
  justify-content: space-between;
  /* align-items: center;*/
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 0px;
  margin-top: 10px;
}

.box-card {
  width: 350px;
  margin-left: 15%;
}
</style>
<template>
  <el-scrollbar class="scrollbar-wrapper">
    <div style="margin: 10px; display: inline-block;" v-for="(item, index) in msg" :key="index">
      <el-card class="box-card" >
        <template #header>
          <div class="tool">
            <el-tooltip effect="light" content="通过" placement="top-start">
              <!-- showView可以获取当前表单的信息 -->
              <el-button icon="el-icon-check" type="success" @click="Allow(item, $event)" />
            </el-tooltip>
            <el-tooltip effect="light" content="不通过" placement="top-start">
              <el-button icon="el-icon-close" type="danger" @click="Refuse(item,$event)" />
            </el-tooltip>
          </div>
          <div class="card-header" ref="a">
            <span>预约时间：{{item.Time}} {{item.Date}}</span>
            &nbsp;
            <span>地址：{{ item.Floor}} {{item.DormNum}}</span>
          </div>
        </template>
        <!-- <div style="height: 120px"> -->
          <el-image :src="('http://localhost:8080/img/'+item.Imgpath)" ></el-image>
        <!-- </div> -->
      </el-card>
    </div>
  </el-scrollbar>
</template>
<script>
export default {
  name: "Stufix",
  data() {
    return {
      msg:[]
    }
  },
  created() {
    this.Loading();
  },
  methods: {
    Loading() {
      this.$axios.get('/stufix', {}).then(res => {
        // console.log(res)
        this.msg = res.data.msg.msg
      })
    },
    Allow(item) {
      // console.log(item.DormNum)
      this.$axios.post('/stuUpdate', {
        Date: item.Date,
        DormNum: item.DormNum,
        Floor: item.Floor,
        Imgpath: item.Imgpath,
        Time: item.Time,
        Type_:"allow"
      }).then(res => {
        // console.log(res)
        if (res.status == 200) {
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
    Refuse(item) {
      // console.log(item.DormNum)
      this.$axios.post('/stuUpdate', {
        Date: item.Date,
        DormNum: item.DormNum,
        Floor: item.Floor,
        Imgpath: item.Imgpath,
        Time: item.Time,
        Type_: "refuse"
      }).then(res => {
        // console.log(res)
        if (res.status == 200) {
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