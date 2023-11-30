import { defineStore } from "pinia";


export const useClientStore = defineStore("clientStore", {
    state: () => ({
        clients: []
    }),
    getters: {
        getClients(state) {
            return state.clients
        }
    },
    actions: {
        addClient(ep, mode, policy, auth, user, password){
            
        }
    }

})