# goimprove
```
goimprove -r fabric8io/fabric8
```

goimprove will retrive useful information and metrics and optionally send them to ElasticSearch in order to perform powerful queries and visualisations with Kibana.

Currently it only supports getting the number of downloads a particular github release has had.  The project can be extended to retrieve anything that can be useful and send to ElasticSearch.

It's aim is to be run from the CLI or in Kubernetes pod that will watch for events

Ideally we'd expose a REST API so that a monitoring tool such as Prometheus can coe along and scrape any metrics instead of pushing to ElasticSearch here.

## Getting started

### Install / Update & run

Get latest download URL from [goimprove releases](https://github.com/fabric8io/goimprove/releases)

```sh
sudo rm /tmp/goimprove
sudo rm -rf /usr/bin/goimprove
mkdir /tmp/goimprove
curl --retry 999 --retry-max-time 0  -sSL [[ADD DOWNLOAD URL HERE]] | tar xzv -C /tmp/goimprove
chmod +x /tmp/goimprove/goimprove
sudo mv /tmp/goimprove/* /usr/bin/
```

### Usage

```
goimprove is used to gather stats and metrics of various types and expose via rest to be scraped by a metrics tool
       								Find more information at http://fabric8.io.

Usage:
  goimprove [flags]
  goimprove [command]

Available Commands:
  gh-downloads retrives the number of downloads of a GitHub project release
  version      Display version & exit

Flags:
  -y, --yes   assume yes

Use "goimprove [command] --help" for more information about a command.
```

## Development

### Prerequisites

Install [go version 1.6](https://golang.org/doc/install)


### Building

```sh
git clone git@github.com:fabric8io/goimprove.git $GOPATH/src/github.com/fabric8io/goimprove
cd $GOPATH/src/github.com/fabric8io/goimprove
make bootstrap
```

Edit *.go files, run `make` and run the generated binary..

e.g.

```sh
make
./build/goimprove -r fabric8io/fabric8

```
