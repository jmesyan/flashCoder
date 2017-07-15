import { combineReducers } from 'redux'
import {REQUEST_LIST, RECEIVE_LIST, REQUEST_ERROR} from '../actions/dataActions';

const lists = function (state = null, action){
	switch(action.type) {
		case REQUEST_LIST:
			return {
				... state,
				fetched:false,
			}
		case RECEIVE_LIST:
			switch (action.kind){
				case 'cronList':
				return {
					...state,
					fetched:true,
					cronList:action.list,
				}
			}
			return state;
		case REQUEST_ERROR:
			 return {
			 	...state,
			 	fetchingError:true,
			 	error:action.error
			 }
		default:
			return state;
	}
		
}


const rootReducer = combineReducers({
  	lists
})

export default rootReducer
