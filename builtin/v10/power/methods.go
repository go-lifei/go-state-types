package power

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/frc0042"
	"github.com/filecoin-project/go-state-types/proof"
)

var Methods = map[abi.MethodNum]builtin.MethodMeta{
	1: {"Constructor", *new(func(*abi.EmptyValue) *abi.EmptyValue)},       // Constructor
	2: {"CreateMiner", *new(func(*CreateMinerParams) *CreateMinerReturn)}, // CreateMiner
	frc0042.GenerateExportedMethodNum("CreateMiner"): {"CreateMinerExported", *new(func(*CreateMinerParams) *CreateMinerReturn)}, // CreateMinerExported
	3: {"UpdateClaimedPower", *new(func(*UpdateClaimedPowerParams) *abi.EmptyValue)},   // UpdateClaimedPower
	4: {"EnrollCronEvent", *new(func(*EnrollCronEventParams) *abi.EmptyValue)},         // EnrollCronEvent
	5: {"CronTick", *new(func(*abi.EmptyValue) *abi.EmptyValue)},                       // CronTick
	6: {"UpdatePledgeTotal", *new(func(*abi.TokenAmount) *abi.EmptyValue)},             // UpdatePledgeTotal
	7: {"OnConsensusFault", nil},                                                       // deprecated
	8: {"SubmitPoRepForBulkVerify", *new(func(*proof.SealVerifyInfo) *abi.EmptyValue)}, // SubmitPoRepForBulkVerify
	9: {"CurrentTotalPower", *new(func(*abi.EmptyValue) *CurrentTotalPowerReturn)},     // CurrentTotalPower
	frc0042.GenerateExportedMethodNum("NetworkRawPower"):     {"NetworkRawPowerExported", *new(func(*abi.EmptyValue) *NetworkRawPowerReturn)},         // NetworkRawPowerExported
	frc0042.GenerateExportedMethodNum("MinerRawPower"):       {"MinerRawPowerExported", *new(func(*MinerRawPowerParams) *MinerRawPowerReturn)},        // MinerRawPowerExported
	frc0042.GenerateExportedMethodNum("MinerCount"):          {"MinerCountExported", *new(func(*abi.EmptyValue) *MinerCountReturn)},                   // MinerCountExported
	frc0042.GenerateExportedMethodNum("MinerConsensusCount"): {"MinerConsensusCountExported", *new(func(*abi.EmptyValue) *MinerConsensusCountReturn)}, // MinerConsensusCountExported
}
