import axios from 'axios';
import Toastify from 'toastify-js';

axios.defaults.baseURL = import.meta.env.VITE_API_URL_DEV

axios.interceptors.response.use((response) => {
  if (
    response.status >= 200 &&
    response.status < 300 &&
    response.config.method !== 'get' &&
    !response.config.url.includes('question')
  ) {
    Toastify({
      avatar: '/check-circle-fill.svg',
      text: response.data.message,
      duration: 3000,
      gravity: 'top',
      position: 'right',
      style: {
        background:
          'linear-gradient(90deg, var(--primary-600) 0%, var(--primary) 100%)',
        color: 'white',
        boxShadow:
          '0px 0px 5px -16px var(--primary-600), 5px 5px 36px -9px var(--primary)',
        display: 'flex',
        alignItems: 'center',
        gap: '.25rem',
      },
      offset: {
        x: 0,
        y: location.href.includes('login') ? 0 : 65,
      },
    }).showToast();
  }
  return response;
}, ({ response }) => {
  Toastify({
    avatar: '/x-circle-fill.svg',
    text: response.data.message,
    duration: 3000,
    gravity: 'top',
    position: 'right',
    style: {
      background:
        'linear-gradient(90deg, var(--others-red-600) 0%, var(--others-red-300) 100%)',
      color: 'white',
      boxShadow:
        '0px 0px 5px -16px var(--others-red-600), 5px 5px 36px -9px var(--others-red-300)',
      display: 'flex',
      alignItems: 'center',
      gap: '.25rem',
    },
    offset: {
      x: 0,
      y: location.href.includes('login') ? 0 : 65,
    },
  }).showToast();

  return response;
});

export default axios;
