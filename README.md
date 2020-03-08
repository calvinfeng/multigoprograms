# Multiple Go Programs

I am trying out running supervisord in a Docker container. This is my small experimental ground.

## Configurations

```text
nodaemon = true | false
```

If true, supervisord will start in the foreground instead of daemonizing.

```text
autostart = true | false
```

If true, this program will start automatically when supervisord is started.

```text
autorestart = true | unexpected | false
```

If true, the process will be unconditionally restarted when it exits, without regard to its exit
code. If unexpected, the process will be restarted when the program exits with an exit code that is
not one of the exit codes associated with this process’ configuration. If false, the process will
not be auto restarted.

```text
exitcodes = 0, 1, 2, 3, ..., N
```

The list of _expected_ exit codes for this program used with autorestart.
