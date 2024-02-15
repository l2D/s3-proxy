// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/response-handler (interfaces: ResponseHandler)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_ResponseHandler.go -package=mocks github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/response-handler ResponseHandler
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	responsehandler "github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/response-handler"
	gomock "go.uber.org/mock/gomock"
)

// MockResponseHandler is a mock of ResponseHandler interface.
type MockResponseHandler struct {
	ctrl     *gomock.Controller
	recorder *MockResponseHandlerMockRecorder
}

// MockResponseHandlerMockRecorder is the mock recorder for MockResponseHandler.
type MockResponseHandlerMockRecorder struct {
	mock *MockResponseHandler
}

// NewMockResponseHandler creates a new mock instance.
func NewMockResponseHandler(ctrl *gomock.Controller) *MockResponseHandler {
	mock := &MockResponseHandler{ctrl: ctrl}
	mock.recorder = &MockResponseHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResponseHandler) EXPECT() *MockResponseHandlerMockRecorder {
	return m.recorder
}

// BadRequestError mocks base method.
func (m *MockResponseHandler) BadRequestError(arg0 func(context.Context, string) (string, error), arg1 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BadRequestError", arg0, arg1)
}

// BadRequestError indicates an expected call of BadRequestError.
func (mr *MockResponseHandlerMockRecorder) BadRequestError(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BadRequestError", reflect.TypeOf((*MockResponseHandler)(nil).BadRequestError), arg0, arg1)
}

// Delete mocks base method.
func (m *MockResponseHandler) Delete(arg0 func(context.Context, string) (string, error), arg1 *responsehandler.DeleteInput) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", arg0, arg1)
}

// Delete indicates an expected call of Delete.
func (mr *MockResponseHandlerMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockResponseHandler)(nil).Delete), arg0, arg1)
}

// FoldersFilesList mocks base method.
func (m *MockResponseHandler) FoldersFilesList(arg0 func(context.Context, string) (string, error), arg1 []*responsehandler.Entry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FoldersFilesList", arg0, arg1)
}

// FoldersFilesList indicates an expected call of FoldersFilesList.
func (mr *MockResponseHandlerMockRecorder) FoldersFilesList(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FoldersFilesList", reflect.TypeOf((*MockResponseHandler)(nil).FoldersFilesList), arg0, arg1)
}

// ForbiddenError mocks base method.
func (m *MockResponseHandler) ForbiddenError(arg0 func(context.Context, string) (string, error), arg1 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ForbiddenError", arg0, arg1)
}

// ForbiddenError indicates an expected call of ForbiddenError.
func (mr *MockResponseHandlerMockRecorder) ForbiddenError(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForbiddenError", reflect.TypeOf((*MockResponseHandler)(nil).ForbiddenError), arg0, arg1)
}

// GetRequest mocks base method.
func (m *MockResponseHandler) GetRequest() *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequest")
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// GetRequest indicates an expected call of GetRequest.
func (mr *MockResponseHandlerMockRecorder) GetRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequest", reflect.TypeOf((*MockResponseHandler)(nil).GetRequest))
}

// InternalServerError mocks base method.
func (m *MockResponseHandler) InternalServerError(arg0 func(context.Context, string) (string, error), arg1 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InternalServerError", arg0, arg1)
}

// InternalServerError indicates an expected call of InternalServerError.
func (mr *MockResponseHandlerMockRecorder) InternalServerError(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InternalServerError", reflect.TypeOf((*MockResponseHandler)(nil).InternalServerError), arg0, arg1)
}

// NotFoundError mocks base method.
func (m *MockResponseHandler) NotFoundError(arg0 func(context.Context, string) (string, error)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotFoundError", arg0)
}

// NotFoundError indicates an expected call of NotFoundError.
func (mr *MockResponseHandlerMockRecorder) NotFoundError(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotFoundError", reflect.TypeOf((*MockResponseHandler)(nil).NotFoundError), arg0)
}

// NotModified mocks base method.
func (m *MockResponseHandler) NotModified() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotModified")
}

// NotModified indicates an expected call of NotModified.
func (mr *MockResponseHandlerMockRecorder) NotModified() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotModified", reflect.TypeOf((*MockResponseHandler)(nil).NotModified))
}

// PreconditionFailed mocks base method.
func (m *MockResponseHandler) PreconditionFailed() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PreconditionFailed")
}

// PreconditionFailed indicates an expected call of PreconditionFailed.
func (mr *MockResponseHandlerMockRecorder) PreconditionFailed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreconditionFailed", reflect.TypeOf((*MockResponseHandler)(nil).PreconditionFailed))
}

// Put mocks base method.
func (m *MockResponseHandler) Put(arg0 func(context.Context, string) (string, error), arg1 *responsehandler.PutInput) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", arg0, arg1)
}

// Put indicates an expected call of Put.
func (mr *MockResponseHandlerMockRecorder) Put(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockResponseHandler)(nil).Put), arg0, arg1)
}

// RedirectTo mocks base method.
func (m *MockResponseHandler) RedirectTo(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RedirectTo", arg0)
}

// RedirectTo indicates an expected call of RedirectTo.
func (mr *MockResponseHandlerMockRecorder) RedirectTo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedirectTo", reflect.TypeOf((*MockResponseHandler)(nil).RedirectTo), arg0)
}

// RedirectWithTrailingSlash mocks base method.
func (m *MockResponseHandler) RedirectWithTrailingSlash() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RedirectWithTrailingSlash")
}

// RedirectWithTrailingSlash indicates an expected call of RedirectWithTrailingSlash.
func (mr *MockResponseHandlerMockRecorder) RedirectWithTrailingSlash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedirectWithTrailingSlash", reflect.TypeOf((*MockResponseHandler)(nil).RedirectWithTrailingSlash))
}

// StreamFile mocks base method.
func (m *MockResponseHandler) StreamFile(arg0 func(context.Context, string) (string, error), arg1 *responsehandler.StreamInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StreamFile indicates an expected call of StreamFile.
func (mr *MockResponseHandlerMockRecorder) StreamFile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamFile", reflect.TypeOf((*MockResponseHandler)(nil).StreamFile), arg0, arg1)
}

// TargetList mocks base method.
func (m *MockResponseHandler) TargetList() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TargetList")
}

// TargetList indicates an expected call of TargetList.
func (mr *MockResponseHandlerMockRecorder) TargetList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TargetList", reflect.TypeOf((*MockResponseHandler)(nil).TargetList))
}

// UnauthorizedError mocks base method.
func (m *MockResponseHandler) UnauthorizedError(arg0 func(context.Context, string) (string, error), arg1 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnauthorizedError", arg0, arg1)
}

// UnauthorizedError indicates an expected call of UnauthorizedError.
func (mr *MockResponseHandlerMockRecorder) UnauthorizedError(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnauthorizedError", reflect.TypeOf((*MockResponseHandler)(nil).UnauthorizedError), arg0, arg1)
}

// UpdateRequestAndResponse mocks base method.
func (m *MockResponseHandler) UpdateRequestAndResponse(arg0 *http.Request, arg1 http.ResponseWriter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateRequestAndResponse", arg0, arg1)
}

// UpdateRequestAndResponse indicates an expected call of UpdateRequestAndResponse.
func (mr *MockResponseHandlerMockRecorder) UpdateRequestAndResponse(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequestAndResponse", reflect.TypeOf((*MockResponseHandler)(nil).UpdateRequestAndResponse), arg0, arg1)
}
