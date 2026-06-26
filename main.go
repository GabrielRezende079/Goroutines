// Example without goroutines

package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Yellow = "\033[35m"
)

// Ping the list of websites inside a For loop.
func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.youtube.com",
		"https://www.microsoft.com",
		"https://www.cursor.com",
		"https://www.chatgpt.com",
		"https://www.linkedin.com",
	}

	inicio := time.Now()

	fmt.Printf("%s🚀 Iniciando Verificações...%s\n\n", Blue, Reset)

	for _, url := range urls {
		verifyServiceStatus(url)
	}

	fmt.Printf("\n%sFinalizado em: %.2fs%s\n", Green, time.Since(inicio).Seconds(), Reset)
}

func verifyServiceStatus(url string) {
	start := time.Now()
	fmt.Printf("%s[START]%s - Iniciando verificação: %s\n", Yellow, Reset, url)

	res, err := http.Get(url)
	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Printf("%s[ERROR] %s - Erro: %v - Tempo: %.2fs%s\n", Red, url, err, elapsed, Reset)
		return
	}
	defer res.Body.Close()

	fmt.Printf("%s[OK] %s - Status: %d - Tempo: %.2fs%s\n", Green, url, res.StatusCode, elapsed, Reset)
}
