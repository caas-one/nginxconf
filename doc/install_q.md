# Install problems 

#### Q1: pkg-config: exec: "pkg-config": executable file not found in $PATH

Answer:
```
brew install pkg-config
```

Relevant answers:
- https://github.com/gopherdata/gophernotes/issues/82

#### Q2: add the directory containing 'python-2.7.pc'

Detailed description Q:
```
# pkg-config --cflags  -- python-2.7
Package python-2.7 was not found in the pkg-config search path.
Perhaps you should add the directory containing `python-2.7.pc'
to the PKG_CONFIG_PATH environment variable
No package 'python-2.7' found
pkg-config: exit status 1
```

Answer:
```
export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/System/Library/Frameworks/Python.framework/Versions/2.7/lib/pkgconfig
```

Relevant answers:
- https://github.com/sbinet/go-python/issues/28
