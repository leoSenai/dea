/** @type { import('@storybook/vue3').Preview } */

import 'quasar/dist/quasar.css';
import '../src/style.css'

import { setup } from '@storybook/vue3';
import { Quasar } from 'quasar';

import { colorTheme } from '../src/assets/styles/theme'

setup(app => {
  app.use(Quasar, {
    plugins: { },
    config: {
      brand: { ...colorTheme }
    }
  });
})

const preview = {
  parameters: {
    actions: { argTypesRegex: "^on[A-Z].*" },
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/,
      },
    },
  },
};

export default preview;
