<script setup>
import { useClientStore } from '../store/clientStore';

const store = useClientStore()


function browseNode(nodeid, id) {
    store.Browse(nodeid, id)
}

</script>

<template>
    <section class="browser">

        <div v-for="node in store.getBrowseResults" class="browse-node" @click="browseNode(node.nodeId, node.id)"
            :style="{ 'margin-left': node.id.split('.').length * 20 + 'px' }">

            <i class="pi pi-chevron-down" v-if="node.isExpanded"></i>
            <i class="pi pi-chevron-right turn" v-else></i>
            <div class="browse-node-content">
                <div>
                    <i :class="'pi ' + node.icon" :style="{'color': node.color}"></i>
                </div>
                <p> <span>{{ node.name }}</span></p>
            </div>

        </div>

    </section>
</template>
<style>
.browser {
    margin-top: 30px;
}

.browse-node {
    display: flex;
    align-items: center;
    width: fit-content;
    height: 2.0em;
    cursor: pointer;
    padding: 0 0 0 0;
    margin: 10px;
    background-color: var(--theme-color-1);
    
    width: 33%;

}

.browse-node>i {
    margin-right: 5px;
}

.browse-node:hover>.turn {
    transform: rotate(90deg);
}

i {
    color: var(--theme-color-3)
}

.browse-node-content {
    border-bottom: 1px solid var(--theme-color-3);
    display: flex;
    align-items: center;
    margin: 0;
    height: 100%;
    min-width: fit-content;
    background: var(--theme-color-2);
    width: 100%;
}

.browse-node-content>div {
    padding: 10px 12px;
    border-top: 1px solid var(--theme-color-3);
    border-right: 1px solid var(--theme-color-3);
    border-left: 1px solid var(--theme-color-3);
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    margin: 0;
}

.browse-node-content>p {
    padding: 0 20px;
}
</style>