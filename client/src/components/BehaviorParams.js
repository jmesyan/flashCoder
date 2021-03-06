var React = require('react');
var TableWithHandle = require('./TableWithHandle');
var ParamsEditor = require('./ParamsEditor');
import { Field, Form, Errors, actions } from 'react-redux-form';
import {connect} from 'react-redux'
const isRequired = (val) => val && val.length > 0;
var BehaviorParams = React.createClass({
	getDefaultProps: function() {
		return {
			paramTypes:[{val:1,name:"文本"}, {val:2,name:"密码"}, {val:3,name:"编辑器"}],
			paramsList:[],
			editKey:null,
			btype:"add",
			onlyEditContent:false,

		};
	},
	getInitialState: function() {
		return {
			paramsList:this.props.paramsList,
			editKey:this.props.editKey !=null ?this.props.editKey:(this.props.paramsList.length>0? 0:null),
			showUpdate:false,
			initial:false
		};
	},
	jsonToString(data){
        return JSON.stringify(data);
    },
    showType:function(row,key){
    	for (var k in this.props.paramTypes){
    		if (row.type == this.props.paramTypes[k].val) return (<td key={key}>{this.props.paramTypes[k].name}</td>);
    	}
    	return (<td key={key}>"-"</td>);
    },
    _getParamsKeyByName:function(name){
			var paramsList = eval(this.props.paramsList)
			for (var k in paramsList){
				if (paramsList[k].name == name) return k;
			}
			return false;
	},
	_getParamsKeyByNameUseState:function(name){
		var paramsList = eval(this.state.paramsList)
		for (var k in paramsList){
			if (paramsList[k].name == name) return k;
		}
		return false;
	},
    setEditer:function(name){
			var key = this._getParamsKeyByName(name);
			if (key){
					this.setState({editKey:key, showUpdate:true});
			} else {
				key = this._getParamsKeyByNameUseState(name);
				if (key) {
					this.setState({editKey:key, showUpdate:true});
				} else {
					alert("错误1111")
				}
			}		
				
    },
    paramsDelete:function(name){
		var key = this._getParamsKeyByName(name);
		if (key){
			var list = this.state.paramsList;
			list.splice(key, 1);
			this.setState({paramsList:list});
		} else {
			key = this._getParamsKeyByNameUseState(name);
			if (key){
				let paramsList = [];
				$.extend(paramsList,  this.state.paramsList)
				paramsList.splice(key, 1);
				this.state.paramsList = paramsList;
				this.setState({paramsList:paramsList,showUpdate:false})
				this.props.dispatch(actions.change('behaviorForm.paramsList', eval(paramsList)))
			} else {
				alert("不存在该参数")
			}
		}
    },
    setEditerLink:function(row, key){
    	return (
				<td key={key}>
				<a href="javascript:;"  onClick={this.setEditer.bind(this,row.name)}>编辑</a>
				{this.props.btype == "add"? <a href="javascript:;" className="ml-20"  onClick={this.paramsDelete.bind(this,row.name)}>删除</a>:""}
				</td>
		);
    },
    saveAddParams:function(refs){
		var param = {
			type:refs.paramType.value,
			name:refs.paramName.value,
			value:refs.paramValue.value,
		};
		var namekey = this._getParamsKeyByNameUseState(param.name);
		if (namekey) {
			alert("参数名已存在!");
		} else {
			// this.state.paramsList.push(param);
			let paramsList = [];
			$.extend(paramsList,  this.state.paramsList)
			paramsList.push(param);
			this.state.paramsList = paramsList;
			this.setState({paramsList:paramsList,showUpdate:false})
			this.props.dispatch(actions.change('behaviorForm.paramsList', eval(paramsList)))
			this.refs.tabHome.className = "active";
			this.refs.home.className = "tab-pane active";
			this.refs.tabSettings.className = "";
			this.refs.settings.className = "tab-pane";
		}
    },
    saveUpdateParams:function(refs){
		var editKey = this.state.editKey;
		if (editKey == null) {
			alert("错误2222");
		} else {
			var param = {
				type:refs.paramType.value,
				name:refs.paramName.value,
				value:refs.paramValue.value,
			};
			var namekey = this._getParamsKeyByName(param.name);
			namekey = namekey || this._getParamsKeyByNameUseState(param.name);
			if (!namekey || (namekey && namekey == editKey)){
				let paramsList = [];
				$.extend(paramsList,  this.state.paramsList)
				paramsList.splice(editKey,1)
				paramsList.splice(editKey, 0, param)
				this.setState({paramsList:paramsList,showUpdate:false});
				this.state.paramsList = paramsList;
				this.props.behaviorParamsUpdate && this.props.behaviorParamsUpdate(paramsList)
				this.props.dispatch(actions.change('behaviorForm.paramsList', eval(paramsList)))
			} else {
				alert("参数名已存在!");
			} 
		}

    },
    undoUpdate:function(){
		this.setState({showUpdate:false});
		},
	componentWillUpdate:function(nextProps, nextState){
		if (!this.state.initial && nextProps.paramsList.length > 0) {
			this.setState({
				paramsList:eval(this.props.paramsList),
				editKey:this.props.editKey !=null ?this.props.editKey:0,
				showUpdate:false,
				initial:true
			});
			this.state.paramsList = eval(this.props.paramsList);
			nextProps.dispatch(actions.change('behaviorForm.paramsList', eval(this.props.paramsList)))
			this.state.editKey = this.props.editKey !=null ?this.props.editKey:0;
		}
	},
	render: function() {
		  //表格
		var titles = ["参数类型", "参数名称", "操作"];
    var disrow = [
    	{name:false, handle:this.showType},
			{name:"name"},
			{name:false, handle:this.setEditerLink},
		];
	return (
			<div>
				<Field model="behaviorForm.paramsList"  validators={{isRequired}} className="inline-block">
						<input type="hidden" className="form-control" id="paramsList" name="paramsList" value={this.jsonToString(this.state.paramsList)}/>
						<Errors
							wrapper="small"
							className="help-block  form-error"
							show={{ touched: true, focus: false }}
							model="behaviorForm.paramsList"
							messages={{isRequired: '参数不能为空'}}
						/>
				</Field>
				<ul className="nav nav-tabs" role="tablist">
		            <li role="presentation" className="active" ref="tabHome"><a href="#home" aria-controls="home" role="tab" data-toggle="tab">参数列表</a></li>
		            {this.props.btype == "add"? <li role="presentation" ref="tabSettings"><a href="#settings" aria-controls="settings" role="tab" data-toggle="tab">添加参数</a></li>:""}
		           
		          </ul>

		          <div className="tab-content">
		            <div role="tabpanel"  className="tab-pane active" ref="home" id="home">
		                 <div className={this.state.showUpdate? "hide":"show"}><TableWithHandle titles={titles} list={this.state.paramsList}  disrow={disrow}/></div>
		                 <div className={this.state.showUpdate? "show":"hide"}>
											<ParamsEditor type="update" paramsList={this.state.paramsList} editKey={this.state.editKey} updateSaveFunc={this.saveUpdateParams} undoUpdate={this.undoUpdate} onlyEditContent={this.props.onlyEditContent}/>
		                 </div>
		            </div>
		            {this.props.btype == "add" ? <div role="tabpanel" className="tab-pane" ref="settings" id="settings"><ParamsEditor type="add" paramsList={this.state.paramsList} addSaveFunc={this.saveAddParams}/></div>:""}
		           
		          </div>
			</div>
		);
	},

});



const mapStateToProps = state => {
    return state;
  }
  
export default connect(mapStateToProps)(BehaviorParams)

