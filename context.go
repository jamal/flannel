package flannel

import (
	"net/http"
	"reflect"
	"sync"
)

var (
	mutex       sync.RWMutex
	reqContexts = make(map[*http.Request]map[reflect.Type]reflect.Value)
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

func deleteContext(r *http.Request) {
	mutex.Lock()
	delete(reqContexts, r)
	mutex.Unlock()
}

// Context returns a context for a http.Request
func Context(r *http.Request, data interface{}) {
	typ := reflect.TypeOf(data)
	dv := reflect.Indirect(reflect.ValueOf(data))

	// Read existing context
	mutex.RLock()
	if sv, ok := reqContexts[r][typ]; ok {
		mutex.RUnlock()
		dv.Set(sv)
		return
	}
	mutex.RUnlock()

	// Write new context
	mutex.Lock()
	if _, ok := reqContexts[r]; !ok {
		reqContexts[r] = make(map[reflect.Type]reflect.Value)
	}
	reqContexts[r][typ] = dv
	mutex.Unlock()
}
