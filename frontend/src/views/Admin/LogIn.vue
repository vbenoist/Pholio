<template>
    <div class="container">
      <div class="container__title">
        Connection
      </div>
      <form
        class="container__login-form"
        @submit.prevent="handleSubmit"
      >
        <div class="container__login-form__ids">
          <label for="form-login-id">
            Identifiant
          </label>
          <input
            id="form-login-id"
            :class="['container__login-form__ids__username', loginError ? 'input-error' : '']"
            type="text"
            required
            v-model="userIds.username"
          />

          <label for="form-login-password">
            Mot de passe
          </label>
          <input
            id="form-login-password"
            :class="['container__login-form__ids__password', loginError ? 'input-error' : '']"
            type="password"
            required
            v-model="userIds.password"
          />
        </div>

        <span
          v-if="loginError"
          class="container__login-form__error"
        >
          Identifiants non reconnus
        </span>

        <button class="container__login-form__validate" type="submit" @click="() => false">
          <v-icon v-if="loginWorking" name="ri-loader-4-line" animation="spin" scale="1.2" />
          <span v-else>Se connecter</span>
        </button>
      </form>

    </div>
</template>

<script setup lang="ts">
import { inject, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import type { User } from "@/models/api/user"
import type { Auth } from '@/plugins/auth'

const auth = inject('$auth') as Auth
const router = useRouter()

const userIds = ref<User>({
  username: '',
  password: ''
})

const loginError = ref<boolean>(false)
const loginWorking = ref<boolean>(false)

const handleSubmit = () => {
  loginWorking.value = true

  auth.login({
    username: userIds.value.username,
    password: userIds.value.password
  }).then(success => {
    loginWorking.value = false
    if(!success) {
      loginError.value = true
      return
    }

    const redirect = auth.getOverwrittenRoute()
    router.replace({ name: redirect?.name ?? 'RECENT' })
  })
}

watch(userIds, () => {
  if(loginError.value) loginError.value = false
}, { deep: true })

</script>

<style lang="scss" scoped>
.container {
  position: relative;
  display: flex;
  align-items: center;
  margin-top: 80px;
  flex-flow: column;

  &__title {
    width: 400px;
  }

  &__login-form {
    width: 400px;
    height: 260px;
    margin-top: 10px;
    padding: 16px;

    display: flex;
    flex-flow: column nowrap;
    justify-content: space-between;

    border: 2px solid #0f0f13;
    border-radius: 5px;
    box-shadow: 1px 1px 14px 12px rgba(15,15,19,0.75);
    -webkit-box-shadow: 1px 1px 14px 12px rgba(15,15,19,0.75);
    -moz-box-shadow: 1px 1px 14px 12px rgba(15,15,19,0.75);

    &__ids {
      display: flex;
      flex-flow: column nowrap;
      justify-content: center;


      &__username, &__password {
        background-color: #585858;
        color: white;
        border: 1px #313131;
        border-radius: 3px;
        height: 30px;
        margin: 4px 0;
      }

      input[type=text], input[type=password] {
        padding: 0px 10px;
        margin: 8px 0;
      }
    }

    &__error {
      color: rgb(247, 101, 101);
    }

    &__validate {
      cursor: pointer;
      margin-left: auto;
      height: 30px;
      width: 100px;
      background-color: #585858;
      border: 1px #313131;
      border-radius: 3px;
      color: white;
    }
  }
}


.input-error {
  border: 1px solid rgb(247, 101, 101);
}
</style>
