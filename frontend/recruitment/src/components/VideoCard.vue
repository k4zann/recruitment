<script setup>
import { onMounted, ref, computed } from 'vue';

const props = defineProps({
    video: String,
});

// Computed property to determine if the URL is a valid Instagram post
const isInstagram = computed(() => props.video.includes('instagram.com'));

// Function to load Instagram embed script
const loadInstagramEmbedScript = () => {
    const script = document.createElement('script');
    script.src = 'https://www.instagram.com/embed.js';
    script.async = true;
    document.head.appendChild(script);
};

// Function to process Instagram embeds
const processInstagramEmbed = () => {
    if (window.instgrm) {
        window.instgrm.Embeds.process();
    }
};

// Load the script and process embeds when the component mounts
onMounted(() => {
    if (isInstagram.value) {
        loadInstagramEmbedScript();
        // Delay processing to ensure script is loaded
        setTimeout(processInstagramEmbed, 1000); // Adjust timeout as necessary
    }
});
</script>

<template>
  <div class="bg-gray-100 rounded-xl shadow-md relative">
    <div class="p-4">
      <div class="aspect-w-16 aspect-h-9">
        <template v-if="isInstagram">
          <blockquote class="instagram-media rounded-lg" :data-instgrm-permalink="video" style="margin: 1px auto;max-width: 500;">
            <a :href="video">View this post on Instagram</a>
          </blockquote>
        </template>
      </div>
    </div>
  </div>
</template>

<style scoped>
.bg-gray-100 {
  display: flex; /* Enable Flexbox */
  justify-content: center; /* Center horizontally */
  align-items: center; /* Center vertically */
}
.aspect-w-16.aspect-h-9 {
  position: relative;
  width: 100%;
  height: 100%;
}

.instagram-media {
  position: relative;
}
</style>
