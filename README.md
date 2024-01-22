# PassGenGo
Password generator in GO with some extra options. Based on the Python Version https://github.com/B0lg0r0v/PassGen
<div align=center>

 _____               _____             _____       
|  __ \             / ____|           / ____|
| |__) |_ _ ___ ___| |  __  ___ _ __ | |  __  ___
|  ___/ _` / __/ __| | |_ |/ _ \ '_ \| | |_ |/ _ \
| |  | (_| \__ \__ \ |__| |  __/ | | | |__| | (_) |
|_|   \__,_|___/___/\_____|\___|_| |_|\_____|\___/


Author: K3res
Inspired by https://github.com/B0lg0r0v/PassGen


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
for -o, -ex and -enxp use the ="" format



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


![grafik](https://github.com/K3res/PassGenGo/assets/89378576/bf987ec3-895d-4eaf-8b7c-faf53c764d80)


*Command:* ` go run .\passgengo.go -enxp="$~ln-)+~kQ+6V}7g9md~"`

![grafik](https://github.com/K3res/PassGenGo/assets/89378576/bf696cb2-3a2e-4500-a18c-5917d97befc6)

*Command:* ` go run .\passgengo.go -de 470f55af3a31e4da6677fc8bede6afdcce96282841370e8a419b075ce7eec3cee866c17b1215f908be91a33a890921707e8ba4da80d7da7f8286f9725e92e40e,28f5f0e701f0f70e406582549c893759efc39f5e6237b1f8840f75e96cf438fe`

![grafik](https://github.com/K3res/PassGenGo/assets/89378576/117db0ed-4756-43a4-9b55-d8c1ae185043)

*Command:* ` go run .\passgengo.go -ex="euT58lkfDfg!#m." -o="Password.txt" -n 5`

![grafik](https://github.com/K3res/PassGenGo/assets/89378576/5248da45-9efa-4197-b2ec-bc9feb82c213)



## Installation

Install GO with Script: <br/>

1. Download go_linux_install.sh in a new dictonary <br/>
2. Make it executable: `chmod +x go_linux_install.sh `<br/>
3. Run the bash: `./go_linux_install.sh`<br/>

Instal missing GO packages:
1. Download the complet passgengo folder in a dictonary<br/>
2. to install the packages run the GO file:  `go run .\passgengo.go`  <br/>

![grafik](https://github.com/K3res/PassGenGo/assets/89378576/449476e0-b999-46e0-b090-cdde800d89a5) <br/>




## Disclaimer
(Dieses Tool war nur ein klieines project für mich um mit der programiersprache GO zu üben. Es ist nicht dafür gedacht die sicherste und schnellste programm zu sein.  )







