import os, parseopt, json

proc writeHelp(): void = 
    echo("dn --config=/path/to/config.json")
    quit("No parameters passed.. exiting")

proc parseConfigFile(file: string): void =
        

if  paramCount() == 0:
    writeHelp()
else:
    var cmd = initOptParser()
    for kind, key, val in cmd.getopt():
        if kind == cmdLongOption:
            if key == "config":
                parseConfigFile(val)
