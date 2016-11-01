# raptor
Minimal terminal sharing tool 

## Install

```
$ go get github.com/agupta666/raptor
```

## Usage

To share a terminal session use `raptor share` command

```
$ raptor share
Token: 6b8ab9d7-b1ea-47f5-ac6d-a2e0e8108875
bash-3.2$ 
```
share the token to anyone who wishes to join the session ...

To join a terminal session, use the token shared with you to join the session and start collaborating

```
$ raptor join -s <remote-ip> -t 6b8ab9d7-b1ea-47f5-ac6d-a2e0e8108875
bash-3.2$
```

to set your own token while sharing a session use the -t flag

```
$ raptor share -t 1011
Token: 1011
bash-3.2$
```


