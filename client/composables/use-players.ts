import {getApiUser, type ModelsPlayer} from "#shared/utils";

export const usePlayers = () => {
  const players = ref<ModelsPlayer[]>([]);

  const fetchPlayers = async () => {
    const response = await getApiUser();
    if (response.data) {
      players.value = response.data;
    }
  }

  onMounted(() => {
    fetchPlayers();
  })

  return {
    players,
    fetchPlayers
  }
}
