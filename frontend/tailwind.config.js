/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        display: ['Syne', 'sans-serif'],
        body: ['Plus Jakarta Sans', 'Sarabun', 'sans-serif'],
      },
      colors: {
        indigo: { DEFAULT: '#4F46E5' },
        purple: { DEFAULT: '#a855f7' },
        cyan: { DEFAULT: '#06b6d4' },
      },
    },
  },
  plugins: [],
}
