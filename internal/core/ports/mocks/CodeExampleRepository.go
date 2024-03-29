// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	domain "speed-typing-text-service/internal/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// CodeExampleRepository is an autogenerated mock type for the CodeExampleRepository type
type CodeExampleRepository struct {
	mock.Mock
}

// GetCodeExampleByUUID provides a mock function with given fields: _a0
func (_m *CodeExampleRepository) GetCodeExampleByUUID(_a0 string) (domain.CodeExample, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetCodeExampleByUUID")
	}

	var r0 domain.CodeExample
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.CodeExample, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) domain.CodeExample); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.CodeExample)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCodeExamples provides a mock function with given fields:
func (_m *CodeExampleRepository) GetCodeExamples() []domain.CodeExample {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCodeExamples")
	}

	var r0 []domain.CodeExample
	if rf, ok := ret.Get(0).(func() []domain.CodeExample); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.CodeExample)
		}
	}

	return r0
}

// GetCodeExamplesByProgrammingLanguageName provides a mock function with given fields: programmingLanguageName
func (_m *CodeExampleRepository) GetCodeExamplesByProgrammingLanguageName(programmingLanguageName string) ([]domain.CodeExample, error) {
	ret := _m.Called(programmingLanguageName)

	if len(ret) == 0 {
		panic("no return value specified for GetCodeExamplesByProgrammingLanguageName")
	}

	var r0 []domain.CodeExample
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]domain.CodeExample, error)); ok {
		return rf(programmingLanguageName)
	}
	if rf, ok := ret.Get(0).(func(string) []domain.CodeExample); ok {
		r0 = rf(programmingLanguageName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.CodeExample)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(programmingLanguageName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProgrammingLanguages provides a mock function with given fields:
func (_m *CodeExampleRepository) GetProgrammingLanguages() []domain.ProgrammingLanguage {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProgrammingLanguages")
	}

	var r0 []domain.ProgrammingLanguage
	if rf, ok := ret.Get(0).(func() []domain.ProgrammingLanguage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ProgrammingLanguage)
		}
	}

	return r0
}

// NewCodeExampleRepository creates a new instance of CodeExampleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCodeExampleRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CodeExampleRepository {
	mock := &CodeExampleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
