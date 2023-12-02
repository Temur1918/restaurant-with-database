package files

import (
	"fmt"
	"os"
	"restaurant/config"
)

func CreateFiles() {
	filePaths := []string{config.FilePathForTable, config.FilePathForOrder, config.FilePathForOrderProducts, config.FilePathForWaiter, config.FilePathForProduct}

	for _, filePath := range filePaths {
		_, err := os.Stat(filePath)
		isNotFileExists := os.IsNotExist(err)

		if isNotFileExists {
			if _, err := os.Create(filePath); err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(filePath, "is created!")

		}
	}
}
