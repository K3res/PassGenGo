package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"

	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"
)

// function for the Password Generator
func generatePassword(length int, chars string, exclude string) (string, error) {
	password := make([]byte, length)

	// Create a map to store excluded characters for faster lookup
	excludeMap := make(map[byte]struct{})
	for i := 0; i < len(exclude); i++ {
		excludeMap[exclude[i]] = struct{}{}
	}

	for i := 0; i < length; i++ {
		var randomChar byte
		for {
			randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
			if err != nil {
				return "", err
			}
			randomChar = chars[randomIndex.Int64()]

			// Check if the character is not in the exclude map
			if _, excluded := excludeMap[randomChar]; !excluded {
				break
			}
		}
		password[i] = randomChar
	}

	return string(password), nil
}

// CopyToClipboard copies the specified text to the clipboard.
func CopyToClipboard(text string) error {
	return clipboard.WriteAll(text)
}

// PKCS7Unpad removes PKCS7 padding from the input
func PKCS7Unpad(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

// hashed the Password
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Generate the GenerateSymmetricKey
func generateSymmetricKey() []byte {
	key := make([]byte, 32) // 32 bytes for AES-256
	if _, err := rand.Read(key); err != nil {
		panic("Error generating random key")
	}
	return key
}

// Encrypt passwords with the given key
func encryptPasswords(passwords []string, key []byte) ([]string, error) {
	var encryptedPasswords []string

	for _, password := range passwords {
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}

		passwordBytes := []byte(password)

		// Pad the password to be a multiple of the block size
		padSize := aes.BlockSize - len(passwordBytes)%aes.BlockSize
		padding := bytes.Repeat([]byte{byte(padSize)}, padSize)
		passwordBytes = append(passwordBytes, padding...)

		ciphertext := make([]byte, aes.BlockSize+len(passwordBytes))
		iv := ciphertext[:aes.BlockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return nil, err
		}

		mode := cipher.NewCBCEncrypter(block, iv)
		mode.CryptBlocks(ciphertext[aes.BlockSize:], passwordBytes)

		encryptedPasswords = append(encryptedPasswords, hex.EncodeToString(ciphertext))
	}

	return encryptedPasswords, nil
}

// Secrypt passwords with the given key
func decryptPasswords(encryptedPasswords []string, key []byte) ([]string, error) {
	var decryptedPasswords []string

	for _, encryptedPassword := range encryptedPasswords {
		ciphertext, err := hex.DecodeString(encryptedPassword)
		if err != nil {
			return nil, err
		}

		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}

		if len(ciphertext) < aes.BlockSize {
			return nil, fmt.Errorf("ciphertext is too short")
		}

		iv := ciphertext[:aes.BlockSize]
		ciphertext = ciphertext[aes.BlockSize:]

		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(ciphertext, ciphertext)

		decryptedPasswords = append(decryptedPasswords, string(PKCS7Unpad(ciphertext)))
	}

	return decryptedPasswords, nil
}

