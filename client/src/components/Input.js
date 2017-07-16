var React = require('react');

var Input = React.createClass({

	render: function() {
		const {value} = this.props
		return (
			<input type="input"  className="form-control fc-p64" onChange={(e)=>{this.refs.second = e.target.value}} ref="second"   id="second" name="second" value={value} placeholder="请输入秒" required/>
		);
	}

});

module.exports = Input;