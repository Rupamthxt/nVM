# nVM (go-evm) 

A from-scratch, Turing-complete implementation of the Ethereum Virtual Machine (EVM) written in Go. 

This project is an exercise in low-level systems architecture, focusing on building a deterministic stack machine capable of executing EVM bytecode, managing dynamic memory, and strictly metering computational resources.

## 🛠️ Getting Started & Running

Because this EVM is completely standalone and isolated from the Ethereum mainnet, you can feed it raw hex bytecode directly to watch the stack and memory state mutate in real-time.

### Prerequisites
* Go 1.20 or higher.

### Installation

Clone the repository and build the engine:

```bash
git clone https://github.com/Rupamthxt/nVM.git
cd nVM
go build -o evm-engine .
```

## ⚙️ Core Architecture

At its heart, this is a stack-based virtual machine operating on 256-bit words. The execution environment is isolated, deterministic, and sandboxed.

### Implemented Features

* **O(1) Opcode Dispatch:** Utilizes a pre-computed jump table for instruction routing. This avoids slow, monolithic `switch` statements and ensures predictable, low-latency execution times.
* **Turing-Complete Control Flow:** Full support for `JUMP`, `JUMPI`, and `PC` operations. Includes dynamic destination validation (`JUMPDEST`) to prevent the execution of arbitrary data payloads.
* **Strict Gas Metering:** An algorithmic accounting system that decrements gas per opcode. This solves the halting problem by automatically triggering an `Out Of Gas` (OOG) fault if a contract enters an infinite loop or exceeds its computational budget.
* **Volatile Memory (`MSTORE` / `MLOAD`):** A byte-addressable, dynamically expanding linear memory space. Memory expansion costs are calculated exactly per the Ethereum Yellow Paper specifications.
* **The Stack:** A strictly enforced 1024-depth limit, 256-bit word stack with complete implementations of the `PUSH`, `POP`, `DUP`, and `SWAP` instruction families.
* **State Interface:** Core execution logic for `SSTORE` and `SLOAD` operations. *(Note: Currently utilizing an ephemeral in-memory Go map for testing purposes).*

## 🏗️ Project Structure

The core execution loop follows a strict Fetch -> Decode -> Execute cycle. Before a contract runs, the bytecode is analyzed to map out valid jump destinations, ensuring safety during runtime. During execution, the Gas pool is updated *before* the opcode mutates the stack or memory, ensuring state remains uncorrupted if execution fails.

## 🚀 Roadmap (WIP)

The execution layer is stable. The next major phase focuses on moving from ephemeral memory to cryptographic state persistence:

* **Recursive Length Prefix (RLP) Serialization:** Implementing Ethereum's hyper-minimalist byte-encoding format.
* **Cryptographic State (MPT):** Replacing the in-memory `SSTORE` map with a mathematically verifiable Modified Merkle Patricia Trie.
* **Native Crypto:** Integrating `KECCAK256` hashing for state roots.
* **Execution Contexts:** Implementing `CALL`, `DELEGATECALL`, and `CREATE` to support cross-contract interactions.

---