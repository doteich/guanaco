import { defineStore, storeToRefs } from "pinia";
import { GetServices } from "../../wailsjs/go/main/App"

export const useServiceStore = defineStore("serviceStroe", {
    state: () => ({
        services: []
    }),
    getters: {
        getServices(state){
            return this.services
        }
    },
    actions: {
        fetchServices() {
            GetServices()
                .then((json) => {
                    this.services = JSON.parse(json)
                })
                .catch((err) => {
                    console.log(err)
                })
        }
    }
})