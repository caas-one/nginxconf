package core

import "errors"

// determine whether it is a event directive
func isEventDirective(directive string) bool {
	if isEqualString(directive, EventDirective) {
		return true
	}
	return false
}

// determine whether it is a accept_mutex directive
func isEventAcceptMutexDirective(directive string) bool {
	if isEqualString(directive, EventAcceptMutexDirective) {
		return true
	}
	return false
}

// determine whether it is a accept_mutex_delay directive
func isEventAcceptMutexDelayDirective(directive string) bool {
	if isEqualString(directive, EventAcceptMutexDelayDirective) {
		return true
	}
	return false
}

// determine whether it is a debug_connection directive
func isEventDebugConnectionDirective(directive string) bool {
	if isEqualString(directive, EventDebugConnectionDirective) {
		return true
	}
	return false
}

// determine whether it is a use directive
func isEventUseDirective(directive string) bool {
	if isEqualString(directive, EventUseDirective) {
		return true
	}
	return false
}

// determine whether it is a worker_aio_requests directive
func isEventWorkerAioRequestsDirective(directive string) bool {
	if isEqualString(directive, EventWorkerAioRequestsDirective) {
		return true
	}
	return false
}

// determine whether it is a worker_connections directive
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
		// process accept_mutex directive
		if isEventAcceptMutexDirective(block.Directive) {
			event.AcceptMutex = processArgsArray(block.Args)
			continue
		}

		// process accept_mutex_delay directive
		if isEventAcceptMutexDelayDirective(block.Directive) {
			event.AcceptMutexDelay = processArgsArray(block.Args)
			continue
		}

		// process use directive
		if isEventUseDirective(block.Directive) {
			event.Use = processArgsArray(block.Args)
			continue
		}

		// process worker_aio_requests directive
		if isEventWorkerAioRequestsDirective(block.Directive) {
			event.WorkerAioRequests = processArgsArray(block.Args)
			continue
		}

		// process worker_connections directive
		if isEventWorkerConnectionsDirective(block.Directive) {
			event.WorkerConnections = processArgsArray(block.Args)
			continue
		}

		// process debug_connection directive
		if isEventDebugConnectionDirective(block.Directive) {
			event.DebugConnection = append(event.DebugConnection, processArgsArray(block.Args))
		}
	}
	return event, nil
}
