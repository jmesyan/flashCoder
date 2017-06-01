var ModelDialog = React.createClass({
  render() {  
    const props = this.props;  
    return (  
	  <div className="modal fade" id={props.id} tabIndex="-1" role="dialog">
	  	<div className="modal-dialog" role="document">
	    <div className="modal-content">
	      <div className="modal-header">
	        <button type="button" className="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
	        <h4 className="modal-title">{props.title}</h4>
	      </div>
	      <div className="modal-body">
	        <p>{props.desc}</p>
	      </div>
	      <div className="modal-footer">
	      { this.props.leftBtn.type == 'close' ? <button type="button" className="btn btn-primary" data-dismiss="modal">{props.leftBtn.text}</button>: <button type="button" className="btn btn-primary" onClick={this.props.leftBtn.func.bind(this)}>{props.leftBtn.text}</button>}
	      { this.props.rightBtn.type == 'close' ? <button type="button" className="btn btn-primary" data-dismiss="modal">{props.rightBtn.text}</button>: <button type="button" className="btn btn-primary" onClick={this.props.rightBtn.func.bind(this)}>{props.rightBtn.text}</button>}
	      </div>
	    </div>
	  </div>
	</div>
    );  
  }  
})  
  
ModelDialog.propTypes = {  
  title: React.PropTypes.string.isRequired,  
  desc: React.PropTypes.string.isRequired, 
};  