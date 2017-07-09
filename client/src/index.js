import React from 'react'
import { render } from 'react-dom'
import { createStore, applyMiddleware } from 'redux'
import { createLogger } from 'redux-logger';
import createSagaMiddleware from 'redux-saga';
import { Provider } from 'react-redux'
import { BrowserRouter, Route } from 'react-router-dom'
import {cronList} from './pages';
import reducer  from './reducers';
import './css/default.css'

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
      	<BrowserRouter>
        	<Route path="/" component={cronList} />
    	</BrowserRouter>
    </Provider>),
    document.getElementById("container")
);

