// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/anggi-susanto/mrt-go/domain"
	mock "github.com/stretchr/testify/mock"
)

// DeviceService is an autogenerated mock type for the DeviceService type
type DeviceService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, w
func (_m *DeviceService) Create(ctx context.Context, w *domain.DeviceRequest) error {
	ret := _m.Called(ctx, w)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.DeviceRequest) error); ok {
		r0 = rf(ctx, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *DeviceService) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, page, limit
func (_m *DeviceService) GetAll(ctx context.Context, page int, limit int) ([]domain.Device, error) {
	ret := _m.Called(ctx, page, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []domain.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]domain.Device, error)); ok {
		return rf(ctx, page, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []domain.Device); ok {
		r0 = rf(ctx, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *DeviceService) GetByID(ctx context.Context, id string) (*domain.Device, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *domain.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.Device, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Device); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, w
func (_m *DeviceService) Update(ctx context.Context, w *domain.Device) error {
	ret := _m.Called(ctx, w)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Device) error); ok {
		r0 = rf(ctx, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDeviceService creates a new instance of DeviceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceService(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceService {
	mock := &DeviceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
