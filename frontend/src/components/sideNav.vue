<script setup>
import { ref } from "vue"
import { useClientStore } from "../store/clientStore"


import Button from 'primevue/button';
import Sidebar from 'primevue/sidebar';
import InputText from 'primevue/inputtext';
import Dropdown from 'primevue/dropdown';
import Chip from 'primevue/chip';

const store = useClientStore()

const modal = ref(false)
const slimNav = ref(true)


const clientOpts = ref({
    session: "",
    ep: "",
    policy: "",
    mode: "",
    authType: "Anonymous",
    username: "",
    password: ""
})

const modes = [{ type: "None" }, { type: "Sign" }, { type: "SignAndEncrypt" },]
const policies = [{ type: "None" }, { type: "Basic256Sha256" }, { type: "Aes256Sha256RsaPss" }, { type: "Aes128Sha256RsaOaep" },]
const auth = [{ type: "Anonymous" }, { type: "User&Password" }]


function showModal() {
    modal.value ? modal.value = false : modal.value = true
}

function addClient() {
    store.addClient(clientOpts.value.session, clientOpts.value.ep, clientOpts.value.mode.type, clientOpts.value.policy.type, clientOpts.value.authType.type, clientOpts.value.username, clientOpts.value.password)
    setTimeout(() => {
        showModal()
    }, 1000)

}

function toggleNav() {
    slimNav.value ? slimNav.value = false : slimNav.value = true
}

function toggleClient(id, status){
    if (status == "disconnected"){
        store.reconnect(id)
    }else{
        store.disconnectClient(id)
    }

   
}



</script>


<template>
    <section :class="[slimNav ? 'side-nav-slim' : 'side-nav']">
        <img src="../assets/images/logo.png" class="logo">
        <p v-if="!slimNav">Clients</p>


        <div v-for="client in store.getClients">

            <Chip class="client-chip" @click="toggleClient(client.id, client.status)" >
                <span class="client-label" :style="[client.status == 'disconnected' ? 'backgroundColor: crimson' : 'backgroundColor: green']">{{ client.name[0] }}</span>
                <span class="client-name" v-if="!slimNav">{{client.name}}</span>
            </Chip>
        </div>


        <Button icon="pi pi-plus" size="small" aria-label="Filter" @click="showModal" text rounded />
        <Button icon="pi pi-angle-double-right" text rounded class="nav-toggle" @click="toggleNav" v-if="slimNav"></Button>
        <Button icon="pi pi-angle-double-left" text rounded class="nav-toggle" @click="toggleNav" v-else></Button>
        <Sidebar v-model:visible="modal" header="Client Options">
            <div class="input-group">
                <label for="session">Config Name</label>
                <InputText id="session" v-model="clientOpts.session" placeholder="OPC XY" />
            </div>
            <div class="input-group">
                <label for="endpoint">Endpoint</label>
                <InputText id="endpoint" v-model="clientOpts.ep" placeholder="opc.tcp://127.0.0.1:4840" />
            </div>
            <div class="input-group">
                <label for="mode">Security Mode</label>
                <Dropdown v-model="clientOpts.mode" :options="modes" showClear optionLabel="type" placeholder="Mode"
                    id="mode" />
            </div>
            <div class="input-group">
                <label for="policy">Security Policy</label>
                <Dropdown v-model="clientOpts.policy" :options="policies" showClear optionLabel="type" placeholder="Policy"
                    id="policy" />
            </div>
            <div class="input-group">
                <label for="auth">Authentication</label>
                <Dropdown v-model="clientOpts.authType" :options="auth" showClear optionLabel="type" placeholder="Auth"
                    id="auth" />
            </div>
            <div class="input-group" v-if="clientOpts.authType.type == 'User&Password'">
                <label for="username">Username</label>
                <InputText id="username" v-model="clientOpts.username" />
            </div>
            <div class="input-group" v-if="clientOpts.authType.type == 'User&Password'">
                <label for="password">Password</label>
                <InputText id="password" v-model="clientOpts.password" type="password" />
            </div>

            <Button icon="pi pi-plus" label="Add" aria-label="Filter" @click="addClient()" />

        </Sidebar>


    </section>
</template>

<style>
.logo {
    width: 60px;
    background: rgb(196, 190, 255);
    margin: 5px;
    border-radius: 5px;
}

.p-chip {
    padding: 0 !important;
}

.side-nav {
    width: 150px;
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 100vh;
}

.side-nav-slim {
    width: 70px;
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 100vh;
}

.client-label {
    background: green;
    padding: 5px 10px;
    border-radius: 70%;
    justify-self: flex-start;
}

.client-name {
    padding: 0 10px;
}

.client-chip{
    cursor: pointer;
    margin: 10px 0; 
}
.client-chip:hover{
    border:1px solid var(--theme-color-3);
    transform: scale(1.05);
}

.input-group {
    text-align: left;
    margin: 20px 0;
    display: flex;
    flex-direction: column;

}

.input-group>label {
    font-family: monospace;
}

.nav-toggle {
    margin-top: auto;
    margin-bottom: 10%;
}
</style>