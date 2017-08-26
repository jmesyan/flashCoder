import React from 'react'
import PropTypes from 'prop-types'
import {requestItem, postForm} from '../../actions/dataActions';
import {TitleWithFunc} from '../../components' 
import {connect} from 'react-redux'
import { Field, Form, Errors, actions } from 'react-redux-form';
const isRequired = (val) => val && val.length > 0;
class operateAdd extends React.Component {

    handleSubmit(operateForm) {
        const { dispatch } = this.props;
        dispatch(postForm('operateFormAdd', operateForm))
    }

    render(){
        return (
            <div>
            <TitleWithFunc title="添加操作" handleName="返回列表" handleUrl="#" handleFunc={()=>history.go(-1)}/>
            <Form model="operateForm"  onSubmit={this.handleSubmit.bind(this)} className="clear cl-fff">
                <div className="form-group">
                <label htmlFor="opname">操作名称：</label>
                <Field model="operateForm.opname"  validators={{isRequired}}>
                        <input type="text" className="form-control" id="opname" name="opname"  placeholder="操作名称"/>
                        <Errors
                            wrapper="small"
                            className="help-block  form-error"
                            show={{ touched: true, focus: false }}
                            model="operateForm.opname"
                            messages={{isRequired: '操作名称不能为空'}}
                        />
                </Field>
                </div>
                <div className="form-group">
                <label htmlFor="optag">操作标识：</label>
                <Field model="operateForm.optag"  validators={{isRequired}}>
                        <input type="text" className="form-control" id="optag" name="optag"  placeholder="操作标识"/>
                        <Errors
                            wrapper="small"
                            className="help-block  form-error"
                            show={{ touched: true, focus: false }}
                            model="operateForm.optag"
                            messages={{isRequired: '操作标识不能为空'}}
                        />
                </Field>
                </div>
                <div className="form-group">
                <label htmlFor="remark">备注：</label>
                <Field model="operateForm.remark"  validators={{isRequired}}>
                        <input type="text" className="form-control" id="remark" name="remark"  placeholder="备注"/>
                        <Errors
                            wrapper="small"
                            className="help-block  form-error"
                            show={{ touched: true, focus: false }}
                            model="operateForm.remark"
                            messages={{isRequired: '备注不能为空'}}
                        />
                </Field>
                </div>
                <div className="form-group text-center">
                <input type="submit" className="btn btn-primary" value="提交操作"  />
                </div>
            </Form>
            </div>
        )
    }
}

const mapStateToProps = state => {
    return state;
}
  
export default connect(mapStateToProps)(operateAdd)