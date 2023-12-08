# Pyg #
Bindings for embedded Python3 in Go.
---

These bindings were built with the Python 3.11 C API as reference.
Hypothetically, it should work with versions as low as 3.9, but not any further
with 3.8 on its way to the chopping block.

The CGO implementation points to the `python3-embed` configuration. To use a
specific version of the Python API, you will need to link the desired version
of embedded Python to this `*.pc` file.