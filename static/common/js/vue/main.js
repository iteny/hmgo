import Vue from 'vue/dist/vue'
import Vuex from 'vuex/dist/vuex'
import Favlist from './components/Favlist'
Vue.use(Vuex)
const store = new Vuex.Store({
	delimiters: ['<{', '}>'],
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
    components: { 
    	Favlist 
    },
    computed: {
    count () {
	    return store.state.count
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