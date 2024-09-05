<script setup>
  import Navbar from '@/components/Navbar.vue'
  import Hero from '@/components/Hero.vue'
  import HomeCards from '@/components/HomeCards.vue'
  import JobListings from './components/JobListings.vue';
</script>

<template>
  <Navbar/>
  <Hero title="Доступная работа – доступные возможности" 
    subtitle="Найдите работу, соответствующую вашим способностям и уникальности"
    />
  <HomeCards/>
  <JobListings :limit="3" :showButton="true"/>
</template>

