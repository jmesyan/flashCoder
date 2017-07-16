import fetch from 'isomorphic-fetch'
import { take, call, put } from 'redux-saga/effects'
export const REQUEST_LIST = 'REQUEST_LIST';
export const RECEIVE_LIST = 'RECEIVE_LIST';
export const REQUEST_ITEM = 'REQUEST_ITEM';
export const RECEIVE_ITEM = 'RECEIVE_ITEM';
export const CHANGE_ITEM = 'CHANGE_ITEM';
export const REQUEST_ERROR  = 'REQUEST_ERROR';

const Api = "http://localhost:6339/api";

export const requestList = function(kind) {
	return {
		type:REQUEST_LIST,
		kind:kind,
	}
}

export const requestItem = function(kind, params) {
	return {
		type:REQUEST_ITEM,
		kind:kind,
		params:params
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

export const receiveItem = function(kind, res){
	return {
		type:RECEIVE_ITEM,
		kind:kind,
		item:res,
		receiveAt:Date.now()
	}
}

export const changeItem = function(kind, res){
	return {
		type:CHANGE_ITEM,
		kind:kind,
		item:res,
		changeAt:Date.now()
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


export function* watchFetchItem(){
	var itemRequest = yield take(REQUEST_ITEM);
	var kind = itemRequest.kind, params = itemRequest.params;
	var paramstr = '?', paramstr2 = '';
	for (var i in params) {
		paramstr2 += ("&"+i+"="+params[i]);
	}
	paramstr += paramstr2.substring(1);
	var fetchUrl = Api;
	switch (kind) {
		case 'cronItem':
		 	fetchUrl += ('/cron/update'+ paramstr);
	}
	try {
		var res = yield call(fetchApi, fetchUrl);
		yield put(receiveItem(kind, res));
	} catch (err){
		yield put(requestError(err))
	}
}
