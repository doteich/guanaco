<script setup>
import { ref } from "vue"

import Dropdown from 'primevue/dropdown';
import Button from 'primevue/button';
import Calendar from 'primevue/calendar';


const props = defineProps({
    services: Array,
    nodeIds: Array,
    nodeNames: Array,  
})

const emits = defineEmits(['on-select', 'on-reload', 'on-resize'])


let start = new Date()
start.setHours(start.getHours() - 1)


const selectedService = ref("")
const dateRange = ref([start, new Date()])
const showMenu = ref(true)
const selectedNodeId = ref("")
const selectedNodeName = ref("")



function toggleMenu() {
    showMenu.value = !showMenu.value
    emits('on-resize', showMenu.value)
}




</script>
<template>
    <div style="background-color: var(--theme-color-2);">
        <Button icon="pi pi-chevron-up" text size="small" @click="toggleMenu()" style="padding: 1.5%" v-if="showMenu" />
        <Button icon="pi pi-chevron-down" text size="small" @click="toggleMenu()" style="padding: 1.5%" v-else />
    </div>
    <Transition>

        <div class="query-menu-bar" v-if="showMenu">
            <div class="logger-selection">
                <div class="logger-selection-selector">
                    <p>Logger Selection</p>
                    <Dropdown v-model="selectedService" :options="services" optionLabel="id"
                        placeholder="Available Loggers" checkmark :highlightOnSelect="true" style="border-radius: 0;"
                        @change="$emit('on-select', selectedService.id)" />
                </div>
                <Button icon="pi pi-refresh" label="Reload" size="small" @click="$emit('on-reload', selectedService.id, selectedNodeId, selectedNodeName, dateRange[0], dateRange[1])"
                    style="margin-left: auto; margin-right: 2%; padding: 1.5%; width: fit-content;"
                    :disabled="selectedService == ''" />
            </div>
            <div class="query-selection">
                <div class="query-selection-selector">
                    <p>Node Id</p>
                    <Dropdown v-model="selectedNodeId" showClear :options="nodeIds"
                        placeholder="Node IDs" checkmark :highlightOnSelect="true" style="border-radius: 0;" />
                </div>
                <div class="query-selection-selector">
                    <p>Node Name</p>
                    <Dropdown v-model="selectedNodeName" showClear :options="nodeNames"
                        placeholder="Node Names" checkmark :highlightOnSelect="true" style="border-radius: 0;" />
                </div>
                <div class="query-selection-selector">
                    <p>Date Range</p>
                    <Calendar v-model="dateRange" selectionMode="range" showTime :manualInput="true"
                        style="border-radius: 0;" />
                </div>
            </div>
        </div>
    </Transition>


</template>

<style>
.query-menu-bar {
    display: flex;
    margin: 0vh 0vw;
    width: 100%;
    flex-direction: column;
    background-color: var(--theme-color-2);
    padding: 10px;
    border-bottom: 1px solid var(--theme-color-3);
}


.logger-selection {
    display: flex;
    width: 99%;
    margin: 2px 5px;
    justify-content: flex-start;
    align-items: flex-end;
    width: 100%;

}

.logger-selection-selector {
    width: calc(33% - 5px);

}

.logger-selection-selector>* {
    width: 100%;
}

.logger-selection-selector>p {
    text-align: left;
}

.logger-selection-selector>*>span {
    text-align: left;
}




.query-selection {
    display: flex;
    width: 100%;
}

.query-selection-selector {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
    margin: 0 5px;
    width: 33%;
}

.query-selection-selector>p {
    text-align: start;
}

.query-selection-selector>* {
    border-radius: 0;
    width: 100%;
}

.query-selection-selector>* input {
    border-radius: 0;
}

.query-selection-selector>*>span {
    text-align: left;
}

.v-enter-active,
.v-leave-active {
    transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
    opacity: 0;
}
</style>