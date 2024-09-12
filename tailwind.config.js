/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./ui/html/**/*.tmpl}", "./**/**/**/*.tmpl"],
  theme: {
    extend: {
      colors: {
        primary: "#14532d",
        secondary: "#854d0e",
      },
    },
  },
  plugins: [],
};
