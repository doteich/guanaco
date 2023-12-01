<script setup>
import { onMounted, watch } from 'vue';
import { useClientStore } from './store/clientStore';
import Toast from 'primevue/toast';
import { useToast } from "primevue/usetoast";
const toast = useToast();

import clientView from './views/clientView.vue';
import sidenav from "./components/sideNav.vue"
import { storeToRefs } from 'pinia';

const store = useClientStore()
const {getToast} = storeToRefs(store)


onMounted(() => {
  store.listen()
})

watch(getToast, (newToast) =>{
  console.log(newToast)
  toast.add({severity: newToast.severity, summary: newToast.summary, detail: newToast.detail, life: newToast.life})
})



</script>
<template>
  <nav>
    <sidenav @click="show"></sidenav>
  </nav>
  <main>
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
