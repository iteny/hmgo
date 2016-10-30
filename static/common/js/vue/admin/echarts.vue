<style scoped>
.my-component h2 {
  color: red;
}
</style>

<template>
  <div style="width: 600px;height:400px;" ref="aa"><a @click="initChart">asdfasdf</a></div>

</template>

<script>

import echarts from 'echarts';
// 导出组件定义
module.exports = {
	// el:'#ss',  	
    data: function() {
        return {
            data: {name:'asdfas'}
        }
    },
    methods:{
    	initChart:function(){
    		 // 基于准备好的dom，初始化echarts实例
	        var myChart = echarts.init(this.$refs.aa);

	        // 指定图表的配置项和数据
	        var option = {
	            title: {
	                text: 'ECharts 入门示例'
	            },
	            tooltip: {},
	            legend: {
	                data:['销量']
	            },
	            xAxis: {
	                data: ["衬衫","羊毛衫","雪纺衫","裤子","高跟鞋","袜子"]
	            },
	            yAxis: {},
	            series: [{
	                name: '销量',
	                type: 'bar',
	                data: [5, 20, 36, 10, 10, 20]
	            }]
	        };

	        // 使用刚指定的配置项和数据显示图表。
	        myChart.setOption(option);
    	}
    	// alert(this.$data.items)
    	
    	
    },    
  	
}
</script>