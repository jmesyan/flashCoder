import "babel-polyfill"
import React from 'react'
import { render } from 'react-dom'
import { createStore, applyMiddleware,compose } from 'redux'
import DevTools from './helpers/devTools';
import createSagaMiddleware from 'redux-saga';
import { Provider } from 'react-redux'
import { Router, Route, browserHistory, IndexRoute } from 'react-router'
import * as pages from './pages';
import reducer  from './reducers';
import {watchFetchList, watchFetchItem, watchPostForm, watchRequestExec} from './actions/dataActions'
import './assets/css/default.css'
import './assets/css/bootstrap.css'
import './assets/css/bootstrap-theme.css'

const sagaMiddleware = createSagaMiddleware()
const middleware = [ sagaMiddleware ]
var enhancer = applyMiddleware(...middleware);

process.env.NODE_ENV = 'dev'
if (process.env.NODE_ENV == 'dev') {
    enhancer = compose(enhancer, DevTools.instrument());
}

const store = createStore(
  reducer,
  enhancer
)
sagaMiddleware.run(watchFetchList)
sagaMiddleware.run(watchFetchItem)
sagaMiddleware.run(watchPostForm)
sagaMiddleware.run(watchRequestExec)

render(
    (<Provider store={store}>
      	<Router history={browserHistory}>
        	<Route path="/" component={pages.app}>
                <IndexRoute component={pages.cronList}/>
                <Route path="/cron/list" component={pages.cronList}/>
                <Route path="/cron/update/:crid" component={pages.cronUpdate}/>
                <Route path="/cron/add" component={pages.cronAdd}/>
                <Route path="/task/list" component={pages.taskList}/>
                <Route path="/task/add" component={pages.taskAdd}/>
                <Route path="/task/editTaskBehaviors/:tid" component={pages.editTaskBehaviors}/>
                <Route path="/task/taskBehaviorParams/:tbid" component={pages.taskBehaviorParams}/>
                <Route path="/behavior/list" component={pages.behaviorList}/>
                <Route path="/behavior/update/:bid" component={pages.behaviorUpdate}/>
                <Route path="/behavior/add" component={pages.behaviorAdd}/>
                <Route path="/operate/list" component={pages.operateList}/>
                <Route path="/operate/add" component={pages.operateAdd}/>
        	</Route>
            <Route path="/jump/:success/:message(/:wait(/:jump))" component={pages.jump}/>
    	</Router>
    </Provider>),
    document.getElementById("container")
);

