var React = require('react');

var TitleWithFunc = React.createClass({

	render: function() {
		return (
			<div>
			  	<div className="h1 text-center lw-80">{this.props.title}</div>
				<div className="h2 text-right lw-20">
					<a className="btn btn-primary" href={this.props.handleUrl} role="button" onClick={this.props.handleFunc}>{this.props.handleName}</a>
				</div>
			  </div>
		);
	}

});

module.exports = TitleWithFunc;

