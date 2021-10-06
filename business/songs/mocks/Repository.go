// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	songs "myuseek/business/songs"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, domain
func (_m *Repository) Add(ctx context.Context, domain songs.Domain) (songs.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 songs.Domain
	if rf, ok := ret.Get(0).(func(context.Context, songs.Domain) songs.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(songs.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, songs.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSongById provides a mock function with given fields: ctx, domain
func (_m *Repository) GetSongById(ctx context.Context, domain songs.Domain) (songs.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 songs.Domain
	if rf, ok := ret.Get(0).(func(context.Context, songs.Domain) songs.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(songs.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, songs.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSongLyrics provides a mock function with given fields: ctx, domain
func (_m *Repository) GetSongLyrics(ctx context.Context, domain songs.Domain) (songs.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 songs.Domain
	if rf, ok := ret.Get(0).(func(context.Context, songs.Domain) songs.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(songs.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, songs.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSongs provides a mock function with given fields: ctx, domain
func (_m *Repository) GetSongs(ctx context.Context, domain songs.Domain) ([]songs.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 []songs.Domain
	if rf, ok := ret.Get(0).(func(context.Context, songs.Domain) []songs.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]songs.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, songs.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
