var React = require('react');

var TableWithHandle = React.createClass({

	_renderRow: function(row, disrow){
		let res =[];
		disrow.map((dis,k)=>{
			if (dis.name) {
				if (typeof(dis.filters) != 'undefined' && dis.filters.length>0){
					var item = row[dis.name];
					for (var i in dis.filters) {
						item = eval(dis.filters[i])(item)
					}
					res.push(
					<td key={k}>{item}</td>
					);
				} else {
					res.push(
					<td key={k}>{row[dis.name]}</td>
					);
				}
 				
			} else if (typeof(dis.content) != 'undefined') {
				res.push(
					<td key={k}>{dis.content}</td>
				);
			} else if (typeof(dis.handle) != 'undefined') {
				res.push(eval(dis.handle)(row, k));
			}
		});
		return res;
	},

	render: function() {
		var titles = this.props.titles? this.props.titles: [];
		var list = this.props.list? eval(this.props.list): [];
		return (
		<table className="table table-striped">
			<tbody>
	 		<tr>
	 			{
	 				titles.map(function(title,k){
	 					return <th key={k}>{title}</th>
	 				})
	 			}
			</tr>
			{
				list.map((row, k)=>{
					return <tr key={k}>{this._renderRow(row,this.props.disrow)}</tr>;
				})
			}
			</tbody>
		 </table>
		);
	},

});

module.exports = TableWithHandle;

