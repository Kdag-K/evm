# EVM-LITE

[![CircleCI](https://circleci.com/gh/mosaicnetworks/evm-lite.svg?style=svg)](https://circleci.com/gh/Kdag-K/evm)
[![Go Report](https://goreportcard.com/badge/github.com/Kdag-K/evm)](https://goreportcard.com/report/github.com/Kdag-K/evm)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## A lean Ethereum node with interchangeable consensus.

We took the [Go-Ethereum](https://github.com/ethereum/go-ethereum)
implementation (Geth) and extracted the EVM and Trie components to create a lean
and modular version with interchangeable consensus.

The EVM is a virtual machine specifically designed to run untrusted code on a
network of computers. Every transaction applied to the EVM modifies the State
which is persisted in a Merkle Patricia tree. This data structure allows to
simply check if a given transaction was actually applied to the VM and can
reduce the entire State to a single hash (merkle root) rather analogous to a
fingerprint.

The EVM is meant to be used in conjunction with a system that broadcasts
transactions across network participants and ensures that everyone executes the
same transactions in the same order. Ethereum uses a Blockchain and a Proof of
Work consensus algorithm. EVM-Lite makes it easy to use any consensus system,
including [kdag](https://github.com/Kdag-K/kdag) .

## ARCHITECTURE

```
                +-------------------------------------------+
+----------+    |  +-------------+         +-------------+  |       
|          |    |  | Service     |         | State       |  |
|  Client  <-----> |             | <------ |             |  |
|          |    |  | -API        |         | -EVM        |  |
+----------+    |  |             |         | -Trie       |  |
                |  |             |         | -Database   |  |
                |  +-------------+         +-------------+  |
                |         |                       ^         |     
                |         v                       |         |
                |  +-------------------------------------+  |
                |  | Engine                              |  |
                |  |                                     |  |
                |  |       +----------------------+      |  |
                |  |       | Consensus            |      |  |
                |  |       +----------------------+      |  |
                |  |                                     |  |
                |  +-------------------------------------+  |
                |                                           |
                +-------------------------------------------+

```

## Usage

EVM is a Go library, which is meant to be used in conjunction with a 
consensus system like Babble, Tendermint, Raft etc.

This repo contains **Solo**, a bare-bones implementation of the consensus 
interface, which is used for testing or launching a standalone node. It relays
transactions directly from Service to State.



