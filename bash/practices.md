
## How to write good bash scripts

### Steps to follow

1. Must check the return code of your commands

```bash
# check return code and handle errors
cp source.txt dest/ || { echo "copy failed (rc=$?)" >&2; exit 1; }
```

2. You can save the status of successfu tasks and then retry from that point

```bash
# record progress and resume
echo 3 > progress.txt   # completed step 3
step=$(cat progress.txt)
if [[ $step -lt 5 ]]; then
	./do_step_4 && echo 4 > progress.txt
fi
```

3. Bash 'trap' is your friend, use it for cleaning and error handling

```bash
# cleanup temporary files on exit or error
tmpdir=$(mktemp -d)
trap 'rm -rf "$tmpdir"; echo "cleaned"' EXIT INT TERM
# ... work with $tmpdir
```

The above command above runs when the script exits for any reason
- avoid complex logic in a trap that itself fail

The `trap` command registers commands (shell command string)
- to run when the shell receives specified signals or when certain shell events happen (like EXIT)

``` bash
trap 'commands' SIGNALS
```

4. When downloading data assume its corrupted, never overwrite your good data set with fresh data.

```bash
# download to a staging file, verify, then move
curl -fSL -o data.csv.part "$url"
if sha256sum -c data.csv.sha256; then
	mv data.csv.part data.csv
else
	echo "download verification failed" >&2; rm -f data.csv.part
fi
```

5. Take advantage of `journalctl` and custom fields

```bash
# write a custom journal message and query it
logger -t myscript -- "stage=fetch status=ok file=data.csv"
journalctl -t myscript --no-pager
```

Logger sends messages to the system logging facility, its a simple CLI for writing an entry into the system log with optional tag, priority, and other metadata


#### Where to read the logs

systemd Linux: use journalctl
- By tag: journalctl -t myscript
- Recent: journalctl -u some-service --since "1 hour ago" (service-specific)

non-systemd Linux (or generic syslog): check /var/log/syslog or /var/log/messages
- e.g. tail -n 100 /var/log/syslog | grep myscript

macOS: open Console.app or use log:
- e.g. log show --predicate 'process == "myscript"' --last 1h (predicate formatting may vary)

6. Check statuses of background tasks, just remember to save the PIDs

```bash
# run background jobs and wait/check
long_task &
pid=$!
# later
if kill -0 "$pid" 2>/dev/null; then echo "running"; else echo "finished"; fi
wait "$pid" || echo "task failed (rc=$?)"
```

THe `$!` expands to the PID of the most recently executed background command in the current shell.

7. Use bash linter like shellcheck

```bash
# lint this script
# shellcheck bash/practices.md || true
shellcheck ./myscript.sh
```