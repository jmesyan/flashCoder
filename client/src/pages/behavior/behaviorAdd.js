import React from 'react'
import PropTypes from 'prop-types'
import {requestItem, postForm} from '../../actions/dataActions';
import {TitleWithFunc, DropMenu, BehaviorParams} from '../../components' 
import {connect} from 'react-redux'
import { Field, Form, Errors, actions } from 'react-redux-form';
import "bootstrap/js/tab";
const isRequired = (val) => val && val.length > 0;
var _dispatch;
class behaviorAdd extends  React.Component{
    constructor(props) {
        super(props)
        _dispatch = props.dispatch
    }

    opChange(value){
        _dispatch(actions.change('behaviorForm.operate', value))	
    }

    handleSubmit(behaviorForm){
        _dispatch(postForm('behaviorAdd', behaviorForm))
    }

    render(){
        return (
            <div>
                <TitleWithFunc title="添加行为" handleName="返回列表" handleUrl="#" handleFunc={()=>history.go(-1)}/>
                <Form model="behaviorForm"  onSubmit={this.handleSubmit}>
                <div className="he-30"></div>
                <div className="form-inline cl-fff">
                    <div className="form-group">
                    <label htmlFor="bname">行为名称：</label>
                    <Field model="behaviorForm.bname"  validators={{isRequired}} className="inline-block">
                                <input type="text" className="form-control" id="bname" name="bname"  placeholder="行为名称"/>
                                <Errors
                                    wrapper="small"
                                    className="help-block  form-error"
                                    show={{ touched: true, focus: false }}
                                    model="behaviorForm.bname"
                                    messages={{isRequired: '行为名称不能为空'}}
                                />
                        </Field>
                    </div>
                    <div className="form-group ml-20">
                    <label htmlFor="operate">操作类型：</label>
                    <span id="operateAdd">
                    <DropMenu fetchType="jsonOperateList"  changeFunc={this.opChange} />    
                    </span> 
                    </div>
                    <div className="form-group ml-50">
                    <label htmlFor="bname">备注：</label>
                    <Field model="behaviorForm.remark"  validators={{isRequired}} className="inline-block">
                            <input type="text" className="form-control" id="remark" name="remark"  placeholder="备注"/>
                            <Errors
                                wrapper="small"
                                className="help-block  form-error"
                                show={{ touched: true, focus: false }}
                                model="behaviorForm.remark"
                                messages={{isRequired: '备注不能为空'}}
                            />
                    </Field>
                    </div>
                    <div className="form-group ml-20 fr">
                    <input type="submit" className="btn btn-primary" value="提交数据" />
                    </div>
                </div>
                <div className="panel panel-primary clear mt-20">
                    <div className="panel-heading">默认参数</div>
                    <div className="panel-body">
                    <div id="params">
                    <BehaviorParams btype="add" />
                    </div>
                    </div>
                </div>
                </Form>
            </div>
        )
    }
}

const mapStateToProps = state => {
    return state;
}
  
export default connect(mapStateToProps)(behaviorAdd)