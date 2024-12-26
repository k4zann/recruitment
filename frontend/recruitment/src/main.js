import './assets/main.css';
import 'primeicons/primeicons.css';
import Toast from 'vue-toastification';
import 'vue-toastification/dist/index.css';

import router from './router';

<script async src="https://www.instagram.com/embed.js"></script>

export default function (Vue, { head }) {
    head.script.push({
      src: 'https://www.instagram.com/embed.js',
      async: true
    })
  
    head.meta.push({
      name: 'keywords',
      content: 'HTML,CSS,XML,JavaScript'
    })
  }
  


import { createApp } from 'vue';
import App from './App.vue';

const app = createApp(App);

app.use(router);
app.use(Toast);

app.mount('#app');