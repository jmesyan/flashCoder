import React from 'react'
import PropTypes from 'prop-types'
import {requestItem, postForm} from '../../actions/dataActions';
import {TitleWithFunc, CronFromCompont} from '../../components' 
import {connect} from 'react-redux'
import { Field, Form, Errors, actions } from 'react-redux-form';
const  regexps = {
	dayExist:/^((([1-9]|([12]\d|3[01]))|([1-9]|([12]\d|3[01]))[-,]([1-9]|([12]\d|3[01])))|(\*\/([1-9]|([12]\d|3[01]))))$/,
	weekExist:/^(([0-6]|[0-6][-,][0-6])|(\*\/[1-6]))$/,
}
 
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
				<div className="form-inline cl-fff">
					<div className="form-group">
						<label htmlFor="bname">任务ID：</label>
						<input type="input" className="form-control" value={cronItem.task.Tid} placeholder="任务ID" />
					</div>
					<div className="form-group ml-50">
						<label htmlFor="bname">任务名称：</label>
						<input type="input" className="form-control" value={cronItem.task.Tname} placeholder="任务名称" />
					</div>
				</div>
				<TitleWithFunc title="添加定时任务" handleName="返回列表" handleUrl="#"  handleFunc={()=>history.go(-1)}/>
					<Form 
						model="cronForm" 
							validators={{
								'': {
									dayOrWeek:(vals) => {
										if (vals.Day != "" && vals.Week != "" && regexps.dayExist.test(vals.Day) && regexps.weekExist.test(vals.Week)) return false;
										return true;
									}
								},
							}}
							onSubmit = {this.handleSubmit}
					>
				<CronFromCompont/>
				<div className="form-group mt-20">
					<input type="submit" className="btn btn-primary" value="提交任务" />
				</div>	
				</Form>
			</div>
	);
  
  }
}

const mapStateToProps = state => {
  return state;
}

export default connect(mapStateToProps)(cronUpdate)