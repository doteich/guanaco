<script setup>
import { onMounted, ref } from "vue"

import { useServiceStore } from "../store/serviceStore"
import { useQueryStore } from "../store/queryStore"

import queryMenuBar from '../components/queryMenuBar.vue';

import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';


const serviceStore = useServiceStore()
const queryStore = useQueryStore()

const tHeight = ref(54)


function getDropdownValues(svc) {
    queryStore.FetchUniqueValues(svc, "nodeName")
    queryStore.FetchUniqueValues(svc, "nodeId")
}

function getData(svc, ni, nn, start, end) {
    console.log(svc, ni, nn, start, end)
    queryStore.FetchTimeSeriesData(svc, ni, nn, start.toISOString(), end.toISOString())
}

function setTableHeight(isExpanded) {
    isExpanded ? tHeight.value = 55 : tHeight.value = 80
}


function saveCSV() {
    let line = ""
    let lines = ""

    queryStore.getResults.forEach(el => {
        for (let key in el) {
            line += el[key] + ";"
        }
        lines += line + "\n"
        line = ""
    })
    queryStore.ExportResults(lines)
}

onMounted(() => {
    serviceStore.fetchServices()
})

</script>


<template>
    <section>
       
        <queryMenuBar :services="serviceStore.getServices" :node-ids="queryStore.getUniqueNodeIds"
            :node-names="queryStore.getUniqueNodeNames" @on-reload="getData" @on-select="getDropdownValues"
            @on-resize="setTableHeight">
        </queryMenuBar>
        <div class="query-table">
            <DataTable :value="queryStore.getResults" :scrollHeight="tHeight + 'vh'" scrollable
                tableStyle="min-width:92vw">
                <template #header>
                    <div style="text-align: left;">
                        <button style="background: none; border: 0" @click="saveCSV" title="Export results as csv" v-if="queryStore.getResults.length > 0"><i class="pi pi-external-link"></i></button>
                    </div>
                </template>
                <Column field="ts" header="Timestamp"></Column>
                <Column field="nodeId" header="Node ID"></Column>
                <Column field="nodeName" header="Node Name"></Column>
                <Column field="value" header="Value"></Column>
            </DataTable>
        </div>
    </section>

</template>

<style>
.query-table {
    display: flex;
    align-items: center;
    padding: 5px;
    width: 93vw;
}

.p-datatable-header {
    padding: 5px;
    background: var(--theme-color-1);
}
</style>