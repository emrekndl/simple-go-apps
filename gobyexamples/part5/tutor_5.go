package main

import (
	"bufio"
	"crypto/sha256"
	"embed"
	b64 "encoding/base64"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func _numberParsing() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println("String to float('1.234 - 64'): ", f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println("String to int('123 - 0 - 64'): ", i)

	i2, _ := strconv.ParseInt("2357", 10, 64)
	fmt.Println("String to int('2357 - 10 - 64'): ", i2)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println("String to int('0x1c8 - 0 - 64'): ", d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println("String to uint('789 - 0 - 64'): ", u)

	// Atoi base10 int parsing
	k, _ := strconv.Atoi("135")
	fmt.Println("String to atoi(135): ", k)

	_, e := strconv.Atoi("wat")
	fmt.Println("String to atoi(wat): ", e)
}

func _urlParsing() {
	// URL, which includes a scheme, authentication info, host, port, path, query params, and query fragment.
	psql := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(psql)
	if err != nil {
		panic(err)
	}
	fmt.Println("URL Scheme: ", u.Scheme)

	fmt.Println("URL User: ", u.User)
	fmt.Println("URL User.Username: ", u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("URL User.Password: ", p)

	fmt.Println("URL Host: ", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("URL Host(sp-host-port): ", host)
	fmt.Println("URL Port(sp-host-port): ", port)

	fmt.Println("URL Path: ", u.Path)
	fmt.Println("URL Fragment: ", u.Fragment)

	fmt.Println("URL RawQuery: ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("URL Query: ", m)
	fmt.Println("URL Query: ", m["k"][0])
}

func _sha256Hashes() {
	str := "sha256 this is a test."

	hash := sha256.New()

	hash.Write([]byte(str))

	bs := hash.Sum(nil)

	fmt.Println("Text: ", str)
	fmt.Printf("SHA256: %x\n", bs)
}

func _base64Encoding() {
	data := "abc123!?$*&()'-=@~"

	standartEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Standart Encoding: ", standartEnc)

	standartDec, _ := b64.StdEncoding.DecodeString(standartEnc)
	fmt.Println("Standart Decoding: ", string(standartDec))
	fmt.Println()

	urlEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("URL Encoding: ", urlEnc)

	urlDec, _ := b64.URLEncoding.DecodeString(urlEnc)
	fmt.Println("URL Decoding: ", string(urlDec))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func _readinFiles() {
	// Create a /tmp/data directory for passing to the error
	// cat "hello" > /tmp/data; cat "go" >> /tmp/data
	data, err := os.ReadFile("/tmp/data")
	check(err)
	fmt.Println("Os ReadFile: ", string(data))

	file, err := os.Open("/tmp/data")
	check(err)

	b1 := make([]byte, 5)
	n1, err := file.Read(b1)
	check(err)
	fmt.Printf("%d bytes %s\n", n1, string(b1[:n1]))

	o2, err := file.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := file.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = file.Seek(4, io.SeekStart)
	check(err)

	// _, err = file.Seek(-10, io.SeekEnd)
	// check(err)

	o3, err := file.Seek(6, io.SeekStart)
	check(err)

	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(file, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = file.Seek(0, io.SeekStart)
	check(err)

	r4 := bufio.NewReader(file)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes %s\n", string(b4))

	file.Chdir()
}

func _writingFiles() {
	data1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/data1", data1, 0644)
	check(err)

	file2, err := os.Create("/tmp/data2")
	check(err)

	defer file2.Close()

	data2 := []byte{115, 111, 109, 101, 10}
	n2, err := file2.Write(data2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := file2.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	file2.Sync()

	w := bufio.NewWriter(file2)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}

func _lineFilters() {
	// "hello lines" > lines.txt
	// cat lines.txt | go run tutor_5.go
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println("Scanner Text: ", ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}

func _filePaths() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("Join Path: ", p)

	fmt.Println("Join (dir1// filename): ", filepath.Join("dir1//", "filename"))
	fmt.Println("Join (dir1/../dir1 filename): ", filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p): ", filepath.Dir(p))
	fmt.Println("Base(p): ", filepath.Base(p))

	fmt.Println("Is absolute path(dir/file): ", filepath.IsAbs("dir/file"))
	fmt.Println("Is absolute path(/dir/file): ", filepath.IsAbs("/dir/file"))

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println("Filename extension: ", ext)

	fmt.Println("File name extension removed(trimsuffix): ", strings.TrimSuffix(filename, ext))

	//Rel, finds a relative path between a base and a target.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println("Relative path found: ", rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println("Relative path found: ", rel)

}

func visit(path string, d fs.DirEntry, err error) error {

	if err != nil {
		return err
	}
	fmt.Printf("Path: %s - isDir: %v\n", path, d.IsDir())
	return nil
}

func _directories() {
	err := os.Mkdir("subdir", 0755)
	check(err)

	// rm -rf
	defer os.RemoveAll("subdir")

	createEmptyFile := func(filename string) {
		d := []byte("")
		check(os.WriteFile(filename, d, 0644))
	}

	createEmptyFile("subdir/file1")
	// mkdirall -> mkdir -p
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir lists directory contents, returning a slice of os.DirEntry objects.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent :")
	for _, entry := range c {
		fmt.Printf("Name: %v - isDir: %v\n ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	check(err)

	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child:")
	for _, entry := range c {
		fmt.Printf("Name: %s - isDir: %v\n", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("(WalkDir- recursively)Visiting subdir: ")
	err = filepath.WalkDir("subdir", visit)
	check(err)
}

func _temporaryFilesAndDirectories() {
	tmpFile, err := os.CreateTemp("", "samplefile")
	check(err)

	fmt.Println("Temporary File Name: ", tmpFile.Name())

	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte{2, 3, 5, 7, 11})
	check(err)

	dirName, err := os.MkdirTemp("", "sampleDir")
	check(err)

	fmt.Println("Temporary Directory Name: ", dirName)

	defer os.RemoveAll(dirName)

	fileName := filepath.Join(dirName, "file1")
	err = os.WriteFile(fileName, []byte("temp dir inside file..."), 0666)
	check(err)

	data, err := os.ReadFile(fileName)
	check(err)
	fmt.Println("file1 data: ", string(data))
}

// //go:embed folder/single_file.txt
var fileString string

// //go:embed folder/single_file.txt
var fileByte []byte

// //go:embed folder/single_file.txt
// //go:embed folder/*.hash
var folder embed.FS

// //go:embed folder/single_file2.txt
var fileString2 string

func _embedDirective() {
	// mkdir -p folder
	// echo "hello go" > folder/single_file.txt
	// echo "test" > folder/single_file2.txt
	// echo "123" > folder/file1.hash
	// echo "456" > folder/file2.hash

	fmt.Println("FileString2: ", fileString2)

	fmt.Println("FileString: ", fileString)
	fmt.Println("FileByte: ", string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	fmt.Println("file1.hash: ", string(content1))
	content2, _ := folder.ReadFile("folder/file2.hash")
	fmt.Println("file2.hash: ", string(content2))

}

func main() {
	// Number Parsing
	_numberParsing()
	fmt.Println("----------------")
	// URL Parsing
	_urlParsing()
	fmt.Println("----------------")
	// SHA256 Hashes
	_sha256Hashes()
	fmt.Println("----------------")
	// Base64 Encoding
	_base64Encoding()
	fmt.Println("----------------")
	// Reading Files
	// _readinFiles()
	fmt.Println("----------------")
	// Writing Files
	_writingFiles()
	fmt.Println("----------------")
	// Line Filters
	_lineFilters()
	fmt.Println("----------------")
	// File Paths
	_filePaths()
	fmt.Println("----------------")
	// Directories
	_directories()
	fmt.Println("----------------")
	// Temporary Files and Directories
	_temporaryFilesAndDirectories()
	fmt.Println("----------------")
	// Embed Directive
	// _embedDirective()
	fmt.Println("----------------")
}
