<script setup>
    import Hero from '@/components/Hero.vue'
    import HomeCards from '@/components/HomeCards.vue'
    import JobListings from '@/components/JobListings.vue';
    import VideoCards from '@/components/VideoCards.vue';
</script>

<template>
    <Navbar/>
    <Hero title="Доступная работа – доступные возможности" 
    subtitle="Найдите работу, соответствующую вашим способностям и уникальности"
    />
    <HomeCards/>
    <VideoCards/>
    <JobListings :limit="3" :showButton="true"/>
</template>