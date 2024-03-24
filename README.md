# File Guard

## Overview

A secure way to upload, store, and verify small files on a server using Merkle trees, ensuring data integrity without retaining local copies.

## Features

- Construct Merkle tree from a set of files
- Generate proof
- Verify proof

## Technologies Used

- Go (Golang)
- Protocol Buffers (Proto3)

## Getting Started

Here's the commands to set up and run the application.

- **start-server**: Launches the server application.

  ```bash
  make start-server
  ```

- **start-client**: Launches the client application.

  ```bash
  make start-client
  ```

- **install**: Installs all necessary dependencies, including protoc-gen-go, protoc-gen-go-grpc, protobuf, clang-format, and grpcurl. It also ensures these tools are added to your PATH. This is typically run once before starting development or after cleaning.

  ```bash
  make install
  ```

- **test**: Runs unit tests for the project, particularly focusing on the Merkle tree implementation to ensure its integrity.

  ```bash
  make test
  ```

- Prerequisites: software or tools that need to be installed before running your application.
- Installation steps: how to clone/download the repository and any necessary build steps.
- How to use Docker Compose to spin up your application.

## Components

- Merkle Tree
- Server
- Client

## Challenges & Limitations

The project kicked off with the implementation of the Merkle tree, which initially posed challenges in handling an odd number of files. Subsequently, a basic server-client framework was established without incorporating networking; in this setup, the client would instantiate the server directly to execute all functions. The next phase involved integrating a network layer via gRPC with proto3, enabling direct interactions between the client and server. However, this setup initially allowed for only a single Merkle tree to be constructed, limiting the system to one client at a time.

To address this, a sessionId concept was introduced to differentiate between Merkle trees and their associated operations, allowing clients to either initiate a new session or join an existing one. Ideally, this sessionId would serve as a unique identifier for clients during their interactions with the server, facilitating a more dynamic and interactive environment.

## Future Improvements

The current implementation of the Merkle tree primarily addresses essential features. Future expansions could include:

- Enhancements to the Merkle tree to allow file additions or removals post-construction.

For the server:

- Implementing a Least Recently Used (LRU) caching mechanism for managing active sessions efficiently.
- Enabling the server to reconstruct the tree based on existing files if the tree is not already loaded.
- Improving handling of concurrent operations to enhance performance and reliability.

For the client:

- The client, currently a basic command-line interface for server interaction, could be enriched with additional functionalities.
- Refining the process for uploading and downloading files to ensure smoother and more reliable file transfers.

Additionally, extending test coverage for both server and client components is crucial for ensuring the robustness and reliability of the application.
