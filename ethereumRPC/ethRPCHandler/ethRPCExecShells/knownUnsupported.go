package ethRPCExecShells

func GenUnsupportedShell(name string) *EthRPCExecShell {
	return &EthRPCExecShell{
		name:        name,
		unsupported: true,
	}
}
