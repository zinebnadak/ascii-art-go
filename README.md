# ascii-art
 
A Go command-line tool that turns text into big ASCII-art banners, for anyone who wants large text in the terminal.
 
## Example
 
Input:
 
```bash
go run . "hi"
```
 
Output:
 
```
 _       _  
| |     (_) 
| |__    _  
|  _ \  | | 
| | | | | | 
|_| |_| |_| 

```
 
## Usage
 
```bash
git clone https://github.com/zinebnadak/ascii-art-go.git
cd ascii-art-go
go run . "Hello\nThere"
```
 
- `\n` in the text starts a new banner line.
- Use SINGLE QUOTES for text with `!` in zsh: `go run . 'Hi!'`

## Tests
 
Compare the output with the examples in the subject (`cat -e` shows a `$` at the end of each line):
 
```bash
go run . "Hello\n\nThere" | cat -e
go vet ./...
```
 
## Tech
 
- Language: Go
- Libraries: standard library only (`fmt`, `os`, `strings`)

## What I learned
 
- Reading a file with `os.ReadFile` and splitting it into lines with `strings.Split`
- Using a character's ASCII code to calculate where its drawing starts in a file
- Building output row by row with nested loops
- Handling `\n` typed on the command line, which arrives as two characters
---
 
Built by Mohamed & Zineb at [grit:lab](https://gritlab.ax), Åland (01-edu peer-learning program).