![Build Status](https://travis-ci.org/spiffe/spire.svg?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/spiffe/spire/badge.svg?branch=master)](https://coveralls.io/github/spiffe/spire?branch=master)

![SPIRE Logo](/doc/spire_logo.png)

SPIRE (the [SPIFFE](https://github.com/spiffe/spiffe) Runtime Environment) provides a toolchain that defines a central registry of
SPIFFE IDs (the Server), and a Node Agent that can be run adjacent to a workload and
exposes a local Workload API. To get a better idea of what SPIRE is, and how it works, here is a [video](https://www.youtube.com/watch?v=uDHNcZ0eGHI) of it in action.


# Installing SPIRE

There are several ways to install the SPIRE binaries:

* Binary releases can be found at https://github.com/spiffe/spire/releases
* [Building from source](/CONTRIBUTING.md)

# Configuring SPIRE

## SPIRE Agent  

SPIRE Agent runs on every node and is responsible for requesting certificates from the spire server,
attesting the validity of local workloads, and providing them SVIDs.

### SPIRE Agent configuration

The following details the configurations for the spire agent. The configurations can be set through
.conf file or passed as command line args, the command line configurations takes precedence.

 |Configuration          | Description                                                          |
 |-----------------------|----------------------------------------------------------------------|
 |BindAddress            |  The GRPC Address where the WORKLOAD API Service is set to listen    |
 |BindPort               |  The GRPC port where the WORKLOAD API Service is set to listen       |
 |DataDir                |  Directory where the runtime data will be stored                     |
 |LogFile                |  Sets the path to log file                                           |
 |LogLevel               |  Sets the logging level \<DEBUG\|INFO\|WARN\|ERROR\>                 |
 |PluginDir              |  Directory where the plugin configuration are stored                 |
 |ServerAddress          |  The GRPC Address where the SPIRE Server is running                  |
 |ServerPort             |  The GRPC port of the SPIRE Service                                  |
 |SocketPath             |  Sets the path where the socket file will be generated               |
 |TrustBundlePath        |  Path to trusted CA Cert bundle                                      |
 |TrustDomain            |  SPIFFE trustDomain of the SPIRE Agent                               |
 |JoinToken              |  A join token which has been generated by the server                 |
 |Umask                  |  Umask to use (default value 0077). **Changing this may expose your signing authority to users other than the SPIRE agent/server**|


[default configuration file](/conf/agent/agent.conf)

```
BindAddress = "127.0.0.1"
BindPort = "8088"
DataDir = "."
LogLevel = "INFO"
PluginDir = "conf/agent/plugin"
ServerAddress = "127.0.0.1"
ServerPort = "8081"
SocketPath ="/tmp/agent.sock"
TrustBundlePath = "conf/agent/carootcert.pem"
TrustDomain = "example.org"
Umask = ""
```


### SPIRE Agent commands

 |Command                   | Action                                                           |
 |--------------------------|------------------------------------------------------------------|
 |`spire-agent run`         |  Starts the SPIRE Agent                                          |

## SPIRE Server  

SPIRE Server is responsible for validating and signing all CSRs in the SPIFFE trust domain.
Validation is performed through platform-specific Attestation plugins, as well as policy enforcement
backed by the SPIRE Server datastore.

### SPIRE Server configuration

The following details the configurations for the spire server. The configurations can be set through
a .conf file or passed as command line args, the command line configurations takes precedence.

 |Configuration          | Description                                                          |
 |-----------------------|----------------------------------------------------------------------|
 |BaseSVIDTTL            |  TTL that defines how long the generated Base SVID is valid          |
 |ServerSVIDTTL          |  TTL that defines how long the generated Server SVID is valid        |
 |BindAddress            |  The GRPC Address where the SPIRE Service is set to listen           |
 |BindPort               |  The GRPC port where the SPIRE Service is set to listen              |
 |BindHTTPPort           |  The HTTP port where the SPIRE Service is set to listen              |
 |LogFile                |  Sets the path to log file                                           |
 |LogLevel               |  Sets the logging level \<DEBUG\|INFO\|WARN\|ERROR\>                 |
 |PluginDir              |  Directory where the plugin configuration are stored                 |
 |TrustDomain            |  SPIFFE trustDomain of the SPIRE Agent                               |
 |Umask                  |  Umask to use (default value 0077). **Changing this may expose your signing authority to users other than the SPIRE agent/server**|

[default configuration file](/conf/server/server.conf)

```
BaseSVIDTTL = 999999
ServerSVIDTTL = 999999
BindAddress = "127.0.0.1"
BindPort = "8081"
BindHTTPPort = "8080"
LogLevel = "INFO"
PluginDir = "conf/server/plugin"
TrustDomain = "example.org"
Umask = ""
```

### SPIRE Server commands

 |Command                       | Action                                                           |
 |------------------------------|------------------------------------------------------------------|
 |`spire-server run`            |  Starts the SPIRE Server                                         |
 |`spire-server token generate` |  Generates a new join token

# Community

The SPIFFE community, and [Scytale](https://scytale.io) in particular, maintain the SPIRE project.
Information on the various SIGs and relevant standards can be found in
https://github.com/spiffe/spiffe.

The SPIFFE and SPIRE governance policies are detailed in
[GOVERNANCE](https://github.com/spiffe/spiffe/blob/master/GOVERNANCE.md)