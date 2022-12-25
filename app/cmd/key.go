package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	logger "github.com/aland20/go-noting/app/loggers"
	"github.com/labstack/gommon/color"
	"github.com/spf13/cobra"
)

func NewKeyCommand() *cobra.Command {

	keyCommand := &cobra.Command{
		Use:   "key",
		Short: "Generate a random base64 32 bit string",
		Run: func(cmd *cobra.Command, args []string) {

			key := generateRandomKey()

			var wg sync.WaitGroup
			wg.Add(1)

			go func() {

				defer wg.Done()
				updateEnvFile(key)
			}()

			wg.Wait()

			fmt.Println(key)
		},
	}

	return keyCommand
}

func generateRandomKey() string {
	cmd := exec.Command("openssl", "rand", "-base64", "32")
	out, err := cmd.Output()

	if err != nil {
		panic(color.Red("Failed to generate key"))
	}

	return string(out)
}

func updateEnvFile(key string) {

	pwd, _ := os.Getwd()
	path := pwd + "/.env"

	inputFile, oErr := os.ReadFile(path)

	if oErr != nil {
		logger.Panic("Failed to read .env file")
	}

	fileLines := strings.Split(string(inputFile), "\n")

	for i, val := range fileLines {

		if strings.Contains(val, "APP_KEY") {

			fileLines[i] = fmt.Sprintf("APP_KEY=%s", key)
			break
		}
	}

	outputFile := strings.Join(fileLines, "\n")

	wErr := os.WriteFile(path, []byte(outputFile), 0644)

	if wErr != nil {
		logger.Panic("Failed to edit .env file")
	}

}
