<script setup>
import { ref } from "vue"
import { useClientStore } from "../store/clientStore"
import Button from 'primevue/button';
import Chip from 'primevue/chip';
import TieredMenu from 'primevue/tieredmenu';

import clientSideBar from "./clientSideBar.vue";


const store = useClientStore()

const slimNav = ref(true)

let selectedClient = {
    id: "",
    status: ""
}

const items = [{
    label: '',
    icon: 'pi pi-refresh',
    command: () => {
        toggleClient()
    }
}]

const menu = ref();

const toggle = (event, id, status) => {
    selectedClient.id = id
    selectedClient.status = status
    if (status == "disconnected") {
        items[0].label = "Connect"
        items[0].icon = "pi pi-play"
    } else {
        items[0].label = "Disconnect"
        items[0].icon = "pi pi-times"
    }
    menu.value.toggle(event);
};

function toggleNav() {
    slimNav.value ? slimNav.value = false : slimNav.value = true
}

function toggleClient() {
    if (selectedClient.status == "disconnected") {
        store.reconnect(selectedClient.id)
    } else {
        store.disconnectClient(selectedClient.id)
    }
}

function sClient(id) {
    store.selectClient(id)
}

function save(){
    store.saveConfig()
}
function load(){
    store.loadConfig()
}



</script>


<template>
    <section :class="[slimNav ? 'side-nav-slim' : 'side-nav']">
        <img src="../assets/images/logo.png" class="logo">
        <p v-if="!slimNav">Clients</p>

        <div v-for="client in store.getClients" aria-haspopup="true" aria-controls="overlay_tmenu"
            @contextmenu.prevent="toggle($event, client.id, client.status)" @click="sClient(client.id)">

            <Chip class="client-chip">
                <span :style="[client.status == 'disconnected' ? 'backgroundColor: crimson' : 'backgroundColor: green']"
                    :class="[client.selected ? 'client-label-selected' : 'client-label']">{{
                        client.name[0] }}</span>
                <span class="client-name" v-if="!slimNav">{{ client.name }}</span>
            </Chip>
        </div>

        <TieredMenu ref="menu" id="overlay_tmenu" :model="items" popup />

        <clientSideBar></clientSideBar>

        <div class="nav-actions">
            <Button icon="pi pi-file-import" text rounded label="Load" v-if="!slimNav" @click="load()"></Button>
            <Button icon="pi pi-file-import" text rounded v-if="slimNav" @click="load()"></Button>
            <Button icon="pi pi-file-export" text rounded label="Save" v-if="!slimNav" @click="save()"></Button>
            <Button icon="pi pi-file-export" text rounded v-if="slimNav" @click="save()"></Button>
            <Button icon="pi pi-angle-double-right" text rounded class="nav-toggle" @click="toggleNav"
                v-if="slimNav"></Button>
            <Button icon="pi pi-angle-double-left" text rounded class="nav-toggle" @click="toggleNav" v-else></Button>
        </div>

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
    height: 100%;
}

.side-nav-slim {
    width: 70px;
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 100%;
}

.client-label {
    padding: 5px 10px;
    border-radius: 70%;
    justify-self: flex-start;
}

.client-label-selected {
    padding: 5px 10px;
    border-radius: 70%;
    transform: scale(1.05);
    justify-self: flex-start;
    border: 1px solid salmon;
}

.client-name {
    padding: 0 10px;
}

.client-chip {
    cursor: pointer;
    margin: 10px 0;
}

.client-chip:hover {
    border: 1px solid var(--theme-color-3);
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

.nav-actions {
    margin-top: auto;
    display: flex;
    flex-direction: column;
    margin-bottom: 10%;
    align-items: center;
}

.nav-toggle {

    margin-bottom: 10%;
}
</style>