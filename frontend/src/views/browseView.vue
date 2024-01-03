<script setup>
import { useClientStore } from '../store/clientStore';

const store = useClientStore()


function browseNode(nodeid, id){
    store.Browse(nodeid, id)
}

</script>

<template>
    <section>

        <div v-for="node in store.getBrowseResults" class="browse-node" @click="browseNode(node.nodeId, node.id)" :style="{ 'margin-left': node.id.length * 5 + 'px' }">
            <div>
                <i :class="'pi ' + node.icon"></i>
                <span>{{ node.id }}</span>
            </div>
            <p> <span>{{ node.name }}</span></p>
        </div>
    </section>
</template>
<style>



.browse-node {
    display: flex;
    align-items: center;
    width: fit-content;
    height: 2.5em;
    cursor: pointer;
    padding: 0 0 0 0;
    margin: 10px; 
    background-color: var(--theme-color-2);
    border-bottom: 1px solid var(--theme-color-3);
    

}

.browse-node>div {
    color: rgb(41, 41, 39);
    background-color: var(--theme-color-1);
    padding: 0 15px;
    height: 100%;
    font-size: 20px;
    display: flex;
    border-top: 1px solid var(--theme-color-3);
    border-left: 1px solid var(--theme-color-3);
    border-right: 1px solid var(--theme-color-3);
    align-items: center;
    justify-content: center;
    flex-direction: column;

}
.browse-node>div>i{
    color: var(--theme-color-3)
}


.browse-node>div>span{
    font-size: 11px;
    color: var(--border-color-1)
}

.browse-node>p {
    margin-left: 5px;
    font-size: 14px;
    padding: 5px;
    min-width: 100px;
}
</style>