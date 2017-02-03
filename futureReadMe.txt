server side only local user can delete anyones notes
client user can only delete his notes


Distrubrited Notes is inspired by theca and task spooler

dn 0 (Displays a sample of DN usages) and stores server config not displayed by default only local
dn [options] <id>--dec decrypt note --json|-j -l # | --limit # will be server side if defined
dn [l|local]  <id>--dec decrypt note --json|-j -l # | --limit # force local listing if servers are defined resort to local listing [#]by default dn lists all notes
dn [a|add]  add alocal  note --enc encrypt the note --password or --keyfile  --label append a label to note --prehook/--posthook
dn [d|del] del a local note
dn [e|edit] [#] edit a local note -- By default edits last additionnd tries to run EDITOR will go to STDIN if not avail
dn [s|search] search a note by subject -b to search body also -l labels
dn [p|push] <#> --all  push note[s] to server defined in  DN_SERVER or DN_SERVER in config or --dnserver ip[:default port 6011]]
dn [g|get] <#> get note[s] and append to my local note json store
dn [c|config] displays the  config

dn e 0 to update config
dn 1 Display note 1 (local note by default)



if  DN_SERVER is defined by default it will be used when listing
DN_SERVER overrides config
###########

I wanted to created a client/server note app that i can just look up commands and not have to leave my
terminal looking therpughout scattered webpages, bland also a simple binary that colleagues can view and utlize notesand update with their
command line notes, we use it mostly for descriping commands to be used operationally

limit subject to 80 chars

(local or ServerName)
id | E | user |  title                   |         last touched |  label
1   app_do cycle application    (+)  2017-02-02 03:55:53
2   LQS list queues and TID     (+)  2016-07-08 16:35:44
3*  Houston commands            (+)  2016-07-14 17:18:49
4   LQS view items              (+)  2016-07-20 15:26:12
5   LQS move items              (+)  2016-07-20 15:31:21
6   LQS view en[de]-queue rates (+)  2016-07-20 16:42:21
7   SI webserver mode           (+)  2016-07-29 10:54:31

dn --config /path/to/dns.json ----server  10.10.10.10 --port 6464 (by default 6464)  --name 'name to be displayed'  --log /path/to/logfile.logi  '
by default server creates a dns.json file in cwd unless specified
only user can add/delete their own notes server side

db is designed to be run inside a secure network while I do use it privately to store encrypted passwords, collectively we place public commands
for wprl for easy retieval without and is not designed to face the  public internet for some reaso you's want to use it over the intber 
use spiped
