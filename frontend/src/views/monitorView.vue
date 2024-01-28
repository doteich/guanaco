<script setup>
import { useClientStore } from '../store/clientStore';

const store = useClientStore()

</script>

<template>

<section>
    <div v-for="value, key in store.getMonitoredItems">
        {{ value.name }}
        <div v-for="item in value.items">
           {{ item.name }} {{ item.value }}
        </div>


    </div>
</section>


</template>

<style scoped></style>

