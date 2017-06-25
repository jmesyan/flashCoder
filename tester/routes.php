<?php

Route::get('/', function(){});

Route::group(['prefix'=>'qw', 'namespace'=>'QW', 'middleware'=>['api']], function () {
	Route::get('/v1/account_auth', 'QWController@account_auth');
	Route::get('/v1/exchange', 'QWController@exchange');
	Route::get('/v1/exchangecoin', 'QWController@exchangecoin');
	Route::get('/v1/exchangeout', 'QWController@exchangeout');
	Route::get('/v1/order', 'QWController@order');
});

Route::group(['prefix'=>'report', 'namespace'=>'Game', 'middleware'=>['api']], function () {
	Route::get('/payment', 'ReportController@payment');
	Route::get('/servers', 'ReportController@servers');
});

Route::group(['prefix'=>'mobile', 'namespace'=>'Game', 'middleware'=>['api']], function () {
	//friend invit award
	Route::get('/friendInvite', 'ActivityController@friendInvite');
	Route::get('/inviteFriendInfo', 'ActivityController@inviteFriendInfo');
	Route::get('/award_pool', 'ActivityController@award_pool');
	Route::get('/inviteAwardRecieve', 'ActivityController@inviteAwardRecieve');

	Route::get('/online', 'MobileController@online');
	Route::get('/time', 'MobileController@Time');
	Route::get('/authorize', 'MobileController@Auth');
	Route::get('/ad', 'MobileController@ad');
	Route::get('/accountLogin', 'MobileController@accountLogin');
	Route::get('/accountInfo', 'MobileController@accountInfo');
	
	Route::get('/shop', 'PayController@Shop');
	Route::get('/shopPaykey', 'PayController@ShopPaykey');
	Route::get('/appstoreCallback', 'PayController@AppStoreCallback');
	Route::post('/appstoreCallback', 'PayController@AppStoreCallback');
	Route::get('/googleCallback', 'PayController@GoogleCallback');
	Route::get('/alipayCallback', 'PayController@AlipayCallback');
	Route::get('/alipayNotify', 'PayController@AlipayNotify');
	Route::get('/payConfirm', 'PayController@PayConfirm');
	Route::get('/setPayChannel', 'PayController@setPayChannel');
	Route::get('/mycardRequest', 'PayController@mycardRequest');
	Route::any('/codapayRequest', 'PayController@codapayRequest');
	Route::any('/mycardCallback', 'PayController@mycardCallback');
	
	Route::any('/game9Callback', 'PayController@game9Callback');
	Route::any('/codapayCallback', 'PayController@codapayCallback');
	
	Route::get('/config', 'ConfigController@Config');
	Route::get('/stonesToGolds', 'UserController@StonesToGolds');
	Route::get('/dayCheckin', 'MobileController@DayCheckIn');
	Route::get('/dayCheckin15', 'MobileController@DayCheckIn15');
	Route::get('/userAvatar', 'UserController@SCAvatar');
	Route::get('/texasAvatar', 'UserController@TexasAvatar');
	Route::get('/rateApp', 'UserController@RateApp');
	Route::get('/userRename', 'UserController@Rename');
	Route::get('/editSign', 'UserController@EditSign');
	Route::get('/editLocale', 'UserController@EditLocale');
	Route::get('/userAchievement', 'UserController@Achievement');
	Route::get('/userAchieve', 'UserController@Achieve');
	Route::post('/editAvatar', 'UserController@UploadAvatar');
	Route::get('/checkProp', 'UserController@CheckProp');
	Route::get('/getChips', 'MobileController@GetChips');
	Route::get('/token', 'UserController@Token');
	Route::get('/ucpay', 'UcpayController@UcCallBack');

	//Route::get('/agentList', 'UserController@AgentList');
	//Route::get('/editAgent', 'UserController@EditAgent');
	//Route::get('/applyAgent', 'UserController@ApplyAgent');
	
	Route::get('/goldsRank', 'UserController@GoldsRank');
	Route::get('/mailList', 'MailController@Mails');
	Route::get('/editAvatar2', 'UserController@EditAvatar');
	
	Route::get('/mailNum', 'MailController@MailNum');
	Route::get('/mailRead', 'MailController@MailRead');
	Route::get('/mailDelete', 'MailController@MailDelete');
	Route::get('/mailSend', 'MailController@MailSend');
	Route::get('/mailSave', 'MailController@MailSave');

	Route::get('/friendList', 'UserController@Friends');
	Route::get('/friendDelete', 'UserController@DeleteFriends');
	Route::get('/friendAdd', 'UserController@AddFriends');
	Route::get('/remindeOne', 'UserController@remindeOne');
	Route::get('/sendAll', 'UserController@sendAll');
	Route::get('/colectAll', 'UserController@colectAll');
	Route::get('/invitQuestList', 'UserController@invitQuestList');
	Route::get('/upFriendState', 'UserController@upFriendState');

	Route::get('/bankInfo', 'UserController@BankInfo');
	Route::get('/bankGet', 'UserController@BankGet');
	Route::get('/bankSave', 'UserController@BankSave');
	Route::get('/giftShop', 'UserController@GiftShop');
	Route::get('/giftExchange', 'UserController@GiftExchange');
	Route::get('/exchangeStones', 'ActivityController@exchangeStones');
	Route::get('/myExchangeGift', 'UserController@myExchangeGift');
	//Route::get('/stonesGolds', 'UserController@StonesToGolds');
	Route::get('/gameTips', 'UserController@GameTips');
	Route::get('/testUsers', 'UserController@TestUsers');
	
	Route::get('/questionConfig', 'QuestionController@Config');
	Route::post('/postQuestion', 'QuestionController@PostQuestion');
	Route::get('/questionList', 'QuestionController@QuestionList');
	Route::get('/questionComplete', 'QuestionController@QuestionComplete');
	Route::get('/questionAnswerList', 'QuestionController@QuestionAnswerList');
	Route::post('/questionAnswerPost', 'QuestionController@QuestionAnswerPost');
	Route::get('/messageCount', 'QuestionController@MessageCount');
	Route::get('/gameHelp', 'QuestionController@gameHelp');
	
	Route::get('/lotteryList', 'ActivityController@LotteryList');
	Route::get('/richList', 'ActivityController@RichList');
	Route::get('/dailyAddRec', 'ActivityController@dailyAddRec');
	Route::get('/dailyResult', 'ActivityController@dailyResult');
	Route::get('/yestodayWinList', 'ActivityController@YestodayWinList');
	Route::get('/winList', 'ActivityController@winList');
	Route::get('/getBigWinBonus', 'ActivityController@getBigWinBonus');
	Route::get('/bonusTimeStart', 'ActivityController@bonusTimeStart');
	Route::get('/bonusTimeEnd', 'ActivityController@bonusTimeEnd');
	Route::get('/bonusTimeGet', 'ActivityController@bonusTimeGet');
	Route::get('/sevenDayBonusGet', 'ActivityController@sevenDayBonusGet');
	Route::get('/dealer', 'ActivityController@dealer');
	
	Route::get('/getFileList', 'MobileController@getFileList');
	Route::get('/getFileList2', 'MobileController@getFileList2');
	Route::get('/test', 'MobileController@Test');
	Route::get('/getFiles', 'MobileController@getFiles');
	
	Route::get('/texas/vipConfig', 'TexasController@vipConfig');
	Route::get('/texas/userVip', 'TexasController@userVip');
	Route::get('/texas/getVipList', 'TexasController@getVipList');
	Route::get('/texas/getVipDailyBonus', 'TexasController@getVipDailyBonus');
	Route::get('/texas/buyVip', 'TexasController@buyVip');
	Route::get('/texas/activitySpec', 'TexasController@activitySpec');
	Route::get('/texas/getUserInfo', 'TexasController@getUserInfo');
	Route::get('/firendList', 'UserController@firendList');
	Route::get('/searchUser', 'UserController@searchUser');
	Route::get('/addFriend', 'UserController@addFriend');
	Route::get('/handleFriendQuest', 'UserController@handleFriendQuest');
	Route::get('/texas/goldsTopList', 'TexasController@goldsTopList');
	Route::get('/texas/weekProList', 'TexasController@weekProList');
	Route::get('/texas/friendGoldsList', 'TexasController@friendGoldsList');
	Route::get('/texas/sharkList', 'TexasController@sharkList');
	Route::get('/texas/luckyCarousel', 'TexasController@luckyCarousel');
	Route::get('/texas/luckyCarouselList', 'TexasController@luckyCarouselList');
	Route::get('/texas/userAddress', 'TexasController@userAddress');
	Route::get('/texas/addAddress', 'TexasController@addAddress');
	Route::get('/texas/getSubsidyGolds', 'TexasController@getSubsidyGolds');
	Route::get('/texas/getFreeGolds', 'TexasController@getFreeGolds');
	
	Route::get('/eggConfig', 'ActivityController@eggConfig');
	Route::get('/getEggAward', 'ActivityController@getEggAward');
	
	Route::get('/fbInviteCallback', 'MobileController@fbInviteCallback');
	
	Route::get('/getVipBonus', 'ActivityController@getVipBonus');
	Route::get('/getVipBonus2', 'ActivityController@getVipBonus2');
	
	Route::get('/specialOrder', 'ActivityController@specialOrder');
	Route::get('/addSpecialOrder', 'ActivityController@addSpecialOrder');
	
	Route::get('/mttList', 'MTTController@mttList');
	Route::get('/mttApply', 'MTTController@mttApply');
	Route::get('/mttUnApply', 'MTTController@mttUnApply');
	
	Route::get('/editInfo', 'ActivityController@editInfo');
	
	Route::get('/subsidyGolds', 'ActivityController@subsidyGolds');
	Route::get('/luckyFruitTreeConfig', 'ActivityController@luckyFruitTreeConfig');
	Route::get('/receiveTree', 'ActivityController@receiveTree');
	Route::get('/handleTree', 'ActivityController@handleTree');
	Route::get('/getTreePrize', 'ActivityController@getTreePrize');
	Route::get('/logDevice', 'MobileController@logDevice');
	Route::get('/setSecretTime', 'ActivityController@setSecretTime');
	
	Route::get('/vipBonusConfig', 'ActivityController@vipBonusConfig');
	
	Route::get('/sendRedPack', 'UserController@sendRedPack');
	Route::get('/sendBlessing', 'UserController@sendBlessing');
	Route::get('/redPackList', 'UserController@redPackList');
	Route::get('/receiveRedPack', 'UserController@receiveRedPack');
	Route::get('/receiveBlessing', 'UserController@receiveBlessing');
	Route::get('/props', 'ActivityController@props');

	Route::get('/sendCode', 'UserController@sendCode');
	Route::get('/bindPhone', 'UserController@bindPhone');
	Route::get('/changePassword', 'UserController@changePassword');
	Route::get('/loseUserCount', 'MobileController@loseUserCount');
	Route::get('/inviteCodeValidate', 'UserController@inviteCodeValidate');
	Route::get('/goodHandList', 'ActivityController@goodHandList');
	Route::get('/prizeList', 'ActivityController@prizeList');
	Route::get('/getDailyCard', 'ActivityController@getDailyCard');
	Route::get('/props', 'ActivityController@props');

	Route::get('/dailiLaba', 'ActivityController@dailiLaba');
	Route::get('/dailiLaba2', 'ActivityController@dailiLaba2');
	//家族
	Route::get('/familyConfig', 'FamilyController@familyConfig');
	Route::get('/createFamily', 'FamilyController@createFamily');
	Route::get('/searchFamily', 'FamilyController@searchFamily');
	Route::get('/applyFamily', 'FamilyController@applyFamily');
	Route::get('/approveFamily', 'FamilyController@approveFamily');
	Route::get('/getFamilyGolds', 'FamilyController@getFamilyGolds');
	Route::get('/familyList', 'FamilyController@familyList');
	Route::get('/quitFamily', 'FamilyController@quitFamily');
	Route::get('/familyMember', 'FamilyController@familyMember');
	Route::get('/offlineMessage', 'FamilyController@offlineMessage');
	Route::get('/getApplyList', 'FamilyController@getApplyList');
	Route::get('/clearOfflineChat', 'FamilyController@clearOfflineChat');
	Route::get('/kickMember', 'FamilyController@kickMember');
	Route::get('/dissolveFamily', 'FamilyController@dissolveFamily');
	Route::get('/codeToGold', 'UserController@codeToGold');
	Route::get('/hammerPig', 'ActivityController@hammerPig');
	Route::get('/hammerNum', 'ActivityController@hammerNum');
	Route::get('/zhuanPan', 'ActivityController@zhuanPan');
	Route::get('/getUserById', 'UserController@getUserById');
});

