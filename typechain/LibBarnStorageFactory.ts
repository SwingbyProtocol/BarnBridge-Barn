/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer } from "ethers";
import { Provider, TransactionRequest } from "@ethersproject/providers";
import { Contract, ContractFactory, Overrides } from "@ethersproject/contracts";

import type { LibBarnStorage } from "./LibBarnStorage";

export class LibBarnStorageFactory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer);
  }

  deploy(overrides?: Overrides): Promise<LibBarnStorage> {
    return super.deploy(overrides || {}) as Promise<LibBarnStorage>;
  }
  getDeployTransaction(overrides?: Overrides): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  attach(address: string): LibBarnStorage {
    return super.attach(address) as LibBarnStorage;
  }
  connect(signer: Signer): LibBarnStorageFactory {
    return super.connect(signer) as LibBarnStorageFactory;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): LibBarnStorage {
    return new Contract(address, _abi, signerOrProvider) as LibBarnStorage;
  }
}

const _abi = [
  {
    inputs: [
      {
        internalType: "bytes32",
        name: "c__0x6c7ee5de",
        type: "bytes32",
      },
    ],
    name: "c_0x6c7ee5de",
    outputs: [],
    stateMutability: "pure",
    type: "function",
  },
];

const _bytecode =
  "0x60e0610025600b82828239805160001a60731461001857fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063b644a27f146038575b600080fd5b604e6004803603810190604a91906066565b6050565b005b50565b6000813590506060816096565b92915050565b600060208284031215607757600080fd5b60006083848285016053565b91505092915050565b6000819050919050565b609d81608c565b811460a757600080fd5b5056fea26469706673582212204d2597a26eb3bbc808e15b1c8084096d732e8ca1d6bbfa5209372cac37a2e65d64736f6c63430007030033";