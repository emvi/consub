# consub

consub is a simple configuration substitution tool.

## Usage

Download the binary from the release section or compile it. Then run:

```
consub .vars.env output.tpl.txt output.txt
```

The `.vars.env` file must be in the following format:

```
VAR=test
```

The `output.tpl.txt` file can then be configured like this:

```
My ${VAR} configuration!
```

The resulting `output.txt` file will look like this:

```
My test configuration!
```

## License

MIT
