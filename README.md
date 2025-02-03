# Build Your Own wc Tool

## From https://codingchallenges.fyi/challenges/challenge-wc

This is a code challenge where we create a clone of the `wc` command in terminal. To create the clone, called ccwc, Golang was used, since Golang has built in utilities which make it ideal for a project like this.

To run locally:

1. Make sure you have golang installed
2. Git clone https://github.com/creativeflow3/code-challenge-01-wc
3. Go to https://codingchallenges.fyi/challenges/challenge-wc and download the test file, or use your own test file if you wish.
3. Go to the directory, and run a command, such as `go run ccwc.go -w test.txt`
4. The options:
    - `-c`: File size in bytes
    - `-w`: Words of text in the file
    - `-m`: Characters in a file
    - `-l`: Lines in a file 
5. You can also build and install the package locally. If you wish to do so, follow the instructions here: https://go.dev/doc/tutorial/compile-install
6. You can then run like `ccwc -w test.txt`