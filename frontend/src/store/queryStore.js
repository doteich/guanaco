import { defineStore } from "pinia"
import { GetUniqueEntries, GetTimeSeries, SaveResultsAsCSV } from "../../wailsjs/go/main/App"
import { useGeneralStore } from "./generalStore"


export const useQueryStore = defineStore("queryStore", {

    state: () => ({
        selectedLogger: "",
        uniqueNodeIds: [],
        uniqueNodeNames: [],
        results: []

    }),
    getters: {
        getUniqueNodeIds(state) {
            return state.uniqueNodeIds
        },
        getUniqueNodeNames(state) {
            return state.uniqueNodeNames
        },
        getResults: state => state.results

    },
    actions: {
        FetchUniqueValues(svc, type) {
            GetUniqueEntries(svc, type)
                .then((res) => {
                    switch (type) {
                        case 'nodeName':
                            this.uniqueNodeNames = res
                            break
                        case 'nodeId':
                            this.uniqueNodeIds = res

                    }
                })
                .catch(err => {
                    console.error(err)
                    useGeneralStore().setToast("error", `Failed to fetch unique values`, err, 3000)
                })
        },
        FetchTimeSeriesData(svc, nodeId, nodeName, start, end) {
            GetTimeSeries(svc, nodeId, nodeName, start, end)
                .then((res) => {
                    this.results = res
                })
                .catch((err) => {
                    console.error(err)
                    useGeneralStore().setToast("error", `Failed to fetch timeseries`, err, 3000)
                })
        },
        ExportResults(csv) {
            SaveResultsAsCSV(csv)
                .then((file) => {
                    useGeneralStore().setToast("success", "Exported", `Results exported to ${file}`, 3000)
                })
                .catch(err => {
                    useGeneralStore().setToast("error", `Failed to export results`, err, 3000)
                })

        }


    }

}


)