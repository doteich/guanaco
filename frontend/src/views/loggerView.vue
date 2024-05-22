<script setup>

import { onMounted, ref, computed } from 'vue';
import { useServiceStore } from "../store/serviceStore"
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import Tag from 'primevue/tag';

import Dialog from 'primevue/dialog';

import { storeToRefs } from 'pinia';

const store = useServiceStore()

const showInfos = ref(false)
const showItems = ref(false)

function toggle(name, cmd) {
    store.toggleService(name, cmd)
}

const { getServiceInfos } = storeToRefs(store)


function info(name) {
    store.fetchServiceInfo(name)
    showInfos.value = true
}
function toggleItems() {
    showItems.value = !showItems.value
}


onMounted(() => {
    store.fetchServices()
})


</script>

<template>
    <section class="logger-config">
        <DataTable :value="store.getServices" class="logger-table">
            <Column field="id" header="Service ID" style="width: 50%;"></Column>

            <Column field="statusName" header="Status" style="width: 30%;">
                <template #body="slotProps">
                    <Tag :value="slotProps.data.statusName" :severity="slotProps.data.severity" />
                </template>
            </Column>
            <Column :exportable="false">
                <template #body="slotProps">
                    <Button icon="pi pi-stop" v-if="slotProps.data.status == 1" text severity="warning"
                        @click="toggle(slotProps.data.id, 'stop')" />
                    <Button icon="pi pi-play" v-else text class="" @click="toggle(slotProps.data.id, 'start')" />
                    <Button icon="pi pi-info" text @click="info(slotProps.data.id)" />
                    <Button icon="pi pi-times" text severity="danger" @click="" />
                </template>
            </Column>
        </DataTable>


        <Dialog v-model:visible="showInfos" :style="{ width: '50rem' }">

            <div class="service-infos" v-if="!showItems">
                <p> <span class="info-heading">ID</span><span class="info-content">{{ getServiceInfos.id }}</span></p>
                <p> <span class="info-heading">Endpoint</span><span class="info-content">{{ getServiceInfos.ep }}</span>
                </p>
                <p> <span class="info-heading">Mode</span><span class="info-content">{{ getServiceInfos.mode }}</span>
                </p>
                <p> <span class="info-heading">Policy</span><span class="info-content">{{ getServiceInfos.policy
                        }}</span></p>
                <p> <span class="info-heading">Authentication</span><span class="info-content">{{ getServiceInfos.auth
                        }}</span>
                </p>
                <p> <span class="info-heading">Items</span><span class="info-content" style="width: 74%;"> {{
            getServiceInfos.monitoredItems.length }} </span> <Button icon="pi pi-info" size="small" text
                        @click="toggleItems" /></p>
                <p> <span class="info-heading">Interval</span><span class="info-content">{{ getServiceInfos.interval
                        }}</span>
                </p>

            </div>
            <div v-else>
                <Button icon="pi pi-arrow-left" size="small" text
                        @click="toggleItems" />
                <ul>
                    <li v-for="item in getServiceInfos.monitoredItems">{{ item }}</li>
                </ul>
            
            </div>

        </Dialog>

    </section>


</template>
<style scoped>
.logger-table {
    font-size: 20px;
}

.service-infos {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    width: 100%;
}

.service-infos>p {
    border: 1px solid var(--theme-color-3);
    height: 30px;
    margin: 2px;
    width: 100%;
    text-align: left;
    display: flex;
    align-items: center;


}

.info-content {
    width: 80%;
    display: inline-block;
    overflow-x: auto;
    overflow-y: auto;
    text-wrap: nowrap;

}

.info-heading {
    display: inline-block;
    height: 100%;
    color: var(--theme-color-2);
    margin-right: 1%;
    padding: 5px;
    width: 20%;
    background-color: var(--theme-color-3);
}
</style>