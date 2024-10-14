package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Proje ismini sor
	fmt.Print("Proje ismi: ")
	projectName, _ := reader.ReadString('\n')
	projectName = strings.TrimSpace(projectName)

	// GitHub modülü kullanmak isteyip istemediğini sor
	fmt.Print("GitHub projesi mi? (y/n): ")
	githubChoice, _ := reader.ReadString('\n')
	githubChoice = strings.TrimSpace(githubChoice)

	var moduleName string
	if githubChoice == "y" || githubChoice == "Y" {
		// GitHub username sor
		fmt.Print("GitHub kullanıcı adı: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		// GitHub modül ismini ayarla
		moduleName = fmt.Sprintf("github.com/%s/%s", username, projectName)
	} else {
		// Yerel modül ismi ayarla
		moduleName = projectName
	}

	// Proje dizin yapısını oluştur
	createDirStructure(projectName)

	// go mod init komutunu çalıştır
	runGoModInit(projectName, moduleName)

	// main.go dosyasını oluştur
	createMainFile(projectName)

	fmt.Printf("%s projesi başarıyla oluşturuldu.\n", projectName)
}

func createDirStructure(projectName string) {
	directories := []string{
		filepath.Join(projectName, "cmd"),
		filepath.Join(projectName, "pkg"),
		filepath.Join(projectName, "internal"),
	}

	for _, dir := range directories {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatalf("Dizin oluşturulamadı: %v", err)
		}
	}
}

func runGoModInit(projectName, moduleName string) {
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = projectName
	if err := cmd.Run(); err != nil {
		log.Fatalf("go mod init başarısız: %v", err)
	}
}

func createMainFile(projectName string) {
	mainFilePath := filepath.Join(projectName, "main.go")
	mainFileContent := fmt.Sprintf(`package main

import "fmt"

func main() {
	fmt.Println("Hello, %s!")
}
`, projectName)

	err := os.WriteFile(mainFilePath, []byte(mainFileContent), 0644)
	if err != nil {
		log.Fatalf("main.go dosyası oluşturulamadı: %v", err)
	}
}
