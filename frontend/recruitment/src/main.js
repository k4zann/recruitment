import './assets/main.css';
import 'primeicons/primeicons.css';
import Toast from 'vue-toastification';
import 'vue-toastification/dist/index.css';

import router from './router';


import { createApp } from 'vue';
import App from './App.vue';

export default function (Vue, { head }) {
    head.script.push({  
        src: 'https://www.instagram.com/embed.js',
        async: true
    })
}
  

const app = createApp(App);

app.use(router);
app.use(Toast);

app.mount('#app');