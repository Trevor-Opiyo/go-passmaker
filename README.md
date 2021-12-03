<img src="passmaker.png" width="600"/>

### A password generator CLI written in Go.

This password generator has sets variability and security as first order priorities. go-passmaker chooses and indeterminate amount of each character type to add to the final password. This increases the variability of the password and presents users with the ability to regenerate the password immediately if it does not fulfill their criteria. (If the password doesn't meet your needs hit 'Y' to regenerate!)

### Prerequisites

Go 1.16 +

### Installation instructions

Clone a copy of the repository.

`git clone https://github.com/Trevor-Opiyo/go-passmaker.git`

Set the repository as the current directory.

`cd go-passmaker`

Produce an executable.

`go install`

If you have not, add the `~/go/bin` folder to your environment

Run the program

`passmaker`
