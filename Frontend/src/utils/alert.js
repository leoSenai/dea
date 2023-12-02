import Toastify from 'toastify-js';

export function alerta(aviso, valid){
    if(valid){
        Toastify({
            avatar: '/x-circle-fill.svg',
            text: aviso,
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
              y: 65,
            },
          }).showToast();
    }else{
        Toastify({
            avatar: '/x-circle-fill.svg',
            text: aviso,
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
              y: 65,
            },
          }).showToast();
    }
}