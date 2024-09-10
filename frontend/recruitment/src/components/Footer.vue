<script setup>
    import { RouterLink } from 'vue-router';
</script>

<template>
    <div class="bg-blue-900 py-4">
        <div class="max-w-screen-lg px-4 sm:px-6 text-gray-800 sm:grid md:grid-cols-4 sm:grid-cols-2 mx-auto">
            <div class="p-5">
                <h3 class="font-bold text-xl text-white">Mavi Jobs</h3>
            </div>
            <div class="p-5">
                <div class="text-sm uppercase text-white font-bold">Для работадателей</div>
                <RouterLink class="my-3 block text-gray-100" to="/jobs/add">Добавить вакансию <span class="text-white text-xs p-1"></span></RouterLink>
            </div>
            <div class="p-5">
                <div class="text-sm uppercase text-white font-bold">Для работников</div>
                <RouterLink class="my-3 block text-gray-100" to="/resume/add">Добавить вакансию <span class="text-white text-xs p-1"></span></RouterLink>
                <RouterLink class="my-3 block text-gray-100" to="/jobs">Вакансии <span class="text-white text-xs p-1"></span></RouterLink>
            </div>
            <div class="p-5">
                <div class="text-sm uppercase text-white font-bold">Связитесь с нами</div>
                <div class="my-3 block text-gray-100" to="/">Астана, Казахстан
                    <span class="text-white text-xs p-1"></span></div><div class="my-3 block text-gray-100" to="/">mavi-jobs@gmail.com
                    <span class="text-white text-xs p-1"></span></div>
            </div>
         </div>
    </div>

    <div class="bg-blue-900 pt-2">
        <div class="flex pb-5 px-3 m-auto pt-5 border-t text-gray-800 text-sm flex-col
        max-w-screen-lg items-center">
            <div class="md:flex-auto md:flex-row-reverse mt-2 flex-row flex">
                
            </div>
            <div class="my-5 text-white">© Copyright 2024. All Rights Reserved.</div>
        </div>
    </div>

</template>