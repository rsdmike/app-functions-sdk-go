// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BackgroundPublisher is an autogenerated mock type for the BackgroundPublisher type
type BackgroundPublisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: payload, correlationID, contentType
func (_m *BackgroundPublisher) Publish(payload []byte, correlationID string, contentType string) {
	_m.Called(payload, correlationID, contentType)
}
