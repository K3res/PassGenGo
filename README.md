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
-de encrypt-password, key 
for -o -enxp use the ="" format



## Usage
```

Password Options:
  -h, -help                                                       Show this help message
  -l LENGTH, --length LENGTH                                      Specify the password length. Default is 20.
  -n NUMBER_PASSWORDS, --number-passwords NUMBER_PASSWORDS        Specify the number of passwords to generate. Default is 1.
  -c, --clipboard                                                 Copy the generated password to the clipboard
  -en, --encrypt                                                  Encrypt a password or passwords with AES-256 and the same key
  -enxp PASSWORD, --encrypt-ext-pass PASSWORD                     Encrypt a password with AES-256. Only works with one password at a time
  -de ENCRYPT-PASSWORD, Key, --decrypt ENCRYPT-PASSWORD, Key      Decrypt a password given the key and the encrypted password.
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
-en 

-enxp

-de

-o

-ex



## Installation

Linux <br/>
Install GO with Script: <br/>

download go_linux_install.sh in a new dictonary <br/>
make it executable chmod +x go_linux_install.sh <br/>
run the bash ./go_linux_install.sh<br/>

download the complet passgengo folder and go in<br/>
to download the package jsut run the file <br/>
go run passgengo.go  <br/>

![grafik](https://github.com/K3res/PassGenGo/assets/89378576/449476e0-b999-46e0-b090-cdde800d89a5) <br/>




Windows <br/>

Mac OS <br/>


## Disclaimer
(Dieses Tool war nur ein klieines project für mich um mit der programiersprache GO zu üben. Es ist nicht dafür gedacht die sicherste und schnellste programm zu sein.  )







