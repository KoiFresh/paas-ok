<template>
  <div class="footer">
    <ULink
      to="/imprint"
      active-class="text-primary"
      inactive-class="text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-200"
    >
      Imprint
    </ULink>

    <ULink
      to="/privacy"
      active-class="text-primary"
      inactive-class="text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-200"
    >
      Privacy Policy
    </ULink>

    <ULink
      to="/contact"
      active-class="text-primary"
      inactive-class="text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-200"
    >
      Contact
    </ULink>
  </div>
</template>

<style scoped>
.footer {
  display: flex;
  gap: 1em;
  justify-content: center;
  margin-top: 2em;
}
</style>
