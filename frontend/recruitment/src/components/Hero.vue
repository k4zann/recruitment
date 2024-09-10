<script setup>
    import { defineProps } from 'vue';

    defineProps({
        title: {
            type: String,
            default: '"Работа для каждого"' 
        },
        subtitle: {
            type: String,
            default: 'Найди работу вместе с нами',
        }
    });
</script>

<template>
  <section class="bg-darker-blue py-20 mb-4">
    <div
      class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex flex-col items-center"
    >
      <div class="text-center">
        <h1
          class="text-3xl sm:text-4xl md:text-5xl font-extrabold text-white"
        >
          {{ title }}
        </h1>
        <p class="my-4 text-lg sm:text-xl text-white">
          {{ subtitle }}
        </p>
      </div>
    </div>
  </section>
</template>

