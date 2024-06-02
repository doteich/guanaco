import { defineStore } from "pinia"
import { GetUniqueEntries } from "../../wailsjs/go/main/App"
import { useGeneralStore } from "./generalStore"


export const useQueryStore = defineStore("queryStore", {

    state: () => ({

    }),
    getters: {},
    actions: {
        getUniqueNodes(svc, type){
            GetUniqueEntries(svc, type)
            .then((res)=>{
                console.log(JSON.parse(res))
            })
            .catch(err =>{
                useGeneralStore().setToast("error", `Failed to fetch unique nodes"`, err, 3000)
            })
        }


    }

}


)