<template>
    <section class="chart">
        <el-row>
        <!--
            <el-col :span="12">
                <div id="chartColumn" style="width:100%; height:400px;"></div>
            </el-col>
            <el-col :span="12">
                <div id="chartBar" style="width:100%; height:400px;"></div>
            </el-col>
            -->
            <el-col :span="12">
                <div id="chartLine" style="width:200%; height:600px;"></div>
            </el-col>
            <!--
            <el-col :span="12">
                <div id="chartPie" style="width:100%; height:400px;"></div>
            </el-col>
            -->
        </el-row>
    </section>
</template>

<script>
    import echarts from 'echarts'
    
    export default{
        data (){
            return {
                chartColumn:null,
                chartBar:null,
                chartLine:null,
                chartPie:null
            }
        },
        mounted:function(){
            var _this=this;
            //基于准备好的dom，初始化echarts实例
            //this.chartColumn = echarts.init(document.getElementById('chartColumn'));
            //this.chartBar = echarts.init(document.getElementById('chartBar'));
            this.chartLine = echarts.init(document.getElementById('chartLine'));
            //this.chartPie = echarts.init(document.getElementById('chartPie'));
            this.chartLine.setOption({
                title: {
                    text: 'Line Chart(单位：亿)'
                },
                tooltip: {
                    trigger: 'axis'
                },
                legend: {
                    data:['营收总收入','营收总收入同比增长','扣非净利润','扣非净利润同步增长']
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
                    data: ['2006','2007','2008','2009','2010','2011','2012','2013','2014','2015']
                },
                yAxis: {
                    type: 'value'
                },
                series: [
                    {
                        name:'营收总收入',
                        type:'line',
                        stack: '总量',
                        data:[0, 0, 0, 44.8, 55.2, 60.9, 70.7,84.0,98.2,113]
                    },
                    {
                        name:'营收总收入同比增长',
                        type:'line',
                        stack: '总量',
                        data:[0, 0, 0, 0, 23.3, 10.43, 16.07,18.84,16.85,15.05]
                    },
                    {
                        name:'扣非净利润',
                        type:'line',
                        stack: '总量',
                        data:[0, 0, 0, 6.82, 6.41, 9.53, 12.0,15.5,20.1,24.4]
                    },
                    {
                        name:'扣非净利润同步增长',
                        type:'line',
                        stack: '总量',
                        data:[0, 0, 0, 0, -6.57, 48.74, 26.17,28.58,29.86,21.49]
                    }
                ]
            });            
            /*
            this.chartColumn.setOption({
                title: { text: 'Column Chart' },
                tooltip: {},
                xAxis: {
                    data: ["衬衫","羊毛衫","雪纺衫","裤子","高跟鞋","袜子"]
                },
                yAxis: {},
                series: [{
                    name: '销量',
                    type: 'bar',
                    data: [5, 20, 36, 10, 10, 20]
                }]
            });
            
            this.chartBar.setOption({
                title: {
                    text: 'Bar Chart',
                    subtext: '数据来自网络'
                },
                tooltip: {
                    trigger: 'axis',
                    axisPointer: {
                        type: 'shadow'
                    }
                },
                legend: {
                    data: ['2011年', '2012年']
                },
                grid: {
                    left: '3%',
                    right: '4%',
                    bottom: '3%',
                    containLabel: true
                },
                xAxis: {
                    type: 'value',
                    boundaryGap: [0, 0.01]
                },
                yAxis: {
                    type: 'category',
                    data: ['巴西','印尼','美国','印度','中国','世界人口(万)']
                },
                series: [
                    {
                        name: '2011年',
                        type: 'bar',
                        data: [18203, 23489, 29034, 104970, 131744, 630230]
                    },
                    {
                        name: '2012年',
                        type: 'bar',
                        data: [19325, 23438, 31000, 121594, 134141, 681807]
                    }
                ]
            });



            this.chartPie.setOption({
                title : {
                    text: 'Pie Chart',
                    subtext: '纯属虚构',
                    x:'center'
                },
                tooltip : {
                    trigger: 'item',
                    formatter: "{a} <br/>{b} : {c} ({d}%)"
                },
                legend: {
                    orient: 'vertical',
                    left: 'left',
                    data: ['直接访问','邮件营销','联盟广告','视频广告','搜索引擎']
                },
                series : [
                    {
                        name: '访问来源',
                        type: 'pie',
                        radius : '55%',
                        center: ['50%', '60%'],
                        data:[
                            {value:335, name:'直接访问'},
                            {value:310, name:'邮件营销'},
                            {value:234, name:'联盟广告'},
                            {value:135, name:'视频广告'},
                            {value:1548, name:'搜索引擎'}
                        ],
                        itemStyle: {
                            emphasis: {
                                shadowBlur: 10,
                                shadowOffsetX: 0,
                                shadowColor: 'rgba(0, 0, 0, 0.5)'
                            }
                        }
                    }
                ]
            });

                        */
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