<script setup>

const messages = ref()

onMounted(() => {
  const socket = connectToSocket('201');
  socket.onopen = () => {
    const handshake = {
      type: "handshake",
      data: {
        content: "data"
      }
    };
    const throw1 = {
      type: "throw",
      data: {
        content: "data"
      }
    };
    socket.send(JSON.stringify(handshake));
    console.log("conneted to websocket");
    socket.send(JSON.stringify(throw1));
  }
  socket.onmessage = (evt) => {
    console.log(evt);
    const data = JSON.parse(evt.data);
    messages.value = data.currentPlayer.score;
  }

  const newGame = {
    name: "testico",
    gameMode: "x01",
    startingScore: 401,
    players: [
      { Id: 1, name: "test" },
      { Id: 1, name: "test" }
    ]
  }
})

const onCreateGame = function () {
  const newGame = {
    name: "testico",
    gameMode: "x01",
    startingScore: 401,
    players: [
      { Id: 1, name: "test" },
      { Id: 2, name: "test2" }
    ]
  }

  fetch('http://localhost:8080/api/game', {
    method: 'POST',
    body: JSON.stringify(newGame)
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
