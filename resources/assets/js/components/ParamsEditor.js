
var ParamsEditor = React.createClass({
	getDefaultProps: function() {
		console.log("getDefaultProps")
		return {
			paramTypes:[{val:1,name:"文本"}, {val:2,name:"密码"}, {val:3,name:"编辑器"}],
			paramsList:[],
			editKey:null,
			type:"add",
			onlyEditContent:false,
		};
	},
	getInitialState: function() {
		console.log(this.props.type, "getInitialState")
		return {
			paramsList:this.props.paramsList,
			editKey:this.props.editKey,
			curType:this.props.editKey !=null? this.props.paramsList[this.props.editKey].type:"1",
			onlyTypeName:null,
		};
	},
	componentWillReceiveProps: function(nextProps) {
		console.log(this.props.type,"componentWillReceiveProps")
		this.setState ({
			paramsList:nextProps.paramsList,
			editKey:nextProps.editKey,
		});
	},
	componentWillUpdate: function(nextProps, nextState) {
			console.log(this.props.type,"componentWillUpdate")
		var editKey = nextState.editKey;
		if (this.props.type == "update" && editKey != null){
				this.refs.paramType.value = nextState.paramsList[editKey].type;
				this.refs.paramName.value = nextState.paramsList[editKey].name;
				this.refs.paramValue.value = nextState.paramsList[editKey].value;
		}
	},
	componentDidMount: function() {
		var editKey = this.state.editKey;
		if(this.props.onlyEditContent && editKey != null){
			for (var k in this.props.paramTypes){
				console.log(this.props.paramTypes[k].val , parseInt(this.state.curType))
				if (this.props.paramTypes[k].val == parseInt(this.state.curType)){
					this.setState({onlyTypeName:this.props.paramTypes[k].name});
					break;
				}
			}
		}
	},
	render: function() {
		console.log(this.props.type,"render")
		return (
			<div>
			 <div className="form-group mt-20">
			      <label>参数类型：</label>
			      {this._renderParamType()}
		    	</div>
	    	   	<div className="form-group">
			      <label >参数名称：</label>
			      {this._renderParamName()}
			    </div>
		      	<div className="form-group">
		     	 <label >参数内容：</label>
		    	 {this._renderParamValue()}
			    </div>
			    <input type="button" className="btn btn-primary" value="提交数据" onClick={this.saveParams}/>
			    {this.props.type == "update"? <input type="button" className="btn btn-primary ml-20" value="返回列表" onClick={this.props.undoUpdate}/>:""}
			</div>
		);
	},

	_renderParamType:function(){
		var curType = this.state.curType
		 if (this.props.onlyEditContent) {
		 	return (
		 		<div>
		 		 <input type="hidden" className="form-control" ref="paramType"   placeholder="参数类型" />
		 		 {this.state.onlyTypeName}
		 		 </div>
		 		);
		 } else {
		 	return (
		 		<select ref="paramType" className="form-control dropdown-toggle"   onChange={this._paramsTypeChange}>
			      	{
			      		this.props.paramTypes.map(function(row,key){
			      			return (
			      			  <option key={key} value={row.val} selected = {row.val == curType}>{row.name}</option>
		      				);
			      		})
			      	}
			      </select>
		 		);
		 }
	},

	_renderParamName:function(){
		if (this.props.onlyEditContent){
			var editKey = this.state.editKey;
			var txtName = editKey != null ? this.props.paramsList[editKey].name : "";
			return (
				<div>
		 		 <input type="hidden" className="form-control" ref="paramName"   placeholder="参数名称" />
		 		  {txtName}
	 		     </div>
			);
		} else {
			return(
				<input type="input" className="form-control" ref="paramName"   placeholder="参数名称" />
			);
		}
	},

	_paramsTypeChange:function(event){
		this.refs.paramType.value = event.target.value;
		this.state.paramsList
		this.setState({curType:event.target.value});
	},

	_renderParamValue:function(){
		if (this.state.curType == "2"){
			return (<input type="password" className="form-control" ref="paramValue"  placeholder="参数内容" />);
		} else if(this.state.curType == "3"){
			return (<textarea className="form-control"  ref="paramValue"  placeholder="参数内容"></textarea> );
		}
		return (<input type="input" className="form-control"  ref="paramValue"  placeholder="参数内容" />);
	},
	saveParams(){
		if (this.props.type == "add" && this.props.addSaveFunc != 'undefined'){
			var func = this.props.addSaveFunc;
			func(this.refs);
			this.refs.paramType.value = "1";
			this.refs.paramName.value = "";
			this.refs.paramValue.value = "";
		}

		if (this.props.type == "update" && this.props.updateSaveFunc != 'undefined'){
			var func = this.props.updateSaveFunc;
			func(this.refs);
		}
	},
});

