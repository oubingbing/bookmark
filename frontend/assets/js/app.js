	
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
	
	var vm = new Vue({
		el: '#app',
	  	data: {
	  		title:"星虫",
	  		searchWorld:"",//搜索词,
	  		normalBookmarkList:normalBookmark//常用书签
	  	},

	  	mounted: function () {
			                  
        },

        methods:{

        	//跳转到搜索引擎
        	openSearchPage:function(){
        		var url = "https://www.baidu.com/s?&word="+this.searchWorld
        		this.searchWorld = ''
        		window.open(url,"_blank");
        	}

        }
	})