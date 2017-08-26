import React from 'react'
import PropTypes from 'prop-types'
import {requestItem, postForm} from '../../actions/dataActions';
import {TitleWithFunc, BehaviorParams} from '../../components' 
import {connect} from 'react-redux'
import { Field, Form, Errors, actions } from 'react-redux-form';
var initial = false, _bname = '', _remark = '';
class taskBehaviorParams extends React.Component {
    static propTypes = {
        dispatch: PropTypes.func.isRequired
    }

    handleSubmit(behaviorForm) {
        const { dispatch, params} = this.props;
        let tbid = params.tbid;
        let submitForm = {tbid:tbid, paramsList:behaviorForm.paramsList}
        dispatch(postForm('taskBehaviorParams', submitForm))
    }


    componentDidMount() {
        const {dispatch, params} = this.props
          var param = {tbid:params.tbid}
          dispatch(requestItem('taskBehaviorItem', param))
    }

    
    shouldComponentUpdate(nextProps, nextState){
        if (nextProps.items != null & !initial) {
            let taskBehaviorItem = nextProps.items.taskBehaviorItem;
            let item = {paramsList:taskBehaviorItem.params}
            nextProps.dispatch(actions.change('behaviorForm', item))
            _bname = taskBehaviorItem.base.Bname;
            _remark = taskBehaviorItem.base.Remark
            initial = true;
            return false;
        }
        return true;
    }

    behaviorParamsUpdate(behaviorParams){
        this.setState({behaviorParams:behaviorParams})
    }



     render() {
        var items = this.props.items, taskBehaviorItem = {};
        if (items != null && 'taskBehaviorItem' in items){
            taskBehaviorItem = items.taskBehaviorItem;
        }
         return (
            <div>
                 <TitleWithFunc title="更新行为" handleName="返回列表" handleUrl="#"handleFunc={()=>history.go(-1)}/>
                 <Form model="behaviorForm"  onSubmit={this.handleSubmit.bind(this)} className="clear">
                    <div className="he-30"></div>
                    <div className="form-inline cl-fff">
                        <div className="form-group">
                        <label htmlFor="bname">行为名称：{_bname}</label>
                        </div>
                        <div className="form-group ml-50">
                        <label htmlFor="remake">备注：{_remark}</label>
                        </div>
                        <div className="form-group ml-20 fr">
                        <input type="submit" className="btn btn-primary" value="提交数据" />
                        </div>
                    </div>
                    <div className="panel panel-primary clear mt-20">
                        <div className="panel-heading">修改参数</div>
                        <div className="panel-body">
                        <div id="params">
                            <BehaviorParams btype="update" paramsList={taskBehaviorItem.params} onlyEditContent={true} behaviorParamsUpdate={this.behaviorParamsUpdate.bind(this)}/>
                        </div>
                        </div>
                    </div>
                </Form>
            </div>
         );
     }
}

const mapStateToProps = state => {
    return state;
  }
  
export default connect(mapStateToProps)(taskBehaviorParams)