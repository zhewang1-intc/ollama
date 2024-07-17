package gpu

import (
	"strings"

	"github.com/klauspost/cpuid/v2"
	"golang.org/x/sys/cpu"
)

func GetCPUCapability() CPUCapability {
	if cpu.X86.HasAVX2 {
		return CPUCapabilityAVX2
	}
	if cpu.X86.HasAVX {
		return CPUCapabilityAVX
	}
	// else LCD
	return CPUCapabilityNone
}

func IsIntelCoreUltraCpus() bool {
	if strings.Contains(cpuid.CPU.BrandName, "Core(TM) Ultra") {
		return true
	}
	return false
}
