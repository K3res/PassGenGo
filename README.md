# PassGenGo
Password generator in GO with some extra options.</br>
Based on the Python Version https://github.com/B0lg0r0v/PassGen. <-----------Check it out!

                          
![grafik](https://github.com/K3res/PassGenGo/assets/89378576/add97f67-24d8-43c7-8809-5726d106ed12)


  

# Table of Content
- [PassGenGo](#passgengo)
  * [Remarks](#note)
  * [Usage](#usage)
  * [Examples](#examples)
  * [Installation](#installation)
  * [Disclaimer](#disclaimer)



## Remarks

In this program, there are two functions for encryption.  
  1. `-en` is used to encrypt the newly generated password(s) within the program.
  2. `-enxp` is used to encrypt any single input password.<br/>

For the decryption option `-de`, you must first input the encrypted password, then set a comma `,` and finally input your key.<br/>
Concerning the three options `-o`, `-ex`, and `-enxp`, you must set the arguments with the `=""` symbol to avoid your terminal misinterpreting some special characters. <br/> 


## Usage
```

Password Options:
  -h, -help                                                       Show this help message
  -l LENGTH, --length LENGTH                                      Specify the password length. Default is 20.
  -n NUMBER-PASSWORDS, --number-passwords NUMBER-PASSWORDS        Specify the number of passwords to generate. Default is 1.
  -c, --clipboard                                                 Copy the generated password to the clipboard
  -en, --encrypt                                                  Encrypt a password or passwords with AES-256 and the same key
  -enxp PASSWORD, --encrypt-ext-pass PASSWORD                     Encrypt a password with AES-256. Only works with one password at a time
  -de ENCRYPT-PASSWORD, Key, --decrypt ENCRYPT-PASSWORD, Key      Decrypt a password given the key and the encrypted password.
  -o OUTPUT-PATH, --output OUTPUT-PATH                            Save the generated password to a file

Options for Exclusion:
  -ex, EXCLUDE-SPECIFIC, --exclude-specific EXCLUDE-SPECIFIC      Exclude specific characters from the password
  -exl, --exclude-lower                                           Exclude lowercase letters from the password
  -exs, --exclude-special                                         Exclude special characters from the password.
  -exu, --exclude-upper                                           Exclude uppercase letters from the password.
  -exd, --exclude-digits                                          Exclude digits from the password.

Additional Options:
  -t, --time                                                      Give the Execution time from the Password(s) back.

``` 

## Examples
*Command:* ` go run ./passgengo.go -en -l 30 -n 2 -t -exs -exd`


![grafik](https://github.com/K3res/PassGenGo/assets/89378576/bf987ec3-895d-4eaf-8b7c-faf53c764d80)


*Command:* ` go run ./passgengo.go -enxp="$~ln-)+~kQ+6V}7g9md~"`

![grafik](https://github.com/K3res/PassGenGo/assets/89378576/bf696cb2-3a2e-4500-a18c-5917d97befc6)

*Command:* ` go run ./passgengo.go -de 470f55af3a31e4da6677fc8bede6afdcce96282841370e8a419b075ce7eec3cee866c17b1215f908be91a33a890921707e8ba4da80d7da7f8286f9725e92e40e,28f5f0e701f0f70e406582549c893759efc39f5e6237b1f8840f75e96cf438fe`

![grafik](https://github.com/K3res/PassGenGo/assets/89378576/117db0ed-4756-43a4-9b55-d8c1ae185043)

*Command:* ` go run ./passgengo.go -ex="euT58lkfDfg!#m." -o="Password.txt" -n 5`

![grafik](https://github.com/K3res/PassGenGo/assets/89378576/5248da45-9efa-4197-b2ec-bc9feb82c213)



## Installation

Preparation for PassGenGo
1. Download the reposoritory from GitHub: `git clone https://github.com/B0lg0r0v/PassGenGo.git`
2. Go into the new Folder with: `cd PassGenGo`

Install GO with the script: <br/>
1. Go to the `LinOG` Folder. <br/>
2. Make it executable: `chmod +x go_linux_install.sh`.<br/>
3. Run the bash script: `./go_linux_install.sh`.<br/>

Install missing GO packages:
1. To install the packages, run the GO file:  `go run ./passgengo.go`<br/>

![grafik](https://github.com/K3res/PassGenGo/assets/89378576/449476e0-b999-46e0-b090-cdde800d89a5) <br/>




## Disclaimer
This tool was just a small project for me to practice with the programming language GO. It is not intended to be the most secure and fastest program.







