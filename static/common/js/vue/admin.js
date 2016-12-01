import Vue from 'vue/dist/vue'
import VueAjax from 'vue-resource/dist/vue-resource'
import VueI18n from 'vue-i18n/dist/vue-i18n'
// import Validator from 'vue-validator/dist/vue-validator'
import input from './admin/form/input'
// import echarts from './admin/echarts'

var locales = {
  cn: {
    message: {
      username: '用户名',
      password: '密码'
    }
  },
  en: {
    message: {
      username: 'username',
      password: 'password',
    }
  }
}
Vue.use(VueAjax)
Vue.use(VueI18n)
Vue.config.lang = 'cn';
Object.keys(locales).forEach(function (lang) {
  Vue.locale(lang, locales[lang])
})
// Vue.use(Validator)
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


