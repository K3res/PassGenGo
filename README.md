# PassGenGo
Password generator in GO with some extra options. Based on the Python Version https://github.com/B0lg0r0v/PassGen
<div align=center>
![grafik](https://github.com/K3res/PassGenGo/assets/89378576/34a00bab-3dab-45b7-9c4c-aee04019dd77)

</div>

# Table of Content
- [PassGenGo](#passgengo)
  * [Remarks](#note)
  * [Usage](#usage)
  * [Examples](#examples)
  * [Installation](#installation)
  * [Disclaimer](#disclaimer)



## Remarks
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
*Command:* ` go run .\passgengo.go -en -l 30 -n 2 -t -exs -exd`

<div align=center>
![grafik](https://github.com/K3res/PassGenGo/assets/89378576/bf987ec3-895d-4eaf-8b7c-faf53c764d80)

</div>

<div align=center>

</div>

<div align=center>

</div>

<div align=center>

</div>

-en 

-enxp

-de

-o

-ex



## Installation


Install GO with Script: <br/>

1. download go_linux_install.sh in a new dictonary <br/>
2. make it executable chmod +x go_linux_install.sh <br/>
3. run the bash ./go_linux_install.sh<br/>

download the complet passgengo folder<br/>
to download the package jsut run the file:  go run .\passgengo.go  <br/>

![grafik](https://github.com/K3res/PassGenGo/assets/89378576/449476e0-b999-46e0-b090-cdde800d89a5) <br/>




## Disclaimer
(Dieses Tool war nur ein klieines project für mich um mit der programiersprache GO zu üben. Es ist nicht dafür gedacht die sicherste und schnellste programm zu sein.  )







