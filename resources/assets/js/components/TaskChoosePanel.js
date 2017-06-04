var TaskChoosePanel = React.createClass({
	getDefaultProps: function() {
		return {
			itemList:[]
		};
	},
	getInitialState: function() {
		let confirmParams = {
	     	  id:"Confirm",  
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
		      }  
	    };
		return {
			confirmParams: confirmParams,
			chooseList:[],
			insertIndex:0
		};
	},
	componentWillMount: function() {
		if (this.props.itemList.length > 0) {
			var itemChecks = [];
			for (var i in this.props.itemList){
				var id = this.props.itemList[i].Bid;
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
		       	   <input type="hidden" name="basicTaskBehaviors" value={this._jsonToString(this.state.chooseList)}/>
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
		this.props.itemList.map(function(row, key){
			items.push(
			   <div key={key} className="col-xs-6 col-md-4">{row.Bname}<input type="checkbox" value={row.Bid} checked={itemChecks[row.Bid].isChecked} name="behaviors" onChange={()=>itemChange(row.Bname, row.Bid)}/></div>
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
			var length = 3;
			var desc = "确认选择？ 选择事件顺序：<select name='order' refs='itemOrder' onChange={this._itemInsertOrder}>";
			if (length > 0){
				for (var k= 0; k <length;k++){
					desc += ("<option value='"+k+"'>"+(k+1)+"<\/option>");
				}
			}else {
				desc += ("<option value='0'>1<\/option>");
			}
			
			desc += "<\/select>";
			this.state.confirmParams.desc = desc;
			var itemChoose = this._itemChoose;
			var itemCancel = this._itemCancel;
			this.state.confirmParams.leftBtn.func = ()=>itemCancel(itemId);
			this.state.confirmParams.rightBtn.func = ()=>itemChoose(itemName, itemId);
			this.setState({confirmParams:this.state.confirmParams});
			$("#Confirm").modal("show");
		}
		
	},

	_itemInsertOrder:function(){
		console.log(this.refs.itemOrder)
	},

	_renderComposite:function(){
		return (
			""
		);
	},

	_itemCancel:function(itemId){
		 this.state.itemChecks[itemId].isChecked = false;
		 this.setState({itemChecks:this.state.itemChecks});
		 $("#Confirm").modal("hide");

	},

	_itemChoose:function(itemName,itemId){
		this.props.chooseList.push(
			{itemId:itemId,itemName:itemName}
		);

		this._itemCancel(itemId);

	},

	_insert:function(arr, index, item){
		arr.splice(index, 0, item);
	}


});

