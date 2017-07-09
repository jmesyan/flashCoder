import { createLogger } from 'redux-logger'
import createSagaMiddleware from 'redux-saga'
const sagaMiddleware = createSagaMiddleware()
const middleware = [ sagaMiddleware ]

if (process.env.npm_lifecycle_event == 'dev') {
  middleware.push(createLogger())
}

// const store = createStore(
//   reducer,
//   applyMiddleware(...middleware)
// )

 function component() {
    var element = document.createElement('div');

    // Lodash, now imported by this script
    element.innerHTML = "hello react client";

    return element;
}

 document.body.appendChild(component());