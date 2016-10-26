import Vue from 'vue/dist/vue'
import VueAjax from 'vue-resource/dist/vue-resource'
import Vuex from 'vuex/dist/vuex'
import Favlist from './admin/favlist'
Vue.use(Vuex)
Vue.use(VueAjax)
Vue.config.devtools = true
const login = {
  state: { count:10111 },  
}
const store = new Vuex.Store({
	delimiters: ['<{', '}>'],
	modules:{
		login:login,
	},
  	state: {
    	count: 0
  	},
  	mutations: {
  		increment: state => state.count++,
    	decrement: state => state.count--
  	}
})
const app = new Vue({
	delimiters: ['<{', '}>'],
    el: '#baby',
    store,
    components: { 
    	Favlist 
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
	    count () {
		    return store.state.count
	    },
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
		increment () {
		  store.commit('increment')
		},
		decrement () {
			store.commit('decrement')
		}
	}
})

