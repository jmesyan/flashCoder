import React from 'react'
import { render } from 'react-dom'
import { createStore, applyMiddleware } from 'redux'
import { createLogger } from 'redux-logger';
import createSagaMiddleware from 'redux-saga';
import { Provider } from 'react-redux'
import { Router, Route, browserHistory, IndexRoute } from 'react-router'
import {app, cronList} from './pages';
import reducer  from './reducers';
import './css/default.css'
import './css/bootstrap.css'
import './css/bootstrap-theme.css'

const sagaMiddleware = createSagaMiddleware()
const middleware = [ sagaMiddleware ]

if (process.env.NODE_ENV == 'dev') {
  middleware.push(createLogger())
}

const store = createStore(
  reducer,
  applyMiddleware(...middleware)
)

render(
    (<Provider store={store}>
      	<Router history={browserHistory}>
        	<Route path="/" component={app}>
            <IndexRoute component={cronList}/>
            <Route path="/cron/list" component={cronList}/>
        	</Route>
    	</Router>
    </Provider>),
    document.getElementById("container")
);

