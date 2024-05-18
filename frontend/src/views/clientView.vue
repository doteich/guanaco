<script setup>
import { useClientStore } from '../store/clientStore';

const store = useClientStore()

function browse() {
    store.Browse()
}

</script>

<template>
    <section class="start-screen">
        <div class="start-screen-image">
            <img src="../assets/images/logo.png" style="width: 20rem;">

        </div>
        <div style="margin: 1%">
            <h2 class="top">OPC</h2>
            <h1>G<span>UA</span>NACO</h1>
            <h2 class="bottom">Client & Logger</h2>
        </div>
    </section>
</template>

<style>
.start-screen {
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}

.start-screen>div>h1 {
    font-size: 3rem;
    margin: 0;
    
}

.start-screen>div>h1>span{
    color: var(--theme-color-3);
}

.start-screen>div>.top {
    margin: 0;
    color: var(--theme-color-3);
    text-align: left;
}

.start-screen>div>.bottom {
    margin: 0;
    padding: 0;
    text-align: right;
    color: var(--theme-color-3);
}

.start-screen-image {
    background: var(--theme-color-4);
    width: fit-content;
    border-radius: 10px;
  
}
</style>