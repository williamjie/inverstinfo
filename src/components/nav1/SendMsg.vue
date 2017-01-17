<template>
	<section>
		<div id="logSections"></div>
		<!--工具条-->
		<el-col :span="24" class="toolbar">
			<el-form :inline="true" :model="formInline" class="demo-form-inline">
				<el-form-item>
					<el-input v-model="formInline.user" placeholder="电话号码"></el-input>
				</el-form-item>
				<el-form-item>
					<el-button @click="handleQuery">查询</el-button>
				</el-form-item>
				<el-form-item>
					<el-button @click="fetchData">刷新数据</el-button>
				</el-form-item>
				<!--
				<el-form-item>
					<el-button @click="handleAdd">新增</el-button>
				</el-form-item>
				-->
			</el-form>
		</el-col>

		<!--列表-->
		<template>
			<el-table :data="tableData" highlight-current-row v-loading="listLoading" style="width: 100%;">
				<el-table-column type="index" width="80" >
				</el-table-column>
				<el-table-column prop="name" label="电话号码" width="180" sortable>
				</el-table-column>
				<el-table-column prop="sex" label="部门" width="100" :formatter="formatSex" sortable>
				</el-table-column>
				<el-table-column prop="age" label="发送情况" width="100" sortable>
				</el-table-column>
				<el-table-column prop="birth" label="原因" width="180" sortable>
				</el-table-column>
				<el-table-column prop="addr" label="发送时间" sortable>
				</el-table-column>
				<!--
				<el-table-column inline-template :context="_self" label="操作" width="100">
					<span>
						<el-button type="text" size="small" @click="handleEdit(row)">编辑</el-button>
						<el-button type="text" size="small" @click="handleDel(row)">删除</el-button>
					</span>
				</el-table-column>
				-->
			</el-table>
		</template>

		<!--分页-->
		<el-col :span="24" class="toolbar" style="padding-bottom:10px;">
			<el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
			:current-page="currentPage"
			:page-size="100"
	        layout="prev, pager, next, jumper"
	        :total="10000">
   		 </el-pagination>

		<!--
		<el-col :span="24" class="toolbar" style="padding-bottom:10px;">
			<el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
			:current-page="currentPage"
			:page-size="100"
	        layout="total, sizes, prev, pager, next, jumper"
	        :total="1000">
   		 </el-pagination>
   		 -->
			<!--
			<el-pagination v-bind:current-Page="pageIndex" v-bind:page-size="pageSize" :total="total"
                               layout="total,sizes,prev,pager,next,jumper" v-bind:page-sizes="pageSizes"
                               v-on:size-change="sizeChange" v-on:current-change="pageIndexChange">

                </el-pagination>
               -->
			<!--
			<el-pagination :current-page="1" :page-sizes="[100, 200, 300, 400]" :page-size="100" layout="total, sizes, prev, pager, next, jumper"
				:total="400" style="float:right">
			</el-pagination>
			-->
		</el-col>

		<!--编辑界面-->
		<el-dialog :title="editFormTtile" v-model="editFormVisible" :close-on-click-modal="false">
			<el-form :model="editForm" label-width="80px" :rules="editFormRules" ref="editForm">
				<el-form-item label="姓名" prop="name">
					<el-input v-model="editForm.name" auto-complete="off"></el-input>
				</el-form-item>
				<el-form-item label="性别">
					<!--<el-select v-model="editForm.sex" placeholder="请选择性别">
						<el-option label="男" :value="1"></el-option>
						<el-option label="女" :value="0"></el-option>
					</el-select>-->
					<el-radio-group v-model="editForm.sex">
						<el-radio class="radio" :label="1">男</el-radio>
						<el-radio class="radio" :label="0">女</el-radio>
					</el-radio-group>
				</el-form-item>
				<el-form-item label="年龄">
					<el-input-number v-model="editForm.age" :min="0" :max="200"></el-input-number>
				</el-form-item>
				<el-form-item label="生日">
					<el-date-picker type="date" placeholder="选择日期" v-model="editForm.birth"></el-date-picker>
				</el-form-item>
				<el-form-item label="地址">
					<el-input type="textarea" v-model="editForm.addr"></el-input>
				</el-form-item>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button @click.native="editFormVisible = false">取 消</el-button>
				<el-button type="primary" @click.native="editSubmit" :loading="editLoading">{{btnEditText}}</el-button>
			</div>
		</el-dialog>
	</section>
</template>