// generateSalt generates a random salt
func generateSalt() ([]byte, error) {
	salt := make([]byte, 25) // You can adjust the length of the salt as needed
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// combineSaltAndPassword combines a salt with a password
func combineSaltAndPassword(salt []byte, password string) string {
	// Combine salt and password and return the result
	combined := hex.EncodeToString(salt) + password
	return combined
}

// message for all sub commands
func customUsage() {

	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	fmt.Println("\n Password Options::")
	fmt.Println(" -h -help\t\t Show this help message")
	fmt.Println("")
	fmt.Println("  -l --length\t\t Specify the password length. Default is 20.")
	fmt.Println("")
	fmt.Println("  -n --number-passwords \tSpecify the number of passwords to generate. Default is 1.")
	fmt.Println("")
	fmt.Println("  -c --clipboard \tCopy the generated password to the clipboard")
	fmt.Println("")
	fmt.Println("  -en --encrypt \tEncrypt a password or passwords with AES-256 and the same key")
	fmt.Println("")
	fmt.Println("  -de --decrypt \tDecrypt a password given the key and the encrypted password.")
	fmt.Println("")
	fmt.Println("  -o --output \tSave the generated password to a file")
	fmt.Println("")
	fmt.Println("  -ex --exclude-specific \tExclude specific characters from the password")
	fmt.Println("")
	fmt.Println("  -exl --exclude-lower \tExclude lowercase letters from the password")
	fmt.Println("")
	fmt.Println("  -exs --exclude-special \tExclude special characters from the password.")
	fmt.Println("")
	fmt.Println("  -exu --exclude-upper \tExclude uppercase letters from the password.")
	fmt.Println("")
	fmt.Println("  -exd --exclude-digits \tExclude digits from the password.")
	fmt.Println("")
	fmt.Println("  -t --time \tGive the Execution time from the Password(s) back.")

}

func handleSubcommands(help_command bool) {
	// Check if the -help flag is provided
	if help_command {
		customUsage()
		os.Exit(0)
	}

}

func main() {

	// Start Timer
	startTime := time.Now()

	intro := " _____               _____             _____       \n" +
		"|  __ \\             / ____|           / ____|      \n" +
		"| |__) |_ _ ___ ___| |  __  ___ _ __ | |  __  ___  \n" +
		"|  ___/ _` / __/ __| | |_ |/ _ \\ '_ \\| | |_ |/ _ \\ \n" +
		"| |  | (_| \\__ \\__ \\ |__| |  __/ | | | |__| | (_) |\n" +
		"|_|   \\__,_|___/___/\\_____|\\___|_| |_|\\_____|\\___/ \n" +
		"                                                  \n" +
		"                                                  \n" +
		"Author: K3res\n" +
		"Inspired by https://github.com/B0lg0r0v/PassGen\n"

	fmt.Println(intro)

	for i := 0; i <= 5; i++ {
		fmt.Println("")
	}

	// Color variable
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	// Variable for the GenPassword
	lowerCaseLetters := "abcdefghijklmnopqrstuvwxyz"
	upperCaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specialCharacters := "!#$%&'()*+,-./:;<=>?@[]^_`{|}~"

	//Declare the flags (Parse Arguments)
	//--------------------------------------------------------

	// Define flags (alternative to Parse arguments)

	// Override the default Usage function
	flag.CommandLine.Usage = customUsage

	help_command := flag.Bool("help", false, "Show this help message.")
	flag.BoolVar(help_command, "h", false, "")

	length := flag.Int("length", 20, "Specify the password length. Default is 20.")
	flag.IntVar(length, "l", 20, "")

	numbPassword := flag.Int("number-passwords", 1, "Specify the number of passwords to generate. Default is 1.")
	flag.IntVar(numbPassword, "n", 1, "")

	// Boolean : Change from false to true if the flag is use

	clipboard := flag.Bool("clipboard", false, "Copy the generated password to the clipboard. Only works with one password at a time.")
	flag.BoolVar(clipboard, "c", false, "")

	encrypt_password := flag.Bool("encrypt", false, "Encrypt a password with AES-256. Only works with one password at a time.")
	flag.BoolVar(encrypt_password, "en", false, "")

	exe_time := flag.Bool("time", false, "Give the Execution time from the Password(s) back.")
	flag.BoolVar(exe_time, "t", false, "")

	// Syntax -de "encryptedPassword,key"
	decrypt_password := flag.String("decrypt", "", "Decrypt a password given the key and the encrypted password.")
	flag.StringVar(decrypt_password, "de", "", "")

	//syntax -o=/.../password.txt
	output := flag.String("output", "", "Save the generated password to a file. (syntax: -o=/.../password.txt)")
	flag.StringVar(output, "o", "", "")

	//Exlude Options

	// String: default value is ""

	// syntax -ex="ABDH/$l940ki"
	exlude_specific := flag.String("exclude-specific", "", "Exclude specific characters from the password. (syntax -ex=ABDH/$l940ki)")
	flag.StringVar(exlude_specific, "ex", "", "")

	// Boolean for Exclude options

	exclude_lower := flag.Bool("exclude-lower", false, "Exclude lowercase letters from the password.")
	flag.BoolVar(exclude_lower, "exl", false, "")

	exclude_special := flag.Bool("exclude-special", false, "Exclude special characters from the password.")
	flag.BoolVar(exclude_special, "exs", false, "")

	exclude_upper := flag.Bool("exclude-upper", false, "Exclude uppercase letters from the password.")
	flag.BoolVar(exclude_upper, "exu", false, "")

	exclude_digits := flag.Bool("exclude-digits", false, "Exclude digits from the password.")
	flag.BoolVar(exclude_digits, "exd", false, "")

	// Parse command-line arguments
	flag.Parse()

	// Set the custom usage function
	flag.Usage = customUsage

	// Handle subcommands
	handleSubcommands(*help_command)

	if *decrypt_password != "" {
		// Split the input using a delimiter
		values := strings.Split(*decrypt_password, ",")

		// Check if both values are provided
		if len(values) == 2 {
			encryptedPassword := values[0]
			key := values[1]

			decodedKey, err := hex.DecodeString(key)
			if err != nil {
				fmt.Println(red("Error decoding key:", err))
				return
			}

			// Decrypt the password
			decryptedPassword, err := decryptPasswords([]string{encryptedPassword}, decodedKey)
			if err != nil {
				fmt.Println(red("Error decrypting password:", err))
				return
			}

			// Join the decrypted passwords into a single string
			decryptedPasswordStr := strings.Join(decryptedPassword, ", ")
			fmt.Printf("Decrypted Passwords: %s\n", decryptedPasswordStr)
			os.Exit(0)
		}

	}

	if *length < 1 || *length > 256 {
		fmt.Println(red("The minimum password length musst between 1 and 256 character"))
		os.Exit(0)
	}

	if *numbPassword >= 1 {

		// Generate a single key for all passwords
		key := generateSymmetricKey()

		// String variable for Clipboard Function
		var generatedPasswords []string

		for i := 1; i <= *numbPassword; i++ {
			fmt.Printf("Password %d\n", i)

			allowedCharacters := ""

			if !*exclude_lower {
				allowedCharacters += lowerCaseLetters
			}

			if !*exclude_upper {
				allowedCharacters += upperCaseLetters
			}

			if !*exclude_digits {
				allowedCharacters += numbers
			}

			if !*exclude_special {
				allowedCharacters += specialCharacters

			}

			if *exlude_specific != "" {

				allowedCharacters += *exlude_specific
			}

			password, err := generatePassword(*length, allowedCharacters, *exlude_specific)
			if err != nil {
				fmt.Println(red("Error during password generation:", err))
				os.Exit(1)
			}

			// define passwords here
			passwords := []string{password}

			if *encrypt_password == true {
				// Encrypt the password
				encryptedPassword, err := encryptPasswords(passwords, key)
				if err != nil {
					fmt.Println(red("Error encrypting password:", err))
					return
				}

				// Join the encrypted passwords into a single string
				encryptedPasswordStr := strings.Join(encryptedPassword, ", ")
				//fmt.Printf("Encrypted Passwords: %s\n", encryptedPasswordStr)

				salt, err := generateSalt()
				if err != nil {
					fmt.Println(red("Error generating salt:", err))
					return
				}

				saltcryptpassword := combineSaltAndPassword(salt, encryptedPasswordStr)
				fmt.Printf("Encrypted Password: %s\n", yellow(saltcryptpassword))

			} else {
				fmt.Println("Generated password:", green(password))
			}

			generatedPasswords = append(generatedPasswords, password)
			fmt.Println("-----------------------------------------")
		}

		// Join the generated passwords into a single string
		allPasswords := strings.Join(generatedPasswords, "\n")

		if *encrypt_password == true {
			// Print the symmetric key
			fmt.Printf("Symmetric Key: %s\n", yellow(hex.EncodeToString(key)))
			fmt.Printf(red("Please save your symmetric key in a secure and accessible location. Without the key, decryption is not possible."))
		}

		if *decrypt_password != "" {
			// Split the input using a delimiter
			values := strings.Split(*decrypt_password, ",")

			// Check if both values are provided
			if len(values) == 2 {
				encryptedPassword := values[0]
				key := values[1]

				decodedKey, err := hex.DecodeString(key)
				if err != nil {
					fmt.Println(red("Error decoding key:", err))
					return
				}

				// Decrypt the password
				decryptedPassword, err := decryptPasswords([]string{encryptedPassword}, decodedKey)
				if err != nil {
					fmt.Println(red("Error decrypting password:", err))
					return
				}

				// Join the decrypted passwords into a single string
				decryptedPasswordStr := strings.Join(decryptedPassword, ", ")
				fmt.Printf("Decrypted Passwords: %s\n", decryptedPasswordStr)
				os.Exit(0)
			}

		}
		// Copy the Password in Clipboard
		if *clipboard == true {
			err := CopyToClipboard(allPasswords)
			if err != nil {
				fmt.Println(red("Error copying to clipboard:", err))
				os.Exit(1)
			} else {
				fmt.Printf(blue("Password(s) copied to clipboard.\n"))
			}
		}

		// Text File for Generated Password
		if *output != "" {
			// Save all generated passwords to a file
			err := os.WriteFile(*output, []byte(allPasswords), 0644)
			if err != nil {
				fmt.Println(red("Error saving passwords to file:", err))
				os.Exit(1)
			}

			fmt.Println(blue("All passwords saved to " + *output + " file."))

		}

	} else {
		println(red("Number muss be bigger than 0"))
	}

	if *exe_time {
		// End Timer
		elapsedTime := time.Since(startTime)
		fmt.Printf("Execution time: %s\n", elapsedTime)
	}
}
