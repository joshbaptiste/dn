## Distributed Notes (DN) - Command line note snippet application

- A simple, Client/Server command line note taking tool written in Nim - Inspired by theca and pet
## Features
- Client/Server Support
- JSON profile format for easy scripting/integration
- Add/edit/delete notes
- Transfer notes between client/server
- Add/edit note body or title using command line arguments, STDIN, or using the editor set via $VISUAL or $EDITOR
- Search notes (title or body using keyword or regex pattern)

##  Client Usage

    dn [command] [id] [-j|--json] [-r|--reverse] [--limit|-l] [--sort d[ate] | t[itle] ]
    dn [-c|--config]
    dn [-l|--local]  [id]  [-j|--json] [--limit|-l] [--sort d[ate] | t[itle] r[everse] ]
    dn [a|add] <title>  [-b BODY|-t| - ] [-l| --local] [--prehook | --posthook] 
    dn [d|del]  [id] [-l| --local] [--prehook | --posthook] [--server ip:port]
    dn [e|edit] [id] [-l| --local] [--prehook | --posthook]
    dn [s|search] [-l| --local] [-b] [--server ip:port] search a note by title -b to search body
    dn [p|push] <id> [--all] [--clobber] [--server ip:port] 
    dn [g|get] <id> [--server ip:port]
    dn --server address[:port]

### Default Output
```bash
josh@node $ dn
id | user |  title                          updated 
--------------------------------------------------------------------------------------
1 | josh | app_do cycle application         2017-02-02 03:55:53
2 | john | LQS list queues and TID          2016-07-08 16:35:44
3 | danny | Houston commands                2016-07-14 17:18:49
4 | curtis |  LQS view items                2016-07-20 15:26:12
5 | josh |  LQS move items                  2016-07-20 15:31:21
6 | fred |  LQS view en[de]-queue rates     2016-07-20 16:42:21
```

### Adding notes
```bash
josh@host $  dn a "First snippet with no body"
josh@host $  dn a "Second snippet with body" -b "Body of my snippet"
josh@host $  dn a "Third snippet with body from EDITOR"  -e 
josh@host $  echo "body from here" | dn a "Fourth snippet with body from STDIN"  -b -
```
### Display notes
```sh
josh@host $ dn # Display notes (if DN_SERVER is defined or in config, dn will display server side notes, add -l to force local note listings)
josh@host $ dn 1 # Display note 1
josh@host $ dn --sort t --reverse # Display notes sorted by titles in reverse order (default sort is date DESC)
josh@host $ dn 0 # Displays sample DN usages and stores server config not displayed by default only local
```
### Editing notes
```bash
josh@host $  dn e # By default edits last node in EDITOR
josh@host $  dn e 12 # Edit note 12 in EDITOR
josh@host $  echo "Replace last note body with this" | dn e - 
josh@host $  echo "Replace note body #2 with this" | dn e 2 - 
josh@host $ dn e 0 # to edit note config
```
### Deleting notes
```bash
josh@host $ dn d       # deletes last note
josh@host $ dn d 12 15 # delete notes 12 and 15
josh@host $ dn d 16-30  # delete range of notes 16 through 30
```

### Search Notes
```sh
josh@client $ dn s linux # incase-sensitive Server title search 
josh@client $ dn s -b linux # incase-sensitive Server body search 
```
## Running a Server
- dn can also run as a server where notes can be sent to a central server (Please read Limitations of server running)

#### Server Usage
    dn --sip <Listening Address>  --name <Unique Server Name> [ --log <path/to/logfile.log> ] [--port 6464] [--allow-write] 
    
#### Server Example 
```sh
josh@server $ dn --sip 10.10.10.10 --port 6464 (default TCP 6464) --name 'Josh Notes' --log /path/to/logfile.log --allow-writes
Starting Server 'Josh Notes' on 10.10.10.10:6464 (Writes allowed) 
```

#### Display server notes
```sh
# If Server is NOT ReadOnly
josh@client $ dn --server 10.10.10.10  #no need to specify port if default is used --server address:port
josh@client $ DN_SERVER=10.10.10.10:6464 dn # if DN_SERVER env variable is set server notes are retrieved
josh@client $ dn e 0 # if "dn_server" is set in config db (dn uses id 0 as it's config) it will be used and conect to server
josh@client $ dn -l # force local note listing regardless of above settings 
```
#### Search Server Notes
```sh
# Assuming DN_SERVER variable is set or dn_server config is set or --server is used
josh@client $ dn s linux # incase-sensitive Server title search 
josh@client $ dn s -b linux # incase-sensitive Server body search
josh@client $ dn s -l linux # force local search regardless of server settings 
```

#### Delete server notes
```sh
# If Server is NOT ReadOnly and your username matches username of the corresponding note[s]
josh@client $ dn d 100 --server 10.10.10.10  # Delete server note 100
josh@client $ DN_SERVER=10.10.10.10:6464 dn d 4 # Delete note 4
josh@client $ dn d 13 # Assuming DN_SERVER variable is set or dn_server config is set
josh@client $ dn d 13 -l  # delete local note 13 regardless of above settings 
```

#### Push notes to Server
```sh
#If Server is NOT ReadOnly and your username matches username of the corresponding note[s]
josh@client $ dn p 100 --server 10.10.10.10  # Push local note 100 to Server appending 
josh@client $ DN_SERVER=10.10.10.10:6464 dn p 4 # Push local note 4 to Server appending
josh@client $ dn p 13 # Assuming DN_SERVER variable is set or dn_server config is set
josh@client $ dn p 13 --clobber 41  # push local note 13 and overwrite server note 41
josh@client $ dn p --all #push all local notes appending to server
```

#### Get notes from server
```sh
josh@client $ dn g 100 --server 10.10.10.10  # Get Server note 100 and append locally
josh@client $ DN_SERVER=10.10.10.10:6464 dn g 4 # get server note 4 and append locally
josh@client $ dn g 13 # Assuming DN_SERVER variable is set or dn_server config is set
josh@client $ dn g 13 --clobber 41  # get server note 13 and overwrite local note 41
josh@client $ dn g --all # Get all Server notes appending to local
```
### Environment Varables
```sh
DN_SERVER - <Host>[:port] # DN_SERVER=10.10.10.5:1234 
DN_CONFIG - overrides default config path (by default client creates $HOME/.config/dn.json store/config file)
```

### Inspiration
- I wanted a simple note taking and command line snippet appplication for saving hundreds of my Linux/BSD command snippets  and notes locally (as we are retricted from using online tools), but with only some of the features of a full blown note taking app and be able to share with colleagues at work who can also add their cool snippets. 

### Server Limitations
- The Server does not provide any type of user authentication or encryption, It's trivial for a user to overwrite other users notes which is why server starts up readonly by default. dn is designed to be used within a trusted network with trusted colleagues. Though internet usage is not recommended, if your server runs on the internet or you're being intercepted by certain 3 letter agencies you can use [spiped](https://www.tarsnap.com/spiped.html) to create encrypted tunnel between client/server.
