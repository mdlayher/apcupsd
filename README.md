# apcupsd [![Linux Test Status](https://github.com/mdlayher/apcupsd/workflows/Linux%20Test/badge.svg)](https://github.com/mdlayher/apcupsd/actions) [![GoDoc](http://godoc.org/github.com/mdlayher/apcupsd?status.svg)](http://godoc.org/github.com/mdlayher/apcupsd) [![Report Card](https://goreportcard.com/badge/github.com/mdlayher/apcupsd)](https://goreportcard.com/report/github.com/mdlayher/apcupsd)

Package `apcupsd` provides a client for the [apcupsd](http://www.apcupsd.org/)
Network Information Server (NIS).  MIT Licensed.

Fork notes: this fork fixes a bug in which alternate values for ALARMDEL break the ability to parse the APC UPS daemon's response. If this bug is fixed in the central repo, there should be no reason not to use it instead. 
