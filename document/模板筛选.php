<?php
//����
public function accumScoreRedpackStat(Request $request) {
	$this->checkFunction("AccumScoreRedpackStat", "edit");
	$uid = $request->get("uid", 0);
	$page = intval($request->get('page'));
	if ($page < 1) $page = 1;
	$pageSize = 50;
	$DATA['title'] = '�û��ۼƷֺͺ��ͳ��';
	$DATA['uid'] = $uid;
	$DATA['page'] = $page;
	$DATA['pageSize'] = $pageSize;
	return view('admin/report/accumScoreRedpackStat', $DATA);	
}