// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package web

import (
	"context"
	"net/http"
	"sync"
)

var (
	lockConnectionMockAddr           sync.RWMutex
	lockConnectionMockListenAndServe sync.RWMutex
	lockConnectionMockListenPath     sync.RWMutex
	lockConnectionMockShutdown       sync.RWMutex
)

// ConnectionMock is a mock implementation of Connection.
//
//     func TestSomethingThatUsesConnection(t *testing.T) {
//
//         // make and configure a mocked Connection
//         mockedConnection := &ConnectionMock{
//             AddrFunc: func() string {
// 	               panic("TODO: mock out the Addr method")
//             },
//             ListenAndServeFunc: func() error {
// 	               panic("TODO: mock out the ListenAndServe method")
//             },
//             ListenPathFunc: func(path string, method string, handler http.HandlerFunc) error {
// 	               panic("TODO: mock out the ListenPath method")
//             },
//             ShutdownFunc: func(ctx context.Context)  {
// 	               panic("TODO: mock out the Shutdown method")
//             },
//         }
//
//         // TODO: use mockedConnection in code that requires Connection
//         //       and then make assertions.
//
//     }
type ConnectionMock struct {
	// AddrFunc mocks the Addr method.
	AddrFunc func() string

	// ListenAndServeFunc mocks the ListenAndServe method.
	ListenAndServeFunc func() error

	// ListenPathFunc mocks the ListenPath method.
	ListenPathFunc func(path string, method string, handler http.HandlerFunc) error

	// ShutdownFunc mocks the Shutdown method.
	ShutdownFunc func(ctx context.Context)

	// calls tracks calls to the methods.
	calls struct {
		// Addr holds details about calls to the Addr method.
		Addr []struct {
		}
		// ListenAndServe holds details about calls to the ListenAndServe method.
		ListenAndServe []struct {
		}
		// ListenPath holds details about calls to the ListenPath method.
		ListenPath []struct {
			// Path is the path argument value.
			Path string
			// Method is the method argument value.
			Method string
			// Handler is the handler argument value.
			Handler http.HandlerFunc
		}
		// Shutdown holds details about calls to the Shutdown method.
		Shutdown []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
}

// Addr calls AddrFunc.
func (mock *ConnectionMock) Addr() string {
	if mock.AddrFunc == nil {
		panic("moq: ConnectionMock.AddrFunc is nil but Connection.Addr was just called")
	}
	callInfo := struct {
	}{}
	lockConnectionMockAddr.Lock()
	mock.calls.Addr = append(mock.calls.Addr, callInfo)
	lockConnectionMockAddr.Unlock()
	return mock.AddrFunc()
}

// AddrCalls gets all the calls that were made to Addr.
// Check the length with:
//     len(mockedConnection.AddrCalls())
func (mock *ConnectionMock) AddrCalls() []struct {
} {
	var calls []struct {
	}
	lockConnectionMockAddr.RLock()
	calls = mock.calls.Addr
	lockConnectionMockAddr.RUnlock()
	return calls
}

// ListenAndServe calls ListenAndServeFunc.
func (mock *ConnectionMock) ListenAndServe() error {
	if mock.ListenAndServeFunc == nil {
		panic("moq: ConnectionMock.ListenAndServeFunc is nil but Connection.ListenAndServe was just called")
	}
	callInfo := struct {
	}{}
	lockConnectionMockListenAndServe.Lock()
	mock.calls.ListenAndServe = append(mock.calls.ListenAndServe, callInfo)
	lockConnectionMockListenAndServe.Unlock()
	return mock.ListenAndServeFunc()
}

// ListenAndServeCalls gets all the calls that were made to ListenAndServe.
// Check the length with:
//     len(mockedConnection.ListenAndServeCalls())
func (mock *ConnectionMock) ListenAndServeCalls() []struct {
} {
	var calls []struct {
	}
	lockConnectionMockListenAndServe.RLock()
	calls = mock.calls.ListenAndServe
	lockConnectionMockListenAndServe.RUnlock()
	return calls
}

// ListenPath calls ListenPathFunc.
func (mock *ConnectionMock) ListenPath(path string, method string, handler http.HandlerFunc) error {
	if mock.ListenPathFunc == nil {
		panic("moq: ConnectionMock.ListenPathFunc is nil but Connection.ListenPath was just called")
	}
	callInfo := struct {
		Path    string
		Method  string
		Handler http.HandlerFunc
	}{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
	lockConnectionMockListenPath.Lock()
	mock.calls.ListenPath = append(mock.calls.ListenPath, callInfo)
	lockConnectionMockListenPath.Unlock()
	return mock.ListenPathFunc(path, method, handler)
}

// ListenPathCalls gets all the calls that were made to ListenPath.
// Check the length with:
//     len(mockedConnection.ListenPathCalls())
func (mock *ConnectionMock) ListenPathCalls() []struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
} {
	var calls []struct {
		Path    string
		Method  string
		Handler http.HandlerFunc
	}
	lockConnectionMockListenPath.RLock()
	calls = mock.calls.ListenPath
	lockConnectionMockListenPath.RUnlock()
	return calls
}

// Shutdown calls ShutdownFunc.
func (mock *ConnectionMock) Shutdown(ctx context.Context) {
	if mock.ShutdownFunc == nil {
		panic("moq: ConnectionMock.ShutdownFunc is nil but Connection.Shutdown was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockConnectionMockShutdown.Lock()
	mock.calls.Shutdown = append(mock.calls.Shutdown, callInfo)
	lockConnectionMockShutdown.Unlock()
	mock.ShutdownFunc(ctx)
}

// ShutdownCalls gets all the calls that were made to Shutdown.
// Check the length with:
//     len(mockedConnection.ShutdownCalls())
func (mock *ConnectionMock) ShutdownCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockConnectionMockShutdown.RLock()
	calls = mock.calls.Shutdown
	lockConnectionMockShutdown.RUnlock()
	return calls
}