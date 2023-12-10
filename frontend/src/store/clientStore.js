import { defineStore } from "pinia";
import { AddClient, DisconnectClient, ReconnectClient, GetClients } from "../../wailsjs/go/main/App"


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
        getToast(state) {
            return state.toast
        }
    },
    actions: {
        listen() {
            window.runtime.EventsOn("client-message", (id, event) => {
                switch (event) {
                    case 'disconnect':
                        this.clients.filter(c => c.id == id)[0].status = "disconnected"
                        break;
                    case 'reconnect':
                        this.clients.filter(c => c.id == id)[0].status = "connected"
                        break;
                    case 'keepalive':
                        console.log(id, event)
                        break;
                }
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
                        status: "connected"
                    })
                })
                .catch((err) => {
                    this.toast = { severity: "error", summary: "Failed to Add OPC UA Client", detail: err, life: 5000 }
                })
        },
        disconnectClient(id) {
            DisconnectClient(id)

        },
        reconnect(id) {
            ReconnectClient(id)
        },
        getActiveConnections() {
            GetClients()
                .then(res => {
                    res.forEach((client) => {
                        this.clients.push({
                            id: client.ClientId,
                            name: client.Name,
                            status: client.Status
                        })
                    })
                })
        }
    }

})


