import fetch from 'isomorphic-fetch'
import { take, call, put } from 'redux-saga/effects'
import { takeLatest } from 'redux-saga'
export const REQUEST_LIST = 'REQUEST_LIST';
export const RECEIVE_LIST = 'RECEIVE_LIST';
export const REQUEST_ITEM = 'REQUEST_ITEM';
export const RECEIVE_ITEM = 'RECEIVE_ITEM';
export const REQUEST_ERROR  = 'REQUEST_ERROR';
export const POST_FORM  = 'POST_FORM';
export const REQUEST_EXEC  = 'REQUEST_EXEC';

const Api = "http://localhost:6339/api";

export const requestList = function(kind, params) {
	return {
		type:REQUEST_LIST,
		params:params,
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

export const requestExec = function(kind, params){
	return {
		type:REQUEST_EXEC,
		kind:kind,
		params:params,
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
		if (data[k] instanceof Array) {
			formData.append(k, JSON.stringify(data[k]));
		} else {
			formData.append(k, data[k]);  
		}
		
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
		case 'cronFormAdd':
			fetchUrl += '/cron/add';
			break;
		case 'taskFormAdd':
			fetchUrl += '/task/add';
			break;
		case 'taskBehaviorParams':
			fetchUrl += '/task/taskBehaviorParams?tbid='+data.tbid;
			 break;
		case 'behaviorAdd':
			fetchUrl += '/behavior/add';
			break;
			case 'behaviorUpdate':
			fetchUrl += '/behavior/update?bid='+data.bid;
			break;
		case 'operateFormAdd':
			fetchUrl += '/operate/add';
			break;
	}
	try {
		var res = yield call(postApi, fetchUrl, data);
		if (res.ret > 0) {
			let jumpUrl = '/jump/0/'+encodeURI(res.msg);
			location.href = jumpUrl;
		} else {
			let jumpUrl;
			switch (kind) {
				case 'cronFormUpdate':
				case 'cronFormAdd':
					jumpUrl = '/jump/1/'+encodeURI(res.msg)+'/1/%2fcron%2flist';
					location.href = jumpUrl;
					break;
				case 'taskFormAdd':
					jumpUrl = '/jump/1/'+encodeURI(res.msg)+'/1/%2ftask%2flist';
					location.href = jumpUrl;
					break;
				case 'taskBehaviorParams':
					jumpUrl = '/jump/1/'+encodeURI(res.msg)+'/1/%2ftask%2feditTaskBehaviors%2f'+res.tid;
					location.href = jumpUrl;
					break;
				case 'behaviorAdd':
				case 'behaviorUpdate':
					jumpUrl = '/jump/1/'+encodeURI(res.msg)+'/1/%2fbehavior%2flist';
					location.href = jumpUrl;
					break;
				case 'operateFormAdd':
					jumpUrl = '/jump/1/'+encodeURI(res.msg)+'/1/%2foperate%2flist';
					location.href = jumpUrl;
					break;
			}
		}
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
	var params = listRequest.params;
	var fetchUrl = Api;
	switch (kind) {
		case 'cronList':
			 fetchUrl += '/cron/list';
			 break;
		case 'taskList':
			 fetchUrl += '/task/list';
			 break;
		case 'taskBehaviors':
			 fetchUrl += '/task/editTaskBehaviors?tid='+params.tid;
			 break;
		case 'behaviorList':
			 fetchUrl += '/behavior/list';
			 break;
		case 'operateList':
			 fetchUrl += '/operate/list';
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
		case 'cronAddTasks':
			fetchUrl += '/cron/add';
			break;
		case 'AddTasks':
			fetchUrl += '/task/add';
			break;
		case 'behaviorItem':
			fetchUrl += ('/behavior/update'+ paramstr);
			break;
		case 'jsonOperateList':
			fetchUrl += '/operate/JsonOperateList';
			break;
		case 'taskBehaviorItem':
			fetchUrl += ('/task/taskBehaviorParams'+ paramstr);
			break;
	}
	try {
		var res = yield call(fetchApi, fetchUrl);
		yield put(receiveItem(kind, res));
	} catch (err){
		yield put(requestError(err))
	}
}

export function* watchRequestExec() {
	var itemRequest = yield take(REQUEST_EXEC);
	var kind = itemRequest.kind, params = itemRequest.params;
	var fetchUrl = Api;
	switch (kind) {
		case 'cronUpdateState':
			fetchUrl += '/cron/updateState?crid='+params.Crid;
			break;
		case 'cronDelete':
			fetchUrl += '/cron/delete?crid='+params.Crid;
			break;
		case 'taskDelete':
			fetchUrl += '/task/delete?tid='+params.tid;
			break;
		case 'taskExecute':
			fetchUrl += '/task/taskExecute?tid='+params.tid;
			break;
		case 'behaviorDelete':
			fetchUrl += '/behavior/delete?bid='+params.bid;
			break;
		case 'operateDelete':
			fetchUrl += '/operate/delete?opid='+params.opid;
			break;
	}
	try {
		var res = yield call(fetchApi, fetchUrl);
		if (res.ret > 0) {
			let jumpUrl = '/jump/0/'+encodeURI(res.msg);
			location.href = jumpUrl;
		} else {
			let jumpUrl;
			switch (kind) {
				default:
				  case 'cronUpdateState':
				  case 'cronDelete':
					jumpUrl = '/jump/1/'+encodeURI(res.msg)+'/1/%2fcron%2flist';
					location.href = jumpUrl;
					break;
				  case 'taskDelete':
				  case 'taskExecute':
					jumpUrl = '/jump/1/'+encodeURI(res.msg)+'/1/%2ftask%2flist';
					location.href = jumpUrl;
					break;
				  case 'behaviorDelete':
					jumpUrl = '/jump/1/'+encodeURI(res.msg)+'/1/%2fbehavior%2flist';
					location.href = jumpUrl;
					break;
				  case 'operateDelete':
					jumpUrl = '/jump/1/'+encodeURI(res.msg)+'/1/%2foperate%2flist';
					location.href = jumpUrl;
					break;
			}
		}
	} catch (err){
		yield put(requestError(err))
	}
}
