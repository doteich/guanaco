<script setup>
import { useClientStore } from '../store/clientStore';

const store = useClientStore()

function browse(){
    store.Browse()
}

</script>

<template>
    <section>
        <h1>HALLO WELT</h1>
        <button @click="browse">BROWSE</button>
    </section>
</template>

<style></style>