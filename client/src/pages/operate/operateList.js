import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {requestList} from '../../actions/dataActions';
import {TableWithHandle, TitleWithFunc} from '../../components' 
import {connect} from 'react-redux'
import {requestExec} from '../../actions/dataActions';
import {timeFormat} from '../../helpers/utils'

class operateList extends Component {
    static propTypes = {
        lastUpdated: PropTypes.number,
        dispatch: PropTypes.func.isRequired
    }

    deleteOperate(opid) {
        const {dispatch} = this.props;
        let data = {opid:opid};
        dispatch(requestExec('operateDelete', data))
    }

    componentDidMount(){
        const {dispatch} = this.props
        dispatch(requestList('operateList'))		
    }
  

    handle(row, k){
    	var opid = row.Opid;
        return (
                <td key={k}><a href="javascript:;"  onClick={this.deleteOperate.bind(this, opid)}>删除</a></td>
        );
    }

    render(){
        var titles = ["编号", "操作名称", "操作标识", "添加时间", "操作", "备注"];
        var disrow = [
            {name:"Opid"},
            {name:"Opname"},
            {name:"Optag"},
            {name:"Addtime",filters:[timeFormat]},
            {name:false, handle:this.handle.bind(this)},
            {name:"Remark"},
        ];
        var lists = this.props.lists, cronList = {list:[], page:''};
        if (lists != null && 'operateList' in lists){
            operateList = lists.operateList;
        } 
        return (
            <div>
                <TitleWithFunc title="操作列表" handleName="添加操作" handleUrl="/operate/add" handleFunc={function(){return false}}/>
                <TableWithHandle titles={titles} list={operateList.list}  disrow={disrow}/>
                <div dangerouslySetInnerHTML={{__html: operateList.page}} />
            </div>
        )
    }
}

const mapStateToProps = state => {
    return state;
  }
  
export default connect(mapStateToProps)(operateList)