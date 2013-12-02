GoPro (Go Programs)
=====
A small repository that I created in order to have some nice flexible data structures and the likes in order
to make my future development a bit easier, as I won't have to start from scratch every time. Feel free to
clone this repository and use my code as you wish, but it will be at your own risk. I won't take any responsibility
if your code misbehaves.

Usage
=====
To include the data structures in your projects:
* Simply clone this repository anywhere you like, then copy as many root folders as you need into your source 
folder (usually in your `src` folder in your Go workbench).
* Then include the directory path in your .go files (i.e. "ds/queue").

To run the test files in the `/test` folder:
* Change to the test folder that should be at the root.
* Build all tests with the `go test` command.
* If you do not wish to have all the tests built at once use `go test <name of test file>`.
* Run your test binary which should now appear in the same folder.  

All data structures should be self-contained so there should not be any dependency issues if you only include 
one or two files.
