import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {requestItem, changeItem} from '../../actions/dataActions';
import {TitleWithFunc} from '../../components' 
import {connect} from 'react-redux'
import  'bootstrap-validator'
let  regexps = {
    second:/^(((\d|([1-5]\d))|(\d|([1-5]\d))[-,](\d|([1-5]\d)))|(\*\/([1-9]|([1-5]\d)))|(\*))$/,
    minute:/^(((\d|([1-5]\d))|(\d|([1-5]\d))[-,](\d|([1-5]\d)))|(\*\/([1-9]|([1-5]\d)))|(\*))$/,
    hour:/^((([0-9]|(1[0-9]|2[0-3]))|([0-9]|(1[0-9]|2[0-3]))[-,]([0-9]|(1[0-9]|2[0-3])))|(\*\/([1-9]|(1[0-9]|2[0-3])))|(\*))$/,
    day:/^((([1-9]|([12]\d|3[01]))|([1-9]|([12]\d|3[01]))[-,]([1-9]|([12]\d|3[01])))|(\*\/([1-9]|([12]\d|3[01])))|([\*?]))$/,
    month:/^((([1-9]|(1[0-2]))|([1-9]|(1[0-2]))[-,]([1-9]|(1[0-2]))?)|(\*\/([1-9]|(1[0-2]?)))|(\*))$/,
    week:/^(([0-6]|[0-6][-,][0-6])|(\*\/[1-6])|([\*?]))$/,
    dayExist:/^((([1-9]|([12]\d|3[01]))|([1-9]|([12]\d|3[01]))[-,]([1-9]|([12]\d|3[01])))|(\*\/([1-9]|([12]\d|3[01]))))$/,
    weekExist:/^(([0-6]|[0-6][-,][0-6])|(\*\/[1-6]))$/,
    period:/^(\d+)-(\d+)$/
}


class cronUpdate extends Component{
  static propTypes = {
    dispatch: PropTypes.func.isRequired
  }

  componentDidMount() {
	  const {dispatch, params} = this.props
	  var param = {crid:params.crid}
	  dispatch(requestItem('cronItem', param))		
  }

  changeCronItem(e, item, key){
  	item.cron[key] = e.target.value; 
  	this.props.dispatch(changeItem('cronItem', item));
  }


  render(){
    var items = this.props.items, cronItem = {task:{}, cron:{}};
  	if (items != null && 'cronItem' in items){
       cronItem = items.cronItem;
    }

	return (
  		<div>
  			 <TitleWithFunc title="更新定时任务" handleName="返回列表" handleUrl="#" handleFunc={()=>history.go(-1)}/>
  			 <form method="post" id="updateForm"  data-toggle="validator">
			  <div className="form-inline clear" >
			    <div className="form-group fl">
			      <label className="form-control">任务ID:{cronItem.task.Tid}</label>
			    </div>
			    <div className="form-group fl ml-20">
			      <label className="form-control">任务名称:{cronItem.task.Tname}</label>
			    </div>
			  </div>
			 <div className="form-inline clear mt-20">
			    <div className="form-group">
			      	<input type="input" className="form-control fc-p64" pattern="/^(((\d|([1-5]\d))|(\d|([1-5]\d))[-,](\d|([1-5]\d)))|(\*\/([1-9]|([1-5]\d)))|(\*))$/"  onChange={(e)=>{this.changeCronItem(e, cronItem, 'Second')}} id="second" name="second" value={cronItem.cron.Second || ''} placeholder="请输入秒" required/>
		         	<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
    				<div class="help-block with-errors">Hey look, this one has feedback icons!</div>
			    </div>
			     <div className="form-group">
			      <input type="input" className="form-control fc-p64" onChange={(e)=>{this.changeCronItem(e, cronItem, 'Minute')}} id="minute" name="minute" value={cronItem.cron.Minute || ''} placeholder="请输入分"/>
			    </div>
			     <div className="form-group">
			      <input type="input" className="form-control fc-p64" onChange={(e)=>{this.changeCronItem(e, cronItem, 'Hour')}} value={cronItem.cron.Hour || ''} placeholder="请输入时"/>
			    </div>
			     <div className="form-group">
			      <input type="input" className="form-control fc-p64" onChange={(e)=>{this.changeCronItem(e, cronItem, 'Day')}} id="day" name="day" value={cronItem.cron.Day || ''} placeholder="请输入月天"/>
			    </div>
			     <div className="form-group">
			      <input type="input" className="form-control fc-p64" onChange={(e)=>{this.changeCronItem(e, cronItem, 'Month')}} id="month" name="month" value={cronItem.cron.Month || ''} placeholder="请输入月"/>
			    </div>
			     <div className="form-group">
			      <input type="input" className="form-control fc-p64" onChange={(e)=>{this.changeCronItem(e, cronItem, 'Week')}} id="week" name="week" value={cronItem.cron.Week || ''} placeholder="请输入周天"/>
			    </div>
			  </div>
			  <div className="form-group mt-20">
			   <input type="submit" className="btn btn-primary" value="提交任务" />
			   </div>
			</form>
  		</div>
	);
  
  }
}

const mapStateToProps = state => {
  return state;
}

export default connect(mapStateToProps)(cronUpdate)