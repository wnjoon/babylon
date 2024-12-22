package btcstaking_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/babylonlabs-io/babylon/btcstaking"
	"github.com/stretchr/testify/require"
)

func TestIsRateValid(t *testing.T) {

	require.True(t, btcstaking.IsRateValid(sdkmath.LegacyMustNewDecFromStr("0.1")))
	require.True(t, btcstaking.IsRateValid(sdkmath.LegacyMustNewDecFromStr("0.01")))
	require.True(t, btcstaking.IsRateValid(sdkmath.LegacyMustNewDecFromStr("0.001")))

}
