<script setup>
import { ref } from 'vue';
import { useClientStore } from '../store/clientStore';


import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';

import InputText from 'primevue/inputtext';
import InputNumber from 'primevue/inputnumber';
import Chip from 'primevue/chip';

import Dialog from 'primevue/dialog';



const store = useClientStore()

function toggleTable(id, status) {
    store.toggleTable(id, status)
}
function stop(id) {
    store.stopNodeMonitor(id)
}

function StartLogger(id) {
    store.createLogger()
}


const showModal = ref(false)
const confValidation = ref("crimson")
const validConfName = ref(false)

function toggleModal(id, items) {

    loggerConf.value.displayName = ""

    loggerConf.value.confName = ""
    sanitizeConfName(loggerConf.value.confName)

    let client = store.getClients.find((c, items) => {
        return c.id == id
    })

    if (!client) {
        return
    }

    loggerConf.value.ep = client.ep
    loggerConf.value.mode = client.mode
    loggerConf.value.policy = client.policy
    loggerConf.value.auth = client.auth
    loggerConf.value.user = client.user
    loggerConf.value.password = client.password
    loggerConf.value.monitoredItems = items.map(e => e.nodeId)

    showModal.value = !showModal.value
}

function sanitizeConfName(val) {
    if (val == "") {
        confValidation.value = "crimson"
        validConfName.value = false
        return
    }
    let regex = /\s/;
    if (regex.test(val)) {
        confValidation.value = "crimson"
        validConfName.value = false
    } else {
        confValidation.value = "green"
        validConfName.value = true
    }

}

const loggerConf = ref({
    displayName: "",
    confName: "",
    ep: "",
    policy: "",
    mode: "",
    auth: "",
    user: "",
    password: "",
    monitoredItems: "",
    interval: 10,
})

function SendServiceData() {
    store.createLogger(loggerConf.value)
}


</script>

<template>
    <section class="monitor">
        <div v-for="value, key in store.getMonitoredItems">
            <div class="category-header">

                <h5>{{ value.name }}</h5>
                <div class="button-bar">
                    <Button icon="pi pi-database" text size="small" @click="toggleModal(key, value.items)"
                        class="log-button" severity="help" title="Setup Logger" />
                    <Button icon="pi pi-times" text size="small" @click="stop(key)" class="log-button" severity="danger"
                        title="Drop Monitor" />

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

        <Dialog v-model:visible="showModal" modal header="Configure Logger" :style="{ width: '25rem' }">
            <div class="input-group">
                <label for="name">Config Name</label>
                <InputText :style="{ borderColor: confValidation }" id="name" v-model="loggerConf.confName"
                    placeholder="SAMPLE_CONF" @input="sanitizeConfName(loggerConf.confName)" />
                <small id="name">Has to be a unique name without whitespaces</small>
            </div>
            <div class="input-group">
                <label for="displayName">Display Name</label>
                <InputText id="displayName" v-model="loggerConf.displayName" placeholder="My Logger" />

            </div>
            <div class="input-group">
                <label for="interval">Interval</label>
                <InputNumber v-model="loggerConf.interval" inputId="interval" suffix=" s" />

            </div>
            <div style="margin:10px">
                <Chip :label="'EP: ' + loggerConf.ep" class="chip" />
                <Chip :label="'Policy: ' + loggerConf.policy" class="chip" />
                <Chip :label="'Mode: ' + loggerConf.mode" class="chip" />
                <Chip :label="'Auth: ' + loggerConf.auth" class="chip" />
            </div>

            <Button label="Setup Logger" title="Create a Service in the OS" :disabled="!validConfName"
                @click="SendServiceData()" />

        </Dialog>

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

.button-bar>button {
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

.chip {
    font-size: 12px;
    padding: 1px 5px !important;
    margin: 2px;
}
</style>
