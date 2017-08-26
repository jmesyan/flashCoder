import React from 'react'
import PropTypes from 'prop-types'
import {requestItem, postForm} from '../../actions/dataActions';
import {TitleWithFunc, TaskChoosePanel} from '../../components' 
import {connect} from 'react-redux'
import { Field, Form, Errors, actions } from 'react-redux-form';
import "bootstrap/js/collapse";

const  regexps = {
	dayExist:/^((([1-9]|([12]\d|3[01]))|([1-9]|([12]\d|3[01]))[-,]([1-9]|([12]\d|3[01])))|(\*\/([1-9]|([12]\d|3[01]))))$/,
	weekExist:/^(([0-6]|[0-6][-,][0-6])|(\*\/[1-6]))$/,
}

class taskAdd extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            taskType:'basicTask',
            taskName:'',
            chooseList:{basicTaskBehaviors:[], compositeTaskBasics:[]}

        }
    }
    taskTypeChange(val) {
        this.setState({taskType:val})
    }
    handleSubmit() {
        var taskName = this.state.taskName;
        if (taskName == '') {
            alert('任务名称不能为空!');
            return false;
        }
        const { dispatch } = this.props;
        let submitForm = {};
        submitForm['taskType'] = this.state.taskType;
        submitForm['taskName'] = this.state.taskName;
        submitForm['basicTaskBehaviors'] = this.state.chooseList.basicTaskBehaviors;
        submitForm['compositeTaskBasics'] = this.state.chooseList.compositeTaskBasics;
        dispatch(postForm('taskFormAdd', submitForm))
    }

    taskChooseList(type, list){
        let chooseList = this.state.chooseList;
        chooseList[type] = list;
        this.setState({chooseList:chooseList});
    }
    
    componentDidMount() {
        const {dispatch, params} = this.props
        var param = {}
        dispatch(requestItem('AddTasks', param))
    }
    render(){
        const {items} = this.props;
        var behaviorList = items&&items.AddTasks? items.AddTasks.behaviorList:[];
        var basicTaskList = items&&items.AddTasks? items.AddTasks.basicTaskList:[];
        return (
            <div>
                <TitleWithFunc title="添加任务" handleName="返回列表" handleUrl="#"  handleFunc={()=>history.go(-1)}/>
                <form>
                <div className="panel-group clear" id="accordion" role="tablist" aria-multiselectable="true">
                <div className="panel panel-primary">
                    <div className="panel-heading" role="tab" id="headingOne">
                    <h4 className="panel-title">
                        <a role="button" data-toggle="collapse" data-parent="#accordion" href="#BasicTask" aria-expanded="true" aria-controls="BasicTask" onClick={this.taskTypeChange.bind(this, 'basicTask')}>
                        基础任务
                        </a>
                    </h4>
                    </div>
                    <div id="BasicTask" className="panel-collapse collapse in" role="tabpanel" aria-labelledby="headingOne">
                        <TaskChoosePanel itemList={behaviorList} saveName="basicTaskBehaviors" itemIdName="Bid" itemDisName="Bname" dialogName="basicConfirm" taskChooseList = {this.taskChooseList.bind(this)}/>
                    </div>
                </div>
                <div className="panel panel-success">
                    <div className="panel-heading" role="tab" id="headingTwo">
                    <h4 className="panel-title">
                        <a className="collapsed" role="button" data-toggle="collapse" data-parent="#accordion" href="#CompositeTask" aria-expanded="false" aria-controls="CompositeTask" onClick={this.taskTypeChange.bind(this, 'compositeTask')}>
                        复合任务
                        </a>
                    </h4>
                    </div>
                    <div id="CompositeTask" className="panel-collapse collapse" role="tabpanel" aria-labelledby="headingTwo">
                    <TaskChoosePanel itemList={basicTaskList} saveName="compositeTaskBasics" itemIdName="Tid" itemDisName="Tname" dialogName="compositeConfirm" taskChooseList = {this.taskChooseList.bind(this)}/>
                    </div>
                </div>
                </div>
                <div className="form-inline cl-fff">
                    <div className="form-group">
                    <label htmlFor="taskName">任务名称：</label>
                    <input type="input" className="form-control" id="taskName" name="taskName" placeholder="请输入任务名称" onChange={(event)=>this.setState({taskName:event.target.value})} />
                    </div>
                    <input type="button" onClick={this.handleSubmit.bind(this)} className="btn btn-primary" value="提交任务" />
                </div>
                </form>
            </div>
        );
    }
}

const mapStateToProps = state => {
    return state;
  }
  
  export default connect(mapStateToProps)(taskAdd)