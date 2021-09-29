import { createStore, useStore as baseUseStore, Store } from "vuex";
import {Torrent} from "@/pb/baseService_pb";
import {InjectionKey} from "vue";

export interface State {
  searchResults: Torrent.AsObject[];
  downloadFolders: string[];
}

export const store = createStore<State>({
  state: {
    searchResults: [],
    downloadFolders: [],
  },
  mutations: {
    setSearchResults(state, results) {
      state.searchResults = results;
    },
    setDownloadFolders(state, folders) {
      state.downloadFolders = folders;
    },
  },
  actions: {},
  modules: {},
});

export const key: InjectionKey<Store<State>> = Symbol();

export function useStore() {
  return baseUseStore(key);
}
