// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	storage_v1 "github.com/jaegertracing/jaeger/proto-gen/storage_v1"
)

// ArchiveSpanReaderPluginClient is an autogenerated mock type for the ArchiveSpanReaderPluginClient type
type ArchiveSpanReaderPluginClient struct {
	mock.Mock
}

// GetArchiveTrace provides a mock function with given fields: ctx, in, opts
func (_m *ArchiveSpanReaderPluginClient) GetArchiveTrace(ctx context.Context, in *storage_v1.GetTraceRequest, opts ...grpc.CallOption) (storage_v1.ArchiveSpanReaderPlugin_GetArchiveTraceClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetArchiveTrace")
	}

	var r0 storage_v1.ArchiveSpanReaderPlugin_GetArchiveTraceClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage_v1.GetTraceRequest, ...grpc.CallOption) (storage_v1.ArchiveSpanReaderPlugin_GetArchiveTraceClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage_v1.GetTraceRequest, ...grpc.CallOption) storage_v1.ArchiveSpanReaderPlugin_GetArchiveTraceClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage_v1.ArchiveSpanReaderPlugin_GetArchiveTraceClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage_v1.GetTraceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewArchiveSpanReaderPluginClient creates a new instance of ArchiveSpanReaderPluginClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewArchiveSpanReaderPluginClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ArchiveSpanReaderPluginClient {
	mock := &ArchiveSpanReaderPluginClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
