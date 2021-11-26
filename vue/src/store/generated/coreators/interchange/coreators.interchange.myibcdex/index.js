import { txClient, queryClient, MissingWalletError } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { BuyOrderBook } from "./module/types/myibcdex/buy_order_book";
import { MyibcdexPacketData } from "./module/types/myibcdex/packet";
import { NoData } from "./module/types/myibcdex/packet";
import { CreatePairPacketData } from "./module/types/myibcdex/packet";
import { CreatePairPacketAck } from "./module/types/myibcdex/packet";
import { SellOrderPacketData } from "./module/types/myibcdex/packet";
import { SellOrderPacketAck } from "./module/types/myibcdex/packet";
import { SellOrderBook } from "./module/types/myibcdex/sell_order_book";
export { BuyOrderBook, MyibcdexPacketData, NoData, CreatePairPacketData, CreatePairPacketAck, SellOrderPacketData, SellOrderPacketAck, SellOrderBook };
async function initTxClient(vuexGetters) {
    return await txClient(vuexGetters['common/wallet/signer'], {
        addr: vuexGetters['common/env/apiTendermint']
    });
}
async function initQueryClient(vuexGetters) {
    return await queryClient({
        addr: vuexGetters['common/env/apiCosmos']
    });
}
function mergeResults(value, next_values) {
    for (let prop of Object.keys(next_values)) {
        if (Array.isArray(next_values[prop])) {
            value[prop] = [...value[prop], ...next_values[prop]];
        }
        else {
            value[prop] = next_values[prop];
        }
    }
    return value;
}
function getStructure(template) {
    let structure = { fields: [] };
    for (const [key, value] of Object.entries(template)) {
        let field = {};
        field.name = key;
        field.type = typeof value;
        structure.fields.push(field);
    }
    return structure;
}
const getDefaultState = () => {
    return {
        SellOrderBook: {},
        SellOrderBookAll: {},
        BuyOrderBook: {},
        BuyOrderBookAll: {},
        _Structure: {
            BuyOrderBook: getStructure(BuyOrderBook.fromPartial({})),
            MyibcdexPacketData: getStructure(MyibcdexPacketData.fromPartial({})),
            NoData: getStructure(NoData.fromPartial({})),
            CreatePairPacketData: getStructure(CreatePairPacketData.fromPartial({})),
            CreatePairPacketAck: getStructure(CreatePairPacketAck.fromPartial({})),
            SellOrderPacketData: getStructure(SellOrderPacketData.fromPartial({})),
            SellOrderPacketAck: getStructure(SellOrderPacketAck.fromPartial({})),
            SellOrderBook: getStructure(SellOrderBook.fromPartial({})),
        },
        _Subscriptions: new Set(),
    };
};
// initial state
const state = getDefaultState();
export default {
    namespaced: true,
    state,
    mutations: {
        RESET_STATE(state) {
            Object.assign(state, getDefaultState());
        },
        QUERY(state, { query, key, value }) {
            state[query][JSON.stringify(key)] = value;
        },
        SUBSCRIBE(state, subscription) {
            state._Subscriptions.add(subscription);
        },
        UNSUBSCRIBE(state, subscription) {
            state._Subscriptions.delete(subscription);
        }
    },
    getters: {
        getSellOrderBook: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.SellOrderBook[JSON.stringify(params)] ?? {};
        },
        getSellOrderBookAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.SellOrderBookAll[JSON.stringify(params)] ?? {};
        },
        getBuyOrderBook: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.BuyOrderBook[JSON.stringify(params)] ?? {};
        },
        getBuyOrderBookAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.BuyOrderBookAll[JSON.stringify(params)] ?? {};
        },
        getTypeStructure: (state) => (type) => {
            return state._Structure[type].fields;
        }
    },
    actions: {
        init({ dispatch, rootGetters }) {
            console.log('Vuex module: coreators.interchange.myibcdex initialized!');
            if (rootGetters['common/env/client']) {
                rootGetters['common/env/client'].on('newblock', () => {
                    dispatch('StoreUpdate');
                });
            }
        },
        resetState({ commit }) {
            commit('RESET_STATE');
        },
        unsubscribe({ commit }, subscription) {
            commit('UNSUBSCRIBE', subscription);
        },
        async StoreUpdate({ state, dispatch }) {
            state._Subscriptions.forEach(async (subscription) => {
                try {
                    await dispatch(subscription.action, subscription.payload);
                }
                catch (e) {
                    throw new SpVuexError('Subscriptions: ' + e.message);
                }
            });
        },
        async QuerySellOrderBook({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.querySellOrderBook(key.index)).data;
                commit('QUERY', { query: 'SellOrderBook', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QuerySellOrderBook', payload: { options: { all }, params: { ...key }, query } });
                return getters['getSellOrderBook']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QuerySellOrderBook', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QuerySellOrderBookAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.querySellOrderBookAll(query)).data;
                while (all && value.pagination && value.pagination.nextKey != null) {
                    let next_values = (await queryClient.querySellOrderBookAll({ ...query, 'pagination.key': value.pagination.nextKey })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'SellOrderBookAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QuerySellOrderBookAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getSellOrderBookAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QuerySellOrderBookAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryBuyOrderBook({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryBuyOrderBook(key.index)).data;
                commit('QUERY', { query: 'BuyOrderBook', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryBuyOrderBook', payload: { options: { all }, params: { ...key }, query } });
                return getters['getBuyOrderBook']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryBuyOrderBook', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryBuyOrderBookAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryBuyOrderBookAll(query)).data;
                while (all && value.pagination && value.pagination.nextKey != null) {
                    let next_values = (await queryClient.queryBuyOrderBookAll({ ...query, 'pagination.key': value.pagination.nextKey })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'BuyOrderBookAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryBuyOrderBookAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getBuyOrderBookAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryBuyOrderBookAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async sendMsgSendCreatePair({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSendCreatePair(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSendCreatePair:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSendCreatePair:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgSendSellOrder({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSendSellOrder(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSendSellOrder:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSendSellOrder:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgSendCreatePair({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSendCreatePair(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSendCreatePair:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSendCreatePair:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgSendSellOrder({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSendSellOrder(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSendSellOrder:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSendSellOrder:Create', 'Could not create message: ' + e.message);
                }
            }
        },
    }
};