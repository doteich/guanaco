import { defineStore } from "pinia";
import { AddClient } from "../../wailsjs/go/main/App"



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
        listen(){
            console.log("start listening to events")
            window.runtime.EventsOn("keepalive-message", (data)=>{
                console.log(data)
            })
        },
        addClient(name, ep, mode, policy, auth, user, password) {
            AddClient(name, ep, mode, policy, auth, user, password)
        }
    }   

})


