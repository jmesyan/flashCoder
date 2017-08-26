import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {requestList} from '../../actions/dataActions';
import {TableWithHandle, TitleWithFunc} from '../../components' 
import {connect} from 'react-redux'
import {requestExec} from '../../actions/dataActions';
import {timeFormat} from '../../helpers/utils'

class taskList extends Component {
    
    taskDelete(tid) {
        const {dispatch} = this.props;
        let data = {tid:tid};
        dispatch(requestExec('taskDelete', data))
    }

    taskExecute(tid) {
        const {dispatch} = this.props;
        let data = {tid:tid};
        dispatch(requestExec('taskExecute', data))
    }
    
    handle(row, k){
    	var tid = row.Tid; 
    	var excuteUrl = "/task/taskExecute?tid="+tid;
    	var deleteUrl = "/task/delete?tid="+tid;
    	var editUrl = "/task/editTaskBehaviors/"+tid;
		return (
			<td key={k}>
			<a href="javascript:void(0)" onClick={this.taskExecute.bind(this, tid)} >执行</a>
			<a className="ml-20" href="javascript:void(0)" onClick={this.taskDelete.bind(this, tid)} >删除</a>
			<a className="ml-20" href={editUrl} >编辑行为</a>
			</td>
		);
    }

    ttFormat(row, k){
		if (row.Tcate == 2){
			return (<td key={k}>复合任务</td>);
		} else {
			return (<td key={k}>基础任务</td>);
		} 
    }

    componentDidMount(){
        const {dispatch} = this.props
        dispatch(requestList('taskList'))		
    }

    render() {
        var titles = ["编号", "任务名称", "任务类型", "更新时间", "操作"];
        var disrow = [
            {name:"Tid"},
            {name:"Tname"},
            {name:false, handle:this.ttFormat.bind(this)},
            {name:"Updtime",filters:[timeFormat]},
            {name:false, handle:this.handle.bind(this)},
        ];

        var lists = this.props.lists, taskList = {list:[], page:''};
        if (lists != null && 'taskList' in lists){
          taskList = lists.taskList;
        } 
        return (
            <div>
                <TitleWithFunc title="任务列表" handleName="添加任务" handleUrl="/task/add" handleFunc={function(){return false}}/>
                <TableWithHandle titles={titles} list={taskList.list}  disrow={disrow}/>
                <div dangerouslySetInnerHTML={{__html: taskList.page}} />
            </div>
        );
    }
}

const mapStateToProps = state => {
    return state;
  }
  
  export default connect(mapStateToProps)(taskList)
