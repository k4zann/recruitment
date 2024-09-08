<script setup>
  import { reactive } from 'vue';
  import axios from 'axios';
  import router from '@/router';
  import { useToast } from 'vue-toastification';
  
  const state = reactive({
    isLoading: false,
    employerId: '',
  });

  state.employerId = localStorage.getItem('employerId');
  console.log('Employer ID:', state.employerId);

  const contactInformationForm = reactive({
    company_name: '',
    sphere: '',
    address: '',
    telephone: '',
    email: '',
    number_of_employees: 0
  });

  const form = reactive({
    employer_id: state.employerId,
    position: '',
    responsibilities: '',
    qualification_reqs: [],
    work_schedule: '',
    working_conditions: {
      accessibility: '',
      equipment: '',
      special_conditions: '',
      safety: ''
    },
    social_benefits: {
      medical_care: '',
      social_compensation: '',
      professional_growth: ''
    },
    employee_expectations: {
      team_interaction: '',
      quality_and_productivity: ''
    }
  });

  const toast = useToast();

  const handleSubmit = async () => {
    const newJob = {
      employer_id: state.employerId,
      position: form.position,
      responsibilities: form.responsibilities, // Convert comma-separated text to an array
      qualification_reqs: form.qualification_reqs.split(',').map(item => item.trim()),
      work_schedule: form.work_schedule,
      working_conditions: {
        accessibility: form.working_conditions.accessibility,
        equipment: form.working_conditions.equipment,
        special_conditions: form.working_conditions.special_conditions,
        safety: form.working_conditions.safety
      },
      social_benefits: {
        medical_care: form.social_benefits.medical_care,
        social_compensation: form.social_benefits.social_compensation,
        professional_growth: form.social_benefits.professional_growth
      },
      employee_expectations: {
        team_interaction: form.employee_expectations.team_interaction,
        quality_and_productivity: form.employee_expectations.quality_and_productivity
      }
    };

    try {
      state.isLoading = true;
      const response = await axios.post('http://localhost:8080/api/vacancy', newJob);
      toast.success('Вакансия добавлена');
      router.push(`/jobs/${response.data.id}`);
    } catch (e) {
      console.error('Ошибка при добавлении вакансии', e);
      toast.error('Произошла ошибка');
    } finally {
      state.isLoading = false;
      router.push('/jobs');
    }
  };

  // submit employer
  const submitEmployer = async () => {
    const newEmployer = {
      company_name: contactInformationForm.company_name,
      sphere: contactInformationForm.sphere,
      address: contactInformationForm.address,
      telephone: contactInformationForm.telephone,
      email: contactInformationForm.email,
      number_of_employees: contactInformationForm.number_of_employees
    };
    
    try {
      state.isLoading = true;
      const response = await axios.post('http://localhost:8080/api/employer', newEmployer);
      state.employerId = response.data.id;
      localStorage.setItem('employerId', response.data.id);
      toast.success('Контактная информация добавлена');
    } catch (e) {
      console.error('Ошибка при добавлении контакной информации', e);
      toast.error('Произошла ошибка');
    } finally {
      console.log('Employer ID:', form.employer_id);
      state.isLoading = false;
    }
  };
</script>


