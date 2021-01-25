module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      colors: {
        urbano: {
          menubg: '#DBDBDB',
          menulinks: '#9D9D9D',
          prev: '#7C7C7C',
          green: '#12B6AD',
        }
      },
      width: {
        '44r': '44rem',
      },
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
