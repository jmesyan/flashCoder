import { combineReducers } from 'redux'
import {RECEIVE_LIST, RECEIVE_ITEM, REQUEST_ERROR, CHANGE_ITEM} from '../actions/dataActions';

const lists = function (state = null, action){
	switch(action.type) {
		case RECEIVE_LIST:
			switch (action.kind){
				case 'cronList':
				return {
					...state,
					cronList:action.list,
				}
			}
			return state;
		case REQUEST_ERROR:
			console.log(action.error)
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
			}
			return state;
		case CHANGE_ITEM:
			switch (action.kind){
				case 'cronItem':
				return {
					...state,
					cronItem:action.item,
				}
			}
			return state;
		case REQUEST_ERROR:
			console.log(action.error)
			return state;
		default:
			return state;
	}
		
}



const rootReducer = combineReducers({
  	lists,
  	items
})

export default rootReducer
