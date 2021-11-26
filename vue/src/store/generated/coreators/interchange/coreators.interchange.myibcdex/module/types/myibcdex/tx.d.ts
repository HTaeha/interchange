import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "coreators.interchange.myibcdex";
export interface MsgSendCreatePair {
    creator: string;
    port: string;
    channelID: string;
    timeoutTimestamp: number;
    sourceDenom: string;
    targetDenom: string;
}
export interface MsgSendCreatePairResponse {
}
export interface MsgSendSellOrder {
    creator: string;
    port: string;
    channelID: string;
    timeoutTimestamp: number;
    amountDenom: string;
    amount: number;
    priceDenom: string;
    price: number;
}
export interface MsgSendSellOrderResponse {
}
export interface MsgSendBuyOrder {
    creator: string;
    port: string;
    channelID: string;
    timeoutTimestamp: number;
    amountDenom: string;
    amount: number;
    priceDenom: string;
    price: number;
}
export interface MsgSendBuyOrderResponse {
}
export declare const MsgSendCreatePair: {
    encode(message: MsgSendCreatePair, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendCreatePair;
    fromJSON(object: any): MsgSendCreatePair;
    toJSON(message: MsgSendCreatePair): unknown;
    fromPartial(object: DeepPartial<MsgSendCreatePair>): MsgSendCreatePair;
};
export declare const MsgSendCreatePairResponse: {
    encode(_: MsgSendCreatePairResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendCreatePairResponse;
    fromJSON(_: any): MsgSendCreatePairResponse;
    toJSON(_: MsgSendCreatePairResponse): unknown;
    fromPartial(_: DeepPartial<MsgSendCreatePairResponse>): MsgSendCreatePairResponse;
};
export declare const MsgSendSellOrder: {
    encode(message: MsgSendSellOrder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendSellOrder;
    fromJSON(object: any): MsgSendSellOrder;
    toJSON(message: MsgSendSellOrder): unknown;
    fromPartial(object: DeepPartial<MsgSendSellOrder>): MsgSendSellOrder;
};
export declare const MsgSendSellOrderResponse: {
    encode(_: MsgSendSellOrderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendSellOrderResponse;
    fromJSON(_: any): MsgSendSellOrderResponse;
    toJSON(_: MsgSendSellOrderResponse): unknown;
    fromPartial(_: DeepPartial<MsgSendSellOrderResponse>): MsgSendSellOrderResponse;
};
export declare const MsgSendBuyOrder: {
    encode(message: MsgSendBuyOrder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendBuyOrder;
    fromJSON(object: any): MsgSendBuyOrder;
    toJSON(message: MsgSendBuyOrder): unknown;
    fromPartial(object: DeepPartial<MsgSendBuyOrder>): MsgSendBuyOrder;
};
export declare const MsgSendBuyOrderResponse: {
    encode(_: MsgSendBuyOrderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendBuyOrderResponse;
    fromJSON(_: any): MsgSendBuyOrderResponse;
    toJSON(_: MsgSendBuyOrderResponse): unknown;
    fromPartial(_: DeepPartial<MsgSendBuyOrderResponse>): MsgSendBuyOrderResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    SendCreatePair(request: MsgSendCreatePair): Promise<MsgSendCreatePairResponse>;
    SendSellOrder(request: MsgSendSellOrder): Promise<MsgSendSellOrderResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    SendBuyOrder(request: MsgSendBuyOrder): Promise<MsgSendBuyOrderResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    SendCreatePair(request: MsgSendCreatePair): Promise<MsgSendCreatePairResponse>;
    SendSellOrder(request: MsgSendSellOrder): Promise<MsgSendSellOrderResponse>;
    SendBuyOrder(request: MsgSendBuyOrder): Promise<MsgSendBuyOrderResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
