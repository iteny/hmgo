<style scoped>
input{font-family: inherit;font-size: inherit;}
.hm_input{
	display: inline-block;
    vertical-align: top;
    margin: 10px 5px;
    position: relative;

    /*width: 180px;*/
}
.hm_input_min{font-size: 14px;}
.hm_input_small{font-size: 16px;}
.hm_input_big{font-size: 18px;}
.hm_input_max{font-size: 24px;}
.hm_input_inner{display: block;padding: 3px 10px;box-sizing: border-box;width: 100%;
    height: 36px;color: #1f2d3d;background-color: #fff;background-image: none;
    border: 1px solid #c0ccda;border-radius: 4px;transition: border-color .2s cubic-bezier(.645,.045,.355,1);-webkit-transition: border-color .2s cubic-bezier(.645,.045,.355,1);outline: none;line-height: normal;}
.hm_input_inner:focus{outline: none; border-color: #20a0ff;}
.hm_input_min .hm_input_inner{height:36px;font-size: 14px;}
.hm_input_small .hm_input_inner{height:40px;font-size: 20px;}
.hm_input_big .hm_input_inner{height: 44px;font-size: 24px;}
.hm_input_max .hm_input_inner{height: 48px;font-size:28px;}
.hm_input i.fa{
	position: absolute;
	/*top:25px;*/
    width: 35px;
    height: 100%;
    text-align: center;
    color: #c0ccda;}

.hm_input i.fa:after {
    content: '';
    height: 100%;
    width: 0;
    display: inline-block;
    vertical-align: middle;
}
.hm_input i.left{left: 0;}
.hm_input i.right{right: 0;}

</style>
<template>
	<div :class="[
		type === 'textarea' ? 'hm_textarea' : 'hm_input',
		size ? 'hm_input_' + size : '',
		{
	      'is_disabled': disabled,
	      'hm_input_group': $slots.prepend || $slots.append
	    }
	]">
		<template v-if="type !== 'textarea'">
			<!-- input 图标 -->
      		<i class="fa" :class="[icon ? 'fa-' + icon : '']" v-if="icon"></i>
      		<!-- <i class="el-input__icon el-icon-loading" v-if="validating"></i> -->
			<input 
				v-if="type === 'text'"
				class="hm_input_inner"
				v-model="inputValue"
				type="text"
				:width="width"
				:name="name"
				:disabled="disabled"
				:maxlength="maxlength"
        		:minlength="minlength"

			>
			<!-- input 右侧图标 -->
      		<!-- <i class="el-input__icon" :class="[icon ? 'el-icon-' + icon : '']" v-if="icon"></i>
      		<i class="el-input__icon el-icon-loading" v-if="validating"></i> -->
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
	    icon: {//图标
	        type: String,
	        default: ''
	    },
	    size: {//长度大小
        	type: String,
        	default: ''
      	},
      	width: {//长度大小
        	type: String,
        	default: '200px'
      	},
      	name: {//form提交用到的name
	        type: String,
	        default: ''
	    },//默认为不禁用
	    disabled: {
	        type: Boolean,
	        default: false
	    },
	    maxlength: Number,
      	minlength: Number

	},
	data(){
		return{
			inputValue:this.value
			
		}
	}
}
</script>