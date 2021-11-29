import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendCreatePair } from "./types/myibcdex/tx";
import { MsgSendBuyOrder } from "./types/myibcdex/tx";
import { MsgCancelSellOrder } from "./types/myibcdex/tx";
import { MsgSendSellOrder } from "./types/myibcdex/tx";
import { MsgCancelBuyOrder } from "./types/myibcdex/tx";
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
    msgSendCreatePair: (data: MsgSendCreatePair) => EncodeObject;
    msgSendBuyOrder: (data: MsgSendBuyOrder) => EncodeObject;
    msgCancelSellOrder: (data: MsgCancelSellOrder) => EncodeObject;
    msgSendSellOrder: (data: MsgSendSellOrder) => EncodeObject;
    msgCancelBuyOrder: (data: MsgCancelBuyOrder) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
