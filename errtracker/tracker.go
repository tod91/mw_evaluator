// Package errtracker defines our global error tracker structure for keeping track of errors that might occur
package errtracker

import (
	"mw_evaluator/models"
	"sync"
)

// Tracker ...
//
// will be used by all requests to take snapshots or different errors that might occur
var Tracker = &errStorage{table: make(map[string]*failedExpression)}

type failedExpression models.ErrResp

type errStorage struct {
	mu    sync.Mutex
	table map[string]*failedExpression
}

func (t *errStorage) Save(expression, endpoint string, e error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, ok := t.table[expression]; !ok {
		t.table[expression] = &failedExpression{}
	}

	t.table[expression].Endpoint = endpoint
	t.table[expression].Frequency = t.table[expression].Frequency + 1
	t.table[expression].Expression = expression
	t.table[expression].ErrType = e.Error()
}

func (t *errStorage) GetAll() []*failedExpression {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.table == nil {
		panic("uninitialized map variable")
	}

	var retVal []*failedExpression
	for _, v := range t.table {
		retVal = append(retVal, v)
	}
	return retVal
}
