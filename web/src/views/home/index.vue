<script setup lang="ts">
import { test } from '@/api/basic'
import { onMounted } from 'vue';


onMounted(() => {
  test()
})

</script>

<template>
  <div class="text-12">home</div>
</template>
