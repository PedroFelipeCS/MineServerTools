package main

import (
	"fmt"
	"os"

	handler "MineServerTools/http"
)

func main() {
	url, err := handler.GetLatestServerVersion()
	if err != nil {
		fmt.Printf("Erro ao buscar a ultima versão: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Baixando URL: %s\n", url)

	file, err := handler.DownloadServer(url)
	if err != nil {
		fmt.Printf("Erro ao fazer download do server: %v\n", err)
		os.Exit(1)
	}

	filename := "downloads/bedrock_server.zip"
	err = handler.SaveFile(filename, file)
	if err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Download completado com sucesso!")
}