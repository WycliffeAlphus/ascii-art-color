## Overview
The ASCII Art Color is a program written in Go that colors a string or it's specified substring.

This program takes optional arguments of color(implemented using flag) and substring to be colored using ANSI color system,RGB or HEXcolor codes.

The format of the input received as arguments is:

go run . --color= **color** substring **"something"**

## NOTE
This program builds on the ASCII Art from  **[ASCII-REPOSITORY](https://learn.zone01kisumu.ke/git/wyonyango/ascii-art.git)**

## Usage
To use the ASCII Art Generator, follow these steps:
1. Clone this repository to your local machine using the comand below.
``` bash
git clone https://learn.zone01kisumu.ke/git/lakoth/ascii-art-color
```

2. Navigate to the directory where the repository is cloned.
```bash
cd ascii-art-color
```

3. Initialize your module to get the required dependancies :Run the command below to initialize your module
```bash
go mod init ascii-art color
```
4. Run the program from the command line, providing a string as an argument and optional color or specific letters to be colored.


## Example 1
Suppose you want to create a colored ASCII representation of the string "HeY GuYs\nwassgood aGuys" and color GuYs in blue.

 You would run the program with the following command:
```bash
go run . --color=blue GuYs "Hey GuYs\n wassgood aGuYs"
```
Output is the graphic representation of the input string using ASCII characters in specified color or specified letter graphics will be colored.

![alt text](<Screenshot from 2024-06-03 17-14-33.png>)

## Example 2
Suppose the substring to be colored is not specified,the whole string will be colored in the given color. 
You would run the program with the following command:
```bash
go run . --color=red "Hello World"
```

Output is the graphic representation of the colored input string.

![alt text](<Screenshot from 2024-06-11 15-06-18.png>)

## Example 3
Suppose the substring does not exist in given string, a non colored ASCII pattern will be displayed. 
You would run the program with the following command:
```bash
go run . --color=blue HoW "Hello World"
```

Output is the graphic representation of the input string .

![alt text](<Screenshot from 2024-06-11 15-16-33.png>)


## Dependencies
This program requires Go (Golang) to be installed on your system. You can download and install it from the [official Go website](https://golang.org/dl/).

## Contributing
Contributions to this project are welcome! If you'd like to contribute, please fork the repository and submit a pull request with your changes.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

