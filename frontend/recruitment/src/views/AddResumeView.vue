<script setup>
    import { reactive } from 'vue';
    import axios from 'axios';
    import router from '@/router';
    import { useToast } from 'vue-toastification';

    const state = reactive({
        employeeId: '',
        isLoading: false
    });

    state.employeeId = localStorage.getItem('employeeId');

    const contactInformationForm = reactive({
        first_name: '',
        last_name: '',
        birth_date: '',
        telephone: '',
        email: '',
        address: '',
    });

    const form = reactive({
        employee_id: '',
        education: {
            institution_name: '',
            location: '',
            specialization: '',
            years_attended: ''
        },
        skills: {
            professional_skills: '',
            personal_qualities: '',
            languages: [
                { language: '', level: '' }
            ]
        },
        experience: [{
            company_name: '',
            position: '',
            work_period: '',
            responsibilities: ''
        }],
        job_expectation: {
            desired_position: '',
            expected_salary: ''
        },
        disability_info: {
            disability_group: '',
            disability_details: '',
            work_conditions: '',
            medical_restrictions: ''
        },
        additional_info: {
            relocation_willingness: false,
            travel_willingness: false,
            hobbies: ''
        }
    });



    const toast = useToast();

    const handleSubmit = async () => {
        const newResume = {
            employee_id: state.employeeId,
            education: {
                institution_name: form.education.institution_name,
                location: form.education.location,
                specialization: form.education.specialization,
                years_attended: form.education.years_attended
            },
            skills: [{
                professional_skills: form.skills.professional_skills.split(',').map(item => item.trim()), 
                personal_qualities: form.skills.personal_qualities.split(',').map(item => item.trim()),
                languages: form.skills.languages
            }],
            experience: form.experience.map(exp => ({
                company_name: exp.company_name,
                position: exp.position,
                work_period: exp.work_period,
                responsibilities: exp.responsibilities
            })),
            job_expectation: {
                desired_position: form.job_expectation.desired_position,
                expected_salary: form.job_expectation.expected_salary
            },
            disability_info: {
                disability_group: form.disability_info.disability_group,
                disability_details: form.disability_info.disability_details,
                work_conditions: form.disability_info.work_conditions,
                medical_restrictions: form.disability_info.medical_restrictions
            },
            additional_info: {
                relocation_willingness: form.additional_info.relocation_willingness,
                travel_willingness: form.additional_info.travel_willingness,
                hobbies: form.additional_info.hobbies
            }
        };

        try {
            state.isLoading = true;
            await axios.post('http://localhost:8080/api/resume', newResume);
            toast.success('Резюме было успешно добавлено');
            router.push(`/`);
        } catch (e) {
            console.error('Ошибка при добавлении резюме', e);
            toast.error('Произошла ошибка');
        } finally {
            state.isLoading = false;
        }
    };

    const addLanguage = () => {
        form.skills.languages.push({ language: '', level: '' });
    };

    const removeLanguage = (index) => {
        form.skills.languages.splice(index, 1);
    };

    const addExperience = () => {
        form.experience.push({
            company_name: '',
            position: '',
            work_period: '',
            responsibilities: ''
        });
    };

    const removeExperience = (index) => {
        form.experience.splice(index, 1);
    };

    const submitContactInformation = async () => {
        try {
            const response = await axios.post('http://localhost:8080/api/employee', contactInformationForm);
            state.employeeId = response.data.id;
            localStorage.setItem('employeeId', state.employeeId);
            toast.success('Контактная информация была успешно добавлена');
        } catch (e) {
            console.error('Error submitting contact information', e);
            toast.error('Произошла ошибка');
        } finally {
            state.isLoading = false;
        }
    }; 
</script>

