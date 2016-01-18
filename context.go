package flannel

import (
	"net/http"
	"sync"
	"time"
)

// RequestContext wraps http.Request
type RequestContext struct {
	Start     time.Time
	RequestID string
}

var (
	mutex       sync.RWMutex
	reqContexts = make(map[*http.Request]*RequestContext)
	reqIDs      = make(map[*http.Request]string)
)

func setReqID(r *http.Request, rid string) {
	mutex.Lock()
	reqIDs[r] = rid
	mutex.Unlock()
}

func reqID(r *http.Request) string {
	mutex.RLock()
	defer mutex.RUnlock()
	return reqIDs[r]
}

// Context returns a context for a http.Request
func Context(r *http.Request) *RequestContext {
	mutex.RLock()
	if ctx, ok := reqContexts[r]; ok {
		mutex.RUnlock()
		return ctx
	}
	mutex.RUnlock()

	mutex.Lock()
	ctx := &RequestContext{}
	reqContexts[r] = ctx
	mutex.Unlock()
	return ctx
}

// DeleteContext deletes the request context for a given http.Request
func DeleteContext(r *http.Request) {
	mutex.Lock()
	delete(reqContexts, r)
	mutex.Unlock()
}
