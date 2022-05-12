package events

var dispatcher *EventDispatcher

// 事件侦听
func Init() {
	dispatcher = NewEventDispatcher()
	ManagerProcesslistener := NewEventListener(ManagerProcessHandler)
	dispatcher.AddEventListener("manager process abnormal", ManagerProcesslistener)
}

// 事件触发
func Trigger(e Event) {
	dispatcher.DispatchEvent(e)
}
