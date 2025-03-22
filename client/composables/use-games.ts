import {getApiGame, type ModelsGame} from "#shared/utils";

export const useGames = () => {
  const games = ref<ModelsGame[]>([])
  const loading = ref(false)

  const fetchGames = async () => {
    loading.value = true
    try {
      const response = await getApiGame();
      if (response.data) {
        games.value = response.data;
      }
    } catch (error) {
      console.error('Error fetching games:', error)
    } finally {
      loading.value = false
    }
  }

  onMounted(() => {
    fetchGames()
  });

  return {
    games,
    loading,
    fetchGames
  }
};
