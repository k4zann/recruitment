<script setup>
    import { ref } from 'vue';

    const props = defineProps({
        video: Object  
    });

    const videoId = ref(props.video);  
</script>


<template>
    <div class="bg-gray-100 rounded-xl shadow-md relative">
      <div class="p-4">
        <div class="aspect-w-16 aspect-h-9">
            <iframe
              title="video"
              :src="'https://www.youtube.com/embed/' + videoId"
              allowfullscreen
              class="w-full h-full rounded-lg"
            ></iframe>
          </div>
      </div>
    </div>
</template>

<style scoped>
    .fixed {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
    }

    .aspect-w-16.aspect-h-9 {
        position: relative;
        width: 100%;
        padding-bottom: 56.25%; 
    }

    iframe {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
    }
</style>