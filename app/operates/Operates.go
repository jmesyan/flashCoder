package operates

var Operates = map[string]interface{}{
	"OpenFile":  new(OpenFile),
	"WriteFile": new(WriteFile),
	"CloseFile": new(CloseFile),
	"MsgTip":    new(MsgTip),
}
