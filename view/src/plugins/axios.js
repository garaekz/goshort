import axios from 'axios';
import store from '@/store';
import router from '@/router';
import Swal from 'sweetalert2';

// Request interceptor
/* eslint-disable no-param-reassign */
axios.interceptors.request.use((request) => {
  const token = store.getters['auth/token'];
  if (token) {
    request.headers.common.Authorization = `Bearer ${token}`;
  }

  const locale = store.getters['lang/locale'];
  if (locale) {
    request.headers.common['Accept-Language'] = locale;
  }

  // request.headers['X-Socket-Id'] = Echo.socketId()

  return request;
});
/* eslint-enable no-param-reassign */

// Response interceptor
axios.interceptors.response.use((response) => response, (error) => {
  const { status } = error.response;

  if (status >= 500) {
    Swal.fire({
      icon: 'error',
      title: 'Ha ocurrido un problema',
      text: '¡Algo salió mal! Inténtalo de nuevo.',
      reverseButtons: true,
      confirmButtonText: 'Ok',
      cancelButtonText: 'Cancel',
    });
  }

  if (status === 401 && store.getters['auth/check']) {
    Swal.fire({
      icon: 'warning',
      title: '!Sesión Expirada!',
      text: 'Por favor inicie sesión de nuevo para continuar.',
      reverseButtons: true,
      confirmButtonText: 'Ok',
      cancelButtonText: 'Cancel',
    }).then(() => {
      store.commit('auth/LOGOUT');

      router.push({ name: 'login' });
    });
  }

  return Promise.reject(error);
});
