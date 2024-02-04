<script setup>
import { onMounted, watch } from 'vue';
import { useClientStore } from './store/clientStore';
import Toast from 'primevue/toast';
import { useToast } from "primevue/usetoast";
import TabMenu from 'primevue/tabmenu';
import { useRouter } from 'vue-router';

const toast = useToast();
const router = useRouter()

import clientView from './views/clientView.vue';
import sidenav from "./components/sideNav.vue"
import { storeToRefs } from 'pinia';

const store = useClientStore()
const { getToast } = storeToRefs(store)


onMounted(() => {
  store.listen()
  store.getActiveConnections()
  // setTimeout(() => {
  //   if (store.getClients.length < 1) {
  //     store.addClient("SESSION", "opc.tcp://192.168.178.108:49320", "SignAndEncrypt", "Basic256Sha256", "Anonymous")
  //   }
  // }, 3000)

})

watch(getToast, (newToast) => {

  toast.add({ severity: newToast.severity, summary: newToast.summary, detail: newToast.detail, life: newToast.life })
})

const items = [{
  label: "Browse",
  icon: "pi pi-list",
  command: (() => {
    router.push("/browse")
  })
  },
  {
    label: "Monitor",
    icon: "pi pi-eye",
    command: (() => {
    router.push("/monitor")
  })
  },
  {
    label: "Log",
    icon: "pi pi-database"
  },
  {
    label: "Query Logs",
    icon: "pi pi-chart-line"
  }, {
    label: "Save/Load Config",
    icon: "pi pi-file"
  },
  {
    label: "Certs",
    icon: "pi pi-credit-card"
  }

]


</script>
<template>
  <nav>
    <sidenav @click="show"></sidenav>
  </nav>
  <main>
    <header>
      <TabMenu :model="items" />
    </header>
    <Toast />
    <RouterView></RouterView>
  </main>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
