/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{html,js,templ}"],
    theme: {
      extend: {},
    },
    plugins: [
        require("@tailwindcss/typography"),
    ],
}
