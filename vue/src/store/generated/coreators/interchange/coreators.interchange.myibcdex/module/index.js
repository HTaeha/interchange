// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendBuyOrder } from "./types/myibcdex/tx";
import { MsgSendCreatePair } from "./types/myibcdex/tx";
import { MsgSendSellOrder } from "./types/myibcdex/tx";
const types = [
    ["/coreators.interchange.myibcdex.MsgSendBuyOrder", MsgSendBuyOrder],
    ["/coreators.interchange.myibcdex.MsgSendCreatePair", MsgSendCreatePair],
    ["/coreators.interchange.myibcdex.MsgSendSellOrder", MsgSendSellOrder],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgSendBuyOrder: (data) => ({ typeUrl: "/coreators.interchange.myibcdex.MsgSendBuyOrder", value: data }),
        msgSendCreatePair: (data) => ({ typeUrl: "/coreators.interchange.myibcdex.MsgSendCreatePair", value: data }),
        msgSendSellOrder: (data) => ({ typeUrl: "/coreators.interchange.myibcdex.MsgSendSellOrder", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
