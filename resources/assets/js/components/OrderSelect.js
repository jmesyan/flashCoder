
var OrderSelect = React.createClass({
	getDefaultProps: function() {
		return {
			chooseLength:0
		};
	},

	render: function() {
		return (
			<div>
				<span>确认选择？ 选择事件插入顺序：</span>
				<select name="order" ref="order"  onChange={this._getItemOrder}>
					{this._renderOrder()}
				</select>
			</div>
		);
	},

	_getItemOrder:function(){
		var order = this.refs.order.value;
		this.props.itemInsertOrder(order);
	},

	_renderOrder:function(){
		var items = [];
		for (var k= 0; k <this.props.chooseLength;k++){
			items.push(<option key={k} value={k}>{k+1}</option>);
		}
		return items;
	}

});