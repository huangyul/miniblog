<script setup lang="ts">
import zhCN from 'ant-design-vue/es/locale/zh_CN'
import 'dayjs/locale/zh-cn'
import dayjs from 'dayjs'
import { ref } from 'vue'

dayjs.locale('zh-cn');
const locale = ref(zhCN)
</script>

<template>
  <AConfigProvider :locale="locale">
    <RouterView />
  </AConfigProvider>
</template>
