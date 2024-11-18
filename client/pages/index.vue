<script setup lang="ts">

const newGame = ref<Game>({} as Game)
onMounted(() => {
  /*const socket = connectToSocket('201');
  socket.onopen = () => {
    const handshake = new SocketMessage('handshake', new HandshakeContent());
    const throw1 = new SocketMessage('throw', new ThrowContent(12, 1));
    socket.send(JSON.stringify(handshake));
    console.log("conneted to websocket");
    socket.send(JSON.stringify(throw1));
  }
  socket.onmessage = (evt) => {
    console.log(evt);
    const data = JSON.parse(evt.data);
    messages.value = data.currentPlayer.score;
  }*/
})

const onCreateGame = function () {
  const players: Player[] = [
    {} as Player
  ];
  const game: Game = {} as Game;
  game.players = players;
  fetch('http://localhost:8080/api/game', {
    method: 'POST',
    body: JSON.stringify(game)
  }).then(response => {
    response.json()
  }).then(data => {
    console.log(data)
  })
}
</script>
<template>
  <h1 class="text-center">Hello world</h1>
  <input class="input" type="text" v-model="newGame.name" />
  <input class="input" type="text" v-model="newGame.gameMode" />
  <button class="btn-primary btn" v-on:click="onCreateGame" />
</template>
