package lighthouse

import (
	"context"
	"fmt"
	"testing"

	"github.com/bloxapp/eth2-key-manager/core"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestStreamDuties(t *testing.T) {
	logger := zaptest.NewLogger(t)

	client := newLighthouseStreamDuties(context.Background(), logger, core.PyrmontNetwork, "http://18.158.194.7:5052", []string{"113584"})

	for {
		duty, err := client.Recv()
		require.NoError(t, err)

		fmt.Println("current epoch duties", duty.GetCurrentEpochDuties())
		fmt.Println("next epoch duties", duty.GetNextEpochDuties())
	}
}