<script>
	import util from '../../common/util'
	import NProgress from 'nprogress'
	import $ from 'jquery'
	import jQuery from 'jquery'
	var myObj = {};
	var pageLogNum = 20;

  export default {
    data() {
      return {
				formInline: {
					user: ''
				},
				currentPage: 1,
				pickerOptions0: {
					disabledDate(time) {
						return time.getTime() < Date.now() - 8.64e7;
					}
				},
				value1:'',
				editFormVisible:false,//编辑界面显是否显示
				editFormTtile:'编辑',//编辑界面标题
				//编辑界面数据
				editForm: {
					id:0,
					name: '',
					sex: -1,
					age: 0,
					birth: '',
					addr: ''
				},
				editLoading:false,
				btnEditText:'提 交',
				editFormRules:{
					name:[
						{ required: true, message: '请输入姓名', trigger: 'blur' }
					]
				},
				tableData: [{
					id:1000,
					name: '13113636636',
					sex: '广发通',
					age: 'success',
					birth:'',
					addr:'2016-12-23-14:23:22:00'
				}, {
					id:1001,
					name: '18575698760',
					sex: '运维部门',
					age: 'failed',
					birth:'发送频率超过每小时限制',
					addr:'2016-12-10-14:23:22:00'
				}				],
				listLoading:false
     		}
    },
    methods: {
			//性别显示转换
			formatSex:function(row,column){
				console.log("dev_env formatSex")
				return row.sex
				//return row.sex==1?'男':row.sex==0?'女':'未知';
			},
			/*
			//删除记录
			handleDel:function(row){
				console.log("dev_env row:",row);
				var _this=this;
				this.$confirm('确认删除该记录吗?', '提示', {
					//type: 'warning'
				}).then(() => {
					_this.listLoading=true;
					NProgress.start();
					setTimeout(function(){
						for(var i=0;i<_this.tableData.length;i++){
							if(_this.tableData[i].id==row.id){
								_this.tableData.splice(i,1);

								_this.listLoading=false;
								NProgress.done();
								_this.$notify({
									title: '成功',
									message: '删除成功',
									type: 'success'
								});

								break;
							}
						}
					},1000);
				}).catch(() => {
				});
			},
			//显示编辑界面
			handleEdit:function(row){
				console.log("dev_env handleEdit row:",row)
				this.editFormVisible=true;
				this.editFormTtile='编辑';
				this.editForm.id=row.id;
				this.editForm.name=row.name;
				this.editForm.sex=row.sex;
				this.editForm.age=row.age;
				this.editForm.birth=row.birth;
				this.editForm.addr=row.addr;
			},
			//编辑 or 新增
			editSubmit:function(){
				var _this=this;
				console.log("dev_env editSubmit",)

				_this.$refs.editForm.validate((valid)=>{
					if(valid){
						_this.$confirm('确认提交吗？','提示',{}).then(()=>{
							_this.editLoading=true;
							NProgress.start();
							_this.btnEditText='提交中';
							setTimeout(function(){
								_this.editLoading=false;
								NProgress.done();
								_this.btnEditText='提 交';
								_this.$notify({
									title: '成功',
									message: '提交成功',
									type: 'success'
								});
								_this.editFormVisible = false;

								if(_this.editForm.id==0){
									//新增
									_this.tableData.push({
										id:new Date().getTime(),
										name: _this.editForm.name,
										sex: _this.editForm.sex,
										age: _this.editForm.age,
										birth: _this.editForm.birth==''?'':util.formatDate.format(new Date(_this.editForm.birth),'yyyy-MM-dd'),
										addr: _this.editForm.addr
									});
								}else{
									//编辑
									for(var i=0;i<_this.tableData.length;i++){
										if(_this.tableData[i].id==_this.editForm.id){
											_this.tableData[i].name=_this.editForm.name;
											_this.tableData[i].sex=_this.editForm.sex;
											_this.tableData[i].age=_this.editForm.age;
											_this.tableData[i].birth=_this.editForm.birth==''?'':util.formatDate.format(new Date(_this.editForm.birth),'yyyy-MM-dd');
											_this.tableData[i].addr=_this.editForm.addr;
											break;
										}
									}
								}
							},1000);
						
						});

					}
				});

			},
			//显示新增界面
			handleAdd:function(){
				var _this=this;
				console.log("dev_env handleAdd ")
				this.editFormVisible=true;
				this.editFormTtile='新增';
				
				this.editForm.id=0;
				this.editForm.name='';
				this.editForm.sex=1;
				this.editForm.age=0;
				this.editForm.birth='';
				this.editForm.addr='';
			},*/
			//查询发送记录
			handleQuery:function(){
				var _this=this;
				if (this.formInline.user == ""){
					alert("请输入电话号码")
				}
				console.log("dev_env phone num:",this.formInline.user)

				this.searchData(_this.formInline.user)
				console.log("dev_env handleQuery:",_this.formInline.user)
			},
			sizeChange: function (pageSize) {
	            this.pageSize = pageSize;
	            //this.fetchData();
	           	console.log("sizeChange pageSize:",pageSize)

	        },
	        pageIndexChange: function (pageIndex) {
	            this.pageIndex = pageIndex;
	            //this.fetchData();
	           	console.log("pageIndexChange pageIndex:",pageIndex)
	        },
	        handleSizeChange(val) {  //页码
	        	console.log("handleSizeChange val:",val)
        		console.log(`${val} items per page handleSizeChange dev_env`);
      		},
      		handleCurrentChange(val) { //选择页码
       			 var _this= this
       			 _this.currentPage = val;
       			 _this.clearRowData(_this.tableData.length);
       			 var start=1;
       			 var end=pageLogNum;
       			 if (val != 1)
       			 {
       			 	start = val*pageLogNum;
       			 	end = (val+1)*pageLogNum;
       			 }

		        for (var j=start;j<=end;j++)
		        {
		        	var dateinfo = $.data(myObj,j.toString()); 
		        	console.log("dateinfo:",dateinfo.phone)
		        	_this.tableData.push({
		        		id:new Date().getTime(),
		        		name:dateinfo.phone,
		        		sex:dateinfo.dept,
		        		age:dateinfo.sendinfo,
		        		birth:dateinfo.reason,
						addr:dateinfo.date,
		        	});
		        }
        		 console.log(`current page: ${val} handleCurrentChange dev_env`);
      		},
      		searchData:function(id_info){//clear 
      			var _this = this
				var count =0 //the number search
				for(var i=0;i<=_this.tableData.length;){
					if (_this.tableData.length == 0 || _this.tableData.length == count){
						break;
					} 

					if(_this.tableData[i].name != id_info){
						console.log("splice")
						_this.tableData.splice(i,1);
						i=0  //delete head always
					}else{
						count++
						i++
					}
				}
      		},
      		clearRowData: function(id_length){
      			var _this = this;
				for (var index=0;index<id_length;index++)
				{
					_this.tableData.splice(0,1);
				}
      		},
		    fetchData: function () {
		    	var _this=this;
		    	console.log("dev_env fetchData xxx")
		    	//时间查找，电话号码查找等条件
				$.ajax({
				    url: 'http://localhost:12345/api/GetInfo',
				    headers: {
				        'Content-Type': 'application/json'
				    },
				    type: "POST", /* or type:"GET" or type:"PUT" */
				    dataType: "json",
				    data: {phone:"13113636636"
				    },
				    success: function (result) {
				    	//var el = document.getElementById('logSections')
				    	//_this.tableData.clear()
				    	var data = new Date().getTime()
				    	for(var i=1;i<1000;i++)
				    	{
				        	//console.log("lengh:",i)
							$.data(myObj,i.toString(),{"phone":"13113636636","dept":"电商部门","sendinfo":"成功","reason":i.toString(),"date":"2016-12-10-14:23:22:00"});
							//$.data(myObj,'2','{phone:"18588886666",dept:"电商部门",sendinfo:"成功",reason:"频率限制",date:"2016-12-10-14:23:22:00"}');
						}
						$.data(myObj,"length",1000)
						_this.handleSizeChange(100);
				    	console.log("data:",$.data(myObj,'1'))
				        console.log("data:",$.data(myObj,'2'))
				        console.log("data:",$.data(myObj,'3'))
				        console.log("data:",$.data(myObj,'4'))
				        console.log("data:",$.data(myObj,'4').date)

				        console.log("result:",result);
						//clear table
						_this.clearRowData(_this.tableData.length)
				        for (var j=1;j<=pageLogNum;j++)
				        {
				        	var dateinfo = $.data(myObj,j.toString()); 
				        	console.log("dateinfo:",dateinfo.phone)
				        	_this.tableData.push({
				        		id:new Date().getTime(),
				        		name:dateinfo.phone,
				        		sex:dateinfo.dept,
				        		age:dateinfo.sendinfo,
				        		birth:dateinfo.reason,
								addr:dateinfo.date,
				        	});
				        }
				        console.log("fetchData end")
				        /*
						//新增
						_this.tableData.push({
							id:new Date().getTime(),
							name: "13388886666",
							sex: "电商",
							age: "success",
							birth: "good",
							addr: new Date().getTime()
						});
						*/
				    },
				    error: function () {
				        console.log("error");
				    }
				});
	       	},
	       	ResizeForm: function () {
		    	var _this=this;
		    	console.log("dev_env ResizeForm xxx")
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