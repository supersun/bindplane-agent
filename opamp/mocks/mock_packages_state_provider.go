// Code generated by mockery v2.31.1. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	protobufs "github.com/open-telemetry/opamp-go/protobufs"

	types "github.com/open-telemetry/opamp-go/client/types"
)

// MockPackagesStateProvider is an autogenerated mock type for the PackagesStateProvider type
type MockPackagesStateProvider struct {
	mock.Mock
}

// AllPackagesHash provides a mock function with given fields:
func (_m *MockPackagesStateProvider) AllPackagesHash() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePackage provides a mock function with given fields: packageName, typ
func (_m *MockPackagesStateProvider) CreatePackage(packageName string, typ protobufs.PackageType) error {
	ret := _m.Called(packageName, typ)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, protobufs.PackageType) error); ok {
		r0 = rf(packageName, typ)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePackage provides a mock function with given fields: packageName
func (_m *MockPackagesStateProvider) DeletePackage(packageName string) error {
	ret := _m.Called(packageName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(packageName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileContentHash provides a mock function with given fields: packageName
func (_m *MockPackagesStateProvider) FileContentHash(packageName string) ([]byte, error) {
	ret := _m.Called(packageName)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(packageName)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(packageName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(packageName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LastReportedStatuses provides a mock function with given fields:
func (_m *MockPackagesStateProvider) LastReportedStatuses() (*protobufs.PackageStatuses, error) {
	ret := _m.Called()

	var r0 *protobufs.PackageStatuses
	var r1 error
	if rf, ok := ret.Get(0).(func() (*protobufs.PackageStatuses, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *protobufs.PackageStatuses); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protobufs.PackageStatuses)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PackageState provides a mock function with given fields: packageName
func (_m *MockPackagesStateProvider) PackageState(packageName string) (types.PackageState, error) {
	ret := _m.Called(packageName)

	var r0 types.PackageState
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (types.PackageState, error)); ok {
		return rf(packageName)
	}
	if rf, ok := ret.Get(0).(func(string) types.PackageState); ok {
		r0 = rf(packageName)
	} else {
		r0 = ret.Get(0).(types.PackageState)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(packageName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Packages provides a mock function with given fields:
func (_m *MockPackagesStateProvider) Packages() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAllPackagesHash provides a mock function with given fields: hash
func (_m *MockPackagesStateProvider) SetAllPackagesHash(hash []byte) error {
	ret := _m.Called(hash)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(hash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLastReportedStatuses provides a mock function with given fields: statuses
func (_m *MockPackagesStateProvider) SetLastReportedStatuses(statuses *protobufs.PackageStatuses) error {
	ret := _m.Called(statuses)

	var r0 error
	if rf, ok := ret.Get(0).(func(*protobufs.PackageStatuses) error); ok {
		r0 = rf(statuses)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPackageState provides a mock function with given fields: packageName, state
func (_m *MockPackagesStateProvider) SetPackageState(packageName string, state types.PackageState) error {
	ret := _m.Called(packageName, state)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, types.PackageState) error); ok {
		r0 = rf(packageName, state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateContent provides a mock function with given fields: ctx, packageName, data, contentHash
func (_m *MockPackagesStateProvider) UpdateContent(ctx context.Context, packageName string, data io.Reader, contentHash []byte) error {
	ret := _m.Called(ctx, packageName, data, contentHash)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, []byte) error); ok {
		r0 = rf(ctx, packageName, data, contentHash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockPackagesStateProvider creates a new instance of MockPackagesStateProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPackagesStateProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPackagesStateProvider {
	mock := &MockPackagesStateProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
