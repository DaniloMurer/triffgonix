<script setup lang="ts">
  import CreateGameDialog from './components/dialogs/create-game.dialog.vue';
  import { ref, useI18n } from '#imports';

  const i18n = useI18n();
  const availableLocales = i18n.locales.value;

  const localeItems = [];

  for (const locale of availableLocales) {
    localeItems.push({
      label: locale.name,
      icon: 'i-lucide-globe',
      onSelect: () => i18n.setLocale(locale.code),
      active: locale.code === i18n.locale.value,
    });
  }
  const items = ref([
    {
      label: 'Triffgonix',
      icon: 'i-lucide-home',
      to: '/',
    },
    {
      label: i18n.t('language'),
      icon: 'i-lucide-globe',
      children: localeItems,
    },
    {
      slot: 'create-game',
    },
  ]);
</script>

<template>
  <UApp>
    <div class="m-3">
      <UNavigationMenu
        color="primary"
        variant="pill"
        :items="items"
        orientation="horizontal"
        class="w-full justify-center"
      >
        <template #create-game>
          <create-game-dialog />
        </template>
      </UNavigationMenu>
      <NuxtRouteAnnouncer />
      <NuxtPage />
    </div>
  </UApp>
</template>
