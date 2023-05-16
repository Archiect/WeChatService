<template>

	<router-view/>

</template>

<script>

	export default {
		name:'App',
		data() {
			return {
			}
		},
	// created(){
	// 	this.$router.push({
	// 		name:'login'
	// 	})
	// }
	
	}
</script>

<style>
*{
    margin: 0;
    padding: 0;
  }
.el-scrollbar .el-scrollbar__wrap {
  overflow-x: hidden;
}

 /* 上面隐藏横向滚动条会导致下拉框底部遮盖问题 */
.el-select-dropdown .el-scrollbar {
  padding-bottom: 17px;
}
</style>
