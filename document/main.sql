/*
Navicat SQLite Data Transfer

Source Server         : flashCoder
Source Server Version : 30808
Source Host           : :0

Target Server Type    : SQLite
Target Server Version : 30808
File Encoding         : 65001

Date: 2017-06-17 23:06:00
*/

PRAGMA foreign_keys = OFF;

-- ----------------------------
-- Table structure for flash_behavior
-- ----------------------------
DROP TABLE IF EXISTS "main"."flash_behavior";
CREATE TABLE "flash_behavior" (
"bid"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL DEFAULT 0,
"opid"  INTEGER NOT NULL DEFAULT 0,
"bname"  TEXT(50) NOT NULL,
"paramsdef"  TEXT NOT NULL,
"remark"  TEXT(255),
"addtime"  INTEGER NOT NULL DEFAULT 0,
"updtime"  INTEGER NOT NULL DEFAULT 0
);

-- ----------------------------
-- Records of flash_behavior
-- ----------------------------
INSERT INTO "main"."flash_behavior" VALUES (7, 1, '打开文件', '[{"type":"1","name":"path","value":"b.txt"}]', '打开一个文件', 1496499317, 1496502218);
INSERT INTO "main"."flash_behavior" VALUES (10, 2, '写入文件', '[{"type":"3","name":"content","value":"today a happy day, i want to solve a problem"}]', '把数据写入一个文件', 1496502347, 1496529218);
INSERT INTO "main"."flash_behavior" VALUES (11, 3, '关闭文件', '[]', '关闭一个文件句柄', 1496502399, 1496502399);
INSERT INTO "main"."flash_behavior" VALUES (15, 4, '提示休息', '[{"type":"3","name":"msg","value":"主人你该休息了，加油哦！"}]', '休息提醒', 1497190209, 1497190209);
INSERT INTO "main"."flash_behavior" VALUES (17, 4, '模板测试', '[{"type":"3","name":"editor","value":" public function {{.list}}(OrderLogic $orderLogic, $rec_id)\n    {\n        //获取商品信息\n        $nowTime = time();\n        $orderTime = $this->tradeConfig[''order_time''];\n        $user = Auth::user();\n        $orderGood = OrderGoods::where(\"seller_id\", $user->id)->where(\"rec_id\", $rec_id)->first();\n        $order = Order::where(\"order_id\", $orderGood[''order_id''])->first();\n        $goods = $orderLogic->sortOrderGood(TRADER_SELLER, $order, $orderGood, true);\n        $condition = array(\n            ''rec_id'' => $goods[''rec_id''],\n            ''order_id'' => $goods[''order_id''],\n            ''goods_id'' => $goods[''goods_id''],\n        );\n        if ($order[''order_state''] == ORDER_STATE_SUCCESS) {//订单均为已完成状态\n            $return = OrderReturn::where($condition)->orderBy(''add_time'', ''desc'')->first();\n            //非数码作品,确认收货7天内\n            if ($goods[''goods_nature''] != TRADE_GOODS_DIGITAL and $nowTime <= $order[''receive_time''] + $orderTime[''return_time'']) {\n                //没有退款退货,定制退款，非定制退款退货,租赁的只能是退款退货，不能仅退款\n                if ($return) {\n                    //发生退款退货详情\n                    $return[''address_detail''] = getAreaName($return[''return_province'']) . \" \" . getAreaName($return[''return_city'']) . \" \" . getAreaName($return[''return_district'']) . \" \" . $return[''return_address''];\n                    $view[''return''] = $return;\n                    $view[''goods''] = $goods;\n                    $view[''order''] = $order;\n                    $view[''return_reason''] = $this->tradeConfig[''return_reason''];\n                    $view[''return_imgs''] = explode(\",\", $return[''return_imgs'']);\n                    $view[''shipping''] = $this->tradeConfig[''shipping''];\n                    return langView(''home.trade.order.seller_order_goods_return_detail'', $view);\n                }\n            }\n            //非数码作品，确认收货7天后,并且没有发生退款退货-如果发生退款退货怎么处理\n            if ($goods[''goods_nature''] != TRADE_GOODS_DIGITAL and !$return and $nowTime > $order[''receive_time''] + $orderTime[''return_time'']) {\n                if ($goods[''goods_nature''] == TRADE_GOODS_RENT) {//如果是租赁作品\n                    $back = OrderBack::where($condition)->orderBy(''add_time'', ''desc'')->first();//归还\n                    if ($back) {\n                        //发生商品返还详情\n                        $back[''address_detail''] = getAreaName($back[''back_province'']) . \" \" . getAreaName($back[''back_city'']) . \" \" . getAreaName($back[''back_district'']) . \" \" . $back[''back_address''];\n                        $view[''back''] = $back;\n                        $view[''order''] = $order;\n                        $view[''goods''] = $goods;\n                        $view[''shipping''] = $this->tradeConfig[''shipping''];\n                        return langView(''home.trade.order.seller_order_goods_back_detail'', $view);\n                    }\n                }\n            }\n        }\n        return $this->error(\"操作有误\");\n    }\n"}]', '测试', 1497709355, 1497709355);

