<script setup>
import { onMounted, ref } from "vue"

import Dropdown from 'primevue/dropdown';
import { useServiceStore } from "../store/serviceStore"
import { useQueryStore } from "../store/queryStore"
import { storeToRefs } from "pinia";
import Sidebar from 'primevue/sidebar';

const serviceStore = useServiceStore()
const queryStore = useQueryStore()

const selectedService = ref("")

onMounted(() => {
    serviceStore.fetchServices()
})


</script>


<template>

    <section>

        <div class="query-menu-bar">
            <div class="logger-selection">
                <p>Logger Selection</p>
                <Dropdown v-model="selectedService" :options="serviceStore.getServices" optionLabel="id"
                    placeholder="Available Loggers" checkmark :highlightOnSelect="true" style="border-radius: 0;" />
            </div>
            <div class="logger-selection">
                <p>Node Id</p>
                <Dropdown v-model="selectedService" :options="serviceStore.getServices" optionLabel="id"
                    placeholder="Available Loggers" checkmark :highlightOnSelect="true" style="border-radius: 0;" />
            </div>
            <div class="logger-selection">
                <p>Node Id</p>
                <Dropdown v-model="selectedService" :options="serviceStore.getServices" optionLabel="id"
                    placeholder="Available Loggers" checkmark :highlightOnSelect="true" style="border-radius: 0;" />
            </div>

        </div>

    </section>

</template>

<style>
.query-menu-bar {
    display: flex;
    margin: 0vh 0vw;
    width: 100%;
}


.logger-selection {
    display: flex;
    flex-direction: column;
    width: fit-content;
    border: 1px solid var(--theme-color-3);
    background: var(--theme-color-3);
    align-items: center;

}

.logger-selection>p {
    margin: 0;
    padding: 4px;
    color: var(--theme-color-1);
    font-weight: bold;
}
</style>