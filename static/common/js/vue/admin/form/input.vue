<style scoped>
.hm_input{
	display: inline-block;
    vertical-align: top;
    margin: 10px 5px;
}
.hm_input_min{font-size: 14px;}
.hm_input_small{font-size: 16px;}
.hm_input_big{font-size: 18px;}
.hm_input_max{font-size: 24px;}
.hm_input_max .hm_input_inner{height: 42px;}
</style>
<template>
	<div :class="[
		type === 'textarea' ? 'hm_textarea' : 'hm_input',
		size ? 'hm_input_' + size : '',
	]">
		<template v-if="type !== 'textarea'">
			<input 
				v-if="type === 'text'"
				class="hm_input_inner"
				v-model="inputValue"
				type="text"
				:name="name"				
			>
		</template>
		<textarea v-else>
			
		</textarea>
	</div>
</template>
<script>
module.exports={
	props:{
		value: {
			type:String,
			default:''
		},
		type: {//类型text,password,textarea之类的
	        type: String,
	        default: 'text'
	    },
	    size: {//长度大小
        	type: String,
        	default: ''
      	},
      	name: {//form提交用到的name
	        type: String,
	        default: ''
	    },

	},
	data(){
		return{
			inputValue:this.value
			
		}
	}
}
</script>