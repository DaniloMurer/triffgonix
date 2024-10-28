<script setup>
import { Game, HandshakeContent, Player, SocketMessage, ThrowContent } from '~/common/data';


const messages = ref()

onMounted(() => {
  const socket = connectToSocket('201');
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
  }
})

const onCreateGame = function () {
  const players = [
    new Player(1, 'test'),
    new Player(2, 'test2')
  ];
  const game = new Game('testico', 'x01', 401, players);

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
  <span>{{ messages }}</span>
  <button class="btn-primary btn" v-on:click="onCreateGame" />
</template>
