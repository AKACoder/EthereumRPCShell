package ethRPCExecShells

func GenUnsupportedShell(name string) *EthRPCExecShell {
	return &EthRPCExecShell{
		Name:        name,
		unsupported: true,
	}
}
