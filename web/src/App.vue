<script setup lang="ts">
import { useRouter } from 'vue-router';


const router = useRouter()

const toPage = (path: string) => {
  router.push({ path: path })
}

</script>

<template>
  <div>
    <p @click="toPage('/home')">
      home
    </p>
    <p @click="toPage('/user/23')">
      user
    </p>
    <RouterView></RouterView>
  </div>
</template>
