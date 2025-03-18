<template>
  <div v-if="draftRecord" class="photo-card">
    <form
      class="photo-card__form"
      @submit.prevent="handleSubmit"
    >
      <div class="photo-card__form__section photo-card__form__section__date">
        <label :for="`draftrecord-date-${draftRecord.draftId}`" @click="triggerCalendar()">
          Date: {{ humanFormattedDate(draftRecord.date) }}
          <v-icon name="bi-calendar-date-fill" scale="1.3" />
        </label>
        <input
          :id="`draftrecord-date-${draftRecord.draftId}`"
          class="photo-card__form__section__date__input"
          ref="draftrecord-date"
          type="date"
          required
          :value="inputFormattedDate(draftRecord.date)"
          v-model="inputDate"
        />
      </div>
      <div class="photo-card__form__section photo-card__form__section__location">
        <label :for="`draftrecord-location-${draftRecord.draftId}`">Lieu:</label>
        <input
          :id="`draftrecord-location-${draftRecord.draftId}`"
          class="photo-card__form__section__location__input"
          type="text"
          required
          v-model="draftRecord.location"
        />
      </div>
      <div class="photo-card__form__section__description">
        <label :for="`draftrecord-description-${draftRecord.draftId}`">Description:</label>
        <textarea
          :id="`draftrecord-description-${draftRecord.draftId}`"
          class="photo-card__form__section__description__input"
          type="text"
          v-model="draftRecord.description"
        ></textarea>
      </div>

      <button class="photo-card__form__validate" type="submit" @click="triggerValidation">
        <v-icon :name="submitIcon" :animation="submitAnimation" :hover="submitAnimationHover" scale="1.3" />
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { computed, defineModel, inject, useTemplateRef } from 'vue'
import { useVuelidate } from '@vuelidate/core'
import { maxLength, minLength, required } from '@vuelidate/validators'
import type { ApiResolver } from '@/plugins/apiResolver'
import type { DraftRecord } from '@/models/record'
import { draftRecordToApiRecord } from '@/transformers/record'

const draftRecord = defineModel<DraftRecord>({ default: null })
const calendarInputRef = useTemplateRef<HTMLInputElement>(`draftrecord-date`)
const apiResolver = inject('$apiResolver') as ApiResolver

const formRules = {
  description: { maxLengthValue: maxLength(100) },
  location: { required, minLengthValue: minLength(4), maxLengthValue: maxLength(50) },
  date: { required }
}

const humanFormattedDate = (date: Date): string => {
  return date.toLocaleDateString()
}

const inputFormattedDate = (date: Date): string => {
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')

  return `${year}-${month}-${day}`
}

const inputDate = computed<string>({
  get() {
    if(!draftRecord.value) return inputFormattedDate(new Date())
    return inputFormattedDate(draftRecord.value.date)
  },
  set(val: string) {
    draftRecord.value!.date = new Date(val)
  }
})

const handleSubmit = async () => {
  draftRecord.value.file.status = 'SENDING'

  const apiFormattedRecord = draftRecordToApiRecord(draftRecord.value)

  /* If first attempt or previous attempted has failed on sending record data */
  if(draftRecord.value.status === 'PENDING' || draftRecord.value.status === 'FAILED') {
    const resId = await apiResolver.addRecord(apiFormattedRecord)

    if(!resId) {
      draftRecord.value.status = 'FAILED'
      return
    } else {
      draftRecord.value.draftId = resId
      draftRecord.value.status = 'SENT'
    }
  }

  /* Linking photo to freshly saved record */
  await saveLinkedImage()
}

const saveLinkedImage = (): Promise<void> => {
  return apiResolver.linkImageRecord(draftRecord.value.draftId, draftRecord.value.file)
    .then(success => {
      draftRecord.value.file.status = success ? 'SENT' : 'FAILED'
    })
}

const triggerValidation = () => {
  v$.value.$touch()
}

const triggerCalendar = () => {
  calendarInputRef.value?.showPicker()
}

const submitIcon = computed<string>(() => {
  if(v$.value.$error) return 'md-error-outlined'
  if(draftRecord.value.status === 'FAILED') return 'bi-cloud-slash-fill'
  if(draftRecord.value.file.status === 'SENDING') return 'md-pending'
  if(draftRecord.value.file.status === 'FAILED') return 'bi-cloud-slash-fill'
  if(draftRecord.value.file.status === 'SENT') return 'bi-cloud-check-fill'
  return 'bi-cloud-arrow-up-fill'
})

const submitAnimation = computed<string>(() => {
  return draftRecord.value.file.status === 'SENDING' ? 'float' : 'wrench'
})

const submitAnimationHover = computed<boolean>(() => {
  return draftRecord.value.file.status !== 'SENDING'
})

const v$ = useVuelidate(formRules, draftRecord)

</script>

<style scoped lang="scss">

.photo-card {
  margin: 0 12px;
  display: flex;
  flex-flow: column nowrap;
  height: 200px;
  width: 400px;

  &__form {
    position: relative;
    display: flex;
    flex-flow: column nowrap;

    input[type=text], textarea {
      background-color: #585858;
      color: white;
      border: 1px #313131;
      border-radius: 3px;
    }

    &__section {
      display: flex;
      width: 100%;

      &__date {
        label {
          cursor: pointer;
        }

        &__input {
          opacity: 0;
          width: 0;
          height: 0;
        }
      }

      &__location {
        &__input {
          margin-left: 10px;
          width: 100%;
        }
      }

      &__description {
        display: flex;
        flex-flow: column;

        &__input {
          resize: none;
          height: 100px;
        }
      }
    }

    &__validate {
      position: absolute;
      right: 0;
      top: 0;
      background: none;
      color: inherit;
      border: none;
      padding: 0;
      cursor: pointer;
      outline: inherit;
    }
  }
}

</style>
