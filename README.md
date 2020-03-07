# go-demo

## Prerequisites

Make sure you have `brew` and a supported `GO` version installed. At the time of writing this guide it's version 1.13.

On mac:

```bash
brew install go@1.13
```

If you have previous version installed, either remove it or add the following to your bash profile and reload it.

```bash
...

export PATH="/usr/local/opt/go@1.13/bin:$PATH"

...

```

## Install TinyGO

Install instruction can be found at the official site, https://tinygo.org/getting-started/.

Here is the abbreviated version for mac

```bash
brew tap tinygo-org/tools
brew install tinygo
```

## Install AVR tools

In this repo I will be using an Arduino Nano, which requires the toolset/programmer for AVR.

```bash
brew tap osx-cross/avr
brew install avr-gcc
brew install avrdude
```

## Test example/blinky

To program the device, the following command is used. The port can be found in the `/dev/` path.

```bash
tinygo flash -target=<DEVICE_TYPE> -port <PORT> examples/blinky1
```

Find your device (in the case of a arduino)

```bash
ls /dev/tty.usbserial*
```

In our case that would be `arduino-nano` and the port `/dev/tty.usbserial-#####`

```bash
tinygo flash -target=arduino-nano -port /dev/tty.usbserial-14130 examples/blinky1
```

If everything went ok, the device you now be blinking :)
