This repository contains simple programs that I'm trying out as I learn Go.

The scratch folder contains various subfolders, each of which is a main
package which can be built and run as a separate command.

I intend to build out a server implemting a RESTish API and also serves static
assests. This code will be organised under the following folders:
app, conf, assets & executables. The goal is to have the core logic implemented
in various packages in the app folder which will be used by a main package in
the executables folder to implement a server.

The various static assets and configuration files will be stored in conf and
assets respectively.

The goal is to finally have a gradle build file that will build the server,
generate a compiled version of the static assets (less, sass, require etc) and
organize it in a timestamped build folder that can then be deployed by just
running the executable.

---
