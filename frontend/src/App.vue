<script setup>
import clientView from './views/clientView.vue';
import sidenav from "./components/sideNav.vue"

</script>
<template>
  <nav>
    <sidenav></sidenav>
  </nav>
  <main>
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
