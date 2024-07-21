### Pomo CLI

A cli tool to time your study sessions in active/rest intervals.

### Quick start 

```bash
$ git clone https://github.com/maxcelant/pomo-cli.git
$ cd pomo-cli
$ chmod +x ./scripts/build.sh
$ ./scripts/build.sh
```

### Commands

Starting it with the default inputs is very simple using `pomo start`

```
$ pomo start

🍎 Time to focus
   State: Active 🟢
   Interval: 1 
   Time Remaining: 24:59s
```

- `-s, --silent`: Will run in the background and only notify when your time is up.
- `-i, --intervals`: Set the number of intervals you want to accomplish.
- `-l, --log`: It will prompt you to log your accomplishments after every rest.


You can set your active, rest times, and more with `pomo config`.

```
$ pomo config -a 45m -r 15m
```

- `-a, --active`: The amount of time to work
- `-r, --rest`: The amount of time to rest
- `-l, --link`: Link your obsidian to pomo so your notes go straight there.
- `-f, --file`: You can optionally send a `pomo.yaml` file with your options.

```yaml
pomo:
  active: 25m
  rest: 15m
  link: /link/to/obsidian
```


