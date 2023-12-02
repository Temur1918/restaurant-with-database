package config

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	FilePathForTable         = "files/Table.json"
	FilePathForOrder         = "files/Order.json"
	FilePathForProduct       = "files/Product.json"
	FilePathForOrderProducts = "files/OrderProducts.json"
	FilePathForWaiter        = "files/Waiter.json"
)

func ServiceFee(sum float64) float64 {
	const pr float64 = 19
	return float64(sum/100) * pr
}

// GPT
func ClearJSONFile(filePath string) error {
	// JSON faylni o'qish
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("faylni o'qishda xatolik yuz berdi: %v", err)
	}
	defer file.Close()

	// Faylni tozalash uchun bo'sh qilish
	err = file.Truncate(0)
	if err != nil {
		return fmt.Errorf("faylni tozalashda xatolik yuz berdi: %v", err)
	}

	return nil
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
