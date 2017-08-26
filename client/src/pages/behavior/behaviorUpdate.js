import React from 'react'
import PropTypes from 'prop-types'
import {requestItem, postForm} from '../../actions/dataActions';
import {TitleWithFunc, BehaviorParams} from '../../components' 
import {connect} from 'react-redux'
import { Field, Form, Errors, actions } from 'react-redux-form';
const isRequired = (val) => val && val.length > 0;
var initial = false, _this;
class behaviorUpdate extends React.Component {
    constructor(props){
        super(props);
        this.handleSubmit = this.handleSubmit.bind(this);
    }
     static propTypes = {
        dispatch: PropTypes.func.isRequired
      }
      
      componentDidMount() {
          const {dispatch, params} = this.props
          var param = {bid:params.bid}
          dispatch(requestItem('behaviorItem', param))
      }

      behaviorParamsUpdate(behaviorParams){
        this.setState({behaviorParams:behaviorParams})
      }

      handleSubmit(){
          const {dispatch, behaviorForm} = this.props;
          dispatch(postForm('behaviorUpdate', behaviorForm))
      }

     shouldComponentUpdate(nextProps, nextState){
            if (nextProps.items != null & !initial) {
                let item = {bid:nextProps.items.behaviorItem.Bid,bname:nextProps.items.behaviorItem.Bname, remark:nextProps.items.behaviorItem.Remark}
                nextProps.dispatch(actions.change('behaviorForm', item))	
                initial = true;
                return false;
            }
            return true;
        }
      
  
      render() {
        var items = this.props.items, behaviorItem = {};
        if (items != null && 'behaviorItem' in items){
               behaviorItem = items.behaviorItem;
        }

        return (
             <div>
                 <TitleWithFunc title="更新行为" handleName="返回列表" handleUrl="#" handleFunc={()=>history.go(-1)} />
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
                        <input type="button" onClick={this.handleSubmit.bind(this)} className="btn btn-primary" value="提交数据" />
                        </div>
                    </div>
                    <div className="panel panel-primary clear mt-20">
                    <div className="panel-heading">默认参数</div>
                    <div className="panel-body">
                    <div id="params">
                        <BehaviorParams btype="update" paramsList={behaviorItem.Paramsdef} behaviorParamsUpdate={this.behaviorParamsUpdate.bind(this)}/>
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
  
export default connect(mapStateToProps)(behaviorUpdate)