<template>
  <div v-if="draftRecord" class="photo-card">
    <form class="photo-card__form" @submit.prevent="handleSubmit">
      <div class="photo-card__form__section photo-card__form__section__date">
        <label :for="`draftrecord-date-${getRecordId(draftRecord)}`" @click="triggerCalendar()">
          Date: {{ humanFormattedDate(draftRecord.date) }}
          <v-icon name="bi-calendar-date-fill" scale="1.3" />
        </label>
        <input
          :id="`draftrecord-date-${getRecordId(draftRecord)}`"
          class="photo-card__form__section__date__input"
          ref="draftrecord-date"
          type="date"
          required
          :value="inputFormattedDate(draftRecord.date)"
          v-model="inputDate"
        />
      </div>
      <div class="photo-card__form__section photo-card__form__section__location">
        <label :for="`draftrecord-location-${getRecordId(draftRecord)}`">Lieu:</label>
        <input
          :id="`draftrecord-location-${getRecordId(draftRecord)}`"
          class="photo-card__form__section__location__input"
          type="text"
          required
          v-model="draftRecord.location"
        />
      </div>
      <div class="photo-card__form__section__description">
        <label :for="`draftrecord-description-${getRecordId(draftRecord)}`">Description:</label>
        <textarea
          :id="`draftrecord-description-${getRecordId(draftRecord)}`"
          class="photo-card__form__section__description__input"
          type="text"
          v-model="draftRecord.description"
        ></textarea>
      </div>

      <button class="photo-card__form__validate" type="submit" @click="triggerValidation">
        <v-icon
          :name="submitIcon"
          :animation="submitAnimation"
          :hover="submitAnimationHover"
          scale="1.3"
        />
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
import type { DetailedRecord } from '@/models/record'
import {
  detailedRecordToApiDetailedRecord,
  draftRecordToApiRecord
} from '@/transformers/record'

const draftRecord = defineModel<DraftRecord | DetailedRecord>({ default: null })
const calendarInputRef = useTemplateRef<HTMLInputElement>(`draftrecord-date`)
const apiResolver = inject('$apiResolver') as ApiResolver

const formRules = {
  description: { maxLengthValue: maxLength(100) },
  location: { required, minLengthValue: minLength(4), maxLengthValue: maxLength(50) },
  date: { required },
}

const getRecordId = (record: DraftRecord | DetailedRecord): string => {
  if('draftId' in record) {
    return record.draftId
  } else {
    return record.id
  }
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
    if (!draftRecord.value) return inputFormattedDate(new Date())
    return inputFormattedDate(draftRecord.value.date)
  },
  set(val: string) {
    draftRecord.value!.date = new Date(val)
  },
})

const handleSubmit = async () => {
  if('draftId' in draftRecord.value) {
    await _submitNewDraft()
  } else {
    await _submitDraft()
  }
}

const _submitDraft = async() => {
  const lcDraftRecord = draftRecord.value as DetailedRecord
  const apiFormattedRecord = detailedRecordToApiDetailedRecord(lcDraftRecord)

  if (draftRecord.value.status === 'PENDING' || draftRecord.value.status === 'FAILED') {
    draftRecord.value.status = 'SENDING'
    const resId = await apiResolver.updateRecord(apiFormattedRecord)

    if (!resId) {
      lcDraftRecord.status = 'FAILED'
      return
    } else {
      lcDraftRecord.status = 'SENT'
    }
  }
}

const _submitNewDraft = async () => {
  const lcDraftRecord = draftRecord.value as DraftRecord
  const apiFormattedRecord = draftRecordToApiRecord(lcDraftRecord)

  /* If draft has already been saved once, juste updating */
  if(lcDraftRecord.saved) {
    draftRecord.value.status = 'SENDING'
    const success = await apiResolver.updateDraftRecord(apiFormattedRecord)

    if(!success) {
      lcDraftRecord.status = 'FAILED'
    } else {
      lcDraftRecord.status = 'SENT'
    }
    return
  }
  /* By default, adding a new record */
  else {
    draftRecord.value.file.status = 'PENDING'
    draftRecord.value.file.status = 'SENDING'
    const resId = await apiResolver.addRecord(apiFormattedRecord)

    if (!resId) {
      lcDraftRecord.status = 'FAILED'
      return
    } else {
      lcDraftRecord.draftId = resId
      lcDraftRecord.status = 'SENT'
      lcDraftRecord.saved = true
    }

    /* Linking photo to freshly saved record */
    await _saveLinkedImage()
  }
}

const _saveLinkedImage = (): Promise<void> => {
  return apiResolver
    .linkImageRecord((draftRecord.value as DraftRecord).draftId, draftRecord.value.file)
    .then((success) => {
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
  if (v$.value.$error) return 'md-error-outlined'

  if (draftRecord.value.status === 'SENDING'
    || draftRecord.value.file.status === 'SENDING'
  ) return 'md-pending'

  if (draftRecord.value.status === 'FAILED'
    || draftRecord.value.file.status === 'FAILED'
  ) return 'bi-cloud-slash-fill'

  /* In case of new record, that's the photo supply which will determine if the process has been completed */
  if('draftId' in draftRecord.value) {
    if (draftRecord.value.file.status === 'SENT') return 'bi-cloud-check-fill'
  } else if (draftRecord.value.status === 'SENT') return 'bi-cloud-check-fill'

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

    input[type='text'],
    textarea {
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
