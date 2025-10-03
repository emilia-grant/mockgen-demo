package internal_test

import (
	"net/http"
	"testing"

	"github.com/emili-grant/mockgen-demo/internal"
	mock_internal "github.com/emili-grant/mockgen-demo/internal/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

// testify/suite doesn't support parallel tests, which may or may not be a deal breaker for this
// particular suite provider

type ServiceSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	// Dependencies are typically stored in the suite
	h   *http.Client
	dep *mock_internal.MockSomeDependencyContract //NOTE: the mock type itself is used here
}

// This gets go test to run our suite
func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

// This gets called ONCE before the suite is run, can be used for things like
// disabling logs or if state needs to be created outside of a subtest
func (s *ServiceSuite) SetupSuite() {}

// This gets called before EACH subtest (s.Run) and should be used
// to create fresh state for the tests. This helps eliminate some flakiness issues
// with tests that get steamrolled by random misbheaving other tests
func (s *ServiceSuite) SetupSubTest() {
	s.h = &http.Client{}
	s.ctrl = gomock.NewController(s.T())
	s.dep = mock_internal.NewMockSomeDependencyContract(s.ctrl)
}

func (s *ServiceSuite) TestDoAThing() {
	// Now we label our tests clearly. Go test's label looks like SuiteName/FuncName/SubtestName
	// so we don't need to be explicit here that it's for our service or module
	s.Run("Success", func() {
		// SetupSubtest was just called, no need to clog my test with general state setup.
		// Only state relative to EXACTLY this test case.
		svc := internal.NewService(s.dep, s.h)
		s.dep.EXPECT().DoNetworkThing(s.T().Context(), s.h).Return(nil)

		err := svc.DoAThing(s.T().Context())
		s.Require().NoError(err)
	})
	// You can actually tell what I'm testing because it's labeled right there !WOW!
	s.Run("Failure - h is nil", func() {
		svc := internal.NewService(s.dep, nil)
		// Note that DoNetworkThing doesn't actually get called, so the value passed for h doesn't matter
		// It just needs to match. In this case, we just care about the return
		s.dep.EXPECT().DoNetworkThing(s.T().Context(), nil).Return(internal.ErrNilClient)

		err := svc.DoAThing(s.T().Context())
		s.Require().ErrorIs(err, internal.ErrUnrecoverable)
	})
	// Other tests cases ...
}
