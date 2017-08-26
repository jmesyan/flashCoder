import React from 'react';
import DevTools from '../helpers/devTools';

var app = React.createClass({

	componentDidMount:function(){
		var href = location.href;
		if (href.indexOf("cron") != -1){
			$(".breadcrumb li a").removeClass("active");
			$(".breadcrumb li:eq(0) a").addClass("active");
		} else if (href.indexOf("task") != -1){
			$(".breadcrumb li a").removeClass("active");
			$(".breadcrumb li:eq(1) a").addClass("active");
		} else if (href.indexOf("behavior") != -1){
			$(".breadcrumb li a").removeClass("active");
			$(".breadcrumb li:eq(2) a").addClass("active");
		} else if (href.indexOf("operate") != -1){
			$(".breadcrumb li a").removeClass("active");
			$(".breadcrumb li:eq(3) a").addClass("active");
		} 
	},

	render: function() {
		return (
			<div className="container bg-image  mt-20">
				<div id="navigator" className="bg-navi">
					<ol className="breadcrumb bgc-none">
					  <li><a href="/cron/list" className="cron active">定时任务</a></li>
					  <li><a href="/task/list" className="task">任务列表</a></li>
					  <li><a href="/behavior/list" className="behavior">行为列表</a></li>
					  <li><a href="/operate/list" className="operate">操作列表</a></li>
					 </ol>
				</div>
			 	{this.props.children}
				<div>
					<DevTools/>
				</div>
			</div>
		);
	}

});

module.exports = app;