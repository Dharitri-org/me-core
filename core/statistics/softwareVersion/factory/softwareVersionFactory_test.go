package factory

import (
	"testing"

	"github.com/Dharitri-org/me-core/core"
	"github.com/Dharitri-org/me-core/core/mock"
	"github.com/stretchr/testify/assert"
)

func TestNewSoftwareVersionFactory_NilStatusHandlerShouldErr(t *testing.T) {
	t.Parallel()

	softwareVersionFactory, err := NewSoftwareVersionFactory(nil, "", 0)

	assert.Equal(t, core.ErrNilAppStatusHandler, err)
	assert.Nil(t, softwareVersionFactory)
}

func TestSoftwareVersionFactory_Create(t *testing.T) {
	t.Parallel()

	statusHandler := &mock.AppStatusHandlerStub{}
	softwareVersionFactory, _ := NewSoftwareVersionFactory(statusHandler, "", 1)
	softwareVersionChecker, err := softwareVersionFactory.Create()

	assert.Nil(t, err)
	assert.NotNil(t, softwareVersionChecker)
}