<template>
	<section>
		<!--工具条-->
		<el-row class="demo-autocomplete">
		  <el-col :span="12">
		    <div class="sub-title">激活即列出输入建议</div>
		    <el-autocomplete
		      class="inline-input"
		      v-model="state1"
		      :fetch-suggestions="querySearch"
		      placeholder="请输入内容"
		      @select="handleSelect"
		    ></el-autocomplete>
		  </el-col>
		  <el-col :span="12">
		    <div class="sub-title">输入后匹配输入建议</div>
		    <el-autocomplete
		      class="inline-input"
		      v-model="state2"
		      :fetch-suggestions="querySearch"
		      placeholder="请输入内容"
		      :trigger-on-focus="false"
		      @select="handleSelect"
		    ></el-autocomplete>
		  </el-col>		
  		</el-row>

	</section>
</template>

<script>
	import util from '../../common/util'
	import NProgress from 'nprogress'

  export default {
    data() { 
    	return {
	        restaurants: [],
	        state1: '',
	        state2: ''
      	};
    },
    methods: {
			handleCommand:function(command){
				var _this = this
				console.log("handleCommand command:",command)
			},
	      	querySearch(queryString, cb) {
		        var restaurants = this.restaurants;
		        var results = queryString ? restaurants.filter(this.createFilter(queryString)) : restaurants;
		        // 调用 callback 返回建议列表的数据
		        cb(results);
	      	},
	      	createFilter(queryString) {
	       	  return (restaurant) => {
	          return (restaurant.value.indexOf(queryString.toLowerCase()) === 0);
	        };
	      	},
		    loadAll() {
		        return [
		          { "value": "三全鲜食（北新泾店）", "address": "长宁区新渔路144号" },
		          { "value": "Hot honey 首尔炸鸡（仙霞路）", "address": "上海市长宁区淞虹路661号" },
		        ];
		    },
		    handleSelect(item) {
		        console.log(item);
		    },
		   	mounted() {
		      	this.restaurants = this.loadAll();
		    },
			//查询发送记录
			handleQuery:function(){
				var _this=this;
				console.log("dev_env handleQuery:",_this.formInline.user)
			}
    }
  }
</script>

<style scoped>
	.toolbar .el-form-item {
		margin-bottom: 10px;
	}
	
	.toolbar {
		background: #fff;
		padding: 10px 10px 0px 10px;
	}
</style>