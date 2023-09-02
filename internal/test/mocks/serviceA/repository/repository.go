// Code generated by mockery v2.14.0. DO NOT EDIT.

package repository

import (
	domain "app/internal/serviceA/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/satori/go.uuid"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]*domain.ItemA, error) {
	ret := _m.Called(ctx)

	var r0 []*domain.ItemA
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.ItemA); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.ItemA)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.ItemA, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.ItemA
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.ItemA); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ItemA)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, item
func (_m *Repository) Insert(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error) {
	ret := _m.Called(ctx, item)

	var r0 *domain.ItemA
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ItemA) *domain.ItemA); ok {
		r0 = rf(ctx, item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ItemA)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.ItemA) error); ok {
		r1 = rf(ctx, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, id
func (_m *Repository) Remove(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, id, item
func (_m *Repository) Update(ctx context.Context, id uuid.UUID, item *domain.ItemA) error {
	ret := _m.Called(ctx, id, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *domain.ItemA) error); ok {
		r0 = rf(ctx, id, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}