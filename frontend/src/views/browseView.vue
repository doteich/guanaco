<script setup>
import { ref } from "vue"

import { useClientStore } from '../store/clientStore';
import Button from 'primevue/button';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Checkbox from 'primevue/checkbox';
import ProgressSpinner from 'primevue/progressspinner';
import Dialog from 'primevue/dialog';


const store = useClientStore()
const loading = ref(false)
const selectedVars = ref([])
const crawlEnabled = ref(false)
const autoSelect = ref(false)
const nodeToDrop = ref()

async function browseNode(nodeid, id, type) {
    if (type == "NodeClassVariable") {
        return
    }
    loading.value = true

    await store.Browse(nodeid, id, crawlEnabled.value)

    if (autoSelect.value){
       selectedVars.value = store.getBrowseResults?.filter(e => e.type === 'NodeClassVariable')
    }

    loading.value = false
}

function selectNode(name, nodeId, id, dataType) {
    selectedVars.value.push({
        name,
        nodeId,
        id,
        dataType
    })

}

function dropNode() {
    selectedVars.value = selectedVars.value.filter(el => el.nodeId != nodeToDrop.value.nodeId)
}

function exportSelection() {
    store.ExportBrowsedNodes(selectedVars.value)
}

function monitorItems() {
    store.CreateNodeMonitor(selectedVars.value)
}

</script>

<template>
    <section class="browser">
      
        <Dialog v-model:visible="loading" modal header="Browsing ..." :style="{ width: '25rem' }" :closable="false">
            <ProgressSpinner />
        </Dialog>
        <div v-for="node in store.getBrowseResults" class="browse-node"
            @click="browseNode(node.nodeId, node.id, node.type)"
            :style="{ 'margin-left': node.id.split('.').length * 20 + 'px' }">

            <i class="pi pi-chevron-down" v-if="node.isExpanded"></i>
            <i class="pi pi-chevron-right turn" v-else></i>
            <div class="browse-node-content">
                <div>
                    <i :class="'pi ' + node.icon" :style="{ 'color': node.color }"></i>
                </div>
                <div v-if="node.dataType != ''" class="variable">
                    <span class="datatype">{{ node.dataType }}</span>
                </div>
                <p> <span>{{ node.name }}</span> </p>

                <Button icon="pi pi-plus" size="small" aria-label="Add"
                    @click="selectNode(node.name, node.nodeId, node.id, node.dataType)" text
                    v-if="node.type == 'NodeClassVariable'" />
            </div>

        </div>
        <div class="browse-actions">
            <div class="crawler">
                <Checkbox inputId="crawl" v-model="crawlEnabled" :binary="true" />
                <label for="crawl">Enable Crawling</label>
                <i class="pi pi-info-circle" title="Recursive expands node until end of a node is reached"></i>
            </div>
            <div class="crawler">
                <Checkbox inputId="auto" v-model="autoSelect" :binary="true" />
                <label for="auto">Auto Select Tags</label>
                
            </div>
            <div v-if="selectedVars.length > 0" class="browse-actions-bar">
                <div class="selection">
                    <DataTable :value="selectedVars" tableStyle="min-width: 5rem" v-model:selection="nodeToDrop"
                        selectionMode="single" dataKey="name" @rowSelect="dropNode()">
                        <Column field="name" header="Selection"></Column>

                    </DataTable>
                </div>

                <Button icon="pi pi-eye" size="small" aria-label="Add" label="Monitor Selection" raised
                    @click="monitorItems()" />
                <Button icon="pi pi-file" size="small" aria-label="Add" @click="exportSelection()"
                    label="Export Node-IDs" raised/>
            </div>
        </div>
    </section>
</template>
<style>
.crawler {
    display: flex;
    align-items: center;
    margin: 5px 0
}

.crawler>label {
    margin: 0 10px
}

.browser {
    margin-top: 30px;

}

.browse-node {
    display: flex;
    align-items: center;
    min-width: fit-content;
    height: 2.0em;
    cursor: pointer;
    padding: 0 0 0 0;
    margin: 10px;
    background-color: var(--theme-color-1);
    width: 60vw;

}

.browse-node>i {
    margin-right: 5px;
}

.variable {
    background: var(--theme-color-3);
    width: 5%;
}



.browse-node:hover>.turn {
    transform: rotate(90deg);
}

i {
    color: var(--theme-color-3);
}

.browse-node-content {
    border-bottom: 1px solid var(--theme-color-3);
    display: flex;
    align-items: center;
    margin: 0;
    height: 100%;
    min-width: fit-content;
    background: var(--theme-color-2);
    width: 100%;
}

.browse-node-content>div {
    padding: 10px 12px;
    border-top: 1px solid var(--theme-color-3);
    border-right: 1px solid var(--theme-color-3);
    border-left: 1px solid var(--theme-color-3);
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    margin: 0;
}

.browse-node-content>p {
    padding: 0 20px;
}

.browse-node-content>button {
    margin-left: auto;
    margin-right: 10px;
    padding: 1px 3px;
}


.browse-actions {
    position: fixed;
    right: 2vw;
    bottom: 2vh;
}

.browse-actions-bar {

    display: flex;
    flex-direction: column;
    max-height: 80vh;


}

.browse-actions-bar>button {
    margin: 5px 0;
    animation-name: fadein;
    animation-duration: 1s;
}

table {
    font-size: 0.75em !important;
    background-color: var(--theme-color-1);
}

th {
    background-color: var(--theme-color-1) !important;
}

td {
    background-color: var(--theme-color-1);
}

.selection {
    max-height: 70vh !important;
    overflow-y: scroll;

}

.datatype {
    font-size: 11px;
    color: var(--theme-color-2);

    font-weight: bold;
    border-radius: 4px;
    padding: 2px;

}

@keyframes fadein {
    from {
        opacity: 0
    }

    to {
        opacity: 1
    }
}
</style>