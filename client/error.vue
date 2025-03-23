<script setup lang="ts">
  import type { NuxtError } from '#app';
  import { clearError } from '#imports';

  const props = defineProps({
    error: Object as () => NuxtError,
  });

  const handleError = () => clearError({ redirect: '/' });
</script>

<template>
  <UContainer class="flex flex-col items-center justify-center mt-60">
    <h2>{{ $t('error-code') }} - {{ props.error!.statusCode }}</h2>
    <h3>{{ $t('error-message') }} - {{ props.error!.statusMessage }}</h3>
    <UButton @click="handleError" class="mt-10">{{ $t('clear-errors') }}</UButton>
  </UContainer>
</template>
