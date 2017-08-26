import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {requestList} from '../../actions/dataActions';
import {TableWithHandle, TitleWithFunc} from '../../components' 
import {connect} from 'react-redux'

class editTaskBehaviors extends Component {

    handle(row, k){
        var tbid = row.Tbid; 
    	var url = "/task/taskBehaviorParams/"+tbid;
		return (
			<td key={k}><a href={url} >编辑参数</a></td>
		);
    }

    componentDidMount(){
        const {dispatch, params} = this.props
        dispatch(requestList('taskBehaviors', params))		
    }

    render() {
        var titles = ["编号", "任务名称", "任务时序", "行为名称", "行为时序", "操作"];
        var disrow = [
            {name:"Tbid"},
            {name:"Tname"},
            {name:"Torder"},
            {name:"Bname"},
            {name:"Border"},
            {name:false, handle:this.handle.bind(this)},
        ];
        var lists = this.props.lists, taskBehaviors = {list:[]};
        if (lists != null && 'taskBehaviors' in lists){
            taskBehaviors = lists.taskBehaviors;
        } 
        return (
            <div>
                <TitleWithFunc title="任务参数" handleName="返回列表" handleUrl="/task/list" handleFunc={function(){}}/>
                <TableWithHandle titles={titles} list={taskBehaviors.list}  disrow={disrow}/>
            </div>
        );
    }
}

const mapStateToProps = state => {
    return state;
  }
  
  export default connect(mapStateToProps)(editTaskBehaviors)