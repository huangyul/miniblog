<script setup lang="ts">
import { ref } from 'vue'
import type { Dayjs } from 'dayjs'
import { SelectProps } from 'ant-design-vue'
import { changeLanguage, type MessageKey } from '@/lang'
const value1 = ref<Dayjs>()


const value2 = ref<string>('')

const options = ref<SelectProps['options']>([
  {
    value: 'zh',
    label: '简体中文'
  },
  {
    value: 'zh_HK',
    label: '繁体中文（中国香港）'
  },
  {
    value: 'en',
    label: '英语'
  },
  {
    value: 'ja',
    label: '日语'
  }
])

const handleChangeLanguage = (value: MessageKey) => {
  changeLanguage(value)
}
</script>

<template>
  <div>home</div>

  <a-select ref="select" v-model:value="value2" :options style="width: 120px" @change="handleChangeLanguage">
  </a-select>

  <a-date-picker v-model:value="value1" />

  <p>{{ $t('dashboard') }}</p>

</template>
