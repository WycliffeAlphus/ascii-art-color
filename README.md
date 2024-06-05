## Overview
The ASCII Art Color is a program written in Go that colors a string or it's specified substring.

This program takes optional arguments of color(implemented using flag) and specific letters to be colored using ANSI color system.

The format of the input received as arguments is:

go run . --color=<color> <letters to be colored> "something"

## NOTE
This program builds on the ASCII Art from  **[ASCII-REPOSITORY](https://learn.zone01kisumu.ke/git/wyonyango/ascii-art.git)**

## Usage
To use the ASCII Art Generator, follow these steps:
1. Clone this repository to your local machine.
2. Navigate to the directory where the repository is cloned.
3. Build the program using the Go compiler:
4. Run the program from the command line, providing a string as an argument and optional color or specific letters to be colored
5. Output is the graphic representation of the input string using ASCII characters in specified color or specified letter graphics will be      colored.

## Example
Suppose you want to create a colored ASCII representation of the word "HeY GuYs wassgood aGuys".

 You would run the program with the following command:
```
go run . --color=blue GuYs "Hey GuYs\n wassgood aGuYs
```

The output would be:

![alt text](<Screenshot from 2024-06-03 17-14-33.png>)


## Dependencies
This program requires Go (Golang) to be installed on your system. You can download and install it from the [official Go website](https://golang.org/dl/).

## Contributing
Contributions to this project are welcome! If you'd like to contribute, please fork the repository and submit a pull request with your changes.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
