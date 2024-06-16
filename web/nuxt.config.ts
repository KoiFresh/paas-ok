// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: true,
  devtools: { enabled: true },
  modules: ["@nuxt/ui", "@nuxt/image"],
  components: true,
  css: ['~/assets/css/main.css'],
})