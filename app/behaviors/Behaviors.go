package behaviors

var Behaviors = map[string]interface{}{
	"OpenFile":  new(OpenFile),
	"WriteFile": new(WriteFile),
	"CloseFile": new(CloseFile),
}
