import React from 'react'
import '../assets/css/jump.css'
import {connect} from 'react-redux'
import Counter from '../helpers/counter'
class dispatchJump extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            timeId : null
        }
    }

    onComplete(jumpUrl){
        location.href = jumpUrl;
    }

    render() {
        const {params} = this.props;
        var jumpLen = params.jump? params.jump.length: 0;
        var waitTime  = parseInt(params.wait) || (params.success == 1? 1: 3);
        var jumpUrl = jumpLen > 0? params.jump:"javascript:history.back(-1)";
        return (
            <div className="container bg-image mt-20">
                <div className="system-message">
                <h1>:)</h1>
                <p 
                   className={params.success == 1?"glyphicon glyphicon glyphicon-flash success":"glyphicon glyphicon glyphicon-flash error"}
                >
                 {params.message}
                </p>
                <p className="detail"></p>
                <p className="jump">
                页面自动&nbsp;&nbsp;
                <a id="href" href={jumpUrl}>跳转</a> 等待时间： <b id="wait">
                    <Counter value={waitTime} onComplete={this.onComplete.bind(this, jumpUrl)}/>
                </b>
                </p>
                </div>
            </div>
        );
    }
}
const mapStateToProps = state => {
    return state;
  }
  
export default connect(mapStateToProps)(dispatchJump);