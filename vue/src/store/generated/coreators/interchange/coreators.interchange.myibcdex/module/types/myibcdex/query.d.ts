import { Reader, Writer } from "protobufjs/minimal";
import { SellOrderBook } from "../myibcdex/sell_order_book";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { BuyOrderBook } from "../myibcdex/buy_order_book";
export declare const protobufPackage = "coreators.interchange.myibcdex";
export interface QueryGetSellOrderBookRequest {
    index: string;
}
export interface QueryGetSellOrderBookResponse {
    sellOrderBook: SellOrderBook | undefined;
}
export interface QueryAllSellOrderBookRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllSellOrderBookResponse {
    sellOrderBook: SellOrderBook[];
    pagination: PageResponse | undefined;
}
export interface QueryGetBuyOrderBookRequest {
    index: string;
}
export interface QueryGetBuyOrderBookResponse {
    buyOrderBook: BuyOrderBook | undefined;
}
export interface QueryAllBuyOrderBookRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllBuyOrderBookResponse {
    buyOrderBook: BuyOrderBook[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetSellOrderBookRequest: {
    encode(message: QueryGetSellOrderBookRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSellOrderBookRequest;
    fromJSON(object: any): QueryGetSellOrderBookRequest;
    toJSON(message: QueryGetSellOrderBookRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetSellOrderBookRequest>): QueryGetSellOrderBookRequest;
};
export declare const QueryGetSellOrderBookResponse: {
    encode(message: QueryGetSellOrderBookResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSellOrderBookResponse;
    fromJSON(object: any): QueryGetSellOrderBookResponse;
    toJSON(message: QueryGetSellOrderBookResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetSellOrderBookResponse>): QueryGetSellOrderBookResponse;
};
export declare const QueryAllSellOrderBookRequest: {
    encode(message: QueryAllSellOrderBookRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSellOrderBookRequest;
    fromJSON(object: any): QueryAllSellOrderBookRequest;
    toJSON(message: QueryAllSellOrderBookRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllSellOrderBookRequest>): QueryAllSellOrderBookRequest;
};
export declare const QueryAllSellOrderBookResponse: {
    encode(message: QueryAllSellOrderBookResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSellOrderBookResponse;
    fromJSON(object: any): QueryAllSellOrderBookResponse;
    toJSON(message: QueryAllSellOrderBookResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllSellOrderBookResponse>): QueryAllSellOrderBookResponse;
};
export declare const QueryGetBuyOrderBookRequest: {
    encode(message: QueryGetBuyOrderBookRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBuyOrderBookRequest;
    fromJSON(object: any): QueryGetBuyOrderBookRequest;
    toJSON(message: QueryGetBuyOrderBookRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetBuyOrderBookRequest>): QueryGetBuyOrderBookRequest;
};
export declare const QueryGetBuyOrderBookResponse: {
    encode(message: QueryGetBuyOrderBookResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBuyOrderBookResponse;
    fromJSON(object: any): QueryGetBuyOrderBookResponse;
    toJSON(message: QueryGetBuyOrderBookResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetBuyOrderBookResponse>): QueryGetBuyOrderBookResponse;
};
export declare const QueryAllBuyOrderBookRequest: {
    encode(message: QueryAllBuyOrderBookRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllBuyOrderBookRequest;
    fromJSON(object: any): QueryAllBuyOrderBookRequest;
    toJSON(message: QueryAllBuyOrderBookRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllBuyOrderBookRequest>): QueryAllBuyOrderBookRequest;
};
export declare const QueryAllBuyOrderBookResponse: {
    encode(message: QueryAllBuyOrderBookResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllBuyOrderBookResponse;
    fromJSON(object: any): QueryAllBuyOrderBookResponse;
    toJSON(message: QueryAllBuyOrderBookResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllBuyOrderBookResponse>): QueryAllBuyOrderBookResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a sellOrderBook by index. */
    SellOrderBook(request: QueryGetSellOrderBookRequest): Promise<QueryGetSellOrderBookResponse>;
    /** Queries a list of sellOrderBook items. */
    SellOrderBookAll(request: QueryAllSellOrderBookRequest): Promise<QueryAllSellOrderBookResponse>;
    /** Queries a buyOrderBook by index. */
    BuyOrderBook(request: QueryGetBuyOrderBookRequest): Promise<QueryGetBuyOrderBookResponse>;
    /** Queries a list of buyOrderBook items. */
    BuyOrderBookAll(request: QueryAllBuyOrderBookRequest): Promise<QueryAllBuyOrderBookResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    SellOrderBook(request: QueryGetSellOrderBookRequest): Promise<QueryGetSellOrderBookResponse>;
    SellOrderBookAll(request: QueryAllSellOrderBookRequest): Promise<QueryAllSellOrderBookResponse>;
    BuyOrderBook(request: QueryGetBuyOrderBookRequest): Promise<QueryGetBuyOrderBookResponse>;
    BuyOrderBookAll(request: QueryAllBuyOrderBookRequest): Promise<QueryAllBuyOrderBookResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
