// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCancelBuyOrder } from "./types/myibcdex/tx";
import { MsgSendBuyOrder } from "./types/myibcdex/tx";
import { MsgSendCreatePair } from "./types/myibcdex/tx";
import { MsgSendSellOrder } from "./types/myibcdex/tx";
import { MsgCancelSellOrder } from "./types/myibcdex/tx";


const types = [
  ["/coreators.interchange.myibcdex.MsgCancelBuyOrder", MsgCancelBuyOrder],
  ["/coreators.interchange.myibcdex.MsgSendBuyOrder", MsgSendBuyOrder],
  ["/coreators.interchange.myibcdex.MsgSendCreatePair", MsgSendCreatePair],
  ["/coreators.interchange.myibcdex.MsgSendSellOrder", MsgSendSellOrder],
  ["/coreators.interchange.myibcdex.MsgCancelSellOrder", MsgCancelSellOrder],
  
];
export const MissingWalletError = new Error("wallet is required");

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCancelBuyOrder: (data: MsgCancelBuyOrder): EncodeObject => ({ typeUrl: "/coreators.interchange.myibcdex.MsgCancelBuyOrder", value: data }),
    msgSendBuyOrder: (data: MsgSendBuyOrder): EncodeObject => ({ typeUrl: "/coreators.interchange.myibcdex.MsgSendBuyOrder", value: data }),
    msgSendCreatePair: (data: MsgSendCreatePair): EncodeObject => ({ typeUrl: "/coreators.interchange.myibcdex.MsgSendCreatePair", value: data }),
    msgSendSellOrder: (data: MsgSendSellOrder): EncodeObject => ({ typeUrl: "/coreators.interchange.myibcdex.MsgSendSellOrder", value: data }),
    msgCancelSellOrder: (data: MsgCancelSellOrder): EncodeObject => ({ typeUrl: "/coreators.interchange.myibcdex.MsgCancelSellOrder", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