Route::group(['prefix'=>'crontab', 'namespace'=>'Crontab', 'middleware'=>['api']], function () {	
	Route::get('/reportDay', 'CrontabController@reportDay');
	Route::get('/reportOnline', 'CrontabController@reportOnline');
	Route::get('/addDaily', 'CrontabController@addDaily');
	这是添加到路由组的内容
});
		
Route::group(['prefix'=>'admin', 'namespace'=>'Admin', 'middleware'=>['web']], function () {
	Route::get('/mobile', 'SysController@mobile');
	Route::get('/', 'SysController@Index');
	Route::get('/index', 'SysController@Index');
	Route::get('/login', 'SysController@LoginView');
	Route::get('/logout', 'SysController@Logout');
	Route::post('/login', 'SysController@LoginPost');
	Route::get('/main', 'SysController@Main');
	Route::get('/modifyPassword', 'SysController@modifyPassword');
	Route::post('/modifyPasswordPost', 'SysController@modifyPasswordPost');
	Route::get('/sys/editpassword', 'SysController@editpassword');
	Route::post('/sys/editpassword_post', 'SysController@editpassword_post');
	
	Route::get('/sys/user_control__', 'SysController@user_control__');
	Route::post('/sys/user_ttt__', 'SysController@user_ttt__');
	Route::post('/sys/user_cg__', 'SysController@user_cg__');
	
	Route::get('/sys/func_list', 'SysController@func_list');
	Route::get('/sys/func_op', 'SysController@func_op');
	Route::post('/sys/func_post', 'SysController@func_post');
	Route::get('/sys/func_delete', 'SysController@func_delete');
	
	Route::get('/sys/app_op', 'SysController@app_op');
	Route::post('/sys/app_post', 'SysController@app_post');
	Route::get('/sys/app_delete', 'SysController@app_delete');
	
	Route::get('/sys/role_list', 'SysController@role_list');
	Route::get('/sys/role_op', 'SysController@role_op');
	Route::post('/sys/role_post', 'SysController@role_post');
	Route::get('/sys/role_delete', 'SysController@role_delete');
	
	Route::get('/sys/admin_list', 'SysController@admin_list');
	Route::get('/sys/admin_op', 'SysController@admin_op');
	Route::post('/sys/admin_post', 'SysController@admin_post');
	Route::get('/sys/admin_delete', 'SysController@admin_delete');
	
	Route::get('/sys/user_list', 'SysController@user_list');
	Route::get('/sys/user_op', 'SysController@user_op');
	Route::post('/sys/user_post', 'SysController@user_post');
	Route::get('/sys/user_delete', 'SysController@user_delete');
	
	Route::get('/sys/admin_check', 'SysController@admin_check');
	Route::get('/sys/user_check', 'SysController@user_check');
	Route::get('/sys/admin_checkFunc', 'SysController@admin_checkFunc');
	Route::post('/sys/user_addvip', 'SysController@user_addvip');
	Route::post('/sys/user_addmoney', 'SysController@user_addmoney');
	Route::post('/sys/agent_addmoney', 'SysController@agent_addmoney');
	Route::get('/sys/unforbid', 'SysController@unforbid');
	Route::post('/sys/user_paygolds', 'SysController@user_paygolds');
	
	Route::get('/setting/texas_rand_motor_avatar', 'SettingController@texas_rand_motor_avatar');
	Route::get('/setting/rand_motor_level', 'SettingController@rand_motor_level');
	Route::get('/setting/flash_config', 'SettingController@flash_config');
	Route::post('/setting/flash_post', 'SettingController@flash_post');
	Route::get('/setting/config', 'SettingController@config');
	Route::get('/setting/configDetail', 'SettingController@configDetail');
	Route::get('/setting/special_config2', 'SettingController@special_config2');
	Route::get('/setting/slot_room_schips', 'SettingController@slot_room_schips');
	Route::post('/setting/config_post', 'SettingController@config_post');
	Route::post('/setting/special_config_post', 'SettingController@special_config_post');
	Route::get('/setting/config_down', 'SettingController@config_down');
	Route::get('/setting/config_view17', 'SettingController@config_view17');
	Route::get('/setting/configDetail', 'SettingController@configDetail');
	Route::get('/setting/game_config_list', 'SettingController@game_config_list');
	Route::get('/setting/game_config_export', 'SettingController@game_config_export');
	Route::get('/setting/game_config_import', 'SettingController@game_config_import');
	Route::post('/setting/game_config_post', 'SettingController@game_config_post');
	Route::post('/setting/game_config_post2', 'SettingController@game_config_post2');
	Route::get('/setting/game_config_view', 'SettingController@game_config_view');
	
	Route::get('/setting/roomstat_list', 'SettingController@roomstat_list');
	Route::get('/setting/roomstat_get', 'SettingController@roomstat_get');
	Route::post('/setting/roomstat_post', 'SettingController@roomstat_post');
	Route::post('/setting/roomstat_post2', 'SettingController@roomstat_post2');
	Route::get('/setting/roomstat_log', 'SettingController@roomstat_log');
	
	Route::get('/setting/task_list', 'SettingController@task_list');
	Route::get('/setting/task_op', 'SettingController@task_op');
	Route::post('/setting/task_post', 'SettingController@task_post');
	Route::get('/setting/task_delete', 'SettingController@task_delete');
	
	Route::get('/setting/table_task_list', 'SettingController@table_task_list');
	Route::get('/setting/table_task_op', 'SettingController@table_task_op');
	Route::post('/setting/table_task_post', 'SettingController@table_task_post');
	Route::get('/setting/table_task_delete', 'SettingController@table_task_delete');
	
	Route::get('/setting/dumb_list', 'SettingController@dumb_list');
	Route::post('/setting/dumb_post', 'SettingController@dumb_post');
	Route::get('/setting/dumb_update', 'SettingController@dumb_update');
	Route::post('/setting/dumb_release', 'SettingController@dumb_release');
	Route::post('/setting/dumb_release_all', 'SettingController@dumb_release_all');
	
	Route::get('/setting/currency_list', 'SettingController@currency_list');
	Route::get('/setting/currency_op', 'SettingController@currency_op');
	Route::post('/setting/currency_post', 'SettingController@currency_post');
	Route::get('/setting/currency_delete', 'SettingController@currency_delete');
	
	Route::get('/setting/room_list', 'SettingController@room_list');
	Route::get('/setting/room_op', 'SettingController@room_op');
	Route::post('/setting/room_post', 'SettingController@room_post');
	Route::get('/setting/room_delete', 'SettingController@room_delete');
	
	Route::get('/setting/fbpay_list', 'SettingController@fbpay_list');
	Route::get('/setting/fbpay_op', 'SettingController@fbpay_op');
	Route::post('/setting/fbpay_post', 'SettingController@fbpay_post');
	Route::get('/setting/fbpay_delete', 'SettingController@fbpay_delete');
	
	Route::get('/setting/vip_list', 'SettingController@vip_list');
	Route::get('/setting/vip_op', 'SettingController@vip_op');
	Route::post('/setting/vip_post', 'SettingController@vip_post');
	Route::get('/setting/vip_delete', 'SettingController@vip_delete');
	
	Route::get('/setting/props_list', 'SettingController@props_list');
	Route::get('/setting/props_op', 'SettingController@props_op');
	Route::post('/setting/props_post', 'SettingController@props_post');
	Route::get('/setting/props_delete', 'SettingController@props_delete');
	
	Route::get('/setting/mail_list', 'SettingController@mail_list');
	Route::get('/setting/mail_op', 'SettingController@mail_op');
	Route::post('/setting/mail_post', 'SettingController@mail_post');
	Route::get('/setting/mail_delete', 'SettingController@mail_delete');
	
	Route::get('/setting/laba_list', 'SettingController@laba_list');
	Route::get('/setting/laba_op', 'SettingController@laba_op');
	Route::post('/setting/laba_post', 'SettingController@laba_post');
	Route::get('/setting/laba_delete', 'SettingController@laba_delete');
	Route::post('/setting/laba_send', 'SettingController@laba_send');
	
	Route::get('/setting/block_ip', 'SettingController@block_ip');
	Route::get('/setting/block_ip_op', 'SettingController@block_ip_op');
	Route::post('/setting/block_ip_post', 'SettingController@block_ip_post');
	Route::get('/setting/block_ip_delete', 'SettingController@block_ip_delete');
	
	Route::get('/setting/game_type', 'SettingController@game_type');
	Route::get('/setting/game_type_op', 'SettingController@game_type_op');
	Route::post('/setting/game_type_post', 'SettingController@game_type_post');
	Route::get('/setting/game_type_delete', 'SettingController@game_type_delete');
	
	Route::get('/setting/activity', 'SettingController@activity');
	Route::get('/setting/activity_op', 'SettingController@activity_op');
	Route::post('/setting/activity_post', 'SettingController@activity_post');
	Route::get('/setting/activity_delete', 'SettingController@activity_delete');
	
	Route::get('/setting/motor', 'SettingController@motor');
	Route::get('/setting/motor_list', 'SettingController@motor_list');
	Route::get('/setting/motor0_list', 'SettingController@motor0_list');
	Route::get('/setting/motor1_list', 'SettingController@motor1_list');
	Route::post('/setting/motor_add_type1', 'SettingController@motor_add_type1');
	Route::post('/setting/motor_add_type2', 'SettingController@motor_add_type2');
	Route::post('/setting/motor_import', 'SettingController@motor_import');
	Route::post('/setting/motor_reset_golds', 'SettingController@motor_reset_golds');
	Route::post('/setting/motor_ttt', 'SettingController@motor_ttt');
	Route::post('/setting/motor_delete_type', 'SettingController@motor_delete_type');
	Route::post('/setting/motor_release', 'SettingController@motor_release');
	Route::post('/setting/motor_release_all', 'SettingController@motor_release_all');
	Route::get('/setting/motor2_list', 'SettingController@motor2_list');
	Route::post('/setting/motor_reload_type2', 'SettingController@motor_reload_type2');
	Route::get('/setting/motor_op', 'SettingController@motor_op');
	Route::post('/setting/motor_post', 'SettingController@motor_post');
	Route::get('/setting/motor_delete', 'SettingController@motor_delete');
	
	Route::get('/setting/agent_list', 'SettingController@agent_list');
	Route::get('/setting/agent_op', 'SettingController@agent_op');
	Route::post('/setting/agent_post', 'SettingController@agent_post');
	Route::get('/setting/agent_delete', 'SettingController@agent_delete');
	
	Route::get('/setting/cache_list', 'SettingController@cache_list');
	Route::post('/setting/cache_post', 'SettingController@cache_post');
	
	Route::get('/setting/gate_list', 'SettingController@gate_list');
	Route::get('/setting/gate_op', 'SettingController@gate_op');
	Route::post('/setting/gate_post', 'SettingController@gate_post');
	Route::get('/setting/gate_delete', 'SettingController@gate_delete');
	
	Route::get('/setting/feed_list', 'SettingController@feed_list');
	Route::get('/setting/feed_op', 'SettingController@feed_op');
	Route::post('/setting/feed_post', 'SettingController@feed_post');
	Route::get('/setting/feed_delete', 'SettingController@feed_delete');
	
	Route::get('/setting/forbid_list', 'SettingController@forbid_list');
	Route::post('/setting/forbid_post', 'SettingController@forbid_post');
	Route::post('/setting/forbid_release', 'SettingController@forbid_release');
	Route::get('/setting/forbid_release_all', 'SettingController@forbid_release_all');
	
	Route::get('/setting/gift_category_list', 'SettingController@gift_category_list');
	Route::get('/setting/gift_category_op', 'SettingController@gift_category_op');
	Route::post('/setting/gift_category_post', 'SettingController@gift_category_post');
	Route::get('/setting/gift_category_delete', 'SettingController@gift_category_delete');
	Route::get('/setting/gift_list', 'SettingController@gift_list');
	Route::get('/setting/gift_op', 'SettingController@gift_op');
	Route::post('/setting/gift_post', 'SettingController@gift_post');
	Route::get('/setting/gift_delete', 'SettingController@gift_delete');
	
	Route::get('/setting/mobilepay_list', 'SettingController@mobilepay_list');
	Route::get('/setting/mobilepay_op', 'SettingController@mobilepay_op');
	Route::post('/setting/mobilepay_post', 'SettingController@mobilepay_post');
	Route::get('/setting/mobilepay_delete', 'SettingController@mobilepay_delete');
	
	Route::get('/setting/table_list', 'SettingController@table_list');
	Route::get('/setting/table_op', 'SettingController@table_op');
	Route::post('/setting/get_game_task', 'SettingController@get_game_task');
	Route::post('/setting/get_game_room', 'SettingController@get_game_room');
	Route::post('/setting/check_table_post', 'SettingController@check_table_post');
	Route::post('/setting/table_post', 'SettingController@table_post');
	Route::get('/setting/table_delete', 'SettingController@table_delete');
	
	Route::get('/setting/world_list', 'SettingController@world_list');
	Route::get('/setting/world_op', 'SettingController@world_op');
	Route::post('/setting/world_post', 'SettingController@world_post');
	Route::get('/setting/world_delete', 'SettingController@world_delete');
	
	Route::get('/setting/activity_order', 'SettingController@activity_order');
	Route::get('/setting/activity_order_op', 'SettingController@activity_order_op');
	Route::get('/setting/activity_order_post', 'SettingController@activity_order_post');
	Route::get('/setting/activity_order_delete', 'SettingController@activity_order_delete');

	Route::get('/setting/game_tips_list', 'SettingController@game_tips_list');
	Route::get('/setting/game_tips_op', 'SettingController@game_tips_op');
	Route::post('/setting/game_tips_post', 'SettingController@game_tips_post');
	Route::get('/setting/game_tips_delete', 'SettingController@game_tips_delete');

	Route::get('/setting/template_list', 'SettingController@template_list');
	Route::get('/setting/template_op', 'SettingController@template_op');
	Route::post('/setting/template_post', 'SettingController@template_post');
	Route::get('/setting/template_delete', 'SettingController@template_delete');
	
	Route::get('/setting/game_audit', 'SettingController@game_audit');
	Route::get('/setting/game_audit_op', 'SettingController@game_audit_op');
	Route::get('/setting/game_audit_delete', 'SettingController@game_audit_delete');
	Route::post('/setting/game_audit_post', 'SettingController@game_audit_post');

	Route::get('/setting/timeBonusConfig', 'SettingController@timeBonusConfig');
	Route::get('/setting/timeBonusConfigOp', 'SettingController@timeBonusConfigOp');
	Route::post('/setting/timeBonusConfigPost', 'SettingController@timeBonusConfigPost');
	Route::get('/setting/timeBonusConfigDelete', 'SettingController@timeBonusConfigDelete');
	Route::get('/setting/crazyCarousel', 'SettingController@crazyCarousel');
	Route::get('/setting/crazyCarouselOp', 'SettingController@crazyCarouselOp');
	Route::post('/setting/crazyCarouselPost', 'SettingController@crazyCarouselPost');
	Route::get('/setting/crazyCarouselDelete', 'SettingController@crazyCarouselDelete');
	Route::get('/setting/hotTest', 'SettingController@hotTest');
	Route::get('/setting/hotTest_op', 'SettingController@hotTest_op');
	Route::post('/setting/hotTest_post', 'SettingController@hotTest_post');
	Route::get('/setting/hotTest_delete', 'SettingController@hotTest_delete');
	Route::get('/setting/gift_order_down', 'SettingController@gift_order_down');
	Route::get('/setting/wblist', 'SettingController@wblist');
	Route::get('/setting/listCeil', 'SettingController@listCeil');
	Route::get('/setting/wblistOp', 'SettingController@wblistOp');
	Route::post('/setting/wblistPost', 'SettingController@wblistPost');
	Route::get('/setting/wblistDelete', 'SettingController@wblistDelete');
	
	Route::get('/report/game_rounds', 'ReportController@game_rounds');
	Route::get('/report/game_report', 'ReportController@game_report');
	Route::get('/report/game_summary', 'ReportController@game_summary');
	Route::get('/report/online_users', 'ReportController@online_users');
	Route::get('/report/online_times', 'ReportController@online_times');
	Route::get('/report/refresh_online', 'ReportController@refresh_online');
	Route::get('/report/onlines', 'ReportController@onlines');
	Route::get('/report/online', 'ReportController@online');
	Route::get('/report/rank', 'ReportController@rank');
	Route::get('/report/gift_order', 'ReportController@gift_order');
	Route::post('/report/gift_order_post', 'ReportController@gift_order_post');
	Route::get('/report/gift_order_delete', 'ReportController@gift_order_delete');
	Route::get('/report/order', 'ReportController@order');
	Route::get('/report/order_post', 'ReportController@order_post');
	Route::get('/report/city_order', 'ReportController@city_order');
	Route::get('/report/all', 'ReportController@all');
	Route::get('/report/all_currency', 'ReportController@all_currency');
	Route::get('/report/day_summary', 'ReportController@day_summary');
	Route::get('/report/hour_currency', 'ReportController@hour_currency');
	Route::get('/report/agent_order', 'ReportController@agent_order');
	Route::get('/report/reg', 'ReportController@reg');
	Route::get('/report/ad_summary', 'ReportController@ad_summary');
	Route::get('/report/ad_summary_channel', 'ReportController@ad_summary_channel');
	Route::get('/report/retention', 'ReportController@retention');
	Route::get('/report/register_pay', 'ReportController@register_pay');
	Route::get('/report/online_stat', 'ReportController@online_stat');
	Route::get('/report/co_user_pay', 'ReportController@co_user_pay');
	Route::get('/report/pay_plat', 'ReportController@pay_plat');
	Route::get('/report/sysnet', 'ReportController@sysnet');
	Route::get('/report/motorLiushui', 'ReportController@motorLiushui');
		
	Route::get('/report/roomChips', 'ReportController@roomChips');
	Route::get('/report/gameCount', 'ReportController@gameCount');
	Route::get('/report/game_distributed', 'ReportController@game_distributed');
	Route::get('/report/userGoldsTotal', 'ReportController@userGoldsTotal');
	
	Route::get('/log/user_liushui_log', 'LogController@user_liushui_log');
	Route::get('/log/chat_log', 'LogController@chat_log');
	Route::get('/log/table_golds_log', 'LogController@table_golds_log');
	Route::get('/log/order_detail_log', 'LogController@order_detail_log');
	Route::get('/log/login_mobile_log', 'LogController@login_mobile_log');
	Route::get('/log/login_game_log', 'LogController@login_game_log');
	Route::get('/log/game_info_log', 'LogController@game_info_log');
	Route::get('/log/tree', 'LogController@tree');
	Route::get('/log/me', 'LogController@me');
	Route::get('/log/admin_op_log', 'LogController@admin_op_log');
	Route::get('/log/bank_op_log', 'LogController@bank_op_log');
	Route::get('/log/user_golds_log', 'LogController@user_golds_log');
	Route::get('/log/user_login_log', 'LogController@user_login_log');
	Route::get('/log/user_online', 'LogController@user_online');
	Route::get('/log/room_tax_log', 'LogController@room_tax_log');
	Route::get('/log/room_stat_log', 'LogController@room_stat_log');
	Route::get('/log/room_record_log', 'LogController@room_record_log');
	Route::get('/log/lottery_log', 'LogController@lottery_log');
	Route::get('/log/round_table_log', 'LogController@round_table_log');
	Route::get('/log/tax_stat_log', 'LogController@tax_stat_log');
	Route::get('/log/round_user_log', 'LogController@round_user_log');
	Route::get('/log/user_achievement_log', 'LogController@user_achievement_log');
	Route::get('/log/logRounds', 'LogController@logRounds');
	Route::get('/log/logRoundsDetail', 'LogController@logRoundsDetail');
	
	Route::get('/file/file_module', 'FileController@file_module');
	Route::post('/file/op_file', 'FileController@op_file');
	Route::get('/file/op_file', 'FileController@op_file');
	Route::get('/file/file_delete', 'FileController@file_delete');
	Route::get('/file/version_manage', 'FileController@version_manage');
	Route::post('/file/version_manage', 'FileController@version_manage');
	Route::get('/file/publish_version_list', 'FileController@publish_version_list');
	Route::post('/file/publish_version_list', 'FileController@publish_version_list');
	Route::get('/file/file_version', 'FileController@file_version');
	Route::post('/file/file_version', 'FileController@file_version');
	Route::get('/file/versionManage', 'FileController@versionManage');
	Route::get('/file/versionManageOp', 'FileController@versionManageOp');
	Route::get('/file/versionManageDelete', 'FileController@versionManageDelete');
	Route::post('/file/versionManagePost', 'FileController@versionManagePost');
	
	Route::get('/question/question_type', 'QuestionController@question_type');
	Route::get('/question/question_ceil', 'QuestionController@question_ceil');
	Route::get('/question/question_type_op', 'QuestionController@question_type_op');
	Route::post('/question/question_type_post', 'QuestionController@question_type_post');
	Route::get('/question/delete_question_type', 'QuestionController@delete_question_type');
	Route::get('/question/question_ceil_op', 'QuestionController@question_ceil_op');
	Route::post('/question/question_ceil_post', 'QuestionController@question_ceil_post');
	Route::get('/question/delete_question_ceil', 'QuestionController@delete_question_ceil');
	Route::get('/question/question_list', 'QuestionController@question_list');
	Route::get('/question/question_answer_list', 'QuestionController@question_answer_list');
	Route::get('/question/get_user_pays', 'QuestionController@get_user_pays');
	Route::post('/question/question_answer_post', 'QuestionController@question_answer_post');
	Route::get('/question/load_answer_list', 'QuestionController@load_answer_list');
	Route::get('/question/question_complete', 'QuestionController@question_complete');
	Route::get('/question/game_help', 'QuestionController@game_help');
	Route::get('/question/game_help_op', 'QuestionController@game_help_op');
	Route::post('/question/game_help_post', 'QuestionController@game_help_post');
	Route::get('/question/delete_game_help', 'QuestionController@delete_game_help');
	
	Route::get('/setting/giftOrder', 'SettingController@giftOrder');
	Route::get('/setting/giftOrderOp', 'SettingController@giftOrderOp');
	Route::post('/setting/giftOrderPost', 'SettingController@giftOrderPost');
	Route::get('/setting/giftOrderDelete', 'SettingController@giftOrderDelete');
	Route::get('/setting/activityConfig', 'SettingController@activityConfig');
	Route::get('/setting/texasActivitySpec', 'SettingController@texasActivitySpec');
	Route::get('/setting/texasActivitySpecOp', 'SettingController@texasActivitySpecOp');
	Route::post('/setting/texasActivitySpecPost', 'SettingController@texasActivitySpecPost');
	Route::get('/setting/texasActivitySpecDelete', 'SettingController@texasActivitySpecDelete');
	
	Route::get('/setting/dealer', 'SettingController@dealer');
	Route::get('/setting/dealerOp', 'SettingController@dealerOp');
	Route::post('/setting/dealerPost', 'SettingController@dealerPost');
	Route::post('/setting/table_edit_post', 'SettingController@table_edit_post');
	
	Route::get('/report/texasUserTax', 'ReportController@texasUserTax');
	Route::get('/report/texasCount', 'ReportController@texasCount');
	
	Route::get('/report/gameDataCount', 'ReportController@gameDataCount');
	Route::get('/setting/mttList', 'SettingController@mttList');
	Route::get('/setting/mtt_op', 'SettingController@mtt_op');
	Route::post('/setting/mtt_post', 'SettingController@mtt_post');
	Route::get('/setting/mtt_delete', 'SettingController@mtt_delete');
	
	Route::get('/saler/list', 'SalerController@salerlist');
	Route::get('/saler/saler_op', 'SalerController@saler_op');
	Route::post('/saler/saler_post', 'SalerController@saler_post');
	Route::get('/saler/saler_delete', 'SalerController@saler_delete');
	Route::get('/saler/friends', 'SalerController@friends');
	Route::get('/saler/mail_op', 'SalerController@mail_op');
	Route::post('/saler/mail_post', 'SalerController@mail_post');
	Route::get('/saler/mail_delete', 'SalerController@mail_delete');
	Route::get('/saler/email', 'SalerController@email');
	Route::get('/saler/salerDetail', 'SalerController@salerDetail');
	Route::post('/saler/salerGolds', 'SalerController@salerGolds');
	Route::post('/saler/userOnline', 'SalerController@userOnline');
	Route::get('/saler/add_saler', 'SalerController@add_saler');
	Route::post('/saler/add_saler_post', 'SalerController@add_saler_post');
	Route::get('/saler/laba', 'SalerController@laba');
	Route::post('/saler/laba_post', 'SalerController@laba_post');
	Route::get('/saler/checkUser', 'SalerController@checkUser');
	Route::post('/saler/closeOrder', 'SalerController@closeOrder');
	
	Route::get('/setting/texasFreeGoldsConfig', 'SettingController@texasFreeGoldsConfig');
	Route::get('/setting/texasFreeGoldsConfigOp', 'SettingController@texasFreeGoldsConfigOp');
	Route::post('/setting/texasFreeGoldsConfigPost', 'SettingController@texasFreeGoldsConfigPost');
	Route::get('/setting/texasSubsidyGoldsConfig', 'SettingController@texasSubsidyGoldsConfig');
	Route::get('/setting/texasSubsidyGoldsConfigOp', 'SettingController@texasSubsidyGoldsConfigOp');
	Route::post('/setting/texasSubsidyGoldsConfigPost', 'SettingController@texasSubsidyGoldsConfigPost');

	Route::get('/setting/super_benefit', 'SettingController@super_benefit');
	Route::get('/setting/super_benefit_op', 'SettingController@super_benefit_op');
	Route::post('/setting/super_benefit_post', 'SettingController@super_benefit_post');
	Route::get('/setting/super_benefit_delete', 'SettingController@super_benefit_delete');

	Route::get('/setting/secret_shop', 'SettingController@secret_shop');
	Route::get('/setting/secret_shop_op', 'SettingController@secret_shop_op');
	Route::post('/setting/secret_shop_post', 'SettingController@secret_shop_post');
	Route::get('/setting/secret_shop_delete', 'SettingController@secret_shop_delete');
	Route::get('/setting/slotGMSet', 'SettingController@slotGMSet');
	Route::post('/setting/slotGMSetPost', 'SettingController@slotGMSetPost');

	Route::get('/setting/g2s', 'SettingController@g2s');
	Route::get('/setting/g2s_op', 'SettingController@g2s_op');
	Route::post('/setting/g2s_post', 'SettingController@g2s_post');
	Route::get('/setting/g2s_delete', 'SettingController@g2s_delete');
	Route::get('/setting/agentChat', 'SettingController@agentChat_op');
	Route::post('/setting/agentChat_post', 'SettingController@agentChat_post');
	
	Route::post('/setting/configField', 'SettingController@configField');
	
	Route::get('/setting/bigWin', 'SettingController@bigWin');
	Route::get('/setting/bigWinOp', 'SettingController@bigWinOp');
	Route::post('/setting/bigWinPost', 'SettingController@bigWinPost');
	Route::get('/setting/bigWinDelete', 'SettingController@bigWinDelete');
	
	Route::post('/setting/giftOrderComplete', 'SettingController@giftOrderComplete');
	Route::post('/setting/handleGiftOrder', 'SettingController@handleGiftOrder');

	Route::get('/setting/activityAudit', 'SettingController@activityAudit');
	Route::get('/setting/activityAuditOp', 'SettingController@activityAuditOp');
	Route::post('/setting/activityAuditPost', 'SettingController@activityAuditPost');
	Route::get('/setting/activityAuditDelete', 'SettingController@activityAuditDelete');
	
	Route::get('/setting/winlist', 'SettingController@winlist');
	
	Route::get('/report/regAreaCount', 'ReportController@regAreaCount');
	Route::get('/report/roomWinning', 'ReportController@roomWinning');
	Route::get('/report/userList', 'ReportController@userList');
	Route::get('/setting/slotRoomData', 'SettingController@slotRoomData');
	Route::get('/report/subsidyGolds', 'ReportController@subsidyGolds');
	Route::get('/report/vipUser', 'ReportController@vipUser');
	Route::get('/report/down_vip_user', 'ReportController@down_vip_user');
	
	Route::get('/p2p/dealerList', 'P2pController@dealerList');
	Route::get('/p2p/dealerOp', 'P2pController@dealerOp');
	Route::get('/p2p/dealerIndex', 'P2pController@dealerIndex');
	Route::get('/p2p/salerIndex', 'P2pController@salerIndex');
	Route::post('/p2p/dealerPost', 'P2pController@dealerPost');
	Route::get('/p2p/dealerDelete', 'P2pController@dealerDelete');
	Route::get('/p2p/manageSaler', 'P2pController@manageSaler');
	Route::get('/p2p/salerOp', 'P2pController@salerOp');
	Route::post('/p2p/salerPost', 'P2pController@salerPost');
	Route::get('/p2p/salerDelete', 'P2pController@salerDelete');
	Route::get('/p2p/addRoomCard', 'P2pController@addRoomCard');
	Route::post('/p2p/addRoomCardPost', 'P2pController@addRoomCardPost');
	Route::post('/p2p/addRoomCardToDealer', 'P2pController@addRoomCardToDealer');
	Route::post('/p2p/addRoomCardToUser', 'P2pController@addRoomCardToUser');
	Route::post('/p2p/getUser', 'P2pController@getUser');
	
	Route::get('/setting/jpush_list', 'SettingController@jpush_list');
	Route::get('/setting/jpush_op', 'SettingController@jpush_op');
	Route::post('/setting/jpush_post', 'SettingController@jpush_post');
	Route::get('/setting/jpush_delete', 'SettingController@jpush_delete');
	Route::post('/setting/jpush_send', 'SettingController@jpush_send');
	Route::post('/setting/jpush_view', 'SettingController@jpush_view');
	//代理商喇叭
	Route::get('/p2p/gameUserLaba', 'P2pController@gameUserLaba');
	Route::get('/p2p/gameUserlabaOp', 'P2pController@gameUserlabaOp');
	Route::post('/p2p/gameUserlabaPost', 'P2pController@gameUserlabaPost');
	Route::get('/p2p/gameUserlabaDelete', 'P2pController@gameUserlabaDelete');
	Route::post('/p2p/UpLabaTime', 'P2pController@UpLabaTime');
	Route::get('/p2p/SalerLaba', 'P2pController@SalerLaba');
	Route::post('/p2p/SalerLabaPost', 'P2pController@SalerLabaPost');

	Route::post('/p2p/getUser2', 'P2pController@getUser2');
	Route::post('/p2p/getUser3', 'P2pController@getUser3');

	Route::get('/report/adjust', 'ReportController@adjust');

	Route::get('/setting/inviteControl', 'SettingController@inviteControl');
	Route::post('/setting/inviteControlPost', 'SettingController@inviteControlPost');
	Route::post('/setting/addGameWblist', 'SettingController@addGameWblist');
	Route::get('/setting/gameWblist', 'SettingController@gameWblist');
	Route::post('/setting/removeGameWblist', 'SettingController@removeGameWblist');
	Route::get('/setting/mailDownload', 'SettingController@mailDownload');
	Route::get('/setting/liushuiDownload', 'SettingController@liushuiDownload');

	Route::get('/setting/exchangeRestrict', 'SettingController@exchangeRestrict');
	Route::get('/setting/exchangeRestrictOp', 'SettingController@exchangeRestrictOp');
	Route::post('/setting/exchangeRestrictPost', 'SettingController@exchangeRestrictPost');
	Route::get('/setting/exchangeRestrictDelete', 'SettingController@exchangeRestrictDelete');
	Route::get('/setting/salerSendGolds', 'SettingController@salerSendGolds');
	Route::post('/setting/salerSendGoldsPost', 'SettingController@salerSendGoldsPost');
	Route::get('/report/codeInvite', 'ReportController@codeInvite');
	Route::get('/report/adNewUser', 'ReportController@adNewUser');
	Route::post('/setting/bigWinTypePost', 'SettingController@bigWinTypePost');
	Route::post('/setting/super_benefit_audit', 'SettingController@super_benefit_audit');
	Route::get('/setting/salerReceiveGolds', 'SettingController@salerReceiveGolds');
	Route::get('/setting/salerReceiveGoldsOp', 'SettingController@salerReceiveGoldsOp');
	Route::post('/setting/salerReceiveGoldsPost', 'SettingController@salerReceiveGoldsPost');
	Route::get('/setting/salerReceiveGoldsDelete', 'SettingController@salerReceiveGoldsDelete');
	Route::get('/report/down_order', 'ReportController@down_order');
	Route::get('/report/down_agent_order', 'ReportController@down_agent_order');
	Route::get('/report/down_gift_order', 'ReportController@down_gift_order');
	Route::get('/sys/down_user_list', 'SysController@down_user_list');
	Route::get('/question/down_question_list', 'QuestionController@down_question_list');
	Route::post('/report/gift_order_complete', 'ReportController@gift_order_complete');
	Route::get('/setting/gameOrder', 'SettingController@gameOrder');
	Route::post('/setting/gameOrderPost', 'SettingController@gameOrderPost');
	//游戏进入条件限制
	Route::get('/setting/clientBet', 'SettingController@clientBet');
	Route::get('/setting/clientBetEdit', 'SettingController@clientBetEdit');
	Route::post('/setting/clientBetUpdate', 'SettingController@clientBetUpdate');
	Route::get('/setting/clientBetDelete', 'SettingController@clientBetDelete');
	//快速充值配置
	Route::get('/setting/speedAccount', 'SettingController@speedAccount');
	Route::get('/setting/speedAccountEdit', 'SettingController@speedAccountEdit');
	Route::post('/setting/speedAccountPost', 'SettingController@speedAccountPost');
	Route::get('/setting/speedAccountDelete', 'SettingController@speedAccountDelete');
	//充值玩转盘配置
	Route::get('/setting/zhuanpan', 'SettingController@zhuanpan');
	Route::get('/setting/zhuanpanEdit', 'SettingController@zhuanpanEdit');
	Route::post('/setting/zhuanpanPost', 'SettingController@zhuanpanPost');
	Route::get('/setting/zhuanpanDelete', 'SettingController@zhuanpanDelete');
	//老虎机内部三种充值配置
	Route::get('/setting/slotChar', 'SettingController@slotChar');
	Route::get('/setting/slotCharEdit', 'SettingController@slotCharEdit');
	Route::post('/setting/slotCharPost', 'SettingController@slotCharPost');
	Route::get('/setting/slotCharDelete', 'SettingController@slotCharDelete');
	//首冲游戏配置
	Route::get('/setting/debutRefill', 'SettingController@debutRefill');
	Route::get('/setting/debutRefillEdit', 'SettingController@debutRefillEdit');
	Route::post('/setting/debutRefillPost', 'SettingController@debutRefillPost');
	Route::get('/setting/debutRefillDelete', 'SettingController@debutRefillDelete');
	//福袋配置
	Route::get('/setting/fudai', 'SettingController@fudai');
	Route::get('/setting/fudaiOp', 'SettingController@fudaiOp');
	Route::post('/setting/fudaiPost', 'SettingController@fudaiPost');
	Route::get('/setting/fudaiDelete', 'SettingController@fudaiDelete');
	//每日累计充值配置
	Route::get('/setting/everydayAccumulateRecharge', 'SettingController@everydayAccumulateRecharge');
	Route::get('/setting/everydayAccumulateRechargeEdit', 'SettingController@everydayAccumulateRechargeEdit');
	Route::post('/setting/everydayAccumulateRechargePost', 'SettingController@everydayAccumulateRechargePost');
	//房间管理
	Route::get('/setting/room_stat_list', 'SettingController@room_stat_list');
	Route::get('/setting/room_stat_op', 'SettingController@room_stat_op');
	Route::post('/setting/room_stat_post', 'SettingController@room_stat_post');
	Route::get('/setting/room_stat_delete', 'SettingController@room_stat_delete');

	Route::get('/report/ajax_real_online_num', 'ReportController@ajax_real_online_num');
	Route::get('/report/ajax_renew_online_line', 'ReportController@ajax_renew_online_line');

	//彩金任务
    Route::get('/setting/mosaic_golds', 'SettingController@mosaicGolds');
    Route::get('/setting/mosaic_golds_op', 'SettingController@mosaicGoldsOp');
    Route::post('/setting/mosaic_golds_post', 'SettingController@mosaicGoldsPost');
    Route::get('/setting/mosaic_golds_delete', 'SettingController@mosaicGoldsDelete');
});