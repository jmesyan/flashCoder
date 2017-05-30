var TitleWithFunc = React.createClass({

	render: function() {
		return (
			<div>
			  	<div className="h1 text-center" style={{float: "left", width: "80%"}}>{this.props.title}</div>
				<div style={{float: "left", width: "20%"}} className="h2 text-right">
					<a className="btn btn-primary" href={this.props.handleUrl} role="button" onClick={this.props.handleFunc}>{this.props.handleName}</a>
				</div>
			  </div>
		);
	},

});

