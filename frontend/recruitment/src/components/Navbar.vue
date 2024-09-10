<script setup>
  import logoMavi from '@/assets/img/mavi_logo.png';
  import {RouterLink, useRoute } from 'vue-router';
  import { ref } from 'vue';
  const mobileMenuOpen = ref(false);
  const isActiveLink = (routePath) => {
    const route = useRoute();
    return route.path === routePath;
  }
</script>

<template>
  <nav class="bg-white  border-b border-darker-blue">
    <div class="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
      <div class="flex h-20 items-center justify-between">
        <!-- Logo -->
        <div class="flex items-center">
          <RouterLink class="flex flex-shrink-0 items-center mr-4" to="/">
            <img class="h-64 w-auto" :src="logoMavi" alt="Recruitment" /> <!-- Increased logo size -->
          </RouterLink>
        </div>

        <div class="hidden md:flex space-x-2 ml-auto">
          <RouterLink
            to="/"
            :class="[isActiveLink('/') 
              ? 'bg-gray-900' 
              : 'hover:bg-gray-900 hover:text-white', isActiveLink('/') ? 'text-white' : 'text-blue-900', 'px-3', 'py-2', 'rounded-md']"
            >Главная</RouterLink
          >
          <RouterLink
            to="/jobs"
            :class="[isActiveLink('/jobs') 
              ? 'bg-gray-900' 
              : 'hover:bg-gray-900 hover:text-white', isActiveLink('/jobs') ? 'text-white' : 'text-blue-900', 'px-3', 'py-2', 'rounded-md']"
            >Вакансии</RouterLink
          >
          <RouterLink
            to="/jobs/add"
            :class="[isActiveLink('/jobs/add') 
              ? 'bg-gray-900' 
              : 'hover:bg-gray-900 hover:text-white', isActiveLink('/jobs/add') ? 'text-white' : 'text-blue-900', 'px-3', 'py-2', 'rounded-md']"
            >Добавить вакансию</RouterLink
          >
          <RouterLink
            to="/resume/add"
            :class="[isActiveLink('/resume/add') 
              ? 'bg-gray-900' 
              : 'hover:bg-gray-900 hover:text-white', isActiveLink('/resume/add') ? 'text-white' : 'text-blue-900', 'px-3', 'py-2', 'rounded-md']"
            >Добавить резюме</RouterLink
          >
        </div>

        <div class="md:hidden flex items-center">
          <button
            type="button"
            class="inline-flex items-center justify-center p-2 rounded-md text-darker-blu focus:outline-none"
            aria-expanded="false"
            @click="mobileMenuOpen = !mobileMenuOpen"
          >
            <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/>
            </svg>
          </button>
        </div>
      </div>

      <div v-if="mobileMenuOpen" class="md:hidden bg-darker-blue">
        <div class="space-y-1 pt-2 pb-3 sm:px-3">
          <RouterLink
            to="/"
            :class="[isActiveLink('/') 
              ? 'bg-darker-blue' 
              : 'hover:bg-gray-900 hover:text-white', 'block', 'text-white', 'px-3', 'py-2', 'rounded-md']"
            >Главная</RouterLink
          >
          <RouterLink
            to="/jobs"
            :class="[isActiveLink('/jobs') 
              ? 'bg-darker-blue' 
              : 'hover:bg-gray-900 hover:text-white', 'block', 'text-white', 'px-3', 'py-2', 'rounded-md']"
            >Вакансии</RouterLink
          >
          <RouterLink
            to="/jobs/add"
            :class="[isActiveLink('/jobs/add') 
              ? 'bg-darker-blue' 
              : 'hover:bg-gray-900 hover:text-white', 'block', 'text-white', 'px-3', 'py-2', 'rounded-md']"
            >Добавить вакансию</RouterLink
          >
          <RouterLink
            to="/resume/add"
            :class="[isActiveLink('/resume/add') 
              ? 'bg-darker-blue' 
              : 'hover:bg-gray-900 hover:text-white', 'block', 'text-white', 'px-3', 'py-2', 'rounded-md']"
            >Добавить резюме</RouterLink
          >
        </div>
      </div>
    </div>
  </nav>
</template>
