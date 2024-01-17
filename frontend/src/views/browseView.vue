<script setup>
import { ref } from "vue"

import { useClientStore } from '../store/clientStore';
import Button from 'primevue/button';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';


const store = useClientStore()

const selectedVars = ref([])

const nodeToDrop = ref()

function browseNode(nodeid, id, type) {
    if (type == "NodeClassVariable") {
        return
    }
    store.Browse(nodeid, id)
}

function selectNode(name, nodeId) {
    selectedVars.value.push({
        name,
        nodeId
    })

}

function dropNode() {
    selectedVars.value = selectedVars.value.filter(el => el.nodeId != nodeToDrop.value.nodeId)
}

function exportSelection(){
    store.ExportBrowsedNodes(selectedVars.value)
}

</script>

<template>
    <section class="browser">

        <div v-for="node in store.getBrowseResults" class="browse-node" @click="browseNode(node.nodeId, node.id, node.type)"
            :style="{ 'margin-left': node.id.split('.').length * 20 + 'px' }">

            <i class="pi pi-chevron-down" v-if="node.isExpanded"></i>
            <i class="pi pi-chevron-right turn" v-else></i>
            <div class="browse-node-content">
                <div>
                    <i :class="'pi ' + node.icon" :style="{ 'color': node.color }"></i>
                </div>
                <p> <span>{{ node.name }}</span></p>
                <Button icon="pi pi-plus" size="small" aria-label="Add" @click="selectNode(node.name, node.nodeId)" text
                    v-if="node.type == 'NodeClassVariable'" />
            </div>

        </div>
        <div class="browse-actions-bar" v-if="selectedVars.length > 0">

            <div class="selection">
                <DataTable :value="selectedVars" tableStyle="min-width: 5rem" v-model:selection="nodeToDrop"
                    selectionMode="single" dataKey="name" @rowSelect="dropNode()">
                    <Column field="name" header="Selection"></Column>

                </DataTable>
            </div>

            <Button icon="pi pi-eye" size="small" aria-label="Add" label="Monitor Selection" raised />
            <Button icon="pi pi-file" size="small" aria-label="Add" @click="exportSelection()"
                label="Export Node-IDs" raised />

        </div>
    </section>
</template>
<style>
.browser {
    margin-top: 30px;
}

.browse-node {
    display: flex;
    align-items: center;
    width: fit-content;
    height: 2.0em;
    cursor: pointer;
    padding: 0 0 0 0;
    margin: 10px;
    background-color: var(--theme-color-1);

    width: 33%;

}

.browse-node>i {
    margin-right: 5px;
}

.browse-node:hover>.turn {
    transform: rotate(90deg);
}

i {
    color: var(--theme-color-3)
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

.browse-actions-bar {
    position: fixed;
    right: 2vw;
    bottom: 2vh;
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


@keyframes fadein {
    from {
        opacity: 0
    }

    to {
        opacity: 1
    }
}
</style>