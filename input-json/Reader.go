package input_json

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

func ReadFile(filename string) (data []byte, err error) {
	// Open our jsonFile
	jsonFile, err := os.Open(filename)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	slog.Info("Successfully Opened " + filename)

	// read our opened jsonFile as a byte array.
	return io.ReadAll(jsonFile)
}
