// +build cli_test

package cli

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/codec/testdata"
	"github.com/stretchr/testify/require"
)

func TestCliQueryConn(t *testing.T) {
	t.Parallel()
	f := NewFixtures(t)

	// start simd server
	proc := f.SDStart()
	t.Cleanup(func() { proc.Stop(false) })

	ctx := client.NewContext()
	queryConn := ctx.QueryConn()
	testClient := testdata.NewTestServiceClient(queryConn)
	res, err := testClient.Echo(context.Background(), &testdata.EchoRequest{Message: "hello"})
	require.NoError(t, err)
	require.Equal(t, "hello", res.Message)
}
