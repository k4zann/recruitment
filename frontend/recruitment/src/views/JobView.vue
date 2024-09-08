<script setup>
    import PulseLoader from 'vue-spinner/src/PulseLoader.vue';
    import { reactive, onMounted, defineProps } from 'vue';
    import { useRoute, RouterLink } from 'vue-router';
    import router from '@/router';
    import axios from 'axios';

    import BackButton from '@/components/BackButton.vue';

    const props = defineProps({
      isAdmin: {
        type: Boolean,
        default: false
      }
    });

    const route = useRoute();
    const jobId = route.params.id;

    const state = reactive({
        job: {},
        isLoading: true,
        contactInformation: {}
    });

    onMounted(async () => {
    try {
        const jobResponse = await axios.get(`http://localhost:8080/api/vacancy/${jobId}`);

        const employerId = jobResponse.data.employer_id._id || jobResponse.data.employer_id.toString();

        const contactResponse = await axios.get(`http://localhost:8080/api/employer/${employerId}`);

        state.job = jobResponse.data;
        state.contactInformation = contactResponse.data;
    } catch (e) {
        console.error('Error fetching job or employer data', e);
    } finally {
        state.isLoading = false;
    }
});
</script>


<template>
  <BackButton/>
  <section v-if="!state.isLoading" class="bg-green-50">
    <div class="container px-6 py-10 m-auto">
      <div class="grid w-full grid-cols-1 gap-6 md:grid-cols-70/30">
        <main>
          <!-- Job Details -->
          <div class="p-6 text-center bg-white rounded-lg shadow-md md:text-left">
            <div class="mb-4 text-gray-500">{{ state.job.work_schedule }}</div>
            <h1 class="mb-4 text-3xl font-bold">{{ state.job.position }}</h1>
            <div class="flex justify-center mb-4 text-gray-500 align-middle md:justify-start">
              <i class="mr-2 text-xl text-orange-700 pi pi-map-marker"></i>
              <p class="text-orange-700">{{state.contactInformation.address}}</p>
            </div>
          </div>

          <div class="p-6 mt-6 bg-white rounded-lg shadow-md">
            <h3 class="mb-6 text-lg font-bold text-green-800">Responsibilities</h3>
            <p class="mb-4">{{ state.job.responsibilities }}</p>

            <h3 class="mb-2 text-lg font-bold text-green-800">Qualification Requirements</h3>
            <ul class="list-disc pl-4 mb-4">
              <li v-for="(req, index) in state.job.qualification_reqs" :key="index">{{ req }}</li>
            </ul>

            <h3 class="mb-2 text-lg font-bold text-green-800">Working Conditions</h3>
            <ul class="list-disc pl-4 mb-4">
              <li>Accessibility: {{ state.job.working_conditions.accessibility }}</li>
              <li>Equipment: {{ state.job.working_conditions.equipment }}</li>
              <li>Special Conditions: {{ state.job.working_conditions.special_conditions }}</li>
              <li>Safety: {{ state.job.working_conditions.safety }}</li>
            </ul>

            <h3 class="mb-2 text-lg font-bold text-green-800">Social Benefits</h3>
            <ul class="list-disc pl-4 mb-4">
              <li>Medical Care: {{ state.job.social_benefits.medical_care }}</li>
              <li>Social Compensation: {{ state.job.social_benefits.social_compensation }}</li>
              <li>Professional Growth: {{ state.job.social_benefits.professional_growth }}</li>
            </ul>

            <h3 class="mb-2 text-lg font-bold text-green-800">Employee Expectations</h3>
            <ul class="list-disc pl-4 mb-4">
              <li>Team Interaction: {{ state.job.employee_expectations.team_interaction }}</li>
              <li>Quality and Productivity: {{ state.job.employee_expectations.quality_and_productivity }}</li>
            </ul>
          </div>
        </main>

        <!-- Sidebar -->
        <aside>
          <!-- Company Info -->
          <div class="p-6 bg-white rounded-lg shadow-md">
            <h3 class="mb-6 text-xl font-bold">Company Information</h3>
            <h2 class="text-2xl">{{ state.contactInformation.company_name }}</h2>
            <p class="my-2">{{ state.contactInformation.sphere }}</p>

            <hr class="my-4" />

            <h3 class="text-xl">Email:</h3>
            <p class="p-2 my-2 font-bold bg-green-100">{{ state.contactInformation.email }}</p>

            <h3 class="text-xl">Phone Number:</h3>
            <p class="p-2 my-2 font-bold bg-green-100">{{ state.contactInformation.telephone }}</p>

            <h3 class="text-xl">Number of Employees:</h3>
            <p class="p-2 my-2 font-bold bg-green-100">{{ state.contactInformation.number_of_employees }}</p>
          </div>

          <!-- Manage -->
          <div v-if="isAdmin" class="p-6 mt-6 bg-white rounded-lg shadow-md">
            <h3 class="mb-6 text-xl font-bold">Manage Job</h3>
            <RouterLink
              :to="`/jobs/edit/${state.job.id}`"
              class="block w-full px-4 py-2 mt-4 font-bold text-center text-white bg-green-500 rounded-full hover:bg-green-600 focus:outline-none focus:shadow-outline"
              >Edit Job</RouterLink
            >
            <button
              class="block w-full px-4 py-2 mt-4 font-bold text-white bg-red-500 rounded-full hover:bg-red-600 focus:outline-none focus:shadow-outline"
            >
              Delete Job
            </button>
          </div>
        </aside>
      </div>
    </div>
  </section>

  <div v-else class="py-6 text-center text-gray-500">
      <PulseLoader/>
  </div>
</template>
