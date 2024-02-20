# ALGE Optic3.NET XML importer utility

Small Golang library and cmdline utility that parses and converts CSV heat series exported by the sports results app (Mantis) used by the Spanish Canoeing Federation (RFEP) and other autonomic federations to the XML format used by the ALGE timing cameras management software.

## Building

This project requires Golang 1.21+, for setup instructions for your system see: https://go.dev/doc/install


### Setup

First, clone the project and setup the module dependencies

```shell
# Clone the project
git clone https://github.com/angiglesias/optic3-importer.git
# Fetch and setup golang vendoring directory
go mod vendor
```

### Linux

This is a pure Golang project, so building the project in linux is quite straighforward:

```shell
go build -v -ldflags "-s -w" -o convert cmd/convert/main.go
```

If you want a static binary for better compatibility accross Linux versions and distros build with `CGO_ENABLED=0`:

```shell
CGO_ENABLED go build -v -ldflags "-s -w" -o convert cmd/convert/main.go
```

To crosscompile the project for other systems and archs, include the `GOOS` and `GOARCH` env vars, as shown in [here](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures):

```shell
# Building the application for ARMv7 linux systems
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 -v -ldflags "-s -w" -o convert cmd/convert/main.go
```

### Windows

Building the project for Windows is similar than in linux:

```shell
go build -v -ldflags "-s -w" -o convert.exe cmd/convert/main.go
```

Or crosscompiling for Windows on another system:

```shell
GOOS=windows go build -v -ldflags "-s -w" -o convert.exe cmd/convert/main.go
```

## Use

The command line utility is pretty simple. Just point to a CSV file with the heat timetable as the one provided in [HCESA.csv](tests/data/HCESA.csv) and pick the formatting options which suit you best:

```shell
$ ./convert --help
Usage of ./convert:
      --ext-names       Extended file names (default true)
      --grp-days        Grouped days
      --inc-index       Include heat order in name (default true)
  -i, --input string    Input CSV file with data to convert
  -o, --output string   Output XML file with conversion
```
