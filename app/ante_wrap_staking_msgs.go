package app

import (
	epochingtypes "github.com/babylonlabs-io/babylon/x/epoching/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// WrapStakingMsgDecorator defines an AnteHandler decorator that rejects all messages that might change the validator set.
type WrapStakingMsgDecorator struct {
}

// NewWrapStakingMsgDecorator creates a new DropValidatorMsgDecorator
func NewWrapStakingMsgDecorator() *WrapStakingMsgDecorator {
	return &WrapStakingMsgDecorator{}
}

// AnteHandle performs an AnteHandler will wrap all the staking msgs that will be sent to epoch.
// It will replace the following types of messages:
// - MsgCreateValidator -> MsgWrappedDelegate
// - MsgDelegate ->
// - MsgUndelegate ->
// - MsgBeginRedelegate ->
// - MsgCancelUnbondingDelegation ->
func (qmd WrapStakingMsgDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// skip if at genesis block, as genesis state contains txs that bootstrap the initial validator set
	if ctx.BlockHeight() == 0 {
		return next(ctx, tx, simulate)
	}

	// after genesis, if validator-related message, reject msg
	for _, msg := range tx.GetMsgs() {
		newMsg, replace := qmd.IsValidatorRelatedMsg(msg)
		if replace {
			msg = newMsg
		}
	}

	return next(ctx, tx, simulate)
}

// IsValidatorRelatedMsg checks if the given message is of non-wrapped type, which should be rejected
func (qmd WrapStakingMsgDecorator) IsValidatorRelatedMsg(msg sdk.Msg) (newMsg sdk.Msg, replace bool) {
	switch msg := msg.(type) {
	case *stakingtypes.MsgDelegate:
		return epochingtypes.NewMsgWrappedDelegate(msg), true
		// Do for others...
	case *stakingtypes.MsgCreateValidator, *stakingtypes.MsgUndelegate, *stakingtypes.MsgBeginRedelegate, *stakingtypes.MsgCancelUnbondingDelegation:
		return msg, false
	default:
		return msg, false
	}
}
