package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type User struct {
	Username string
	Password string
}

func main() {

	decodingJSON()
	encodingURL()
	hash()
}

func decodingJSON() {
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@example.com",
	}

	encoded, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(encoded))
}

func encodingURL() {
	originalURL := "https://example.com/search?q=hello world&category=books"
	encodedURL := url.QueryEscape(originalURL)

	fmt.Println("Original URL:", originalURL)
	fmt.Println("Encoded URL:", encodedURL)

}

func hash() {
	data := "Hello, world!"

	// Membuat objek hash dari algoritma SHA-256
	hash := sha256.New()

	// Mengupdate hash dengan data yang ingin di-hash
	hash.Write([]byte(data))

	// Mengambil nilai hash sebagai array byte
	hashBytes := hash.Sum(nil)

	// Mengubah array byte menjadi representasi heksadesimal
	hashString := hex.EncodeToString(hashBytes)

	fmt.Println("Data:", data)
	fmt.Println("Hash:", hashString)
}

func crypto() {
	// Data awal
	data := "Hello, world!"

	// Hash data awal
	initialHash := generateHash(data)

	// Simulasikan perubahan data
	modifiedData := "Hello, modified!"

	// Hash data yang diubah
	modifiedHash := generateHash(modifiedData)

	// Verifikasi integritas data
	isValid := verifyIntegrity(data, initialHash)
	fmt.Println("Data integrity:", isValid) // Output: true

	isValid = verifyIntegrity(modifiedData, initialHash)
	fmt.Println("Data integrity:", isValid) // Output: false

	isValid = verifyIntegrity(modifiedData, modifiedHash)
	fmt.Println("Data integrity:", isValid) // Output: true

}

func generateHash(data string) string {
	// Membuat objek hash dari algoritma SHA-256
	hash := sha256.New()

	// Mengupdate hash dengan data yang ingin di-hash
	hash.Write([]byte(data))

	// Mengambil nilai hash sebagai array byte
	hashBytes := hash.Sum(nil)

	// Mengubah array byte menjadi representasi heksadesimal
	hashString := fmt.Sprintf("%x", hashBytes)

	return hashString
}

func verifyIntegrity(data, expectedHash string) bool {
	// Menghasilkan hash dari data yang diberikan
	hash := generateHash(data)

	// Membandingkan hash yang dihasilkan dengan hash yang diharapkan
	return hash == expectedHash
}

//	func handler() {
//		registerUser := User
//	}
//
// func registerUser(user *user) {
//
// }
func generateSalt() string {
	salt := make([]byte, 16)
	rand.Read(salt)
	return hex.EncodeToString(salt)
}

//
//func hotpImplementation() {
//	// Membuat kunci rahasia baru untuk HOTP
//	key, err := hotp.Generate(hotp.GenerateOpts{
//		Issuer:      "MyApp",
//		AccountName: "user@example.com",
//	})
//	if err != nil {
//		fmt.Println("Gagal menghasilkan kunci rahasia:", err)
//		return
//	}
//
//	// Menghasilkan kode sandi HOTP dengan counter 0
//	code := hotp.GenerateCode(key.Secret(), 0)
//	fmt.Println("Kode Sandi HOTP:", code)
//
//	// Memverifikasi kode sandi HOTP dengan counter 0
//	valid := hotp.ValidateCustom(code, 0, key.Secret())
//	if valid {
//		fmt.Println("Kode sandi HOTP valid.")
//	} else {
//		fmt.Println("Kode sandi HOTP tidak valid.")
//	}
//
//}
