package internal_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/emili-grant/mockgen-demo/internal"
	mock_internal "github.com/emili-grant/mockgen-demo/internal/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// This is the raw testing functions pattern. It's a bit quicker to set up but intermingles a bunch of setup code with
// your actual testing logic for non-trivial tests

func TestDoAThing(t *testing.T) {
	require := require.New(t)
	ctrl := gomock.NewController(t)
	// Mock dependency is created here
	comp := mock_internal.NewMockSomeDependencyContract(ctrl)
	ctx := context.Background()
	h := &http.Client{}

	// svc gets instantiated with the mock dependency
	svc := internal.NewService(comp, h)

	// the code under test should call comp.DoNetworkThing exactly once, it should use ctx and h, and it will receive nil
	// from that call
	// This forces a sneaky good habit: in order to define what DoNetworkThing should receive, I have to be able to
	// inject it into my code under test
	comp.EXPECT().DoNetworkThing(ctx, h).Return(nil)

	// But you can also allow anything to be passed when you can't and won't inject the value into the code
	// under test like
	// comp.EXPECT().DoNetworkThing(ctx, gomock.Any()).Return(nil)

	// Call the code under test.
	// Try commenting this out and see what the output is
	// Try duplicating the call, too
	err := svc.DoAThing(ctx)
	require.NoError(err)

	// There's a couple of glaring issues here though:
	// - There's implicit state rolling through this test function. If you do multiple tests in this one function
	// failure in one test could cause failure in another if you're not careful with your Asserts/Requires
	//
	// - Test setup is mixed with test logic and is required in each function. Lots of developers solve this
	// by homerolling a testcase type and then running through an array of test cases and calling t.Run() for subtests
	// but what if there was a better way? Like maybe in service_suite_test.go?
}
