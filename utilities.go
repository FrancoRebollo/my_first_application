package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	"strings"
	"unicode"
)

// Generar una clave de 16 bytes (128 bits) a partir de la semilla
func generateKey(seed string) []byte {
	key := []byte(seed)
	if len(key) > 16 {
		key = key[:16] // Truncar si la clave es mayor a 16 bytes
	} else if len(key) < 16 {
		for len(key) < 16 {
			key = append(key, '0') // Rellenar con ceros si es menor a 16 bytes
		}
	}
	return key
}

// Función para cifrar el texto con AES
func encrypt(text, key string) (string, error) {

	keyResized := generateKey(key)

	block, err := aes.NewCipher([]byte(keyResized))
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Función para descifrar el texto con AES
func decrypt(cryptoText, key string) (string, error) {

	keyResized := generateKey(key)

	block, err := aes.NewCipher([]byte(keyResized))
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func passVerifyFunction(username, password, oldPassword string) error {
	var (
		errors []string
	)

	// Contraseña igual al usuario
	if strings.EqualFold(username, password) {
		errors = append(errors, "La contraseña no puede ser igual que el usuario.")
	}

	// Contraseña contiene al usuario
	if strings.Contains(strings.ToLower(password), strings.ToLower(username)) {
		errors = append(errors, "La contraseña no puede contener el usuario.")
	}

	// Longitud menor a 6 caracteres
	if len(password) < 6 {
		errors = append(errors, "La longitud de la contraseña no puede ser menor a 6 caracteres.")
	}

	// Revisamos palabras comunes
	commonPasswords := []string{"bienvenido", "contraseña", "contrasena", "contrasenia", "password", "computer", "qwerty", "asdasd", "123456", "usuario"}
	for _, common := range commonPasswords {
		if strings.EqualFold(password, common) {
			errors = append(errors, "La contraseña es muy simple.")
			break
		}
	}

	// Revisamos que contenga al menos un dígito
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		errors = append(errors, "La contraseña debe contener al menos un dígito.")
	}

	// Revisamos que contenga al menos un carácter
	hasChar := false
	for _, char := range password {
		if unicode.IsLetter(char) {
			hasChar = true
			break
		}
	}
	if !hasChar {
		errors = append(errors, "La contraseña debe contener al menos un carácter.")
	}

	// Revisamos diferencia mínima con la contraseña anterior
	if oldPassword != "" {
		difference := 0
		minLength := len(password)
		if len(oldPassword) < minLength {
			minLength = len(oldPassword)
		}
		for i := 0; i < minLength; i++ {
			if password[i] != oldPassword[i] {
				difference++
			}
		}
		if difference < 3 {
			errors = append(errors, "La contraseña nueva debe diferir de la anterior en al menos 3 caracteres.")
		}

		if oldPassword == password {
			errors = append(errors, "La nueva contraseña no puede ser igual a la contraseña anterior.")
		}

	}

	// Si hay errores, devolvemos un error con los mensajes acumulados
	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, " "))
	}

	// Si no hay errores, devolvemos nil
	return nil
}

func getLastNCharacters(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[len(s)-n:]
}

func generateVerificationCode(length int) (string, error) {
	const charset = "0123456789"
	code := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := range code {
		num, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		code[i] = charset[num.Int64()]
	}
	return string(code), nil
}
