<script setup>
import { useClientStore } from '../store/clientStore';

import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';


const store = useClientStore()

function toggleTable(id, status) {
    store.toggleTable(id, status)
}
function stop(id){
    store.stopNodeMonitor(id)
}



</script>

<template>
    <section class="monitor">

        <div v-for="value, key in store.getMonitoredItems">
            <div class="category-header">

                <h5>{{ value.name }}</h5>
                <div class="button-bar">
                    <Button icon="pi pi-database" text size="small"
                        @click="toggleTable(key, !value.isExpanded)" class="log-button"  severity="help" title="Setup Logger" />
                    <Button icon="pi pi-times"  text size="small"
                        @click="stop(key)" class="log-button" severity="danger" title="Drop Monitor" />

                    <Button icon="pi pi-chevron-right" text rounded size="small" class="turn"
                        @click="toggleTable(key, !value.isExpanded)" />
                </div>

            </div>

            <DataTable :value="value.items" class="table" v-if="value.isExpanded">

                <Column field="name" header="Tag" style="width: 30%"></Column>
                <Column field="value" header="Value"></Column>
                <Column field="ts" header="Source Timestamp"></Column>
            </DataTable>


        </div>
    </section>
</template>

<style scoped>
.monitor {
    display: flex;
    flex-direction: column;
    margin: 1vw;
}

.monitor>div {
    margin-top: 10px;
}

.category-header {
    background-color: var(--theme-color-2);
    display: flex;
    align-items: center;
}

h5 {
    color: var(--theme-color-3);
    margin: 0;
    padding-left: 2vw;
    font-size: 16px;
}

.button-bar {
    margin-left: auto;
    display: flex;
    justify-content: center;
    align-items: center;
}

.button-bar > button{
    margin: 0 10px;
}



.turn:hover {

    animation-name: turn;
    animation-duration: 0.1s;
    transform: rotate(90deg);
}

@keyframes turn {
    from {
        transform: rotate(0deg);
    }

    to {
        transform: rotate(90deg);
    }
}

.table {
    padding: 0 1vw;
    font-size: 17px;
}
</style>

