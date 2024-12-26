<script setup>  
import { ref, computed } from 'vue';

    const props = defineProps({
    video: String, 
    });

    const videoEmbedUrl = computed(() => {
    if (props.video.includes('youtube.com') || props.video.includes('youtu.be')) {
        const videoIdMatch = props.video.match(/(?:v=|youtu\.be\/|embed\/)([^&?/]+)/);
        return videoIdMatch ? `https://www.youtube.com/embed/${videoIdMatch[1]}` : '';
    } else if (props.video.includes('instagram.com')) {
        return `https://www.instagram.com/embed.js`;
    }
    return '';
    });

    const isYouTube = computed(() => props.video.includes('youtube.com') || props.video.includes('youtu.be'));
    const isInstagram = computed(() => props.video.includes('instagram.com'));
</script>

<template>
  <div class="bg-gray-100 rounded-xl shadow-md relative">
    <div class="p-4">
      <div class="aspect-w-16 aspect-h-9">
        <template v-if="isYouTube">
          <iframe
            title="YouTube Video"
            :src="videoEmbedUrl"
            allowfullscreen
            class="w-full h-full rounded-lg"
          ></iframe>
        </template>
        <template v-if="isInstagram">
          <iframe
            title="Instagram Video"
            :src="videoEmbedUrl"
            allowfullscreen
            class="w-full h-full rounded-lg"
          ></iframe>
        </template>
      </div>
    </div>
  </div>
</template>

<style scoped>
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
