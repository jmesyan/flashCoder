import React from 'react'
import PropTypes from 'prop-types'
import {requestItem, postForm} from '../../actions/dataActions';
import {TitleWithFunc, CronFromCompont} from '../../components' 
import {connect} from 'react-redux'
import {actions} from 'react-redux-form';
 
var initial = false;

class cronUpdate extends React.Component{
  constructor(props) {
    super(props);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit(cronForm) {
		const { dispatch } = this.props;
		dispatch(postForm('cronFormUpdate', cronForm))
	}
	
  static propTypes = {
    dispatch: PropTypes.func.isRequired
  }

  componentDidMount() {
	  const {dispatch, params} = this.props
		var param = {crid:params.crid}
		dispatch(requestItem('cronItem', param))
		
	}

	shouldComponentUpdate(nextProps, nextState){
		if (nextProps.items != null & !initial) {
			nextProps.dispatch(actions.change('cronForm', nextProps.items.cronItem.cron))	
			initial = true;
			return false;
		}
		return true;
	}


  render(){
    var items = this.props.items, cronItem = {task:{}, cron:{}};
  	if (items != null && 'cronItem' in items){
			 cronItem = items.cronItem;
    }

	return (
			<div>
				<TitleWithFunc title="更新定时任务" handleName="返回列表" handleUrl="#" handleFunc={()=>history.go(-1)}/>
				<CronFromCompont handleSubmit={this.handleSubmit} cronItem = {cronItem} />
			</div>
	);
  
  }
}

const mapStateToProps = state => {
  return state;
}

export default connect(mapStateToProps)(cronUpdate)