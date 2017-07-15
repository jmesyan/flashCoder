import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {requestList} from '../actions/dataActions';
import {TableWithHandle, TitleWithFunc} from '../components' 
import {connect} from 'react-redux'

class cronList extends Component{
	static propTypes = {
    lastUpdated: PropTypes.number,
    dispatch: PropTypes.func.isRequired
  }

  crontab(v, k){
      var sep = " ";
      var cron = v.Second + sep + v.Minute + sep + v.Hour + sep + v.Day + sep + v.Month + sep + v.Week;
      return (
        <td key={k}>{cron}</td>
      );
  }

  stateDs(row, k){
      var stateDesc = row.state == 1?"关闭":"开启";
      return (
        <td key={k}>{stateDesc}</td>
      );
  }

  handle(row, k){
      var crid = row.Crid, url = "/cron/update?crid="+crid;
      var updateDesc = row.State == 1? "开启":"关闭";
      var updateUrl = "/cron/updateState?crid="+crid;
      var deleteUrl = "/cron/delete?crid="+crid;
      return (
        <td key={k}>
        <a href={url} >修改</a>
        <a href={updateUrl} className="ml-20">{updateDesc}</a>
        <a href={deleteUrl} className="ml-20">删除</a>
        </td>
      );
  }

  componentDidMount(){
  	const {dispatch} = this.props
	  dispatch(requestList('cronList'))		
  }

  render(){
    var titles = ["编号", "任务ID", "执行计划", "状态", "操作"];
    var disrow = [
      {name:"Crid"},
      {name:"Tid"},
      {name:false, handle:this.crontab},
      {name:false, handle:this.stateDs},
      {name:false, handle:this.handle},
    ];

    var lists = this.props.lists, cronList = [];
    if (lists != null && lists.fetched){
      cronList = lists.cronList;
      return (
        <div>
           <TitleWithFunc title="定时任务列表" handleName="添加定时任务" handleUrl="/cron/add" handleFunc={function(){return false}}/>
           <TableWithHandle titles={titles} list={cronList.list}  disrow={disrow}/>
           <div dangerouslySetInnerHTML={{__html: cronList.page}} />
        </div>
      );
    } else {
      return (
        <div>
           <TitleWithFunc title="定时任务列表" handleName="添加定时任务" handleUrl="/cron/add" handleFunc={function(){return false}}/>
           暂无数据
        </div>
      );
    }

		
  }
}

const mapStateToProps = state => {
  return state;
}

export default connect(mapStateToProps)(cronList)