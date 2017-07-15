import fetch from 'isomorphic-fetch'
import { take, call, put } from 'redux-saga/effects'
export const REQUEST_LIST = 'REQUEST_LIST';
export const RECEIVE_LIST = 'RECEIVE_LIST';
export const REQUEST_ERROR  = 'REQUEST_ERROR';

const Api = "http://localhost:6339/";

export const requestList = function(kind) {
	return {
		type:REQUEST_LIST,
		kind:kind,
	}
}

export const receiveList = function(kind, res){
	return {
		type:RECEIVE_LIST,
		kind:kind,
		list:res,
		receiveAt:Date.now()
	}
}

export const requestError = function(err){
	return {
		type:REQUEST_ERROR,
		err:err
	}
}

function fetchApi(url){
	return fetch(url)
            .then(response => response.json() )
            .then(json => json)
}



export function* watchFetchList(){
	var listRequest = yield take(REQUEST_LIST);
	var kind = listRequest.kind;
	var fetchUrl = Api;
	switch (kind) {
		case 'cronList':
		 	fetchUrl += '/cron/list';
	}
	try {
		var res = yield call(fetchApi, fetchUrl);
		yield put(receiveList(kind, res));
	} catch (err){
		yield put(requestError(err))
	}
}
