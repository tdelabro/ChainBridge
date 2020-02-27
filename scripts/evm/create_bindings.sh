#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


set -e

# Compile contracts
./scripts/evm/compile.sh

# Switch to directory
pushd ./on-chain/evm-contracts

# Generate bindings
node ./scripts/compileAbiBin.js

# Go back to root
popd

# Generate go bindings
./scripts/evm/create_go_bindings.sh
