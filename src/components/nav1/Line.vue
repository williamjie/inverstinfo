<template>
    <section class="chart">
        <el-row>
            <el-col :span="24" class="toolbar">
                    <el-button @click="fetchData">刷新数据</el-button>
            </el-col>

            <el-col :span="12">
                <div id="chartLine" style="width:1200px; height:400px;"></div>
            </el-col>

        </el-row>
    </section>
</template>

<script>
    import echarts from 'echarts'
    import $ from 'jquery'
    import jQuery from 'jquery'
    Date.prototype.Format = function (fmt) { //author: meizz 
        var o = {
            "M+": this.getMonth() + 1, //月份 
            "d+": this.getDate(), //日 
            "h+": this.getHours(), //小时 
            "m+": this.getMinutes(), //分 
            "s+": this.getSeconds(), //秒 
            "q+": Math.floor((this.getMonth() + 3) / 3), //季度 
            "S": this.getMilliseconds() //毫秒 
        };
        if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
        for (var k in o)
        if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
        return fmt;
    }

    var base = +new Date(2014, 9, 3);
    var dataAll = [Math.random() * 150];
    var dataSucced =  [Math.random() * 150];
    var date=[];
    var now = new Date().Format("yyyy-MM-dd hh:mm:ss");
    date.push(now)
    export default{
        data (){
            return {
                chartLine:null
            }
        },
        created : function(){
            console.log("create")
        },
        mounted:function(){
            var _this=this;
            //基于准备好的dom，初始化echarts实例
            this.chartLine = echarts.init(document.getElementById('chartLine'));
            
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
                    console.log("line success");
                    dataAll.push(Math.random() * 150);
                    dataSucced.push(Math.random()*110);
                    var now = new Date().Format("yyyy-MM-dd hh:mm:ss");
                    date.push(now)
                    _this.chartLine.setOption({
                        title: {
                            text: 'Line Chart'
                        },
                        tooltip: {
                            trigger: 'axis'
                        },
                        legend: {
                            data:['发送总量','发送成功数量']
                        },
                        grid: {
                            left: '3%',
                            right: '4%',
                            bottom: '3%',
                            containLabel: true
                        },
                        xAxis: {
                            type: 'category',
                            boundaryGap: false,
                            data:date
                            //data: ['周一','周二','周三','周四','周五','周六','周日']
                        },
                        yAxis: {
                            type: 'value'
                        },
                        series: [
                            {
                                name:'发送总量',
                                type:'line',
                                stack: '总量',
                                data:dataAll
                                ///data:[220, 182, 191, 234, 290, 330, 310]
                            },
                            {
                                name:'发送成功数量',
                                type:'line',
                                stack: '成功量',
                                data:dataSucced
                                //data:[120, 132, 101, 134, 90, 230, 210]

                            }
                        ]
                    });
                },
                error: function () {
                    console.log("error");
                }
            });

        },
        methods: {
            fetchData:function(){
                console.log("line fetchData")
                dataAll.push(Math.random() * 150);
                dataSucced.push(Math.random()*110);
                var now = new Date().Format("yyyy-MM-dd hh:mm:ss");
                date.push(now)
                this.chartLine.setOption({
                        title: {
                            text: 'Line Chart'
                        },
                        tooltip: {
                            trigger: 'axis'
                        },
                        legend: {
                            data:['发送总量','发送成功数量']
                        },
                        grid: {
                            left: '3%',
                            right: '4%',
                            bottom: '3%',
                            containLabel: true
                        },
                        xAxis: {
                            type: 'category',
                            boundaryGap: false,
                            data:date
                            //data: ['周一','周二','周三','周四','周五','周六','周日']
                        },
                        yAxis: {
                            type: 'value'
                        },
                        series: [
                            {
                                name:'发送总量',
                                type:'line',
                                stack: '总量',
                                data:dataAll
                                ///data:[220, 182, 191, 234, 290, 330, 310]
                            },
                            {
                                name:'发送成功数量',
                                type:'line',
                                stack: '成功量',
                                data:dataSucced
                                //data:[120, 132, 101, 134, 90, 230, 210]

                            }
                        ]
                    });

            }

        }

    }
</script>

<style scoped>
    .chart {
        width: 100%;
        float: left;
    }
    /*.chart div {
        height: 400px;
        float: left;
    }*/
    .el-col {
        padding: 30px 20px;
    }
</style>