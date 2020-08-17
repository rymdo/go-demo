# goroutine

## compile/flash

```bash
tinygo flash -scheduler coroutines -target=<DEVICE> -port <PORT> .
```

eg.

```bash
tinygo flash -scheduler coroutines -target=arduino-nano -port /dev/tty.usbserial-11111 .
```
