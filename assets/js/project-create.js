var vue = new Vue({
	el: '#page',
	data:{
		project:{
			"name":"",
			"memo":"",
		}
	},
	methods:{
		createproject:function(){
			if(!this.project.name){
				xlutil.error("请输入项目名称");
				return ;
			}
			if(!this.project.name){
				xlutil.error("请输入项目描述");
				return ;
			}
			xlutil.post("project/create",this.project).then(function(r){
					console.log("project/create",r);
			})
		},
		test:function(){
			this.createproject()
		}
	},
	mounted:function(){
		mui.init()
	}
})
