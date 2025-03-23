import { type DtoPlayer, getApiUser } from '#shared/utils';

export const usePlayerService = () => {
  const players = ref<DtoPlayer[]>([]);

  const fetchPlayers = async () => {
    const response = await getApiUser();
    if (response.data) {
      return response.data;
    }
    return [];
  };

  return {
    players,
    fetchPlayers,
  };
};