-- ----------------------------
-- Table structure for flash_cron
-- ----------------------------
DROP TABLE IF EXISTS "main"."flash_cron";
CREATE TABLE "flash_cron" (
"crid"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL DEFAULT 0,
"second"  TEXT(10) NOT NULL,
"minute"  TEXT(10) NOT NULL,
"hour"  TEXT(10) NOT NULL,
"day"  TEXT(10) NOT NULL,
"month"  TEXT(10) NOT NULL,
"week"  TEXT(10) NOT NULL,
"tid"  INTEGER NOT NULL DEFAULT 0,
"state"  INTEGER NOT NULL DEFAULT 0,
"addtime"  INTEGER NOT NULL DEFAULT 0
);

-- ----------------------------
-- Records of flash_cron
-- ----------------------------

-- ----------------------------
-- Table structure for flash_operate
-- ----------------------------
DROP TABLE IF EXISTS "main"."flash_operate";
CREATE TABLE "flash_operate" (
"opid"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL DEFAULT 0,
"opname"  TEXT(30) NOT NULL,
"optag"  TEXT(100) NOT NULL,
"remark"  TEXT(255),
"addtime"  INTEGER NOT NULL DEFAULT 0
);

-- ----------------------------
-- Records of flash_operate
-- ----------------------------
INSERT INTO "main"."flash_operate" VALUES (1, '打开文件', 'OpenFile', '打开文件', 1497180376);
INSERT INTO "main"."flash_operate" VALUES (2, '写入文件', 'WriteFile', '写入文件', 1496236427);
INSERT INTO "main"."flash_operate" VALUES (3, '关闭文件', 'CloseFile', '关闭文件句柄', 1496236427);
INSERT INTO "main"."flash_operate" VALUES (4, '提示框', 'MsgTip', '提示内容', 1497190102);

-- ----------------------------
-- Table structure for flash_task
-- ----------------------------
DROP TABLE IF EXISTS "main"."flash_task";
CREATE TABLE "flash_task" (
"tid"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL DEFAULT 0,
"tname"  TEXT(200) NOT NULL,
"tcate"  INTEGER NOT NULL DEFAULT 0,
"tsubs"  TEXT,
"bids"  TEXT,
"addtime"  INTEGER NOT NULL DEFAULT 0,
"updtime"  INTEGER NOT NULL DEFAULT 0
);

-- ----------------------------
-- Records of flash_task
-- ----------------------------
INSERT INTO "main"."flash_task" VALUES (4, '提醒休息', 1, null, '[{"ItemId":15,"ItemName":"提示休息"}]', 1497190228, 1497190228);

-- ----------------------------
-- Table structure for flash_task_behavior
-- ----------------------------
DROP TABLE IF EXISTS "main"."flash_task_behavior";
CREATE TABLE "flash_task_behavior" (
"tbid"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL DEFAULT 0,
"bid"  INTEGER NOT NULL DEFAULT 0,
"tid"  INTEGER NOT NULL DEFAULT 0,
"ctid"  INTEGER NOT NULL DEFAULT 0,
"border"  INTEGER NOT NULL DEFAULT 0,
"torder"  INTEGER NOT NULL DEFAULT 0,
"paramsin"  TEXT NOT NULL
);

-- ----------------------------
-- Records of flash_task_behavior
-- ----------------------------
INSERT INTO "main"."flash_task_behavior" VALUES (13, 15, 4, 0, 0, 0, '[{"type":"3","name":"msg","value":"主人你该休息了，加油哦！ 哈哈，搞笑"}]');

-- ----------------------------
-- Table structure for sqlite_sequence
-- ----------------------------
DROP TABLE IF EXISTS "main"."sqlite_sequence";
CREATE TABLE sqlite_sequence(name,seq);

-- ----------------------------
-- Records of sqlite_sequence
-- ----------------------------
INSERT INTO "main"."sqlite_sequence" VALUES ('flash_task_behavior', 15);
INSERT INTO "main"."sqlite_sequence" VALUES ('flash_task', 5);
INSERT INTO "main"."sqlite_sequence" VALUES ('flash_operate', 5);
INSERT INTO "main"."sqlite_sequence" VALUES ('flash_cron', 2);
INSERT INTO "main"."sqlite_sequence" VALUES ('flash_behavior', 17);
