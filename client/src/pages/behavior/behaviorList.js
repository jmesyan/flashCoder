import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {requestList} from '../../actions/dataActions';
import {TableWithHandle, TitleWithFunc} from '../../components' 
import {connect} from 'react-redux'
import {requestExec} from '../../actions/dataActions';
import {timeFormat} from '../../helpers/utils'

class behaviorList extends Component {
    
    behaviorDelete(bid) {
        const {dispatch} = this.props;
        let data = {bid:bid};
        dispatch(requestExec('behaviorDelete', data))
    }

    taskExecute(tid) {
        const {dispatch} = this.props;
        let data = {tid:tid};
        dispatch(requestExec('taskExecute', data))
    }
    
    handle(row, k){
    	var bid = row.Bid, url = "/behavior/update/"+bid;
    	var deleteUrl = "/behavior/delete?bid="+bid;
		return (
			<td key={k}>
				<a href={url} >修改</a>
				<a href="javascript:void(0)" onClick={this.behaviorDelete.bind(this, bid)} className="ml-20">删除</a>
			</td>
		);
    }


    componentDidMount(){
        const {dispatch} = this.props
        dispatch(requestList('behaviorList'))		
    }


    render() {
        var titles = ["编号", "行为名称", "操作名称", "备注", "更新时间", "操作"];
        var disrow = [
            {name:"Bid"},
            {name:"Bname"},
            {name:"Opname"},
            {name:"Remark"},
            {name:"Updtime",filters:[timeFormat]},
            {name:false, handle:this.handle.bind(this)},
        ];

        var lists = this.props.lists, behaviorList = {list:[], page:''};
        if (lists != null && 'behaviorList' in lists){
          behaviorList = lists.behaviorList;
        } 
        return (
            <div>
                <TitleWithFunc title="行为列表" handleName="添加行为" handleUrl="/behavior/add" handleFunc={function(){return false}}/>
                <TableWithHandle titles={titles} list={behaviorList.list}  disrow={disrow}/>
                <div dangerouslySetInnerHTML={{__html: behaviorList.page}} />
            </div>
        );
    }
}

const mapStateToProps = state => {
    return state;
  }
  
  export default connect(mapStateToProps)(behaviorList)
