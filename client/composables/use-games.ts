import { type DtoGame, getApiGame, type ModelsGame, postApiGame } from '#shared/utils';

export const useGames = () => {
  const games = ref<ModelsGame[]>([]);
  const loading = ref(false);

  const fetchGames = async () => {
    loading.value = true;
    try {
      const response = await getApiGame();
      if (response.data) {
        games.value = response.data;
      }
    } catch (error) {
      console.error('Error fetching games:', error);
    } finally {
      loading.value = false;
    }
  };

  onMounted(() => {
    void fetchGames();
  });

  return {
    games,
    loading,
    fetchGames,
  };
};

export const createGame = async (game: DtoGame) => {
  const response = await postApiGame({ body: game });
  console.log(response);
  if (response.response.ok) {
    return response.data;
  } else {
    throw new Error('error while creating game');
  }
};
