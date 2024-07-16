### Pomo CLI

A cli tool to time your study sessions in active/rest intervals.

Starting it with the default inputs is very simple.

```bash
$ pomo start

🍎 Let's get started!
   State: Active 🟢
   Interval: 1 
   Time Remaining: 24:59s
```

- `-d, --detach`: Will run in the background and only notify when your time is up.

For more detailed and specific sessions, you can use `pomo session`.

```bash
$ pomo session 

🍎 What are you studying? Biomed
🍎 How many intervals do you plan to do? 3 


🍎 Let's get started!
   State: Active 🟢
   Interval: 1/3
   Time Remaining: 24:59s
   ...

🍎 Time is up!
🍎 Create a log entry for this interval? [y/n]: 
```

You can set your active, rest times, and more with `pomo config`.

```bash
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


