<script setup>
    import Hero from '@/components/Hero.vue'
    import HomeCards from '@/components/HomeCards.vue'
    import JobListings from '@/components/JobListings.vue';
    import VideoCards from '@/components/VideoCards.vue';
</script>

<template>
    <Navbar/>
    <Hero title="Cайт для трудоустройства лиц с инвалидностью" 
    subtitle="Найти работу, соискателя и создать проект"
    />
    <HomeCards/>
    <VideoCards/>
    <JobListings :limit="3" :showButton="true"/>
</template>