import React from 'react';
export default class Counter extends React.Component {
    // 定义属性
    static propTypes={
       
        onStep:React.PropTypes.func,
        onComplete:React.PropTypes.func,
        value:React.PropTypes.number,
        step:React.PropTypes.number
    }
 
 
    //这里面的操作可以移动到componentWillMount()里面去
    constructor(...pa) {
        super(...pa);
        this.initValue = this.props.value || 10;
        this.state     = {count: this.initValue}
        this.interval  = 0;
        this.step      = this.props.step || 1;
       
 
    }
    
    stop() {
        clearInterval(this.interval);
    }
    
    start() {
        this.stop();
        this.interval = setInterval(()=> {
            var count = this.state.count - this.step;
            if (this.props.onStep) {
                this.props.onStep(count);
            }
            if (count ==0) {
                this.props.onComplete && this.props.onComplete();
                this.stop();
            }else{
                this.setState({count});
            }
           
        }, 1000);
    }
    
    restart() {
        this.stop();
        this.setState({count: this.initValue});
        this.start();
    }
    componentDidMount(){
        this.start();
    }
    componentWillUnmount(){
        this.stop();
    }
    
    render() {
        return (<span>{this.state.count}</span>)
    }
}