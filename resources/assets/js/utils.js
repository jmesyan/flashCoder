
function back(level) {
	history.back(level);
}

function timeFormat(timestamp) {
	var now = new Date(parseInt(timestamp) * 1000);
	var year=now.getYear(); 
	var month=now.getMonth()+1; 
	var date=now.getDate(); 
	var hour=now.getHours(); 
	var minute=now.getMinutes(); 
	var second=now.getSeconds(); 
	return "20"+year+"-"+month+"-"+date+" "+hour+":"+minute+":"+second; 
} 