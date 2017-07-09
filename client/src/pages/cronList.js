import {TableWithHandle, TitleWithFunc} from '../components' 
var React = require('react');

var cronList = React.createClass({

	render: function() {
		return (
			<div>
				 <TitleWithFunc title="定时任务列表" handleName="添加定时任务" handleUrl="/cron/add" handleFunc={function(){return false}}/>
			</div>
		);
	}

});

export default cronList