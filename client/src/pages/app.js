var React = require('react');

var app = React.createClass({

	render: function() {
		return (
			<div>
				<div id="navigator" className="bg-navi">
					<ol className="breadcrumb bgc-none">
					  <li><a href="/cron/list" className="cron active">定时任务</a></li>
					  <li><a href="/task/list" className="task">任务列表</a></li>
					  <li><a href="/behavior/list" className="behavior">行为列表</a></li>
					  <li><a href="/operate/list" className="operate">操作列表</a></li>
					 </ol>
				</div>
			 	{this.props.children}
			</div>
		);
	}

});

module.exports = app;