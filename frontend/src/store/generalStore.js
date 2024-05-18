import {defineStore} from "pinia"


export const useGeneralStore = defineStore("generalStore", {
    state: () => ({
        toast: {
            severity: "",
            summary: "",
            detail: "",
            life: 3000,
        }, 
    }),
    getters:{
        getToast(state) {
            return state.toast
        },
    },
    actions:{
        setToast(severity, summary, detail, life){
            this.toast = {
                severity,
                summary,
                detail,
                life
            }
        }
    }
})