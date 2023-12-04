package config

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ServiceFee(sum float64) float64 {
	const pr float64 = 19
	return float64(sum/100) * pr
}

// GPT
func Clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported operating system. Cannot clear the screen.")
	}
}
