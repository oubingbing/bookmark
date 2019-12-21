	const config={
		dev:{//开发环境
			domain:"http://127.0.0.1:8080",//后台接口地址
		},
		prod:{//生产环境
			domain: "https://lianyan.kucaroom.com/api/wechat",
		}
	}

	const domain = config.dev.domain;
	//const domain = config.prod.domain;

	var normalBookmark = [
		{icon:"https://cdn.jsdelivr.net/npm/@bootcss/www.bootcss.com@0.0.3/dist/ico/apple-touch-icon-precomposed.png","title":"百度翻译"},
		{icon:"https://github.com/fluidicon.png","title":"github"},
		{icon:"https://cdn.jsdelivr.net/npm/@bootcss/www.bootcss.com@0.0.3/dist/ico/apple-touch-icon-precomposed.png","title":"新浪军事"},
		{icon:"https://cdn.jsdelivr.net/npm/@bootcss/www.bootcss.com@0.0.3/dist/ico/apple-touch-icon-precomposed.png","title":"github"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"百度翻译"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"github"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"百度翻译"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"github"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"百度翻译"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"github"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"百度翻译"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"github"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"百度翻译"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"github"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"百度翻译"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"github"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"百度翻译"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"github"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"百度翻译"},
		{icon:"https://fanyi-cdn.cdn.bcebos.com/static/translation/img/favicon/favicon-16x16_e1883cf.png","title":"github"}
	]
	var registerData = {
		nickname:"",
		email:"",
		password:"",
		confirm_password:""
	}
	
	var vm = new Vue({
		el: '#app',
	  	data: {
	  		title:"星虫",
	  		searchWorld:"",//搜索词,
	  		normalBookmarkList:normalBookmark,//常用书签
	  		showCreateForm:false,//是否显示书签表单
			showAuth:true,
            showRegisterDiv:false,
			registerForm:registerData
	  	},

	  	mounted: function () {
			console.log(domain)
        },

        methods:{
        	//跳转到搜索引擎
        	openSearchPage:function(){
        		var url = "https://www.baidu.com/s?&word="+this.searchWorld
        		this.searchWorld = ''
        		window.open(url,"_blank");
        	},

        	//显示或隐藏新建书签报表
        	ShowHiddenCreateForm:function(){
        		console.log("是否显示")
        		if (this.showCreateForm) {
        			this.showCreateForm = false
        		}else{
        			this.showCreateForm = true
        		}
        	},
			
			//登录
            login:function () {
				
            },
			
			//注册
            register:function () {
				var formData = this.registerForm

                axios.post(domain+"/auth/register?data=123",Qs.stringify(formData),{
                    headers : {
                        'Content-Type':'application/x-www-form-urlencoded;charset=UTF-8'
                    }
				}).then(function (res) {
                	console.log(res)
				}).catch(function (error) {
					console.log(error);
				});




			},

			//判断是否显示注册框
            showRegister:function () {
        		if (this.showRegisterDiv){
                    this.showRegisterDiv = false
                }else{
                    this.showRegisterDiv = true
                }
            }
        }
	})