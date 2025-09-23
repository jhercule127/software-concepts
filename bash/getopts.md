
## How to use getopts

This note summarizes how to use the builtin `getopts` for parsing short options (and a common pattern for supporting long options).

Key variables
- `OPTARG` — value of the current option's argument (set by `getopts` when the option requires an argument).
- `OPTIND` — index of the next argument to be processed; initialize or reset it to 1 when reusing `getopts` in the same shell.

Pattern characters in the option string
- A leading `:` in the optstring makes `getopts` behave in "silent" mode where you handle errors yourself (`:` and `?` cases in `case`).
- A colon after an option letter (e.g. `r:`) means that option requires an argument.

Basic short-option example

```bash
#!/usr/bin/env bash

while getopts ":ab:c" opt; do
    case $opt in
        a) echo "flag a set" ;;
        b) echo "option b with arg: $OPTARG" ;;
        c) echo "flag c set" ;;
        \?) echo "invalid option -$OPTARG" >&2 ; exit 1 ;;
        :) echo "option -$OPTARG requires an argument" >&2 ; exit 1 ;;
    esac
done

# shift off the processed options
shift $((OPTIND - 1))
echo "Remaining args: $@"
```

Supporting GNU-style long options (common pattern)
- `getopts` only parses single-character (short) options. Use a special case for `-` to parse `--long` options when you include `-:` in the optstring.
- When `getopts` returns `opt='-'`, the long option name is in `$OPTARG`; if the long option takes an argument, it will be the next positional parameter at `${!OPTIND}`. You must then increment `OPTIND`.

Full example: short + long options, help, and validation

```bash
#!/usr/bin/env bash

usage() {
    cat <<EOF
Usage: $0 -r <run-id> -t <token> [-i <instance>] [--help]

    -r <run-id>     Required. Run identifier.
    -t <token>      Required. Authentication token.
    -i <instance>   Optional. Instance name (short or long: --instance).
    -h, --help      Show this help message.
EOF
    exit 1
}

RUN_ID=""
TOKEN=""
INSTANCE=""
HELP=0

# leading colon -> we handle errors in the case branches
while getopts ":r:t:i:h-:" opt; do
    case $opt in
        r) RUN_ID=$OPTARG ;;
        t) TOKEN=$OPTARG ;;
        i) INSTANCE=$OPTARG ;;
        h) HELP=1 ;;
        -) # long option detected
                case "$OPTARG" in
                    help) HELP=1 ;;
                    instance)
                        # next arg is the long option's value
                        INSTANCE="${!OPTIND}"; OPTIND=$((OPTIND + 1))
                        ;;
                    *) echo "Unknown option --$OPTARG" >&2; usage ;;
                esac ;;
        \?) echo "Invalid option: -$OPTARG" >&2; usage ;;
        :) echo "Option -$OPTARG requires an argument." >&2; usage ;;
    esac
done

# if no options were provided, OPTIND remains 1
if [ "$OPTIND" -eq 1 ]; then
    echo "No options provided" >&2
    usage
fi

if [ "$HELP" -eq 1 ]; then
    usage
fi

if [ -z "$RUN_ID" ] || [ -z "$TOKEN" ]; then
    echo "Error: -r and -t are required." >&2
    usage
fi

echo "RUN_ID=$RUN_ID TOKEN=$TOKEN INSTANCE=$INSTANCE"
```

Tips
- Reset `OPTIND=1` if you call `getopts` multiple times in the same shell session/function.
- Prefer `getopts` for simple short-option parsing. For complex argument grammars or true GNU-style long options, consider `getopt` (with care for portability) or a higher-level parser in another language.
- Use a leading `:` in the optstring if you want to customize error messages instead of relying on `getopts`'s default diagnostics.
