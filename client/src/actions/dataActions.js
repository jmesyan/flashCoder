import fetch from 'isomorphic-fetch'
import { take, call, put } from 'redux-saga/effects'
import { takeLatest } from 'redux-saga'
export const REQUEST_LIST = 'REQUEST_LIST';
export const RECEIVE_LIST = 'RECEIVE_LIST';
export const REQUEST_ITEM = 'REQUEST_ITEM';
export const RECEIVE_ITEM = 'RECEIVE_ITEM';
export const REQUEST_ERROR  = 'REQUEST_ERROR';
export const POST_FORM  = 'POST_FORM';

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

export const postForm = function(kind, data){
	return {
		type:POST_FORM,
		kind:kind,
		data:data,
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

function postApi(url, data){
	var formData = new URLSearchParams();  
	for(let k in data){  
		formData.append(k, data[k]);  
	} 
	var request = new Request(url,{
		method: "POST",
		headers: {
			'Content-Type':'application/x-www-form-urlencoded',
		},
		body:formData
	});
	return fetch(request)
	.then(response => response.json() )
	.then(json => json)
}

function* handleForm(action) {
	var kind = action.kind;
	var data = action.data;
	var fetchUrl = Api;
	switch (kind) {
		case 'cronFormUpdate':
			fetchUrl += '/cron/update?crid='+data.Crid;
			break;
	}
	try {
		var res = yield call(postApi, fetchUrl, data);
		alert(res.msg)
	} catch (err){
		yield put(requestError(err))
	}	
}


export function* watchPostForm(){
	yield* takeLatest(POST_FORM, handleForm)
}

export function* watchFetchList(){
	var listRequest = yield take(REQUEST_LIST);
	var kind = listRequest.kind;
	var fetchUrl = Api;
	switch (kind) {
		case 'cronList':
			 fetchUrl += '/cron/list';
			 break;
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
			 break;
	}
	try {
		var res = yield call(fetchApi, fetchUrl);
		yield put(receiveItem(kind, res));
	} catch (err){
		yield put(requestError(err))
	}
}