<template>
    <section class="bg-green-50">
      <div class="container max-w-2xl py-24 m-auto">
        <div class="px-6 py-8 m-4 mb-4 bg-white border rounded-md shadow-md md:m-0">
          <form v-if="state.employeeId === '' || state.employeeId === null">
            <h2 class="mb-6 text-3xl font-semibold text-center">Добавить контактную информацию</h2>
            <div class="mb-4">
              <label for="first_name" class="block mb-2 font-bold text-gray-700">Имя</label>
              <input v-model="contactInformationForm.first_name" type="text" id="first_name" class="w-full px-3 py-2 border rounded" placeholder="Введите ваше имя" />
            </div>
            <div class="mb-4">
              <label for="last_name" class="block mb-2 font-bold text-gray-700">Фамилия</label>
              <input v-model="contactInformationForm.last_name" type="text" id="last_name" class="w-full px-3 py-2 border rounded" placeholder="Введите вашу фамилию" />
            </div>
            <div class="mb-4">
              <label for="birth_date" class="block mb-2 font-bold text-gray-700">Дата рождения</label>
              <input v-model="contactInformationForm.birth_date" type="date" id="birth_date" class="w-full px-3 py-2 border rounded" />
            </div>
            <div class="mb-4">
              <label for="telephone" class="block mb-2 font-bold text-gray-700">Телефон</label>
              <input v-model="contactInformationForm.telephone" type="text" id="telephone" class="w-full px-3 py-2 border rounded" placeholder="Введите ваш номер телефона" />
            </div>
            <div class="mb-4">
              <label for="email" class="block mb-2 font-bold text-gray-700">Email</label>
              <input v-model="contactInformationForm.email" type="email" id="email" class="w-full px-3 py-2 border rounded" placeholder="Введите ваш email" />
            </div>
            <div class="mb-4">
              <label for="address" class="block mb-2 font-bold text-gray-700">Адрес</label>
              <input v-model="contactInformationForm.address" type="text" id="address" class="w-full px-3 py-2 border rounded" placeholder="Введите ваш адрес" />
            </div>
            <div>
              <button
                @click.prevent="submitContactInformation"
                class="w-full px-4 py-2 font-bold text-white bg-green-500 rounded-full hover:bg-green-600 focus:outline-none focus:shadow-outline"
              >
                Добавить контактную информацию
              </button>
            </div>
          </form>

          <form v-else @submit.prevent="handleSubmit">
            <h2 class="mb-6 text-3xl font-semibold text-center">Добавить резюме</h2>

            <!-- Education Section -->
            <h3 class="mb-4 text-2xl font-semibold">Образование</h3>
            <div class="mb-4">
              <label for="institution_name" class="block mb-2 font-bold text-gray-700">Название учебного заведения</label>
              <input v-model="form.education.institution_name" type="text" id="institution_name" class="w-full px-3 py-2 border rounded" placeholder="Название университета, школы и т.д." />
            </div>
            <div class="mb-4">
              <label for="location" class="block mb-2 font-bold text-gray-700">Местоположение</label>
              <input v-model="form.education.location" type="text" id="location" class="w-full px-3 py-2 border rounded" placeholder="Город, страна" />
            </div>
            <div class="mb-4">
              <label for="specialization" class="block mb-2 font-bold text-gray-700">Специализация</label>
              <input v-model="form.education.specialization" type="text" id="specialization" class="w-full px-3 py-2 border rounded" placeholder="Укажите специализацию" />
            </div>
            <div class="mb-4">
              <label for="years_attended" class="block mb-2 font-bold text-gray-700">Годы обучения</label>
              <input v-model="form.education.years_attended" type="text" id="years_attended" class="w-full px-3 py-2 border rounded" placeholder="Пример: 2010-2014" />
            </div>

            <!-- Skills Section -->
            <h3 class="mb-4 text-2xl font-semibold">Навыки</h3>
            <div class="mb-4">
              <label for="professional_skills" class="block mb-2 font-bold text-gray-700">Профессиональные навыки</label>
              <input v-model="form.skills.professional_skills" type="text" id="professional_skills" class="w-full px-3 py-2 border rounded" placeholder="Укажите навыки через запятую" />
            </div>
            <div class="mb-4">
              <label for="personal_qualities" class="block mb-2 font-bold text-gray-700">Личные качества</label>
              <input v-model="form.skills.personal_qualities" type="text" id="personal_qualities" class="w-full px-3 py-2 border rounded" placeholder="Например: коммуникабельность, ответственность" />
            </div>
            <div class="mb-4">
                <label for="languages" class="block mb-2 font-bold text-gray-700">Языки</label>
                <div v-for="(language, index) in form.skills.languages" :key="index" class="flex items-center mb-4">
                    <input v-model="language.language" type="text" class="w-1/2 px-3 py-2 border rounded" placeholder="Язык" />
                    <input v-model="language.level" type="text" class="w-1/2 px-3 py-2 border rounded" placeholder="Уровень" />
                    <button @click.prevent="removeLanguage(index)" class="ml-2 text-red-500">Удалить</button>
                </div>
                <button @click.prevent="addLanguage" class="text-green-500">Добавить язык</button>
            </div>

            <!-- Experience Section -->
            <h3 class="mb-4 text-2xl font-semibold">Опыт работы</h3>
            <div v-for="(exp, index) in form.experience" :key="index" class="mb-4">
            <label for="company_name" class="block mb-2 font-bold text-gray-700">Название компании</label>
            <input v-model="exp.company_name" type="text" id="company_name" class="w-full px-3 py-2 border rounded" placeholder="Укажите название компании" />

            <label for="position" class="block mt-4 mb-2 font-bold text-gray-700">Должность</label>
            <input v-model="exp.position" type="text" id="position" class="w-full px-3 py-2 border rounded" placeholder="Укажите вашу должность" />

            <label for="work_period" class="block mt-4 mb-2 font-bold text-gray-700">Период работы</label>
            <input v-model="exp.work_period" type="text" id="work_period" class="w-full px-3 py-2 border rounded" placeholder="Пример: 2010-2014" />

            <label for="responsibilities" class="block mt-4 mb-2 font-bold text-gray-700">Обязанности</label>
            <textarea v-model="exp.responsibilities" id="responsibilities" class="w-full px-3 py-2 border rounded" rows="4" placeholder="Опишите ваши обязанности"></textarea>

            <button @click.prevent="removeExperience(index)" class="mt-2 text-red-500">
                <i class="pi pi-minus mr-1"></i> Удалить
            </button>
            </div>

            <button @click.prevent="addExperience" class="mt-4 text-blue-500 flex items-center">
            <i class="pi pi-plus mr-1"></i> Добавить опыт работы
            </button>

            <!-- Job Expectation Section -->
            <h3 class="mb-4 text-2xl font-semibold">Ожидания от работы</h3>
            <div class="mb-4">
              <label for="desired_position" class="block mb-2 font-bold text-gray-700">Желаемая должность</label>
              <input v-model="form.job_expectation.desired_position" type="text" id="desired_position" class="w-full px-3 py-2 border rounded" placeholder="Укажите должность, на которую вы претендуете" />
            </div>
            <div class="mb-4">
              <label for="expected_salary" class="block mb-2 font-bold text-gray-700">Ожидаемая зарплата</label>
              <input v-model="form.job_expectation.expected_salary" type="text" id="expected_salary" class="w-full px-3 py-2 border rounded" placeholder="Укажите зарплату, которую вы ожидаете" />
            </div>

            <!-- Disability Information Section -->
            <h3 class="mb-4 text-2xl font-semibold">Информация об инвалидности</h3>
            <div class="mb-4">
              <label for="disability_group" class="block mb-2 font-bold text-gray-700">Группа инвалидности</label>
              <input v-model="form.disability_info.disability_group" type="text" id="disability_group" class="w-full px-3 py-2 border rounded" placeholder="Укажите группу инвалидности" />
            </div>
            <div class="mb-4">
              <label for="disability_details" class="block mb-2 font-bold text-gray-700">Детали инвалидности</label>
              <input v-model="form.disability_info.disability_details" type="text" id="disability_details" class="w-full px-3 py-2 border rounded" placeholder="Опишите вашу инвалидность" />
            </div>
            <div class="mb-4">
              <label for="work_conditions" class="block mb-2 font-bold text-gray-700">Условия труда</label>
              <input v-model="form.disability_info.work_conditions" type="text" id="work_conditions" class="w-full px-3 py-2 border rounded" placeholder="Опишите необходимые условия труда" />
            </div>
            <div class="mb-4">
              <label for="medical_restrictions" class="block mb-2 font-bold text-gray-700">Медицинские ограничения</label>
              <input v-model="form.disability_info.medical_restrictions" type="text" id="medical_restrictions" class="w-full px-3 py-2 border rounded" placeholder="Укажите медицинские ограничения" />
            </div>

            <!-- Additional Information Section -->
            <h3 class="mb-4 text-2xl font-semibold">Дополнительная информация</h3>
            <div class="mb-4">
              <label for="relocation_willingness" class="block mb-2 font-bold text-gray-700">Готовность к переезду</label>
              <input v-model="form.additional_info.relocation_willingness" type="checkbox" id="relocation_willingness" class="mr-2 leading-tight" />
              <label for="relocation_willingness">Я готов(а) к переезду</label>
            </div>
            <div class="mb-4">
              <label for="travel_willingness" class="block mb-2 font-bold text-gray-700">Готовность к командировкам</label>
              <input v-model="form.additional_info.travel_willingness" type="checkbox" id="travel_willingness" class="mr-2 leading-tight" />
              <label for="travel_willingness">Я готов(а) к командировкам</label>
            </div>
            <div class="mb-4">
              <label for="hobbies" class="block mb-2 font-bold text-gray-700">Хобби</label>
              <textarea v-model="form.additional_info.hobbies" id="hobbies" class="w-full px-3 py-2 border rounded" rows="4" placeholder="Укажите ваши увлечения и хобби"></textarea>
            </div>

            <!-- Submit Button -->
            <div>
              <button
                class="w-full px-4 py-2 font-bold text-white bg-green-500 rounded-full hover:bg-green-600 focus:outline-none focus:shadow-outline"
                type="submit"
              >
                Добавить резюме
              </button>
            </div>
          </form>
        </div>
      </div>
    </section>
</template>
