var TaskChoosePanel = React.createClass({
	getDefaultProps: function() {
		return {
			itemList:[]
		};
	},
	getInitialState: function() {
		let confirmParams = {
	     	  id:this.props.dialogName,  
		      title: '提醒',  
		      desc: '确认选择?',  
		      leftBtn: {  
		        text: '取消',
		        type: 'func',
		        func:function(){}  
		      },  
		      rightBtn: {  
		        text: '确认',
		        type:'func',
		        func:function(){}  
		      },
	    };
		return {
			confirmParams: confirmParams,
			chooseList:[],
			insertIndex:0,
			deleteIndex:null,
		};
	},
	componentWillMount: function() {
		if (this.props.itemList.length > 0) {
			var itemChecks = [];
			var idName = this.props.itemIdName;
			for (var i in this.props.itemList){
				var id = this.props.itemList[i][idName];
				itemChecks[id]={isChecked:false};
			}
			this.setState({itemChecks:itemChecks});
		}
	},
	render: function() {
		return (
			<div>
				<div className="panel-body">
		        <div className="container-fluid">
		        	{this._renderItems()}  
				</div>
		      </div>
		       <div className="panel-footer">
		       	   <input type="hidden" name={this.props.saveName} value={this._jsonToString(this.state.chooseList)}/>
			       <div>{this._renderComposite()}</div>
		       </div>
		       <ModelDialog ref="confirm"  {...this.state.confirmParams}/>  
			</div>
		);
	},

	_jsonToString(data){
        return JSON.stringify(data);
    },

	_renderItems:function(){
		var items =[];
		var itemChange = this._itemsChange;
		var itemChecks = this.state.itemChecks;
		var idName = this.props.itemIdName;
		var disName = this.props.itemDisName;
		this.props.itemList.map(function(row, key){
			items.push(
			   <div key={key} className="col-xs-6 col-md-4">{row[disName]}<input type="checkbox" value={row[idName]} checked={itemChecks[row[idName]].isChecked} name="behaviors" onChange={()=>itemChange(row[disName], row[idName])}/></div>
			);
		});
		return items;
	},

	_itemsChange:function(itemName, itemId){
		var isChecked = this.state.itemChecks[itemId].isChecked ? false:true;
		var itemInsertOrder = this._itemInsertOrder;
		this.state.itemChecks[itemId].isChecked = isChecked;
		this.setState({itemChecks:this.state.itemChecks});
		if (isChecked){
			var length = this.state.chooseList.length+1;
			this.state.confirmParams.desc = <OrderSelect chooseLength={length} itemInsertOrder={itemInsertOrder} />;
			var itemChoose = this._itemChoose;
			var itemCancel = this._itemCancel;
			this.state.confirmParams.leftBtn.func = ()=>itemCancel(itemId);
			this.state.confirmParams.rightBtn.func = ()=>itemChoose(itemName, itemId);
			this.setState({confirmParams:this.state.confirmParams});
			$("#"+this.props.dialogName).modal("show");
		}
		
	},

	_itemInsertOrder:function(order){
		this.setState({insertIndex:order})
	},

	_deleteChoose:function(){
		var deleteIndex = this.state.deleteIndex;
		if (deleteIndex != null) {
			this.state.chooseList.splice(deleteIndex,1);
		}	
		this.setState({chooseList:this.state.chooseList, deleteIndex:null});
		$("#"+this.props.dialogName).modal("hide");
	},

	_deleteCancel:function(){
		this.setState({deleteIndex:null});
		$("#"+this.props.dialogName).modal("hide");
	},

	_itemDelete:function(order){
		var deleteChoose = this._deleteChoose;
		var deleteCancel = this._deleteCancel;
		this.state.confirmParams.desc = "确认删除该事件？";
		this.state.confirmParams.leftBtn.func = ()=>deleteCancel();
		this.state.confirmParams.rightBtn.func = ()=>deleteChoose();
		this.setState({confirmParams:this.state.confirmParams, deleteIndex:order});
		$("#"+this.props.dialogName).modal("show");
	},

	_renderComposite:function(){
		var composites = [];
		var itemDelete = this._itemDelete;
		this.state.chooseList.map(function(row,key){
			composites.push(<span key={key}><a href="javascript:;" onClick={()=>itemDelete(key)} >{row.itemName}</a>=></span>);
		});
		return composites;
	},

	_itemCancel:function(itemId){
		 this.state.itemChecks[itemId].isChecked = false;
		 this.setState({itemChecks:this.state.itemChecks});
		$("#"+this.props.dialogName).modal("hide");
		 this.setState({insertIndex:0})

	},

	_itemChoose:function(itemName,itemId){
		this._insert(this.state.chooseList, this.state.insertIndex, {itemId:itemId,itemName:itemName});
		this.setState({chooseList:this.state.chooseList});
		this._itemCancel(itemId);
	},

	_insert:function(arr, index, item){
		arr.splice(index, 0, item);
	}


});

TaskChoosePanel.propTypes = {  
  saveName: React.PropTypes.string.isRequired,  
  itemIdName: React.PropTypes.string.isRequired,  
  itemDisName: React.PropTypes.string.isRequired,  
  dialogName: React.PropTypes.string.isRequired,  
};  

