import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendSellOrder } from "./types/myibcdex/tx";
import { MsgSendCreatePair } from "./types/myibcdex/tx";
import { MsgCancelSellOrder } from "./types/myibcdex/tx";
import { MsgCancelBuyOrder } from "./types/myibcdex/tx";
import { MsgSendBuyOrder } from "./types/myibcdex/tx";
export declare const MissingWalletError: Error;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => Promise<import("@cosmjs/stargate").BroadcastTxResponse>;
    msgSendSellOrder: (data: MsgSendSellOrder) => EncodeObject;
    msgSendCreatePair: (data: MsgSendCreatePair) => EncodeObject;
    msgCancelSellOrder: (data: MsgCancelSellOrder) => EncodeObject;
    msgCancelBuyOrder: (data: MsgCancelBuyOrder) => EncodeObject;
    msgSendBuyOrder: (data: MsgSendBuyOrder) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
