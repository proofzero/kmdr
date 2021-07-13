```
 _  ____  __ _____  _____     _____ _      _____ 
| |/ /  \/  |  __ \|  __ \   / ____| |    |_   _|
| ' /| \  / | |  | | |__) | | |    | |      | |  
|  < | |\/| | |  | |  _  /  | |    | |      | |  
| . \| |  | | |__| | | \ \  | |____| |____ _| |_ 
|_|\_\_|  |_|_____/|_|  \_\  \_____|______|_____|
```
 ![build](https://github.com/proofzero/kmdr/actions/workflows/bazel.yaml/badge.svg)

> **NOTE: This project is a work in progress**

The CLI for developing, collaborating, and delivering data-centric applications.

# Table of Contents
- [Overview](#overview)
- [Concepts](#concepts)
- [Use Cases](#usecases)
- [Tutorials](#tutorials)
- [Documentation](#documentation)
- [Roadmap](#roadmap)

# Install Instructions

### Go
```bash
go install github.com/proofzero/kmdr@latest
```

# Overview

KMDR is the CLI for [kubelt](https://kubelt.com) but, KMDR also works in offline contexts.

KMDR provides a kubernetes-like experience for defining a schema driven graph of data containers, a git-like experience for working with data against those data containers, and a package manager-like experience for materializing/synthensizing datasets and running transforms. With KMDR, users can discover, source, publish and subscribe to data to collaborate with peers across the globe.

## Concepts

[kubelt](https://kubelt.com) is a distrubuted platform that operates on an IPFS data plane. The tools [kubelt](https://kubelt.com) publishes provides users with a familiar developer experience and abstractions for working with datasets and data pipelines.

The KMDR CLI is paired with a daemon named KTRL and as the names suggest, KMDR instructs KTRL to perform specific actions using declarative manifests. KTRL manages an offline IPFS node, auth, and more (for more details checkout the [KTRL](https://github.com/proofzero/ktrl) repo).

Some further concepts are below.

### Manifests

KMDR uses declarative configurations to apply changes to the underlying graph and linked data. These declartive configrations can be supplied to KMDR as manifests written in (cue)[https://cuelang.org].

> TODO: provide examples

### Graph and Content Addresses

[kubelt](https://kubelt.com) maintains a graph of data containers that define users, datasets, and more. The graph informs KMDR and KTRL the metadata and addresses for data stored in the data plane.

### Datasets

Information stored on the data plane is encrypted into opaque blocks of data chuncked based on the schema provided by the user.

#### Types and Schemas

KMDR encourges defining types and schemas for your datasets in order to optimize the utility of the underlying platforms.

Schemas help KMDR reason about your data to validate it, clean it, share it, materalize it, synthesize it, and otherwise improve it. 

> TODO: provide examples

#### Immutability and Versioning

All data published to a specific content address is immutably stored in a directed acyclic graph. Just like with git, changes can be rolled back leaving unwanted changes in a detached state.

### Contexts

Similar to Kubernetes contexts, KMDR supports multiple contexts and contexts are linked to specific enviroments.

Offline only contexts are currently supported however, online public, managed, and on-prem backed contexts will be rolled out over time.

## Use Cases

> TODO

## Tutorials

> TODO

## Documentation

> TODO

## Roadmap
- Documentation portal
- Subscribing
- Forking
- Authentication/Auhtorization & Encryption
- Data generation & DLP
- Programs
- Program flows (ELT/ETL pipelines)
- Events / journaling