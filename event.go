package nginx

import "errors"

// 判断是否是event指令
func isEventDirective(directive string) bool {
	if isEqualString(directive, EventDirective) {
		return true
	}
	return false
}

// 判断是否是accept_mutex指令
func isEventAcceptMutexDirective(directive string) bool {
	if isEqualString(directive, EventAcceptMutexDirective) {
		return true
	}
	return false
}

// 判断是否是accept_mutex_delay指令
func isEventAcceptMutexDelayDirective(directive string) bool {
	if isEqualString(directive, EventAcceptMutexDelayDirective) {
		return true
	}
	return false
}

// 判断是否是debug_connection指令
func isEventDebugConnectionDirective(directive string) bool {
	if isEqualString(directive, EventDebugConnectionDirective) {
		return true
	}
	return false
}

// 判断是否是use指令
func isEventUseDirective(directive string) bool {
	if isEqualString(directive, EventUseDirective) {
		return true
	}
	return false
}

// 判断是否是worker_aio_requests指令
func isEventWorkerAioRequestsDirective(directive string) bool {
	if isEqualString(directive, EventWorkerAioRequestsDirective) {
		return true
	}
	return false
}

// 判断是否是worker_connections指令
func isEventWorkerConnectionsDirective(directive string) bool {
	if isEqualString(directive, EventWorkerConnectionsDirective) {
		return true
	}
	return false
}

// ProcessEvent process event directive
func ProcessEvent(parsed *Parsed) (*Event, error) {
	if !isEventDirective(parsed.Directive) {
		return nil, errors.New("Not event direction")
	}
	event := NewEvent()
	for _, block := range parsed.Blocks {
		// 处理accept_mutex指令
		if isEventAcceptMutexDirective(block.Directive) {
			event.AcceptMutex = processArgsArray(block.Args)
			continue
		}

		// 处理accept_mutex_delay指令
		if isEventAcceptMutexDelayDirective(block.Directive) {
			event.AcceptMutexDelay = processArgsArray(block.Args)
			continue
		}

		// 处理use指令
		if isEventUseDirective(block.Directive) {
			event.Use = processArgsArray(block.Args)
			continue
		}

		// 处理worker_aio_requests指令
		if isEventWorkerAioRequestsDirective(block.Directive) {
			event.WorkerAioRequests = processArgsArray(block.Args)
			continue
		}

		// 处理worker_connections指令
		if isEventWorkerConnectionsDirective(block.Directive) {
			event.WorkerConnections = processArgsArray(block.Args)
			continue
		}

		// 处理debug_connection指令
		if isEventDebugConnectionDirective(block.Directive) {
			event.DebugConnection = append(event.DebugConnection, processArgsArray(block.Args))
		}
	}
	return event, nil
}
