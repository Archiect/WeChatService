<template>
    <div class="mainBodyOfHome">
<!--    数据展示-->
        <div class="dataDisplayOfHome">
<!--          标题  -->
            <div class="headOfDataDisplay">
                <span style="flex: 9;font-size: 30px;font-weight: bold">
                Demo
                </span>
                <a style="font-size: 16px; cursor: pointer;" class="exit" @click="exit">退出</a>
            </div>
        </div>
        <!-- 新增学号 -->
      <div class="dataTableOfHome" scope="scope">
        <h2>查询学生</h2>
        <el-input placeholder="请输入内容" v-model="input" class="input-with-select">
          <el-select v-model="select" slot="prepend" placeholder="请选择">
            <el-option label="按学号查询" value="学号"></el-option>
            <el-option label="按名字查询" value="名字"></el-option>
          </el-select>
          <el-button slot="append" icon="el-icon-search" @click="StuFind">
          </el-button>
        </el-input>
      </div>
      <!-- <div class="dataTableOfHome">
      <h2>查询学生</h2>
      <el-input placeholder="请输入内容" v-model="input" class="input-with-select">
        <el-select v-model="select" slot="prepend" placeholder="请选择">
          <el-option label="按学号查询" value="学号"></el-option>
          <el-option label="按名字查询" value="名字"></el-option>
        </el-select>
        <el-button slot="append" icon="el-icon-search" @click="StuFind"></el-button>
      </el-input>
      </div> -->
      <!--   批量-->
      <div class="visualizationOfHome">
        <h2>新增学生表</h2>
        <el-tag style="font-size: 20px;margin-top: 20px; margin-bottom: 18px;">单独</el-tag>
        <div>
          <el-input v-model="stuname" placeholder="请输入姓名" style="width: 100px;"></el-input>
          <el-input v-model="stusid" placeholder="请输入学号" style="width: 100px; margin-left: 5px;"></el-input>
          <el-dropdown style="margin-left: 5px;" @command="selectFaculty($event)">
            <el-button type="primary" ref="faculty">
              请选择院系<i class="el-icon-arrow-down el-icon--right"></i>
            </el-button>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item command="信息工程系">信息工程系</el-dropdown-item>
              <el-dropdown-item command="电气工程系">电气工程系</el-dropdown-item>
              <el-dropdown-item command="车辆工程系">车辆工程系</el-dropdown-item>
              <el-dropdown-item command="财贸管理系">财贸管理系</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>

          <el-dropdown style="margin-left: 5px;" @command="selectGrade($event)">
            <el-button type="primary" ref="grade">
              请选择年级<i class="el-icon-arrow-down el-icon--right"></i>
            </el-button>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item command="20级">20级</el-dropdown-item>
              <el-dropdown-item command="21级">21级</el-dropdown-item>
              <el-dropdown-item command="22级">22级</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>

          <el-dropdown style="margin-left: 5px;" @command="selectClass($event)">
            <el-button type="primary" ref="class">
              请选择班级<i class="el-icon-arrow-down el-icon--right"></i>
            </el-button>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item command="一班">一班</el-dropdown-item>
              <el-dropdown-item command="二班">二班</el-dropdown-item>
              <el-dropdown-item command="三班">三班</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
            <el-button type="success" icon="el-icon-check" circle style="margin-left: 5px;"  @click="submit($event)"></el-button>
        </div>
        <!-- <el-tag style="font-size: 18px;margin-top: 20px;">批量</el-tag>
        <div></div> -->
      </div>
    </div>
</template>

<script>
    export default {
        name: "Home",
        data(){
            return {
                // inforOfDaPeng:{},
              input: '',
              select: '',
              stuname: '',
              stusid: '',
              // faculty: '',
              faculty_: '',
              // grade: '',
              grade_: '',
              // class: '',
              class_: ''
            }
        },
        methods: {
          exit() {
            this.$router.push("/login")
          },
          StuFind() {
            // console.log(this.input,this.select)
            if (this.input == ''||this.select == '') {
              this.$message({
                duration: 900,
                message: "请输入正确的查询内容",
                type: 'error',
              })
            } else {
              this.$axios.post('/Find', {
                input: this.input,
                select: this.select
              }).then(res => {
                if (!res.data.msg) {
                  this.$message({
                    duration: 900,
                    message: "未查询到该学生",
                    type: 'error',
                  })
                } else {
                  console.log(res)
                  this.$alert(res.data.msg.message, '查询到以下内容', {
                    confirmButtonText: '确定',
                    // callback: action => {
                    //   this.$message({
                    //     // type: 'info',
                    //     // message: `action: ${action}`
                    //   });
                    // }
                  });
                }
              })
            }
          },
          selectFaculty(res) {
            this.faculty_ = res
            this.faculty = res
            var faculty = this.faculty
            this.faculty = this.$refs.faculty.$el.innerText;
            this.$refs.faculty.$el.innerText = faculty
          },
            selectGrade(res) {
              // alert(res)
              this.grade_ = res
              this.grade = res
              var grade = this.grade
              this.grade = this.$refs.grade.$el.innerText;
              this.$refs.grade.$el.innerText = grade
          },
          selectClass(res) {
            this.class_ = res
            this.class = res
            var class_ = this.class
            this.class = this.$refs.class.$el.innerText;
            this.$refs.class.$el.innerText = class_
          },
          submit() {
          //   stuname: '',
          //     stusid: '',
          //       faculty: '',
          //         grade: '',
          //           class: ''
            // console.log(this.stuname,this.stusid,this.faculty_,this.grade_,this.class_)

            // 请求后端
            this.$axios.post('/InsertNewStu', {
              stuname: this.stuname,
              stusid: this.stusid,
              faculty: this.faculty_,
              grade: this.grade_,
              class:this.class_
            }).then(res => {
              // console.log(res)
              if (res.status == 200) {
                this.$message({
                  duration: 900,
                  message: "新增成功",
                  type: 'success',
                })
              }
              if (res.status==0) {
                this.$message({
                  duration: 900,
                  message: "新增失败",
                  type: 'error',
                })
              }
            })
          }
        },
    }
</script>

<style scoped>
.el-select .el-input {
  width: 10px;
}

.input-with-select .el-input-group__prepend {
  background-color: #fff;
  width: 10px;
}
.Sid-Box{
    margin-top: 20px;
    margin-left: 20px;
    font-size: 20px;
    color: white;
  }
.Sid-Btn{
    font-size: 10px;
    color: black;
  }
.exit:hover{
    color: red;
  }
    .mainBodyOfHome{
        display: flex;
        width: 100%;
        height: 100%;
        /* background-image: radial-gradient(circle farthest-side at 10% 90%, #FFE8EA, #EDF3FF 70%, #EDF2FB); */
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }
    .dataTableOfHome {
      flex: 5;
      height: 100%;
      width: 70%;
      /* background-color: red; */
    }
    .dataDisplayOfHome{
        flex: 3;
        height: 100%;
        width: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: 0 50px;
    }
    .visualizationOfHome{
        flex: 13;
        height: 100%;
        width: 70%;
        /* background-color: pink; */
    }
    .headOfDataDisplay{
        flex: 2;
        width: 100%;
        height: 100%;
        display: flex;
        margin-top: 20px;
        align-items:flex-start;
    }
</style>