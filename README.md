# PassGenGo
Password generator in GO with some extra options. Based on the similiar Python Version https://github.com/B0lg0r0v/PassGen

![grafik](https://github.com/K3res/PassGenGo/assets/89378576/34a00bab-3dab-45b7-9c4c-aee04019dd77)



# Table of Content
- [PassGenGo](#passgengo)
  * [Differences from passgen.py V1.1.1](#note)
  * [Usage](#usage)
  * [Examples](#examples)
  * [Installation](#installation)
  * [Disclaimer](#disclaimer)



## Differences from passgen.py V1.1.1
-t to show the execuritv time
-en to encrypt directly the generate password



## Usage
```

usage: go run passgengo.go [Options]

Password Options:
  -h, -help                                                       Show this help message
  -l LENGTH, --length LENGTH                                      Specify the password length. Default is 20.
  -n NUMBER_PASSWORDS, --number-passwords NUMBER_PASSWORDS        Specify the number of passwords to generate. Default is 1.
  -c, --clipboard                                                 Copy the generated password to the clipboard
  -en, --encrypt                                                  Encrypt a password or passwords with AES-256 and the same key
  -de KEY PASSWORD, --decrypt KEY PASSWORD                        Decrypt a password given the key and the encrypted password.
  -o OUTPUT-PATH, --output OUTPUT-PATH                            Save the generated password to a file

Options for Exclusion:
  -ex, EXCLUDE_SPECIFIC, --exclude-specific EXCLUDE_SPECIFIC      Exclude specific characters from the password
  -exl, --exclude-lower                                           Exclude lowercase letters from the password
  -exs, --exclude-special                                         Exclude special characters from the password.
  -exu, --exclude-upper                                           Exclude uppercase letters from the password.
  -exd, --exclude-digits                                          Exclude digits from the password.

Additional Options:
  -t, --time                                                      Give the Execution time from the Password(s) back.


``` 

## Examples

## Installation

## Disclaimer








