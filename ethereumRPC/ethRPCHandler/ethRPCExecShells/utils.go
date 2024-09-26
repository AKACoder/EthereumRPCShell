package ethRPCExecShells

import (
	"github.com/AKACoder/EthereumRPCShell/common/dataStructure/types"
	"github.com/AKACoder/EthereumRPCShell/common/shellErrors"
	"github.com/AKACoder/EthereumRPCShell/common/utils"
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPCProvider"
)

func extractAddressAndBlock(pAddr any, pBlk any) (*types.Address, *EthBlockNumString, error) {
	var addr types.Address

	err := addr.FromParam(pAddr)
	if err != nil {
		return nil, nil, shellErrors.InvalidParameterType
	}

	var blk EthBlockNumString
	if !blk.FromAny(pBlk) {
		return nil, nil, shellErrors.InvalidParameterType
	}

	return &addr, &blk, nil
}

func extractHash(param any) (*types.Hash, error) {
	var hash types.Hash
	err := hash.FromParam(param)
	if err != nil {
		return nil, shellErrors.InvalidParameterType
	}

	return &hash, nil
}

func extractHashAndInt(pHash any, pInt any) (*types.Hash, uint64, error) {
	var hash types.Hash
	err := hash.FromParam(pHash)
	if err != nil {
		return nil, 0, shellErrors.InvalidParameterType
	}

	i, err := utils.HexParamToUint64(pInt)
	if err != nil {
		return nil, 0, shellErrors.InvalidParameterType
	}

	return &hash, i, nil
}

func extractBlkNumAndInt(pBlkNum any, pInt any) (*EthBlockNumString, uint64, error) {
	var blk EthBlockNumString

	if !blk.FromAny(pBlkNum) {
		return nil, 0, shellErrors.InvalidParameterType
	}

	i, err := utils.HexParamToUint64(pInt)
	if err != nil {
		return nil, 0, shellErrors.InvalidParameterType
	}

	return &blk, i, nil
}
