var DropMenu = React.createClass({
	getInitialState: function() {
		return {
			options:[]
		};
	},

	componentDidMount: function() {
		var url = this.props.url || "";
		if (url.length > 0){
			 fetch(url).then((response) => {
	            return response.json();
	        }).then((responseJson) => {
	                this.setState({options:responseJson});
	        }).catch((error) => console.error(error));
		}
	},

	render: function() {
		return (
			 <select name="operate" className="form-control dropdown-toggle" onChange={this.props.changeFunc} >
		       {
		       	this.state.options.map(function(option,k){
		       		return (<option key={k} value={option.value}>{option.name}</option>);
		       	})
		       }
		     </select>
		);
	}

});
