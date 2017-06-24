package tester

import (
	"fmt"
	html "html/template"
	// "io"
	// "os"
	// "reflect"
	"flashCoder/supplier/file"
	"testing"
	text "text/template"
)

var textTmpl = text.Must(text.New("").Parse(`<script type="text/script">
		var list = {{.list}}
		console.log(list)
		</script>`))

var htmlTmpl = html.Must(html.New("").Parse(`111public function {{.list}}Award(){ return true;}`))

type Grade struct {
	Gid   int64
	Gname string
}

type BlankWriter interface {
	Write(p []byte) (n int, err error)
}

type BlankIo struct {
	Content string
}

func (b *BlankIo) Write(p []byte) (n int, err error) {
	b.Content = string(p)
	return len(p), nil
}

func TestTemplate(t *testing.T) {

	// list := []Grade{
	// 	{1, "Name1"},
	// 	{2, "Name2"},
	// }

	data := map[string]interface{}{
		"list": "goodsSALE",
	}

	// fmt.Printf("text/template:\n\n")
	// textTmpl.Execute(os.Stdout, data)

	// fmt.Printf("\n\nhtml/template:\n\n")
	// htmlTmpl.Execute(os.Stdout, data)
	// var dd BlankWriter
	// dd = new(BlankIo)
	// var writer io.Writer = dd
	// path := "./tmptest.html"
	// fd, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	// htmlTmpl.Execute(fd, data)
	itemplate := " public function {{.list}}(OrderLogic $orderLogic, $rec_id)\n    {\n        //获取商品信息\n        $nowTime = time();\n        $orderTime = $this->tradeConfig['order_time'];\n        $user = Auth::user();\n        $orderGood = OrderGoods::where(\"seller_id\", $user->id)->where(\"rec_id\", $rec_id)->first();\n        $order = Order::where(\"order_id\", $orderGood['order_id'])->first();\n        $goods = $orderLogic->sortOrderGood(TRADER_SELLER, $order, $orderGood, true);\n        $condition = array(\n            'rec_id' => $goods['rec_id'],\n            'order_id' => $goods['order_id'],\n            'goods_id' => $goods['goods_id'],\n        );\n        if ($order['order_state'] == ORDER_STATE_SUCCESS) {//订单均为已完成状态\n            $return = OrderReturn::where($condition)->orderBy('add_time', 'desc')->first();\n            //非数码作品,确认收货7天内\n            if ($goods['goods_nature'] != TRADE_GOODS_DIGITAL and $nowTime <= $order['receive_time'] + $orderTime['return_time']) {\n                //没有退款退货,定制退款，非定制退款退货,租赁的只能是退款退货，不能仅退款\n                if ($return) {\n                    //发生退款退货详情\n                    $return['address_detail'] = getAreaName($return['return_province']) . \" \" . getAreaName($return['return_city']) . \" \" . getAreaName($return['return_district']) . \" \" . $return['return_address'];\n                    $view['return'] = $return;\n                    $view['goods'] = $goods;\n                    $view['order'] = $order;\n                    $view['return_reason'] = $this->tradeConfig['return_reason'];\n                    $view['return_imgs'] = explode(\",\", $return['return_imgs']);\n                    $view['shipping'] = $this->tradeConfig['shipping'];\n                    return langView('home.trade.order.seller_order_goods_return_detail', $view);\n                }\n            }\n            //非数码作品，确认收货7天后,并且没有发生退款退货-如果发生退款退货怎么处理\n            if ($goods['goods_nature'] != TRADE_GOODS_DIGITAL and !$return and $nowTime > $order['receive_time'] + $orderTime['return_time']) {\n                if ($goods['goods_nature'] == TRADE_GOODS_RENT) {//如果是租赁作品\n                    $back = OrderBack::where($condition)->orderBy('add_time', 'desc')->first();//归还\n                    if ($back) {\n                        //发生商品返还详情\n                        $back['address_detail'] = getAreaName($back['back_province']) . \" \" . getAreaName($back['back_city']) . \" \" . getAreaName($back['back_district']) . \" \" . $back['back_address'];\n                        $view['back'] = $back;\n                        $view['order'] = $order;\n                        $view['goods'] = $goods;\n                        $view['shipping'] = $this->tradeConfig['shipping'];\n                        return langView('home.trade.order.seller_order_goods_back_detail', $view);\n                    }\n                }\n            }\n        }\n        return $this->error(\"操作有误\");\n    }\n"
	flashfile := new(file.FlashFile)
	res := flashfile.ParseTemplate(itemplate, data)
	fmt.Print(res)
	// ab := reflect.ValueOf(writer).Elem()
	// fmt.Println(ab)
}
