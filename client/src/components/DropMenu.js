var React = require('react');
import {requestItem} from '../actions/dataActions';
import {connect} from 'react-redux'
var DropMenu = React.createClass({
	getInitialState: function() {
		return {
			options:[],
			value:"",
			initial:false
		};
	},

	componentDidMount: function() {
		const {fetchType, fetchParams, dispatch} = this.props;
		if (fetchType.length > 0) {
			dispatch(requestItem(fetchType, fetchParams))
		}
		if(this.props.value != 'undefined'){
			this.setState({value:this.props.value})
		}
	},

	componentWillUpdate:function(nextProps, nextState){
		if (nextProps.items && !this.state.initial){
			this.setState({options:eval(nextProps.items.jsonOperateList), initial:true})
			this.state.options = eval(nextProps.items.jsonOperateList)
			this.state.options && this.props.changeFunc(this.state.options[0].value+"")
		}
	},

	render: function() {
		return (
			 <select name="operate" className="form-control dropdown-toggle"  onChange={(event)=>this.props.changeFunc( event.target.value)} >
		       {
		       	this.state.options.map(function(option,k){
		       		return (<option key={k} value={option.value}>{option.name}</option>);
		       	})
		       }
		     </select>
		);
	}

});


const mapStateToProps = state => {
    return state;
  }
  
export default connect(mapStateToProps)(DropMenu)