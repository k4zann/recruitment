<script setup>
  import Navbar from '@/components/Navbar.vue'
  import {RouterView} from 'vue-router';
  import Footer from './components/Footer.vue';
</script>

<template>
  <Navbar/>
  <RouterView />
  <footer>
    <Footer/>
  </footer>
</template>

