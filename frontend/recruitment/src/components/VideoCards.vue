<script setup>
import VideoCard from './VideoCard.vue';

const videos = [
  'https://www.instagram.com/p/DCq1s1toxmL/?igsh=MWYwdXQwY2dkb2hpag==',
  'https://www.instagram.com/p/DCRHkYaIjdW/?igsh=a2tkYmtneWFsbmc4',
  'https://www.instagram.com/p/DCKhdmhuYUW/?igsh=MWhrNzNocTlqbm02ZA==',
];
</script>

<template>
  <section class="px-4 py-10">
    <div class="m-auto container-xl lg:container">
      <h2 class="mb-6 text-3xl font-bold text-center text-darker-blue">
        Полезное
      </h2>
      <div class="grid grid-cols-1 gap-6 md:grid-cols-2">
        <VideoCard
          v-for="video in videos"
          :key="video"
          :video="video"
        />
      </div>
    </div>
  </section>
</template>
