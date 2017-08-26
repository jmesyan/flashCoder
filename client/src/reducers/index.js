import { combineReducers } from 'redux'
import {RECEIVE_LIST, RECEIVE_ITEM, REQUEST_ERROR} from '../actions/dataActions';
import reduceForm from './form.js'
import { actions } from 'react-redux-form';

const lists = function (state = null, action){
	switch(action.type) {
		case RECEIVE_LIST:
			switch (action.kind){
				case 'cronList':
				return {
					...state,
					cronList:action.list,
				}
				case 'taskList':
				return {
					...state,
					taskList:action.list,
				}
				case 'taskBehaviors':
				return {
					...state,
					taskBehaviors:action.list,
				}
				case 'behaviorList':
				return {
					...state,
					behaviorList:action.list,
				}
				case 'operateList':
				return {
					...state,
					operateList:action.list,
				}
			}
			return state;
		case REQUEST_ERROR:
			console.log(action.error)
			alert(action.error)
			return state;
		default:
			return state;
	}
		
}

const items = function (state = null, action){
	switch(action.type) {
		case RECEIVE_ITEM:
			switch (action.kind){
				case 'cronItem':
				return {
					...state,
					cronItem:action.item,
				}
				case 'cronAddTasks':
				return {
					...state,
					cronAddTasks:action.item,
				}
				case 'cronAddTasks':
				return {
					...state,
					cronAddTasks:action.item,
				}
				case 'AddTasks':
				return {
					...state,
					AddTasks:action.item,
				}

				case 'behaviorItem':
				return {
					...state,
					behaviorItem:action.item,
				}
				
				case 'jsonOperateList':
				return {
					...state,
					jsonOperateList:action.item,
				}
				case 'taskBehaviorItem':
				return {
					...state,
					taskBehaviorItem:action.item,
				}		

			}
			return state;
		case REQUEST_ERROR:
			console.error(action.error)
			return state;
		default:
			return state;
	}
		
}



const rootReducer = combineReducers({
  	lists,
	items,
	...reduceForm
})

export default rootReducer
