# goscripts
A small program to execute scripts hourly/daily using Docker

You need to have a folder containing hourly/ and daily/ folder, where you store your scripts you want to have executed periodically. e.g.

```bash
> tree /home/User/mountedData
mountedData
├── daily
│   └── backup.sh
└── hourly
    └── dnsupdate.sh
```

## To run
```bash
docker build -t elangenhan/goscripts github.com/elangenhan/goscripts && docker run -d --name elangenhan/goscripts -v /path/to/folder:/data elangenhan/goscripts
```
