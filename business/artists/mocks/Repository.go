// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	artists "myuseek/business/artists"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetArtistById provides a mock function with given fields: ctx, domain
func (_m *Repository) GetArtistById(ctx context.Context, domain artists.Domain) (artists.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 artists.Domain
	if rf, ok := ret.Get(0).(func(context.Context, artists.Domain) artists.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(artists.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, artists.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArtists provides a mock function with given fields: ctx, domain
func (_m *Repository) GetArtists(ctx context.Context, domain artists.Domain) ([]artists.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 []artists.Domain
	if rf, ok := ret.Get(0).(func(context.Context, artists.Domain) []artists.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]artists.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, artists.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, domain
func (_m *Repository) Register(ctx context.Context, domain artists.Domain) (artists.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 artists.Domain
	if rf, ok := ret.Get(0).(func(context.Context, artists.Domain) artists.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(artists.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, artists.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
