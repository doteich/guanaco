<script setup>

import { onMounted, ref } from 'vue';
import { useServiceStore } from "../store/serviceStore"
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import Tag from 'primevue/tag';

import Dialog from 'primevue/dialog';

import { storeToRefs } from 'pinia';

const store = useServiceStore()

const showInfos = ref(false)

function toggle(name, cmd) {
    store.toggleService(name, cmd)
}

function info(name) {
    store.fetchServiceInfo(name)
    showInfos.value = true
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
            <p>{{ store.getServiceInfos }}</p>
        </Dialog>

    </section>


</template>
<style scoped>
.logger-table {
    font-size: 20px;
}
</style>