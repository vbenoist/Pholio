<template>
  <div class="container">
    <img class="container__background" src="@/assets/img/IMG_2813.jpg" />
    <img class="container__image" src="@/assets/img/IMG_2813.jpg" />

    <span class="container__title">Portfolio Photo Vincent BENOIST</span>
  </div>
</template>

<style lang="scss" scoped>
.container {
  display: flex;
  height: 680px;
  width: 100%;
  margin: auto;
  justify-content: center;
  overflow: hidden;

  &__image {
    position: absolute;
    z-index: 1;
    height: inherit;
  }

  &__background {
    position: absolute;
    z-index: 0;
    height: inherit;
    width: 98vw;
    filter: blur(15px);
    -webkit-filter: blur(15px);
  }

  &__title {
    margin: auto auto 36px;
    font-family: 'JostSemiBold';
    font-size: 16px;
    color: #ffffff;
  }

  span {
    z-index: 2;
  }
}
</style>
