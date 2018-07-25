import os, parseopt, json

proc writeHelp(): void = 
    echo("""
    dn [command] [id] [-j|--json] [-r|--reverse] [--limit|-l] [--sort d[ate] | t[itle] ]
    dn [-c|--config]
    dn [-l|--local]  [id]  [-j|--json] [--limit|-l] [--sort d[ate] | t[itle] r[everse] ]
    dn [a|add] <title>  [-b BODY|-t| - ] [-l| --local] [--prehook | --posthook] 
    dn [d|del]  [id] [-l| --local] [--prehook | --posthook] [--server ip:port]
    dn [e|edit] [id] [-l| --local] [--prehook | --posthook]
    dn [s|search] [-l| --local] [-b] [--server ip:port] search a note by title -b to search body
    dn [p|push] <id> [--all] [--clobber] [--server ip:port] 
    dn [g|get] <id> [--server ip:port]
    dn --server address[:port]""")
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
