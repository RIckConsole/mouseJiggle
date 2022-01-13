# Mouse Jiggler
Is your hand getting tired from waking your PC? Do you want to appear online while you take a nap? If so, keep reading. 

### Download the latest release [HERE](https://github.com/RIckConsole/mouseJiggle/releases)

Running the binary without arguments will start jiggling your cursor. Press CTRL+C to stop the program and get back control of your mouse.

There are also some command line arguments you can add to slightly change its functionality:

`.\mouseJiggle.exe --help`

```
  Flags:
       --version    Displays the program version string.
    -h --help       Displays help with available flag, subcommand, and positional value parameters.
    -r --random     Enable random mouse movements
    -i --interval   Set the time (in seconds) between mouse movements (default: 1)
```

## Building

The nature of the robotgo package requires it to be built on windows. However, you still need GCC in order for it to work. You can download an installer from here: [TDM-GCC](https://jmeubank.github.io/tdm-gcc/articles/2021-05/10.3.0-release)

At this point, if you try to build, youll get an error regarding `zlib.h`. To fix this, download [ZLIBx64](https://sourceforge.net/projects/mingw-w64/files/External%20binary%20packages%20%28Win64%20hosted%29/Binaries%20%2864-bit%29/zlib-1.2.5-bin-x64.zip/download), and then perform the following operations:

```
copy\zlib\bin to \TDM\bin
copy \zlib\include to \TDM\include
copy \zlib\lib to \TDM\lib
```

Add TDM-GCC's `/bin` directory to your PATH, and then you should be able to build the binary using `go build`