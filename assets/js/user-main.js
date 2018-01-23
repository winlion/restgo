apiready = function(){
        api.parseTapmode();
}
//页面切换tab
var panel=["panel-project","panel-project","panel-todo","panel-wode"]
var tabfooter = new auiTab({
        element:document.getElementById("footer")
    },function(ret){
    	if(!panel[ret.index]){
    		ret.index = 0;
    	}
    	
    	for(var i in panel){
    		if(ret.index==i){
    			document.getElementById(panel[i]).classList.add("active")
    		}else{
    			document.getElementById(panel[i]).classList.remove("active")
    		}
    	}
});

//项目状态tab
var tabprj= new auiTab({
    element:document.getElementById("tab-prj"),
    repeatClick:true
},function(ret){
    
});

//添加待办事项
var tabtodo = new auiTab({
    element:document.getElementById("tabtodo"),
},function(ret){
    console.log(ret)
});
