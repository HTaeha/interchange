import { SellOrderBook } from "../myibcdex/sell_order_book";
import { BuyOrderBook } from "../myibcdex/buy_order_book";
import { DenomTrace } from "../myibcdex/denom_trace";
import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "coreators.interchange.myibcdex";
/** GenesisState defines the myibcdex module's genesis state. */
export interface GenesisState {
    portId: string;
    sellOrderBookList: SellOrderBook[];
    buyOrderBookList: BuyOrderBook[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    denomTraceList: DenomTrace[];
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
