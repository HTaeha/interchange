// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendSellOrder } from "./types/myibcdex/tx";
import { MsgSendCreatePair } from "./types/myibcdex/tx";
import { MsgCancelSellOrder } from "./types/myibcdex/tx";
import { MsgCancelBuyOrder } from "./types/myibcdex/tx";
import { MsgSendBuyOrder } from "./types/myibcdex/tx";
const types = [
    ["/coreators.interchange.myibcdex.MsgSendSellOrder", MsgSendSellOrder],
    ["/coreators.interchange.myibcdex.MsgSendCreatePair", MsgSendCreatePair],
    ["/coreators.interchange.myibcdex.MsgCancelSellOrder", MsgCancelSellOrder],
    ["/coreators.interchange.myibcdex.MsgCancelBuyOrder", MsgCancelBuyOrder],
    ["/coreators.interchange.myibcdex.MsgSendBuyOrder", MsgSendBuyOrder],
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
        msgSendSellOrder: (data) => ({ typeUrl: "/coreators.interchange.myibcdex.MsgSendSellOrder", value: data }),
        msgSendCreatePair: (data) => ({ typeUrl: "/coreators.interchange.myibcdex.MsgSendCreatePair", value: data }),
        msgCancelSellOrder: (data) => ({ typeUrl: "/coreators.interchange.myibcdex.MsgCancelSellOrder", value: data }),
        msgCancelBuyOrder: (data) => ({ typeUrl: "/coreators.interchange.myibcdex.MsgCancelBuyOrder", value: data }),
        msgSendBuyOrder: (data) => ({ typeUrl: "/coreators.interchange.myibcdex.MsgSendBuyOrder", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
