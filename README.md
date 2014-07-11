goblin
======

plan9 inspired command line commands implemented in pure Go


### Installation:
`go get github.com/mortdeus/goblin && goblin install`

While running the `goblin install` command, the goblin tool will
chdir into each './cmd/$tool' directory in order to run 
os.Exec("go build") and compile each tool from source. 
If the tool is sucessfully compiled into an executable,
the go tool then moves that binary to $GOTOOLDIR.
##### Caution:
At the moment there isn't a mechanism in place that checks 
whether the tools in $GOTOOLDIR, (specifically the tools that don't belong to
goblin), have filenames that clash with goblin's tools.

Therefore `goblin install`'s current behavior is to just throw out the old tools
and replace them with the new. In most cases this shouldn't be an issue 
because goblin will not install any goblin tools that would clobber a
standard go tool. (i.e `"*g", "*l", "*a", "*c", "yacc", "objdump" etc)



### Usage:
`go tool $cmd`

 run `go tool` to get the list of tools you can call with go tool. 



### License:
Same license used for Go.
http://golang.org/LICENSE
##### Notice:
While this code is definitely inspired by plan9's tools, these
tools are not a direct c to go source **port** of the original plan9 tools.

All goblin code is written from scratch the ground up, Therefore 
goblin software is not legally tied the Lucent Public License, or any
other licenses that plan9/inferno software has been licensed and distributed
under.

Futhermore, any code that is contributed to goblin may be rejected if
it's blantantly obvious that the code being commited to goblin is a 
derivation of plan9's source code.


To learn more about the reasoning behind this policy:

http://directory.fsf.org/wiki/License:LucentPLv1.02

