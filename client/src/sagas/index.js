import fetch from 'isomorphic-fetch'
import {takeEvery, takeLast} from 'redux-saga'

const 
function* fetchData(kind){

}

export function* watchFetch(action, kind){
	yield* takeLast(action, fetchData(kind))

}