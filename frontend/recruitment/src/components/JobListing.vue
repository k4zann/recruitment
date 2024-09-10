<script setup>
    import { RouterLink } from 'vue-router';
    import { defineProps, ref, computed } from 'vue';

    const props = defineProps({
        job: Object
    });

    const showFullDescription = ref(false);

    const toggleDescription = () => {
        showFullDescription.value = !showFullDescription.value;
    };

    const truncatedDescription = computed(() => {
        let description = props.job.responsibilities;
        if (!showFullDescription.value) {
            description = description.substring(0, 90) + '...';
        }
        return description;
    });
</script>

<template>
    <div class="bg-white rounded-xl shadow-md relative">
            <div class="p-4">
              <div class="mb-6">
                <div class="text-gray-600 my-2">{{ job.work_schedule }}</div>
                <h3 class="text-xl font-bold">{{ job.position }}</h3>
              </div>

              <div class="mb-2">
                <div>
                    {{ truncatedDescription }}                    
                </div>
                <button @click="toggleDescription" class="text-blue-500 hover:text-blue-600 mb-5">
                    {{ showFullDescription ? 'Меньше' : 'Еще' }}
                </button>
              </div>

              <div class="border border-gray-100 mb-5"></div>

              <div class="flex flex-col lg:flex-row justify-between mb-4">

                <RouterLink
                  :to="'/jobs/' + job.id"
                  class="h-[36px] bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-lg text-center text-sm"
                >
                  Подробнее
                </RouterLink>
              </div>
            </div>
          </div>
</template>