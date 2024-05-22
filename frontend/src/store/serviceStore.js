import { defineStore, storeToRefs } from "pinia";
import { GetServices, ToggleService, GetServiceInfo } from "../../wailsjs/go/main/App"
import { useGeneralStore } from "./generalStore"


export const useServiceStore = defineStore("serviceStroe", {
    state: () => ({
        services: [],
        serviceInfos: {},
    }),
    getters: {
        getServices(state) {
            return state.services
        },
        getServiceInfos(state) {
            return state.serviceInfos
        }
    },
    actions: {
        fetchServices() {
            GetServices()
                .then((json) => {
                    let svc = JSON.parse(json)
                    svc.forEach((s) => {
                        switch (s.status) {
                            case 0:
                                s.severity = "warning"
                                break
                            case 1:
                                s.severity = "success"
                                break
                            case 2:
                                s.severity = "danger"
                        }

                    })
                    this.services = svc
                })
                .catch((err) => {
                    useGeneralStore().setToast("error", "Failed to fetch Services", err, 3000)
                })
        },
        toggleService(name, cmd) {
            ToggleService(name, cmd)
                .then(() => {
                    this.fetchServices()
                })
                .catch((err) => {
                    useGeneralStore().setToast("error", `Failed to ${cmd} Service"`, err, 3000)
                })
        },
        fetchServiceInfo(name) {
            GetServiceInfo(name)
                .then((json) => {
                    this.serviceInfos = JSON.parse(json)
                })
                .catch(err => {
                    this.serviceInfos = {}
                    useGeneralStore().setToast("error", "Failed to fetch service info", err, 3000)
                })
        }
    }
})