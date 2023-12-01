import { defineStore } from "pinia";
import { AddClient } from "../../wailsjs/go/main/App"


export const useClientStore = defineStore("clientStore", {
    state: () => ({
        clients: [],
        toast: {
            severity: "",
            summary: "",
            detail: "",
            life: 3000,
        }

    }),
    getters: {
        getClients(state) {
            return state.clients
        },
        getToast(state){
            return state.toast
        }
    },
    actions: {
        listen() {
            console.log("start listening to events")
            window.runtime.EventsOn("keepalive-message", (data) => {
                console.log(data)
            })
        },
        addClient(name, ep, mode, policy, auth, user, password) {
            AddClient(name, ep, mode, policy, auth, user, password)
                .then((data) => {
                    this.clients.push({
                        id: data,
                        name: name,
                        mode,
                        policy,
                        auth,
                    })
                })
                .catch((err) => {
                    this.toast = { severity: "error", summary: "Failed to Add OPC UA Client", detail: err, life: 5000 }
                    console.log(this.toast)
                })
        }
    }

})


