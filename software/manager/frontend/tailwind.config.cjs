/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    styled: true,
    themes: [
      {
        "mytheme": {
          "primary": "#00ff00",
          "secondary": "rgb(100, 100, 0)",
          "accent": "rgb(200, 200, 2000)",
          "neutral": "#3d4451",
          "base-100": "rgb(20, 20, 20)",
          "base-200": "rgb(10, 10, 10)",
          "base-300": "rgb(5, 5, 5)",
          "--rounded-btn": "0rem",
        }
      },
      "business"
    ],
    base: true,
    utils: true,
    logs: true,
    rtl: false,
    prefix: "",
    darkTheme: "",
  },
}
