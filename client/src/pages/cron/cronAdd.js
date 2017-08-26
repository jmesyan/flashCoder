import React from 'react'
import PropTypes from 'prop-types'
import {requestItem, postForm} from '../../actions/dataActions';
import {TitleWithFunc, CronFromCompont,TaskCrontabPanel} from '../../components' 
import {connect} from 'react-redux'
import { Field, Form, Errors, actions } from 'react-redux-form';
import "bootstrap/js/collapse";

const  regexps = {
	dayExist:/^((([1-9]|([12]\d|3[01]))|([1-9]|([12]\d|3[01]))[-,]([1-9]|([12]\d|3[01])))|(\*\/([1-9]|([12]\d|3[01]))))$/,
	weekExist:/^(([0-6]|[0-6][-,][0-6])|(\*\/[1-6]))$/,
}

class cronAdd extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            taskType:'basicTask',
            chooseList:{basicTask:[], compositeTask:[]}
        }
    }
    taskTypeChange(val) {
        this.setState({taskType:val})
    }
    handleSubmit(cronForm) {
        const { dispatch } = this.props;
        let submitForm = {};
        $.extend(submitForm, cronForm)
        submitForm['taskType'] = this.state.taskType;
        submitForm['basicTask'] = this.state.chooseList.basicTask;
        submitForm['compositeTask'] = this.state.chooseList.compositeTask;
        dispatch(postForm('cronFormAdd', submitForm))
    }

    taskChooseList(type, list){
        let chooseList = this.state.chooseList;
        chooseList[type] = list;
        this.setState({chooseList:chooseList});
    }
    
    componentDidMount() {
        const {dispatch, params} = this.props
        var param = {}
        dispatch(requestItem('cronAddTasks', param))
    }
    render(){
        const {items} = this.props;
        var basicTaskList = items&&items.cronAddTasks? items.cronAddTasks.basicTaskList:[];
        var compositeTaskList = items&&items.cronAddTasks? items.cronAddTasks.compositeTaskList:[];
        return (
            <div>
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
                    onSubmit = {this.handleSubmit.bind(this)}
                >
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
                        <TaskCrontabPanel itemList={basicTaskList} saveName="basicTask" itemIdName="Tid" itemDisName="Tname" dialogName="basicConfirm" taskChooseList = {this.taskChooseList.bind(this)}/>
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
                        <TaskCrontabPanel itemList={compositeTaskList} saveName="compositeTask" itemIdName="Tid" itemDisName="Tname" dialogName="compositeConfirm" taskChooseList = {this.taskChooseList.bind(this)}/>
                    </div>
                </div>
                </div>
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
  
  export default connect(mapStateToProps)(cronAdd)