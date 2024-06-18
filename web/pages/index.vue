<template>
  <div>
    <WelcomeHeader />
    <div class="sections">
      <WelcomeAboutUs />
      <WelcomeOurPrints />
      <WelcomePricing />
      <WelcomeGettingStarted />
    </div>
    <PageFooter />
  </div>
</template>

<style scoped>
.sections {
  display: flex;
  flex-direction: column;
  gap: 3em;
}
</style>
