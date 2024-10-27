const { addIconSelectors } = require('@iconify/tailwind')

/** @type {import('tailwindcss').Config} */
export default {
  content: [],
  theme: {
    extend: {},
  },
  plugins: [require('daisyui'), addIconSelectors(['logos', 'carbon'])],
  daisyui: {
    themes: ['aqua']
  }
}

