<script setup>
    import JobListing from '@/components/JobListing.vue';

    import { RouterLink } from 'vue-router';

    import {reactive, defineProps, onMounted } from 'vue';
    import axios from 'axios';
    import PulseLoader from 'vue-spinner/src/PulseLoader.vue';

    defineProps({
        limit: Number,
        showButton: {
            type: Boolean,
            default: false,
        }
    });

    const state = reactive({
        jobs: [],
        isLoading: true
    });

    onMounted(async () => {
        try {
            const response = await axios.get('https://mavi-0svz.onrender.com/api/vacancies');
            state.jobs = response.data;
        } catch(e) {
            console.error('error fetching jobs', e);
        } finally {
            state.isLoading = false;
        }
    });
</script>

<template>
    <section class="px-4 py-10 bg-blue-50">
        <div class="m-auto container-xl lg:container">
            <h2 class="mb-6 text-3xl font-bold text-center text-darker-blue">
                Найти работу
            </h2>
            <div v-if="state.isLoading" class="py-6 text-center text-gray-500">
                <PulseLoader/>
            </div>

            <div v-else class="grid grid-cols-1 gap-6 md:grid-cols-3">
                <JobListing
                v-for="job in state.jobs.slice(0, limit || state.jobs.length)" 
                :key="job.id" 
                :job="job"
                />
            </div>
        </div>
    </section>
    
    <section v-if="showButton" class="max-w-lg px-6 m-auto my-10">
      <RouterLink
        to="/jobs"
        class="block px-6 py-4 text-center text-white bg-black rounded-xl hover:bg-gray-700"
        >Просмотреть все вакансии</RouterLink
      >
    </section>
</template>