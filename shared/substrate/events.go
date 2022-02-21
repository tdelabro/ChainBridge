// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	events "github.com/capsule-corp-ternoa/chainbridge-substrate-events"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// EventFeeChanged is emitted when a fee for a given key is changed.
type EventFeeChanged struct {
	Phase    types.Phase
	Key      types.Hash
	NewPrice types.U128
	Topics   []types.Hash
}

// EventNewMultiAccount is emitted when a multi account has been created.
// First param is the account that created it, second is the multisig account.
type EventNewMultiAccount struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// EventMultiAccountUpdated is emitted when a multi account has been updated. First param is the multisig account.
type EventMultiAccountUpdated struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventMultiAccountRemoved is emitted when a multi account has been removed. First param is the multisig account.
type EventMultiAccountRemoved struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventNewMultisig is emitted when a new multisig operation has begun.
// First param is the account that is approving, second is the multisig account.
type EventNewMultisig struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// TimePoint contains height and index
type TimePoint struct {
	Height types.U32
	Index  types.U32
}

// EventMultisigApproval is emitted when a multisig operation has been approved by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigApproval struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

// EventMultisigExecuted is emitted when a multisig operation has been executed by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigExecuted struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Result    types.DispatchResult
	Topics    []types.Hash
}

// EventMultisigCancelled is emitted when a multisig operation has been cancelled by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigCancelled struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

type EventTreasuryMinting struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventRadClaimsClaimed is emitted when RAD Tokens have been claimed
type EventRadClaimsClaimed struct {
	Phase  types.Phase
	Who    types.AccountID
	Value  types.U128
	Topics []types.Hash
}

// EventRadClaimsRootHashStored is emitted when RootHash has been stored for the correspondent RAD Claims batch
type EventRadClaimsRootHashStored struct {
	Phase    types.Phase
	RootHash types.Hash
	Topics   []types.Hash
}

// EventRegistryTmp is emitted only for testing
type EventRegistryTmp struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

// EventDeposit is emitted when an account withdraw some free balance
type EventBalancesWithdraw struct {
	Phase   types.Phase
	Who     types.AccountID
	Balance types.U128
	Topics  []types.Hash
}

type Events struct {
	types.EventRecords
	events.Events
	Balances_Withdraw                []EventBalancesWithdraw               //nolint:stylecheck,golint
	Council_Proposed                 []types.EventCollectiveProposed       //nolint:stylecheck,golint
	Council_Voted                    []types.EventCollectiveVoted          //nolint:stylecheck,golint
	Council_Approved                 []types.EventCollectiveApproved       //nolint:stylecheck,golint
	Council_Disapproved              []types.EventCollectiveDisapproved    //nolint:stylecheck,golint
	Council_Executed                 []types.EventCollectiveExecuted       //nolint:stylecheck,golint
	Council_MemberExecuted           []types.EventCollectiveMemberExecuted //nolint:stylecheck,golint
	Council_Closed                   []types.EventCollectiveClosed         //nolint:stylecheck,golint
	Fees_FeeChanged                  []EventFeeChanged                     //nolint:stylecheck,golint
	MultiAccount_NewMultiAccount     []EventNewMultiAccount                //nolint:stylecheck,golint
	MultiAccount_MultiAccountUpdated []EventMultiAccountUpdated            //nolint:stylecheck,golint
	MultiAccount_MultiAccountRemoved []EventMultiAccountRemoved            //nolint:stylecheck,golint
	MultiAccount_NewMultisig         []EventNewMultisig                    //nolint:stylecheck,golint
	MultiAccount_MultisigApproval    []EventMultisigApproval               //nolint:stylecheck,golint
	MultiAccount_MultisigExecuted    []EventMultisigExecuted               //nolint:stylecheck,golint
	MultiAccount_MultisigCancelled   []EventMultisigCancelled              //nolint:stylecheck,golint
	TreasuryReward_TreasuryMinting   []EventTreasuryMinting                //nolint:stylecheck,golint
	RadClaims_Claimed                []EventRadClaimsClaimed               //nolint:stylecheck,golint
	RadClaims_RootHashStored         []EventRadClaimsRootHashStored        //nolint:stylecheck,golint
	Registry_RegistryTmp             []EventRegistryTmp                    //nolint:stylecheck,golint
}