<template>
  <section class="bg-green-50">
    <div class="container max-w-2xl py-24 m-auto">
      <div class="px-6 py-8 m-4 mb-4 bg-white border rounded-md shadow-md md:m-0">
        <form v-if="state.employerId === '' || state.employerId === null" @submit.prevent="submitEmployer">
          <h2 class="mb-6 text-3xl font-semibold text-center">Добавить контактную информацию</h2>
          <!-- Company Name -->
          <div class="mb-4">
            <label for="company_name" class="block mb-2 font-bold text-gray-700">Название компании</label>
            <input v-model="contactInformationForm.company_name" type="text" id="company_name" class="w-full px-3 py-2 border rounded" placeholder="Укажите название компании" required />
            <p class="text-sm text-gray-500">Пример: ООО "Рога и копыта"</p>
          </div>

          <!-- Sphere -->
          <div class="mb-4">
            <label for="sphere" class="block mb-2 font-bold text-gray-700">Сфера деятельности</label>
            <input v-model="contactInformationForm.sphere" type="text" id="sphere" class="w-full px-3 py-2 border rounded" placeholder="Укажите сферу деятельности" required />
            <p class="text-sm text-gray-500">Пример: IT, медицина, образование</p>
          </div>

          <!-- Address -->
          <div class="mb-4">
            <label for="address" class="block mb-2 font-bold text-gray-700">Адрес</label>
            <input v-model="contactInformationForm.address" type="text" id="address" class="w-full px-3 py-2 border rounded" placeholder="Укажите адрес" required />
            <p class="text-sm text-gray-500">Пример: г. Москва, ул. Ленина, д. 1</p>
          </div>

          <!-- Telephone -->
          <div class="mb-4">
            <label for="telephone" class="block mb-2 font-bold text-gray-700">Телефон</label>
            <input v-model="contactInformationForm.telephone" type="tel" id="telephone" class="w-full px-3 py-2 border rounded" placeholder="Укажите телефон" required />
            <p class="text-sm text-gray-500">Пример: +7 (999) 123-45-67</p>
          </div>

          <!-- Email -->
          <div class="mb-4">
            <label for="email" class="block mb-2 font-bold text-gray-700">Email</label>
            <input v-model="contactInformationForm.email" type="email" id="email" class="w-full px-3 py-2 border rounded" placeholder="Укажите email" required />
            <p class="text-sm text-gray-500">Пример: name@gmail.com</p>
          </div>

          <!-- Number of Employees -->
          <div class="mb-4">
            <label for="number_of_employees" class="block mb-2 font-bold text-gray-700">Количество сотрудников</label>
            <input v-model="contactInformationForm.number_of_employees" type="number" id="number_of_employees" class="w-full px-3 py-2 border rounded" placeholder="Укажите количество сотрудников" required />
            <p class="text-sm text-gray-500">Пример: 100</p>
          </div>

          <!-- Submit Button -->
          <div>
            <button class="w-full px-4 py-2 font-bold text-white bg-green-500 rounded-full hover:bg-green-600 focus:outline-none focus:shadow-outline" type="submit">
              Добавить контактную информацию
            </button>
          </div>
        </form>

        <form v-else @submit.prevent="handleSubmit">
          <h2 class="mb-6 text-3xl font-semibold text-center">Добавить вакансию</h2>
          <!-- Position -->
          <div class="mb-4">
            <label for="position" class="block mb-2 font-bold text-gray-700">Должность</label>
            <input v-model="form.position" type="text" id="position" class="w-full px-3 py-2 border rounded" placeholder="Укажите должность" required />
            <p class="text-sm text-gray-500">Пример: Программист</p>
          </div>

          <!-- Responsibilities -->
          <div class="mb-4">
            <label for="responsibilities" class="block mb-2 font-bold text-gray-700">Обязанности</label>
            <textarea v-model="form.responsibilities" id="responsibilities" class="w-full px-3 py-2 border rounded" rows="4" placeholder="Перечислите обязанности через запятую" required></textarea>
            <p class="text-sm text-gray-500">Пример: Разработка ПО, тестирование, поддержка пользователей</p>
          </div>

          <!-- Qualification Requirements -->
          <div class="mb-4">
            <label for="qualification_reqs" class="block mb-2 font-bold text-gray-700">Требования к квалификации</label>
            <textarea v-model="form.qualification_reqs" id="qualification_reqs" class="w-full px-3 py-2 border rounded" rows="4" placeholder="Перечислите требования через запятую" required></textarea>
            <p class="text-sm text-gray-500">Пример: Опыт работы 3 года, знание JavaScript</p>
          </div>

          <!-- Work Schedule (Switch to select for predefined schedules) -->
          <div class="mb-4">
            <label for="work_schedule" class="block mb-2 font-bold text-gray-700">График работы</label>
            <select v-model="form.work_schedule" id="work_schedule" class="w-full px-3 py-2 border rounded" required>
              <option disabled value="">Выберите график работы</option>
              <option value="Полный день">Полный день</option>
              <option value="Частичная занятость">Частичная занятость</option>
              <option value="Сменный график">Сменный график</option>
            </select>
          </div>

          <!-- Working Conditions -->
          <h3 class="mb-4 text-2xl font-semibold">Условия труда</h3>
          
          <!-- Accessibility (Switch to a checkbox for availability) -->
          <div class="mb-4">
            <label for="accessibility" class="block mb-2 font-bold text-gray-700">Доступность</label>
            <select v-model="form.working_conditions.accessibility" id="accessibility" class="w-full px-3 py-2 border rounded">
              <option disabled value="">Выберите доступность</option>
              <option value="Да">Да</option>
              <option value="Нет">Нет</option>
            </select>
          </div>

          <!-- Equipment (Text input for specifics) -->
          <div class="mb-4">
            <label for="equipment" class="block mb-2 font-bold text-gray-700">Оборудование</label>
            <input v-model="form.working_conditions.equipment" type="text" id="equipment" class="w-full px-3 py-2 border rounded" placeholder="Перечислите оборудование" />
          </div>

          <!-- Special Conditions -->
          <div class="mb-4">
            <label for="special_conditions" class="block mb-2 font-bold text-gray-700">Особые условия</label>
            <input v-model="form.working_conditions.special_conditions" type="text" id="special_conditions" class="w-full px-3 py-2 border rounded" placeholder="Укажите особые условия" />
          </div>

          <!-- Safety -->
          <div class="mb-4">
            <label for="safety" class="block mb-2 font-bold text-gray-700">Безопасность</label>
            <input v-model="form.working_conditions.safety" type="text" id="safety" class="w-full px-3 py-2 border rounded" placeholder="Перечислите меры безопасности" />
          </div>

          <!-- Social Benefits -->
          <h3 class="mb-4 text-2xl font-semibold">Социальные льготы</h3>
          
          <!-- Medical Care -->
          <div class="mb-4">
            <label for="medical_care" class="block mb-2 font-bold text-gray-700">Медицинское обслуживание</label>
            <input v-model="form.social_benefits.medical_care" type="text" id="medical_care" class="w-full px-3 py-2 border rounded" placeholder="Перечислите медицинские льготы" />
          </div>

          <!-- Social Compensation -->
          <div class="mb-4">
            <label for="social_compensation" class="block mb-2 font-bold text-gray-700">Социальная компенсация</label>
            <input v-model="form.social_benefits.social_compensation" type="text" id="social_compensation" class="w-full px-3 py-2 border rounded" placeholder="Опишите компенсации" />
          </div>

          <!-- Professional Growth -->
          <div class="mb-4">
            <label for="professional_growth" class="block mb-2 font-bold text-gray-700">Профессиональный рост</label>
            <input v-model="form.social_benefits.professional_growth" type="text" id="professional_growth" class="w-full px-3 py-2 border rounded" placeholder="Возможности для карьерного роста" />
          </div>

          <!-- Employee Expectations -->
          <h3 class="mb-4 text-2xl font-semibold">Ожидания от сотрудника</h3>
          
          <!-- Team Interaction -->
          <div class="mb-4">
            <label for="team_interaction" class="block mb-2 font-bold text-gray-700">Взаимодействие с командой</label>
            <input v-model="form.employee_expectations.team_interaction" type="text" id="team_interaction" class="w-full px-3 py-2 border rounded" placeholder="Ожидания по взаимодействию с командой" />
          </div>

          <!-- Quality and Productivity -->
          <div class="mb-4">
            <label for="quality_and_productivity" class="block mb-2 font-bold text-gray-700">Качество и производительность</label>
            <input v-model="form.employee_expectations.quality_and_productivity" type="text" id="quality_and_productivity" class="w-full px-3 py-2 border rounded" placeholder="Ожидания по производительности" />
          </div>

          <!-- Submit Button -->
          <div>
            <button class="w-full px-4 py-2 font-bold text-white bg-green-500 rounded-full hover:bg-green-600 focus:outline-none focus:shadow-outline" type="submit">
              Добавить вакансию
            </button>
          </div>
        </form>
      </div>
    </div>
  </section>
</template>


