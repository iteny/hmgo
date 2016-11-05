import Vue from 'vue/dist/vue'
import VueAjax from 'vue-resource/dist/vue-resource'
import Validator from 'vue-validator/dist/vue-validator'
import input from './admin/form/input'
// import echarts from './admin/echarts'

Vue.use(VueAjax)
Vue.use(Validator)
Vue.config.devtools = true
const app = new Vue({
	delimiters: ['<{', '}>'],
    el: '#baby',    
    components: { 
    	'hm-input':input,
    	// 'hm-echarts':echarts 
    },
    data:{
    	items:[
    		{name:'安全手套',value:100},
    		{name:'安全套',value:200},
    		{name:'手机',value:300}
    	],
    	isA:true,
    	isB:true
    },
    computed: {	    
	    total(){
	    	self = this;
	    	var tmp=0
	    	for(var i=0;i<self.items.length;++i){
	    		tmp = tmp+parseInt(self.items[i].value);
	    	}
	    	return tmp;
	    }
	},
	methods: {
		
	}
})


