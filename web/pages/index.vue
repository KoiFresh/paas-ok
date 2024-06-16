<template>
  <div>
    <BillboardPane />
    <div class="sections">
      <AboutUs />
      <OurPrints />
      <PricingInformationPane />
      <GettingStartedPane />
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
