import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '@/views/HomeView.vue';
import JobsView from '@/views/JobsView.vue';
import NotFoundView from '@/views/NotFoundView.vue';
import JobView from '@/views/JobView.vue';
import AddJobView from '@/views/AddJobView.vue';
import AddResumeView from '@/views/AddResumeView.vue';
import AdminView from '@/views/AdminView.vue';

const router = createRouter({
    history: createWebHistory(
        import.meta.env.BASE_URL
    ),
    routes: [{
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/jobs',
            name: 'jobs',
            component: JobsView
        },
        {
            path: '/jobs/:id',
            name: 'job',
            component: JobView
        },
        {
            path: '/jobs/add',
            name: 'add-job',
            component: AddJobView
        }, {
            path: '/:catchAll(.*)',
            name: 'not-found',
            component: NotFoundView
        },
        {
            path: '/resume/add',
            name: 'add-resume',
            component: AddResumeView
        },
        {
            path: '/admin',
            name: 'admin',
            component: AdminView
        },
    ]
});


export default router;