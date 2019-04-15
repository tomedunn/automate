// Code generated by mockery v1.0.0. DO NOT EDIT.

package events

import deployment "github.com/chef/automate/api/interservice/deployment"
import mock "github.com/stretchr/testify/mock"

// MockEventSender is an autogenerated mock type for the EventSender type
type MockEventSender struct {
	mock.Mock
}

// Backup provides a mock function with given fields: backup
func (_m *MockEventSender) Backup(backup deployment.DeployEvent_Backup) {
	_m.Called(backup)
}

// Deploy provides a mock function with given fields: status
func (_m *MockEventSender) Deploy(status deployment.DeployEvent_Status) {
	_m.Called(status)
}

// Phase provides a mock function with given fields: status, phaseID
func (_m *MockEventSender) Phase(status deployment.DeployEvent_Status, phaseID deployment.DeployEvent_PhaseID) {
	_m.Called(status, phaseID)
}

// PhaseStep provides a mock function with given fields: status, phaseID, stepName, errStr
func (_m *MockEventSender) PhaseStep(status deployment.DeployEvent_Status, phaseID deployment.DeployEvent_PhaseID, stepName string, errStr string) {
	_m.Called(status, phaseID, stepName, errStr)
}

// Stop provides a mock function with given fields:
func (_m *MockEventSender) Stop() {
	_m.Called()
}

// StreamTo provides a mock function with given fields: _a0
func (_m *MockEventSender) StreamTo(_a0 func(*deployment.DeployEvent) error) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*deployment.DeployEvent) error) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TaskComplete provides a mock function with given fields:
func (_m *MockEventSender) TaskComplete() {
	_m.Called()
}
