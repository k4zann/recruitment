<script setup>
    import { RouterLink } from 'vue-router';
</script>

<template>
    <section>
      <div class="container px-6 py-6 m-auto">
        <RouterLink
          to="/jobs"
          class="flex items-center text-darker-blue hover:text-blue-600"
        >
          <i class="mr-2 pi pi-arrow-circle-left"></i> Вернуться к списку
        </RouterLink>
      </div>
    </section>
</template>