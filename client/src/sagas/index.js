import fetch from 'isomorphic-fetch'
import {takeEvery, takeLast} from 'redux-saga'

const ApiUrl = "http://localhost:6339"

function* fetchData(kind){
	var fetchUrl = ApiUrl;
	switch(kind) {
		case 'CRON_LIST':
			fetchUrl +="/cron/index";
			break;
	}

	try {
		var res = yield fetch(fetchUrl);
		return {'ret'=>0, 'res'=>res}
	} catch(error) {
		return {'ret'=>1, 'err'=>error}
	}

}

export function* watchFetch(action, kind){
	yield* takeLast(action, fetchData(kind))

}