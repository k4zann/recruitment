<script setup>
  import Card from '@/components/Card.vue';
  import {RouterLink} from 'vue-router';
</script>

<template>
    <section class="py-4">
      <div class="container-xl lg:container m-auto">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 p-4 rounded-lg">
          <Card>
            <h2 class="text-2xl font-bold">Для работников</h2>
            <p class="mt-2 mb-4">
                Ознакомьтесь с вакансиями и начните свою карьеру уже сегодня
            </p>
            <RouterLink
              to="/resume/add"
              class="inline-block bg-black text-white rounded-lg px-4 py-2 hover:bg-gray-700"
            >
              Добавить резюме
            </RouterLink>
          </Card>
          <Card bg="bg-gray-200">
            <h2 class="text-2xl font-bold">Для работодателей</h2>
            <p class="mt-2 mb-4">
                Разместите объявление о вакансии, чтобы найти идеального кандидата на эту роль
            </p>
            <RouterLink
              to="/jobs/add"
              class="inline-block bg-darker-blue text-white rounded-lg px-4 py-2 hover:bg-gray-600"
            >
              Добавить вакансию
            </RouterLink>
          </Card>
        </div>
      </div>
    </section>
</template>