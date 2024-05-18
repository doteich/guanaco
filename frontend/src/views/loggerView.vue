<script setup>

import { onMounted } from 'vue';
import {useServiceStore} from "../store/serviceStore"

import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import { storeToRefs } from 'pinia';

const store = useServiceStore()

onMounted(()=>{
    store.getServices()
})


</script>

<template>
    <section class="logger-config">
    
    </section>

</template>
<style scoped></style>