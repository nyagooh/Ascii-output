# <span style= "color:cyan">[ASCII-ART](https://www.zone01kisumu.ke/)</span>

![Masterpiece Generator](https://i.pinimg.com/564x/de/04/ad/de04ad9ac6c366872710b53871751517.jpg)
Here is a masterpiece generator that takes text input and transforms it into stunning ASCII using the specific bannar file passed.
With the FS program which is built on ASCII ART program, you can turn a simple text into a mesmerizing graphic masterpiece by specifying the banner file to be used. 
 ## <span style="color:#80A4ED">Files featured</span>
The program works with 3 potent files:
 1. standard.txt
 2. shadow.txt
 3. thinkertoy.txt
 
 The default file used is standard textual file.

## <span style="color:#80A4ED">How it Works</span> 
- The program takes a string input and the banner filefrom the command line.
```bash
go run . hello standard
```
- It checks if the input string contains any unprintable characters (outside the ASCII range 32-126). If it does, it prints an error message and exits.eg.
```bash
go run . 你好
The program will exit in this casee
```
- It checks if the banner file is valid.If not it prints an error e.g
``` bash
go run . hello stan
no such file or directory
```
- It then checks if the input string contains any tabs (\t). If it does, it replaces them with four spaces ( ) eg.
```bash
go run . hello\\tworld
 _                                         
| |                                        
| |__                           __      __ 
|  _ \                          \ \ /\ / / 
| | | |                          \ V  V /  
|_| |_|                           \_/\_/   
```
- It then checks if the input string contains any newlines (\n). If it does,The program splits the input string into individual lines using the newline character (\n). eg.
```bash
go run . h\\nw
```
- For each line, it applies a set of predefined functions to convert each character to its corresponding ASCII art representation. 
- The program then concatenates the resulting ASCII art representations for each line to form the final output. 

## <span style="color:#80A4ED">Installation</span>

To install the program from Gitea, follow these steps:

- Clone the repository from Gitea using the provided URL in your terminal.
``` bash
$git clone https://learn.zone01kisumu.ke/git/rogwel/ascii-art
```
Ensure you have Go installed on your system. If not you could follow the link below:
```bash
https://go.dev/doc/install
```
Navigate to the project directory using the command:
    
```
$cd ascii-art
```

and run the program using ***go run .*** ,input your text  and the banner file you want to use.eg
```
go run . Helloworld standard
```
The program will output :
``` bash
 _    _          _   _                                     _       _  
| |  | |        | | | |                                   | |     | | 
| |__| |   ___  | | | |   ___   __      __   ___    _ __  | |   __| | 
|  __  |  / _ \ | | | |  / _ \  \ \ /\ / /  / _ \  | '__| | |  / _` | 
| |  | | |  __/ | | | | | (_) |  \ V  V /  | (_) | | |    | | | (_| | 
|_|  |_|  \___| |_| |_|  \___/    \_/\_/    \___/  |_|    |_|  \__,_| 
                                                                      
                                                                      
```
## <span style = "color: #80A4ED">Files featured</span>
The program works with 3 potent files:

 1. standard.txt
 2. shadow.txt
 3. thinkertoy.txt
 - The default file used is standard textual file.

 ## <span style="color:#80A4ED">The Final Stretch</span>

I hope you've found the information helpful in getting started with our project.Don't hesitate to reach out if you have any questions or for pull requests.

Happy coding!😊

## <span style = "color:#80A4ED">License</span> 
[license](/home/nymaina/Desktop/fs-art/LICENSE)

