# mpw

This is an implementation of the [MasterpasswordApp](http://masterpasswordapp.com/) algorithm by [Maarten Billemont](http://www.lhunath.com/) in Go. 
His original source code can be found [here on GitHub](https://github.com/Lyndir/MasterPassword/).

As he only provides a C CLI program and a Java application I ported the algorithm to Go to build a GTK3 desktop app for my Linux system.

## Usage

Generating a password for a site is a 3-step process.

### Master Key generation

Call `mpw.GenerateIdenticon(userName, password string) (string, int)` to generate the identicon to show to the user while he enters his password.
When he has entered his master password, call `mpw.GenerateMasterKey(userName, masterPassword string) (*MasterKey, error)`. This gives you the `MasterKey` that will be used to generate passwords.
**Never store this key anywhere. Do not send it. "Forget" the master password immediately and keep this key until it is not needed anymore.**

### Password initialization

Call `mpw.GenerateMasterPasswordWithType(masterKey *MasterKey, siteName string, counter uint32, mpType MasterPasswordType) *MasterPassword` 
(or `GenerateMasterPassword(masterKey *MasterKey, siteName string, counter uint32) *MasterPassword` and `SetType` later) to obtain a `MasterPassword`. This is a password for a specific site.
The site name or counter cannot be changed anymore, generate a new one if these values should be changed. The type can be changed anytime, even after obtaining the password string.

### Password generation

call `MasterPassword.GetPassword() string` to obtain the password string. Display this to the user (or put it in his clipboard). 
**Do not save it anywhere permanently. It can be regenerated at will.**

## Demo

```Go
    masterKey, err := mpw.GenerateMasterKey("John Smith", "Super Secret Password")
	if err != nil {
		// Error handling here. Should only happen if scrypt fails for some reason.
	}
	
	masterPass := mpw.GenerateMasterPassword(masterKey, "google.com", 1)
    masterPass.SetType(MaximalPassword)
    fmt.Printf("pass is %s\n", masterPass.GetPassword())
```

## The Algorithm

The documentation of the used algorithm can be found on the [website](http://masterpasswordapp.com/algorithm.html), but sadly this documentation is not up to date.
Here are some clarifications and additions.

### Encoding of numbers

> Numbers (eg. `name length`) are encoded as 32-bit unsigned integers in network byte order.

This means, that whenever you should put a number into a byte array (for the master key or seed generation) you have to put 4 bytes (4 * 8 = 32) into the array.
In Go this can be done with the `binary.BigEndian.PutUint32(bs, num)` function

### Password Type Templates

2 templates are missing from the documentation but included in the official client:

| Template Name  | Templates |
|---|---|
| Name password  | `{"cvccvcvcv"}` |
| Phrase password | `{ "cvcc cvc cvccvcv cvc", "cvc cvccvcvcv cvcv", "cv cvccv cvc cvcvccv" }`|

### Identicon Generation

The official Java and C clients show an identicon that is derived from the user name and master password so you can check if you typed your password correctly easily.
This generation algorithm is undocumented and I extracted it from the `mpw-util.c` file from the original C code.

The identicon consist of 4 parts: a left arm, a body, a right arm and an accessory. These unicode characters also have a color associated with them.

The left arm can consist of `{ "╔", "╚", "╰", "═" }`. 

The body can consist of `{ "█", "░", "▒", "▓", "☺", "☻" }`.

The right arm can consist of `{ "╗", "╝", "╯", "═" }`.

The accessories can be `{"◈", "◎", "◐", "◑", "◒", "◓", "☀", "☁", "☂", "☃", "☄", "★", "☆", "☎", "☏", "⎈", 
"⌂", "☘", "☢", "☣", "☕", "⌚", "⌛", "⏰", "⚡", "⛄", "⛅", "☔", "♔", "♕", "♖", "♗", "♘", "♙", "♚", "♛", "♜",
 "♝", "♞", "♟", "♨", "♩", "♪", "♫", "⚐", "⚑", "⚔", "⚖", "⚙", "⚠", "⌘", "⏎", "✄", "✆", "✈", "✉", "✌" }`.
 
The identicon is generated from a seed. This seed is computed by using a HMAC with SHA256. The password of the user is the secret key and the user name is the message.

The left arm is then defined as `leftArmChar = leftArmPossibilities[ seed[0] % len(leftArmPossibilities)]`.

The body is defined as `bodyChar = bodyPossibilities[ seed[1] % len(bodyPossibilities)]`.

The right arm is defined as `rightArmChar = rightArmPossibilities[ seed[2] % len(rightArmPossibilities)]`

The accessory is defined as `accessoryChar = accessoryPossibilities[ seed[3] % len(accessoryPossibilities)]`

The unix shell color indetifies is defined as `color = seed[4] % 7 + 1`

#### Colors

The following tale contains the color numbers and their color

| Number  | Color |
|---|---|
|  1 | Red |
|  2 | Green |
|  3 | Yellow |
|  4 | Blue |
|  5 | Pink |
|  6 | Cyan |
|  7 | Grey |
 

