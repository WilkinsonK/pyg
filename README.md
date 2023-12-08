# Pyg #
Bindings for embedded Python3 in Go.
---

These bindings were built with the Python 3.11 C API as reference.
Hypothetically, it should work with versions as low as 3.9, but not any further
with 3.8 on its way to the chopping block at time of writing.

The CGO implementation points to the `python3-embed` configuration. To use a
specific version of the Python API, you will need to link the desired version
of embedded Python to this `*.pc` file.

## Getting Started ##
If thre are multiple versions of Python3 installed on you system, and want
to use a different version, or pkg-config can't find `python3-embed`, you need
to locate the `*.pc` files for your Python dev packages.

This command should make it easier to locate where your pkg-config configs are
placed on your system.
```bash
$ sudo find / -type d -iname *pkgconfig*
/path/to/pkgconfig/files
```

If you don't find any configurations pointing to a Python C API, it is most
likely the needed development packages are not installed.
```base
# For RHEL/Fedora/CentOS systems
$ sudo dnf install python3-devel

# For Debian based systems
$ sudo apt-get install python3-dev
```
