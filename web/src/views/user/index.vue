<script setup lang="ts">
import { useRoute } from 'vue-router';


const route = useRoute()
</script>

<template>
  <div>user</div>
  <div>id: {{ route.params.id }}</div>
</template>
